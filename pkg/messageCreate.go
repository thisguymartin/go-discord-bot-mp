package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!wick" {
		var giphyResponseJSON = new(GiphyResponse)

		var giphyToken = os.Getenv("GIPHY_API_KEY")
		response, err := http.Get("https://api.giphy.com/v1/gifs/random?api_key=" + giphyToken + "&tag=puppies")
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		body, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			panic(err.Error())
		}

		err3 := json.Unmarshal(body, &giphyResponseJSON)
		if err3 != nil {
			fmt.Println("whoops:", err3)
			//outputs: whoops: <nil>
		}

		s.ChannelMessageSend(m.ChannelID, giphyResponseJSON.Data.URL)

	} else {
		fmt.Print("not wick")

	}

}
