
# Simple Web Server

このプロジェクトは、Goで実装されたシンプルなWebサーバーです。

## プロジェクト構造

- `entities`: コアエンティティを定義
- `frontend`: フロントエンドアプリケーション

## 依存関係

- `github.com/gin-gonic/gin`: Webフレームワーク

## 実行方法

1.  依存関係をインストール: `go mod tidy`
2.  アプリケーションを実行: `go run main.go`
3.  ブラウザで `http://localhost:8080/ping` にアクセス

## API エンドポイント

現在利用可能なエンドポイント:

- `GET /ping` - サーバーの動作確認

## フロントエンド

フロントエンドアプリケーションは `frontend/` ディレクトリに配置されています。
