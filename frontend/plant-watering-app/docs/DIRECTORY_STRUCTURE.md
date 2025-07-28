# ディレクトリ構成ドキュメント

## 概要

このドキュメントでは、植物水やり管理フロントエンドアプリケーションのディレクトリ構成とファイル構成について詳しく説明します。

## プロジェクト全体構成

```
frontend/plant-watering-app/
├── docs/                          # ドキュメント
│   ├── INFRASTRUCTURE.md          # インフラ構成ドキュメント
│   └── DIRECTORY_STRUCTURE.md     # このファイル
├── public/                        # 静的ファイル
├── src/                          # ソースコード
├── dist/                         # ビルド成果物（自動生成）
├── node_modules/                 # 依存関係（自動生成）
├── package.json                  # プロジェクト設定・依存関係
├── package-lock.json            # 依存関係ロックファイル
├── vite.config.ts               # Viteビルド設定
├── tsconfig.json                # TypeScript設定
├── tsconfig.app.json            # アプリ用TypeScript設定
├── tsconfig.node.json           # Node.js用TypeScript設定
├── tailwind.config.js           # Tailwind CSS設定
├── postcss.config.js            # PostCSS設定
├── eslint.config.js             # ESLint設定
├── components.json              # shadcn/ui設定
├── Dockerfile                   # Docker設定
├── docker-compose.yml           # Docker Compose設定
├── nginx.conf                   # nginx設定
├── .dockerignore               # Docker除外ファイル
├── .gitignore                  # Git除外ファイル
├── .env                        # 環境変数
├── index.html                  # HTMLエントリーポイント
└── README.md                   # プロジェクト説明
```

## src/ ディレクトリ詳細

### 全体構成

```
src/
├── components/                   # Reactコンポーネント
│   ├── ui/                      # 再利用可能UIコンポーネント
│   ├── Header.tsx               # ヘッダーコンポーネント
│   ├── LoginForm.tsx            # ログインフォーム
│   ├── PlantCard.tsx            # 植物カードコンポーネント
│   ├── PlantList.tsx            # 植物一覧コンポーネント
│   ├── PlantRegistration.tsx    # 植物登録フォーム
│   └── WateringButton.tsx       # 水やりボタンコンポーネント
├── contexts/                    # React Context
│   └── AuthContext.tsx          # 認証状態管理
├── hooks/                       # カスタムフック
│   ├── use-mobile.tsx           # モバイル判定フック
│   └── use-toast.ts             # トースト通知フック
├── lib/                         # ユーティリティライブラリ
│   └── utils.ts                 # 共通ユーティリティ関数
├── services/                    # 外部サービス連携
│   ├── api.ts                   # API通信サービス
│   ├── auth.ts                  # 認証サービス
│   └── s3.ts                    # S3アップロードサービス
├── types/                       # TypeScript型定義
│   └── index.ts                 # 共通型定義
├── assets/                      # 静的アセット
│   └── react.svg                # Reactロゴ
├── App.tsx                      # メインアプリケーションコンポーネント
├── App.css                      # アプリケーション固有CSS
├── main.tsx                     # アプリケーションエントリーポイント
├── index.css                    # グローバルCSS・Tailwind設定
└── vite-env.d.ts               # Vite環境型定義
```

## 各ディレクトリの役割

### `/components` - Reactコンポーネント

アプリケーションのUI要素を構成するReactコンポーネントを格納。

#### `/components/ui` - 再利用可能UIコンポーネント

shadcn/uiライブラリのコンポーネントを格納。デザインシステムの基盤。

```
ui/
├── accordion.tsx               # アコーディオンコンポーネント
├── alert-dialog.tsx           # アラートダイアログ
├── alert.tsx                  # アラート表示
├── avatar.tsx                 # アバター表示
├── badge.tsx                  # バッジ表示
├── button.tsx                 # ボタンコンポーネント
├── card.tsx                   # カードレイアウト
├── checkbox.tsx               # チェックボックス
├── dialog.tsx                 # ダイアログ
├── form.tsx                   # フォーム要素
├── input.tsx                  # 入力フィールド
├── label.tsx                  # ラベル
├── select.tsx                 # セレクトボックス
├── textarea.tsx               # テキストエリア
├── toast.tsx                  # トースト通知
├── toaster.tsx                # トースト管理
└── ...                        # その他UIコンポーネント
```

**特徴:**
- **一貫性**: 統一されたデザインシステム
- **再利用性**: アプリケーション全体で使用可能
- **カスタマイズ性**: Tailwind CSSによるスタイル調整
- **アクセシビリティ**: WAI-ARIA準拠

#### アプリケーション固有コンポーネント

```
components/
├── Header.tsx                  # アプリケーションヘッダー
├── LoginForm.tsx              # ログインフォーム
├── PlantCard.tsx              # 植物情報カード
├── PlantList.tsx              # 植物一覧表示
├── PlantRegistration.tsx      # 植物新規登録フォーム
└── WateringButton.tsx         # 水やり記録ボタン
```

**役割:**
- **Header.tsx**: ナビゲーション、ユーザー情報表示
- **LoginForm.tsx**: 認証フォーム（モック実装）
- **PlantCard.tsx**: 個別植物の情報表示・操作
- **PlantList.tsx**: 植物コレクションの一覧表示
- **PlantRegistration.tsx**: 新規植物登録フォーム
- **WateringButton.tsx**: ワンクリック水やり記録

### `/contexts` - React Context

アプリケーション全体の状態管理を担当。

