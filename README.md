# about

mfaの素振り

# 方針
いくつかの方式で2faを行えるということにするが、今回はnoneとemailのみ実装  
今回は.env等のsecretの漏洩に対しては考えない
# ユーザー作成フロー
- email addressだけを入力させる
- email addressとexpire日時を付与したtokenを発行し、emailでユーザーに送信する。もしemailがすでにDBにあれば409
- tokenを使った認証で、メールアドレス以外の詳細情報(password,2FAType)を入力、保存する(ここで受け取るtokenはemailの実在性確認のため) 。もしemailがすでにDBにあれば409
# ユーザー認証フロー
- email,passwordで認証
- 認証が成功したら、6桁のcodeを作成、user idと紐づけ保存する、one time tokenがついたURLをユーザーに302で返却
- 指定された方式でcodeをユーザーに送信する
- ユーザーはone time tokenを使ってアクセスし、codeを入力する
- 正しい入力ならばauthorization tokenを生成しCookieを使って302で返す。codeとone time tokenはどちらにせよ破棄

# test
## send activate email
```shell script
curl -XPOST localhost:8888/activate/email -H "content-type: application/json" --data '{"email":"test@example.com"}'
```
## create user with jwt
```shell script
curl -XPOST localhost:8888/activate/user?t=<jwt> -H "content-type: application/json" --data  '{"password":"testpass","MFAType":"email"}'
```
## login
```shell script
curl -XPOST localhost:8888/user/login -H "content-type: application/json" --data '{"email":"test@example.com","password":"testpass"}'
```
## mfa
```shell script
curl -XPOST localhost:8888/user/mfa/<one_time_token> -H "content-type: application/json" --data '{"mfacode":"123456"}'
```
## with cred
```shell script
curl localhost:8888/test -H "X-Authorization-Token: <token>"
```