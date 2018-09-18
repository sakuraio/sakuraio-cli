# sakura.io CLI Tool

コマンドライン上でsakura.ioの各種操作を行うためのツールです。
モジュール登録の自動化・蓄積されたメッセージの一括取得など、運用において便利な機能が実装されています。

内部でのAPI呼び出しの際に、各種サービスごとの使用によって課金が発生する可能性が有ります。詳しくは、ドキュメントをご確認ください。 https://sakura.io/docs/pages/spec/platform/index.html

## Install

### Linux/BSD

(version)と(arch) の部分は[リリースページ](https://github.com/sakuraio/sakuraio-cli/releases/)
から適切なOS/アーキテクチャを選択し、読み替えてください。

``` bash
sudo wget https://github.com/sakuraio/sakuraio-cli/releases/download/(version)/sakuraio-cli_(arch) -O /usr/local/bin/sakuraio
sudo chmod +x /usr/local/bin/sakuraio
```

### Windows

[リリースページ](https://github.com/sakuraio/sakuraio-cli/releases/)
からダウンロード後、exeファイルをパスの通った位置に移動させてください

### Go

goの開発環境が有る方向けです。
アップデートをする際も、同じ手順で行えます。

```bash
go get -u github.com/sakuraio/sakuraio-cli
sakuraio-cli --help # `export PATH=$PATH:$GOPATH/bin` is required.
$GOPATH/bin/sakuraio-cli --help
```

### Build

開発に参加したい方は、リポジトリをクローン後ビルドを行ってください。
下記のコマンドでは、他のプラットフォーム向けへのバイナリも生成されます。

```bash
./build.sh
```

## Uninstall

```bash
rm /usr/local/bin/sakuraio
rm -r ~/.sakuraio
```

ホームディレクトリに `.sakuraio` ディレクトリを作成して設定を保存しています。

## 使い方

簡単にコマンドの使用方法を紹介します。
すべての機能は、

### APIキーを設定する

管理APIのAPIキーを取得する必要があります。
https://sakura.io/docs/pages/guide/api-guide/test-management-api.html

APIキーを取得した後、`auth`コマンドで設定を行います。

```
# sakuraio auth [key] [secret]
```

コマンドを実行することで`~/.sakuraio/config.yml`が作成され、認証情報とAPIのBaseURLがセットされます。

```
apitoken: ********-****-****-****-************
apisecret: apisecret: ****************************************************************
baseurl: "https://api.sakura.io/"
```

### プロジェクトを表示する

最初に`project list`プロジェクトのリストを取得します。

```
$ sakuraio project list
ID     ProjectName
664    sample project
663    テストプロジェクト
```

では、ID`664`のプロジェクトの詳細を`project show`コマンドを使用して表示してみます。
実際のプロジェクトIDは、利用者によって異なるので、読み替えて実行してください。

```
$ sakuraio project show 664
ID    ProjectName
664   sample project

ID       Project    Type                 Token                                   ServiceName
14014    664        websocket            7d3de3a9-b0e9-48e7-b799-************
14054    664        datastore            4cbb48ac-d122-4fc9-8aef-************    検証用データストア

ID               Project    Online    ModuleName
emulator-001     664        false
```

プロジェクトに設定しているサービス・モジュールが表示されます。


### データストアの表示

先程の例で表示したプロジェクトに設定されているDataStoreのサービスを使用し、メッセージの取得を行ってみます。
トークンは、プロジェクトに設定したdatastoreのトークンとして読み替えてください。

```
$ service datastore messages --token 4cbb48ac-d122-4fc9-8aef-************
Module          Datetime                          Type          Payload
emulator-001    2018-08-22T01:02:54.867058121Z    connection    {"is_online":false}
emulator-001    2018-08-22T01:02:53.658258856Z    channels      {"channels":[{"channel":0,"datetime":"2018-08-22T01:02:53.658260725Z","type":"l","value":10}]}
emulator-001    2018-08-22T01:02:53.658874553Z    channels      {"channels":[{"channel":1,"datetime":"2018-08-22T01:02:48.658876312Z","type":"l","value":10}]}
emulator-001    2018-08-22T01:02:52.657918955Z    channels      {"channels":[{"channel":1,"datetime":"2018-08-22T01:02:47.657920377Z","type":"l","value":9}]}
emulator-001    2018-08-22T01:02:52.657601773Z    channels      {"channels":[{"channel":0,"datetime":"2018-08-22T01:02:52.657603379Z","type":"l","value":9}]}
...略
```

正常に表示されました。次は、表示形式をJSONにし、最近の3件のみ取得してみましょう。

```
$ sakuraio service datastore messages --token 4cbb48ac-d122-4fc9-8aef-************ --size 3 --raw
{"id":"349493631385635840","module":"emulator-001","datetime":"2018-08-22T01:02:54.867058121Z","type":"connection","payload":{"is_online":false}}
{"id":"349493627195891712","module":"emulator-001","datetime":"2018-08-22T01:02:53.658258856Z","type":"channels","payload":{"channels":[{"channel":0,"datetime":"2018-08-22T01:02:53.658260725Z","type":"l","value":10}]}}
{"id":"349493627195960320","module":"emulator-001","datetime":"2018-08-22T01:02:53.658874553Z","type":"channels","payload":{"channels":[{"channel":1,"datetime":"2018-08-22T01:02:48.658876312Z","type":"l","value":10}]}}
```

他にも柔軟なオプション指定が行えるので、`sakuraio service datastore messages --help`を実行して確認してください。


## コマンドリファレンス

### サンプルコマンド

```
$ sakuraio auth [key] [secret]  # APIキーで認証
$ sakuraio project ls  # プロジェクトの一覧を取得
$ sakuraio service add datastore 7 name=test-service  # プロジェクトにDataStoreサービスを追加
$ sakuraio service datastore channels --project=7  # プロジェクト7のチャンネルデータを閲覧
$ sakuraio service add websocket 7 name=test-service  # Websocketサービスを追加
$ sakuraio service websocket listen --project=7  # リアルタイムにデータを受信
```

### --help

```
Commands:
  help [<command>...]
  auth [<token>] [<secret>]
  project
    list
    show <id>...
    add <name>
    remove [<flags>] <ID>
  module
    list
    show <ID>...
    add <id> <password> <project-id> [<name>]
  service
    list
    show <id>...
    remove [<flags>] <ID>
    add <type> <project id> [<option>...]
    datastore
      channels [<flags>]
      messages [<flags>]
    websocket
      listen [<flags>]
```

### --help-long

```
>go run main.go --help-long
usage: sakuraio [<flags>] <command> [<args> ...]

sakuraio client command

Flags:
  --help                   Show context-sensitive help (also try --help-long and
                           --help-man).
  --api-token=API-TOKEN    API Token
  --api-secret=API-SECRET  API Secret
  --version                Show application version.

Commands:
  help [<command>...]
    Show help.


  auth [<token>] [<secret>]
    Authentication


  project list
    List of projects


  project show <id>...
    Lookup projects


  project add <name>
    Add Project


  project remove [<flags>] <ID>
    Remove Project

    -f, --force  Project force remove

  module list
    List of modules


  module show <ID>...
    Lookup modules


  module add <id> <password> <project-id> [<name>]
    Add module


  service list
    List of services


  service show <id>...
    show service


  service remove [<flags>] <ID>
    Remove Project

    -f, --force  Force remove

  service add <type> <project id> [<option>...]
    Add Service


  service datastore channels [<flags>]
    Get channel data

    -m, --module=""        Module ID
    -s, --size=100         Total fetch count. When 0 is specified fetch
                           unlimited
        --order=ORDER      Order asc/desc
        --token=TOKEN      Service Token
        --cursor=CURSOR    Cursor
        --after=AFTER      Datetime range from
        --before=BEFORE    Datetime range to
        --channel=CHANNEL  Channel
        --project=PROJECT  Project ID
        --raw              Raw JSON output
        --no-rec           Un use recursive fetch
        --max-req=100      Max recursive request count
        --batch-size=100   Fetch size per request

  service datastore messages [<flags>]
    Get message data

    -m, --module=""        Module ID
    -s, --size=100         Total fetch count. When 0 is specified fetch
                           unlimited
        --order=ORDER      Order asc/desc
        --cursor=CURSOR    Cursor
        --after=AFTER      Datetime range from
        --before=BEFORE    Datetime range to
        --project=PROJECT  Project ID
        --raw              Raw JSON output
        --token=TOKEN      Service Token
        --no-rec           Un use recursive fetch
        --max-req=100      Max recursive request count
        --batch-size=100   Fetch size per request

  service websocket listen [<flags>]
    Listen to Websocket

    --project=PROJECT  Project ID
    --token=TOKEN      Service Token
```