```
contexts/
└── AuthContext.tsx            # 認証状態管理
```

**AuthContext.tsx の役割:**
- ユーザー認証状態の管理
- ログイン・ログアウト機能
- 認証情報の永続化（localStorage）
- 子コンポーネントへの認証状態提供

### `/hooks` - カスタムフック

再利用可能なロジックを提供するカスタムフック。

```
hooks/
├── use-mobile.tsx             # モバイルデバイス判定
└── use-toast.ts               # トースト通知管理
```

**特徴:**
- **use-mobile.tsx**: レスポンシブデザイン対応
- **use-toast.ts**: ユーザー通知の統一管理

### `/lib` - ユーティリティライブラリ

共通で使用される関数やユーティリティを格納。

```
lib/
└── utils.ts                   # 共通ユーティリティ関数
```

**utils.ts の内容:**
- CSS クラス名の結合（clsx, tailwind-merge）
- 共通バリデーション関数
- 日付・時刻フォーマット関数

### `/services` - 外部サービス連携

外部API、認証、ストレージとの連携を担当。

```
services/
├── api.ts                     # バックエンドAPI通信
├── auth.ts                    # 認証サービス
└── s3.ts                      # 画像アップロードサービス
```

#### api.ts - API通信サービス

```typescript
class ApiService {
  async getPlants(): Promise<Plant[]>
  async createPlant(plantData: PlantCreate): Promise<Plant>
  async getPlantStatus(plantId: string): Promise<PlantStatus>
  async addWateringRecord(plantId: string, recordData: WateringRecordCreate): Promise<WateringRecord>
  async getWateringRecords(plantId: string): Promise<WateringRecord[]>
}
```

**現在の実装:**
- localStorageベースのモック実装
- 将来的にGoバックエンドAPIとの統合予定

#### auth.ts - 認証サービス

```typescript
class AuthService {
  async login(email: string, password: string): Promise<User>
  async logout(): Promise<void>
  async getCurrentUser(): Promise<User | null>
  async refreshToken(): Promise<string>
}
```

**現在の実装:**
- AWS Cognito OAuth2.0のモック実装
- JWTトークン管理のシミュレーション

#### s3.ts - S3アップロードサービス

```typescript
class S3Service {
  static async uploadImage(file: File, userId: string): Promise<string>
  static async deleteImage(imageUrl: string): Promise<void>
}
```

**現在の実装:**
- AWS S3アップロードのモック実装
- ユーザーID別ディレクトリ構造のシミュレーション

### `/types` - TypeScript型定義

アプリケーション全体で使用される型定義を格納。

```
types/
└── index.ts                   # 共通型定義
```

**主要な型定義:**
```typescript
interface Plant {
  id: string;
  user_id: string;
  name: string;
  description?: string;
  image_url?: string;
  created_at: string;
  updated_at: string;
}

interface WateringRecord {
  id: string;
  plant_id: string;
  watered_at: string;
  notes?: string;
  created_at: string;
}

interface User {
  id: string;
  email: string;
  name: string;
}
```

## 設定ファイル詳細

### package.json - プロジェクト設定

```json
{
  "name": "plant-watering-app",
  "scripts": {
    "dev": "vite",
    "build": "tsc -b && vite build",
    "preview": "vite preview",
    "lint": "eslint ."
  },
  "dependencies": {
    "react": "^18.3.1",
    "react-dom": "^18.3.1",
    "@radix-ui/react-*": "^1.1.2",
    "lucide-react": "^0.468.0",
    "tailwindcss": "^3.4.17"
  }
}
```

### vite.config.ts - ビルド設定

```typescript
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
    },
  },
})
```

**特徴:**
- パスエイリアス設定（`@/` = `src/`）
- React プラグイン統合
- 開発サーバー設定

### tailwind.config.js - CSS設定

```javascript
module.exports = {
  content: ["./index.html", "./src/**/*.{ts,tsx}"],
  theme: {
    extend: {
      colors: {
        // カスタムカラーパレット
      }
    }
  },
  plugins: [require("tailwindcss-animate")]
}
```

## ファイル命名規則

### コンポーネント
- **PascalCase**: `PlantCard.tsx`, `WateringButton.tsx`
- **拡張子**: `.tsx` (JSXを含むTypeScript)

### サービス・ユーティリティ
- **camelCase**: `api.ts`, `auth.ts`, `utils.ts`
- **拡張子**: `.ts` (純粋なTypeScript)

### 設定ファイル
- **kebab-case**: `vite.config.ts`, `tailwind.config.js`
- **ドット記法**: `.env`, `.gitignore`

## 開発ワークフロー

### 新機能追加時の手順

1. **型定義**: `src/types/index.ts` に必要な型を追加
2. **サービス**: `src/services/` に外部連携ロジックを実装
3. **コンポーネント**: `src/components/` にUIコンポーネントを作成
4. **統合**: `src/App.tsx` でルーティング・状態管理を統合

### コード品質管理

- **TypeScript**: 型安全性の確保
- **ESLint**: コード品質チェック
- **Prettier**: コードフォーマット統一
- **Git hooks**: コミット前の自動チェック

## パフォーマンス考慮事項

### コード分割
- **動的インポート**: 大きなコンポーネントの遅延読み込み
- **Tree Shaking**: 未使用コードの自動削除

### バンドル最適化
- **Vite**: 高速ビルドとHMR
- **依存関係最適化**: 必要最小限のライブラリ使用
- **アセット最適化**: 画像・CSS・JSの圧縮
