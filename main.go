package main

import (
	"database/sql"

	_ "github.com/lib/pq"

	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	pkg "thisguymartin/go-discord-bot-mp/pkg"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var DiscordSecret string
var PG_URI string

func main() {

	godotenv.Load()

	val, ok := os.LookupEnv("DISCORD_SECRET")
	if !ok {
		log.Fatal("Missing Discord Secret")
	} else {
		DiscordSecret = val
	}

	db, err := sql.Open("postgres", os.Getenv("PG_URI"))
	if err != nil {
		log.Fatal("error connecting with postgres", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully created connection to database")
	defer db.Close()

	dg, err := discordgo.New("Bot " + DiscordSecret)
	if err != nil {
		log.Fatal("error connecting with postgres", err)
	}

	dg.AddHandler(pkg.MessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		pkg.CheckError(err)
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

}
