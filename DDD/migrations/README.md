# go-migration
- インストール
  - go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
- マイグレーションファイルの生成
  - migrate create -ext sql -dir migrations create_users_table   
- マイグレーションファイルの編集
  - up,down両方
- 