# ログインする

## 基本コース

- ユーザはログイン画面にアクセスする。ユーザはアカウント名とパスワードを入力して、「ログイン」ボタンをクリックする。
- システムはユーザ認証する。アカウント名とパスワードが正しければ、名言登録画面に遷移する。

## 代替コース

### ユーザが存在しない

- システムはユーザが入力したアカウント名をDBで検索し、一致するアカウント名が存在しなかった場合は、「ユーザが存在しません」とエラーメッセージを表示する。

### パスワードが誤っている

- システムは、ユーザが入力したアカウントが存在した場合はパスワードを照合する。パスワードが誤っていた場合は、「パスワードが違います」とエラーメッセージを表示する。