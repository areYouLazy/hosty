# Hosty

## Description

Hosty is a command-line utility that allows for fast inspection and editing of `/etc/hosts`-like files. It is written in [golang](https://golang.org) and uses [libhosty](https://github.com/areYouLazy/libhosty) as the underlying library to operate on files

## Table of Content

* [Main Features](#Main-Features)
* [Installation](#Installation)
  * [Precompiled packages](#Precompiled-packages)
  * [Build](#Build)
* [Usage](#Usage)
* [Examples](#Examples)
* [Commands and Aliases](#Commands-and-Aliases)
  * [show](#show)
  * [add](#add)
  * [delete](#delete)
  * [comment](#comment)
  * [uncomment](#uncomment)
  * [export](#export)
  * [restore](#restore)
* [License](#License)

## Main Features

* Fast Add/Delete/Comment/Uncomment entries
* JSON output for easy parsing
* Backup files before editing

## Installation

### Precompiled packages

You can install `hosty` with one of the precompiled packages in the `release` section

### Build

Ensure you have go on your system

```console
> go version
go version go1.16.3 linux/amd64
```

pull the app

```console
> go get github.com/areYouLazy/hosty
```

build and install

```console
> cd github.com/areYouLazy/hosty
> go build
> go install
```

You should now be able to use `hosty`. Remember that `Add/Delete/Comment/Uncomment` commands needs `root` privileges.

```console
> hosty show localhost
127.0.0.1       localhost
::1             localhost

> hosty comment 127.0.0.1
done
```

## Usage

```console
root@localhost$ hosty -h
Hosty is a command-line tool to interact with the /etc/hosts file.
It allows for fast inspect and edit of the file. Main goals of this tool are to be fast, reliable and scriptable.
Hosty uses libhosty to manipulate the file. You can find more about libhosty at https://github.com/areYouLazy/libhosty

Usage:
  hosty [command]

Available Commands:
  add         Add file data
  comment     Comment file data
  delete      Delete file data
  help        Help about any command
  show        Show file data
  uncomment   Uncomment file data

Flags:
  -f, --file string   use a custom /etc/hosts-like file
  -h, --help          help for hosty
  -j, --json          output in json format for easy parsing
  -q, --quiet         suppress every output except for errors
  -v, --version       version for hosty

Use "hosty [command] --help" for more information about a command.
```

## Examples

```console
root@localhost$ # show entries with hostname equale to localhost
root@localhost$ hosty show localhost
127.0.0.1         localhost
::1               localhost

root@localhost$ # show entries with ip equal to 127.0.0.1
root@localhost$ hosty show 127.0.0.1
127.0.0.1         localhost

root@localhost$ # add an entrie with ip 1.2.3.4 and hostname my.custom.dns
root@localhost$ hosty add 1.2.3.4 my.custom.dns
done

root@localhost$ # show entries with ip equal to 1.2.3.4
root@localhost$ hosty show 1.2.3.4
1.2.3.4         my.custom.dns

root@localhost$ # comment entries with ip equal to 1.2.3.4
root@localhost$ hosty comment 1.2.3.4
done

root@localhost$ # show entries with ip equal to 1.2.3.4
root@localhost$ hostyc-li show 1.2.3.4
# 1.2.3.4         my.custom.dns

root@localhost$ # uncomment entries with ip equal to 1.2.3.4 suppress output
root@localhost$ hosty uncomment 1.2.3.4 --quiet

root@localhost$ # show entries with ip equal to 1.2.3.4 output in json
root@localhost$ hosty show 1.2.3.4 --json | jq .
{
  "raw": "1.2.3.4         my.custom.dns"
}

root@localhost$ # add another entry with same ip and different hostname output in json
root@localhost$ hosty add 1.2.3.4 my.custom2.dns --json
{"done":true}

root@localhost$ # show entries with ip equal to 1.2.3.4 output in json with details
root@localhost$ hosty show 1.2.3.4 --json --details | jq .
{
  "action": "show",
  "number": 12,
  "type": "address",
  "address": "1.2.3.4",
  "hostnames": [
    "my.custom.dns",
    "my.custom2.dns"
  ],
  "comment": "",
  "is_commented": false,
  "raw": "127.0.0.1       my.custom.dns my.custom2.dns"
}

root@localhost$ # delete entries with ip equal to 1.2.3.4
root@localhost$ hosty delete 1.2.3.4
done

root@localhost$ # show entries with ip equal to 1.2.3.4
root@localhost$ hosty show 1.2.3.4
nothing found for ip 1.2.3.4
```

## Commands and Aliases

Here's a little explanation of every command.

Every Hosty command supports several aliases, most of which are just truncated version of the command.

### show

```console
root@localhost$ hosty show [PARAMETER]
```

Show entries based on the given parameter.
Parameter can be both IP or FQDN.

```console
root@localhost$ # Aliases
root@localhost$ hosty sho localhost
root@localhost$ hosty sh localhost
root@localhost$ hosty s localhost
```

### add

```console
root@localhost$ hosty add [IP] [FQDN] [COMMENT]
```

Add an entry to hosts file.
Required parameters are IP and FQDN, optionally you can pass a comment for the line.

```console
root@localhost$ # Aliases
root@localhost$ hosty ad 1.2.3.4 my.custom.dns "DNS#1"
root@localhost$ hosty a 1.2.3.4 my.custom.dns
```

### delete

```console
root@localhost$ hosty delete [PARAMETER]
```

Delete entries from hosts file based on the given parameter.
Parameter can be both IP or FQDN.

```console
root@localhost$ # Aliases
root@localhost$ hosty delet 1.2.3.4
root@localhost$ hosty dele 1.2.3.4
root@localhost$ hosty del 1.2.3.4
root@localhost$ hosty de 1.2.3.4
root@localhost$ hosty d 1.2.3.4
root@localhost$ hosty rem 1.2.3.4
root@localhost$ hosty rm 1.2.3.4
```

### comment

```console
root@localhost$ hosty comment [PARAMETER]
```

Comment entries on hosts file based on given parameter.
Parameter can be both IP or FQDN.

```console
root@localhost$ # Aliases
root@localhost$ hosty commen 1.2.3.4
root@localhost$ hosty comme 1.2.3.4
root@localhost$ hosty comm 1.2.3.4
root@localhost$ hosty com 1.2.3.4
root@localhost$ hosty co 1.2.3.4
root@localhost$ hosty c 1.2.3.4
```

### uncomment

```console
root@localhost$ hosty uncomment [PARAMETER]
```

Uncomment entries on hosts file based on given parameter.
Parameter can be both IP or FQDN.

```console
root@localhost$ # Aliases
root@localhost$ hosty uncommen 1.2.3.4
root@localhost$ hosty uncomme 1.2.3.4
root@localhost$ hosty uncomm 1.2.3.4
root@localhost$ hosty uncom 1.2.3.4
root@localhost$ hosty unco 1.2.3.4
root@localhost$ hosty unc 1.2.3.4
root@localhost$ hosty un 1.2.3.4
root@localhost$ hosty u 1.2.3.4
```

### export

```console
root@localhost$ hosty export [PARAMETER]
```

Export file content to a given location
Parameter must be a writable file

```console
root@localhost$ # Aliases
root@localhost$ hosty expor /home/sonica/hosts-export.txt
root@localhost$ hosty expo /home/sonica/hosts-export.txt
root@localhost$ hosty exp /home/sonica/hosts-export.txt
root@localhost$ hosty ex /home/sonica/hosts-export.txt
root@localhost$ hosty e /home/sonica/hosts-export.txt
```

### restore

```console
root@localhost$ hosty restore [PARAMETER]
```

Restore default hosts file for given OS
If Paramter is omitted, hosty will try to guess the OS and restore the appropriate file

```console
root@localhost$ # Aliases
root@localhost$ hosty restor darwin
root@localhost$ hosty resto darwin
root@localhost$ hosty rest darwin
root@localhost$ hosty res darwin
root@localhost$ hosty re darwin
root@localhost$ hosty r darwin
```

## License

Released under Apache 2.0 license
