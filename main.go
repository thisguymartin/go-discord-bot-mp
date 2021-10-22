package main

import (
	"database/sql"
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

var (
	PG_HOST     = os.Getenv("PG_HOST")
	PG_DB       = os.Getenv("PG_PASSWORD")
	PG_USER     = os.Getenv("PG_USER")
	PG_PASSWORD = os.Getenv("PG_PASSWORD")
)

func main() {

	godotenv.Load()

	val, ok := os.LookupEnv("DISCORD_SECRET")
	if !ok {
		log.Fatal("Missing Discord Secret")
	} else {
		DiscordSecret = val
	}

	var psqlInfo = fmt.Sprintf("host=%s port=%s  user=%s "+
		"password=%s dbname=%s sslmode=disable",
		PG_HOST, 5432, PG_USER, PG_PASSWORD, PG_PASSWORD)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("error connecting with postgres", err)
	}

	err = db.Ping()
	pkg.CheckError(err)
	fmt.Println("Successfully created connection to database")

	defer db.Close()

	dg, err := discordgo.New("Bot " + DiscordSecret)
	if err != nil {
		pkg.CheckError(err)
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
