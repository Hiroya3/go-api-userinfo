# go-api-userinfo
ユーザー情報（name,age）をGET/PUT/POST/DELETEできるAPI

## メソッド
GET   : 特定のIDのユーザー情報を取得する<br>
POST  : 新規ユーザー情報を作成する<br>
PUT   : 特定のユーザーの情報を更新する<br>
DELETE: 特定のユーザー情報を削除する<br>

## DBの準備
### 1.PostgreSQLのインストール
Postgres.app( http://postgresapp.com )にアクセスし,PostgreSQLをインストール

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

### 4.権限付与
以下のコマンドを実行し `dbuser` に権限を付与する

```
psql -c "grant all privileges on database user_info to dbuser"
```

## webサーバーの立ち上げ
ソースをダウンロード or cloneして`go_api_userinfo` を実行

## 各メソッドを実行
### PUT
以下のコマンドでデータ作成({name},{age}は適宜入れてください)

```
curl -i -X POST -H "Content-Type : application/json" -d '{"name":"{name}","age":{age}}' http://localhost:8080/user_info/
```

### GET
以下のコマンドでデータを取得({id}は適宜入れてください)

```
curl -i -X GET http://localhost:8080/user_info/{id}
```

### PUT
以下のコマンドでデータ更新({name},{age},{id}は適宜入れてください)

```
curl -i -X PUT -H "Content-Type : application/json" -d '{"name":"{name}","age":{age}}' http://localhost:8080/user_info/{id}
```
### DELETE
以下のコマンドでデータを削除({id}は適宜入れてください)

```
curl -i -X DELETE http://localhost:8080/user_info/{id}
```
