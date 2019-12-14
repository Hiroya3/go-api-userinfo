# go-api-userinfo
ユーザー情報（name,age）をGET/PUT/POST/DELETEできるAPI

## メソッド
GET   : 特定のIDのユーザー情報を取得する<br>
POST  : 新規ユーザー情報を作成する<br>
PUT   : 特定のユーザーの情報を更新する<br>
DELETE: 特定のユーザー情報を削除する<br>

## DBの準備
### 1.PostgreSQLのインストール
Postgres.app(http://postgresapp.com)にアクセスし,PostgreSQLをインストール

### 2.データベースユーザー作成
以下のコマンドを実行し `dbuser` という名のユーザを作る

```
createuser -P -d dbuser
```
コマンド実行後passwordが求められるので `pass` と入力

### 3.データベース作成
以下のコマンドを実行し `user_info` というDBを作成する

```
createdb user_info
```

### webサーバーの立ち上げ
`go_api_userinfo` を実行

### 各メソッドを実行
