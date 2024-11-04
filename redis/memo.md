`

redisのコネクションに必要な情報は？
- DBname
- usernama
- pass


redisには何を保存している？
Keyがあるかどうかが大事
- redis 
  - session
    - Key
    - user_info 

セッションの有効期限って？何によって決まる？
- cookieにsessionIDを保存しているので
  - クライアント側｜cookieが削除
  - サーバ側｜セッションIDの保持期限