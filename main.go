package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var DiscordSecret string

func main() {

	godotenv.Load()

	val, ok := os.LookupEnv("DISCORD_SECRET")
	if !ok {
		log.Fatal("Missing Discord Secret")
	} else {
		DiscordSecret = val
	}

	dg, err := discordgo.New("Bot " + DiscordSecret)
	if err != nil {
		log.Fatal("error creating Discord session,", err)
	}

	dg.AddHandler(MessageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		log.Fatal("error opening connection")
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

}
