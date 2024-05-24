## 概要

UR賃貸住宅から部屋のデータを収集して、DBに保存するアプリ (rewrote in golang)

https://www.ur-net.go.jp/chintai/


## 動かし方

### 動作環境

- Docker が動く環境

### gRPCサーバを起動

```
git clone https://github.com/shibu1x/ur_v3.git
cd ur_v3
docker compose up -d app_server
```

### データ収取を実行する

https://www.ur-net.go.jp/chintai/ から公開されているデータを取得し、 postgres に登録する

```
docker compose run --rm app crawl
```

### gRPCクライアントを実行

gRPCサーバからデータを取得して、標準出力する

```
docker compose run --rm app client
```

### 終了 & データを消す

```
docker compose down
```

## その他

### なんでこんなことしてるの？

- UR賃貸は人気高いので、空きが出てもすぐに埋まってしまう
- 埋まってしまうと部屋の情報が見れなくなってしまい、どのような条件だったのかも分からない
- どのような条件で募集が出ているか知りたいので、埋まる前の情報をbotで集めたい


## 備忘録

SQLBoiler モデルを生成
```
sqlboiler psql --add-global-variants
```

gRPC Protocol Buffer Compile
```
protoc --go_out=. --go-grpc_out=. pkg/proto/rooms.proto
```