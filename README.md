# Sakura IoT Platform CLI Tool

コマンドライン上で各種操作を行うためのツールです。
開発中につき、機能が未完成です。

## Install

### Linux/BSD

~~~~~~ の部分は[リリースページ](https://github.com/sakura-internet/sakuraio-cli/releases/)
から適切なOS/アーキテクチャを選択し、読み替えてください。

``` bash
sudo wget ~~~~~~ -O /usr/local/bin/sakuraio
sudo chmod +x /usr/local/bin/sakuraio
```

### Windows

[リリースページ](https://github.com/sakura-internet/sakuraio-cli/releases/)
からダウンロード後、exeファイルをパスの通った位置に移動させてください

### Go

goの開発環境が有る方向けです。
アップデートをする際も、同じ手順で行えます。

```bash
go get -u github.com/sakura-internet/sakuraio-cli
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
    -s, --size="100"       Fetch Size
        --order=ORDER      Order asc/desc
        --token=TOKEN      Service Token
        --cursor=CURSOR    Cursor
        --after=AFTER      Datetime range from
        --before=BEFORE    Datetime range to
        --channel=CHANNEL  Channel
        --project=PROJECT  Project ID
        --raw              Raw JSON output

  service datastore messages [<flags>]
    Get message data

    -m, --module=""        Module ID
    -s, --size="100"       Fetch Size
        --order=ORDER      Order asc/desc
        --cursor=CURSOR    Cursor
        --after=AFTER      Datetime range from
        --before=BEFORE    Datetime range to
        --project=PROJECT  Project ID
        --raw              Raw JSON output
        --token=TOKEN      Service Token

  service websocket listen [<flags>]
    Listen to Websocket

    --project=PROJECT  Project ID
    --token=TOKEN      Service Token
```
