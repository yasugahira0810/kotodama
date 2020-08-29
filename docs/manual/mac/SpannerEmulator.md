# Spanner Emulator インストール手順

## gcloudコマンドインストール

```bash
# 実行する
curl https://sdk.cloud.google.com | bash
exec -l $SHELL

# 実行しない（MacでSpanner Emulatorいじる分には不要）
gcloud init
```

[参考](https://cloud.google.com/sdk/docs/downloads-interactive?hl=ja)

## Spanner Emulator起動

### 起動

```bash
gcloud beta emulators spanner start
~メッセージ省略。以下のメッセージが出力されたらSpannerは起動している~
[cloud-spanner-emulator] 2020/08/29 06:46:33 gateway.go:136: REST server listening at 0.0.0.0:9020
[cloud-spanner-emulator] 2020/08/29 06:46:33 gateway.go:137: gRPC server listening at 0.0.0.0:9010
```

### 確認

```bash
docker ps
CONTAINER ID        IMAGE                                          COMMAND                  CREATED             STATUS              PORTS                                                NAMES
8ce06cc86d0d        gcr.io/cloud-spanner-emulator/emulator:1.0.0   "./gateway_main --ho…"   10 minutes ago      Up 10 minutes       127.0.0.1:9010->9010/tcp, 127.0.0.1:9020->9020/tcp   competent_yonath

```

[参考](https://cloud.google.com/spanner/docs/emulator?hl=ja)

