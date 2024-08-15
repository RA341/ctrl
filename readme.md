# CTRL srv

A helpful utility server, to control and monitor my home server. 

## Install and setup

TODO

## Features

System
* Power on
* Restart
* Shutdown

Qbit
* Check client status ( i.e. connected/disconnected/firewalled)
* Check for stalled torrents

Notifications via
* Discord

## Planned

* add docker control
* implement auto updater
* implement install script

## Preferences

Controlled by an ini file with the following usage

**Note: on first run the program creates an empty ini file and then exits, fill this out before starting**

```ini
[General]
; DO NOT TOUCH THIS AS THIS WILL MESS WITH UPDATES
Version = 1.0 

auto_update = true (default) (required)

; how often to check for updates in hours (default is weekly)
update_interval = 168 (default) (required)

[Network]
; BE CAREFUL WHEN CHANGING THIS, IT MAY CAUSE THE SERVER TO BECOME INACCESSIBLE
Host = http://0.0.0.0 (default) (required)  
Port = 9220 (default) (required)  

[Qbit]
enable = true (default)
; ip or hostname and port of your qbittorrent instance ( remember to add https or http accordingly)
host = http://127.0.0.1 (required)  
port = 8085 (required)

; username and password of your qbittorrent instance
user = ....  (required)
pass = ....  (required)

[notifications.Discord]
enable = true (default)

; more info https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks
discord_webhook_url = 'https://discord.com/api/webhooks/.....' (required)

username = CTRL Bot (default) (required)
avatar_url = https://i.imgur.com/KEungv8.png (default) (optional) 
```

# Development

To go grpc files

install pre-requisites

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

generate

```bash
protoc --go_grpc_out=. --go_out=. .\protos\filesystem.proto
```

To generate Dart grpc client files

```bash
protoc --dart_out=grpc:client/lib/ -I. .\protos\filesystem.proto
```