package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JokeResponse struct {
	Error    bool   `json:"error"`
	Category string `json:"category"`
	Type     string `json:"type"`
	Joke     string `json:"joke"`
	Delivery string `json:"delivery"`
	Setup    string `json:"setup"`
	Flags    struct {
		Nsfw      bool `json:"nsfw"`
		Religious bool `json:"religious"`
		Racist    bool `json:"racist"`
		Sexist    bool `json:"sexist"`
		Political bool `json:"political"`
		Explicit  bool `json:"explicit"`
	} `json:"flags"`
	ID   int    `json:"id"`
	Safe bool   `json:"safe"`
	Lang string `json:"lang"`
}

func getJoke() (string, error) {
	response := &JokeResponse{}
	res, err := http.Get("https://v2.jokeapi.dev/joke/Any")
	if err != nil {
		print(err)
	}

	decod := json.NewDecoder(res.Body)
	if err := decod.Decode(response); err != nil {
		fmt.Printf("error decoding JSON: %v\n", err)
		return "", err
	}

	if len(response.Joke) > 0 {
		return response.Joke, nil
	}

	return response.Setup + " " + response.Delivery, nil

}
