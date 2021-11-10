# Go Discord Bot


 [![Discord Gophers](https://img.shields.io/badge/Discord%20Gophers-%23info-blue.svg)](https://discord.gg/0f1SbxBZjYq9jLBk)
<img align="right" src="https://media.giphy.com/media/1AN73nSrDYhiM/giphy.gif">

Dis**go**rd is an example of or starting point for creating an easy to use and 
extensible Discord bot using the [DiscordGo](https://github.com/bwmarrin/discordgo) 
library.

This is just a simple Go Discord Bot project in order to both learn go and discord api.

## Getting Started

The below assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

### Installing

Use go get to download the repository into your Go work environment.

```sh
go get github.com/bwmarrin/disgord
```

### Usage
```sh
cd $GOPATH/src/github.com/thisguymartin/go-discord-bot-mp
```

## ENV  
You will need 2 api keys one for [Giphy](https://giphy.com) key and [Discord](https://discord.com/developers/docs/intro) secret Key

Create .env file
```
DISCORD_SECRET=""
GIPHY_API_KEY=""
```


## Simply Run 
```
go run .
```

## Simply Build 
```
go build .
```

## Heroku Deploy Script
```
./deploy.sh
```

