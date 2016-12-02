# Sakura IoT Platform CLI Tool

コマンドライン上で各種操作を行うためのツールです。
開発中につき、機能が未完成です。

## Install

todo

### Linux

```bash
wget ~~~~~~ -O /usr/local/bin/sakuraio
chmod +x /usr/local/bin/sakuraio
```

### Windows

リリースページからダウンロード後、exeファイルをパスの通った位置に移動させてください


## Uninstall

```bash
rm /usr/local/bin/sakuraio
rm -r ~/.sakuraio
```

ホームディレクトリに `.sakuraio` ディレクトリを作成して設定を保存しています。

## コマンドリファレンス

### サンプルコマンド

```
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
    Get data

    -s, --size="100"       Fetch Size
        --unit=UNIT        Unit channel/message
        --order=ORDER      Order asc/desc
        --token=TOKEN      Service Token
        --cursor=CURSOR    Cursor
        --after=AFTER      Datetime range from
        --before=BEFORE    Datetime range to
        --channel=CHANNEL  Channel
        --project=PROJECT  Project ID
        --raw              Raw JSON output

```