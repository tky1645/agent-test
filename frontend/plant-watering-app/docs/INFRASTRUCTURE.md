# インフラ構成ドキュメント

## 概要

このドキュメントでは、植物水やり管理フロントエンドアプリケーションのインフラ構成について説明します。

## アーキテクチャ概要

```
┌─────────────────┐    ┌─────────────────┐    ┌─────────────────┐
│   ユーザー      │    │   nginx         │    │   React App     │
│   (ブラウザ)    │◄──►│   (Webサーバー) │◄──►│   (静的ファイル) │
└─────────────────┘    └─────────────────┘    └─────────────────┘
                                │
                                ▼
                       ┌─────────────────┐
                       │   Docker        │
                       │   Container     │
                       └─────────────────┘
```

## Docker構成

### Dockerfile

本番環境用のマルチステージビルド構成：

```dockerfile
# ビルドステージ
FROM node:18-alpine AS builder
WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production
COPY . .
RUN npm run build

# 本番ステージ
FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
```

**特徴:**
- Node.js 18 Alpine Linuxベース（軽量）
- マルチステージビルドによる最適化
- 本番用依存関係のみインストール
- nginx Alpine Linuxで静的ファイル配信

### docker-compose.yml

開発環境用の構成：

```yaml
version: '3.8'
services:
  frontend:
    build: .
    ports:
      - "3000:80"
    environment:
      - NODE_ENV=production
```

**用途:**
- ローカル開発環境でのコンテナテスト
- 本番環境構成の検証
- CI/CDパイプラインでの自動テスト

## nginx構成

### nginx.conf

```nginx
events {
    worker_connections 1024;
}

http {
    include       /etc/nginx/mime.types;
    default_type  application/octet-stream;

    server {
        listen 80;
        server_name localhost;
        root /usr/share/nginx/html;
        index index.html;

        # SPA用のフォールバック設定
        location / {
            try_files $uri $uri/ /index.html;
        }
    }
}
```

**設定のポイント:**
- **SPA対応**: `try_files`でReact Routerのクライアントサイドルーティングをサポート
- **静的ファイル配信**: `/usr/share/nginx/html`から配信
- **MIME Type対応**: 適切なContent-Typeヘッダーを設定
- **軽量構成**: 必要最小限の設定でパフォーマンス重視

## ビルドプロセス

### 開発環境

```bash
# 依存関係インストール
npm install

# 開発サーバー起動（ホットリロード対応）
npm run dev
# → http://localhost:5173 でアクセス可能

# 型チェック
npm run type-check

# リンター実行
npm run lint
```

### 本番環境

```bash
# 本番ビルド
npm run build
# → dist/ ディレクトリに最適化されたファイルを生成

# プレビュー（本番ビルドの確認）
npm run preview
# → http://localhost:4173 でアクセス可能

# Dockerイメージビルド
docker build -t plant-watering-app .

# コンテナ起動
docker run -p 3000:80 plant-watering-app
# → http://localhost:3000 でアクセス可能
```

## デプロイメント戦略

### ローカル開発

1. **開発サーバー**: Vite開発サーバーでホットリロード
2. **モックAPI**: localStorageベースのデータ永続化
3. **型安全性**: TypeScriptによるコンパイル時チェック

### ステージング環境

1. **Dockerコンテナ**: 本番環境と同等の構成
2. **nginx**: 静的ファイル配信とSPAルーティング
3. **環境変数**: `.env`ファイルによる設定管理

### 本番環境

1. **コンテナオーケストレーション**: Kubernetes/Docker Swarm対応
2. **CDN統合**: CloudFront等との連携可能
3. **ヘルスチェック**: nginx statusページ対応
4. **ログ管理**: nginx アクセスログ・エラーログ

## パフォーマンス最適化

### ビルド最適化

- **Tree Shaking**: 未使用コードの自動削除
- **Code Splitting**: 動的インポートによる分割読み込み
- **Asset Optimization**: 画像・CSS・JSの圧縮
- **Bundle Analysis**: `npm run build`でバンドルサイズ確認

### 実行時最適化

- **Gzip圧縮**: nginxによる自動圧縮
- **キャッシュ戦略**: 静的アセットの長期キャッシュ
- **レスポンシブ画像**: 適切なサイズでの画像配信

## セキュリティ

### コンテナセキュリティ

- **非rootユーザー**: nginxプロセスの権限最小化
- **Alpine Linux**: 軽量で脆弱性の少ないベースイメージ
- **最小権限**: 必要最小限のパッケージのみインストール

### Webセキュリティ

- **HTTPS対応**: 本番環境でのTLS終端
- **セキュリティヘッダー**: CSP、HSTS等の設定可能
- **静的ファイル**: 実行可能ファイルの配信防止

## 監視・ログ

### アプリケーション監視

- **ヘルスチェック**: `/health`エンドポイント（実装可能）
- **メトリクス**: Prometheusメトリクス対応
- **エラートラッキング**: Sentry等の統合可能

### インフラ監視

- **コンテナ監視**: Docker stats、cAdvisor
- **リソース監視**: CPU、メモリ、ディスク使用量
- **ネットワーク監視**: nginx アクセスログ分析

## トラブルシューティング

### よくある問題

1. **ルーティングエラー**: nginx設定の`try_files`確認
2. **ビルドエラー**: Node.jsバージョンとnpm依存関係確認
3. **コンテナ起動失敗**: Dockerログとポート競合確認

### デバッグ方法

```bash
# コンテナログ確認
docker logs <container_id>

# nginx設定テスト
nginx -t

# ビルド詳細ログ
npm run build -- --verbose
```
