package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/bwmarrin/discordgo"
)

type GiphyResponse struct {
	Data Data `json:"data"`
	Meta Meta `json:"meta"`
}

type Data struct {
	Type                         string `json:"type"`
	ID                           string `json:"id"`
	URL                          string `json:"url"`
	Slug                         string `json:"slug"`
	BitlyGIFURL                  string `json:"bitly_gif_url"`
	BitlyURL                     string `json:"bitly_url"`
	EmbedURL                     string `json:"embed_url"`
	Username                     string `json:"username"`
	Source                       string `json:"source"`
	Title                        string `json:"title"`
	Rating                       string `json:"rating"`
	ContentURL                   string `json:"content_url"`
	SourceTLD                    string `json:"source_tld"`
	SourcePostURL                string `json:"source_post_url"`
	IsSticker                    int64  `json:"is_sticker"`
	ImportDatetime               string `json:"import_datetime"`
	TrendingDatetime             string `json:"trending_datetime"`
	Images                       Images `json:"images"`
	User                         User   `json:"user"`
	ImageOriginalURL             string `json:"image_original_url"`
	ImageURL                     string `json:"image_url"`
	ImageMp4URL                  string `json:"image_mp4_url"`
	ImageFrames                  string `json:"image_frames"`
	ImageWidth                   string `json:"image_width"`
	ImageHeight                  string `json:"image_height"`
	FixedHeightDownsampledURL    string `json:"fixed_height_downsampled_url"`
	FixedHeightDownsampledWidth  string `json:"fixed_height_downsampled_width"`
	FixedHeightDownsampledHeight string `json:"fixed_height_downsampled_height"`
	FixedWidthDownsampledURL     string `json:"fixed_width_downsampled_url"`
	FixedWidthDownsampledWidth   string `json:"fixed_width_downsampled_width"`
	FixedWidthDownsampledHeight  string `json:"fixed_width_downsampled_height"`
	FixedHeightSmallURL          string `json:"fixed_height_small_url"`
	FixedHeightSmallStillURL     string `json:"fixed_height_small_still_url"`
	FixedHeightSmallWidth        string `json:"fixed_height_small_width"`
	FixedHeightSmallHeight       string `json:"fixed_height_small_height"`
	FixedWidthSmallURL           string `json:"fixed_width_small_url"`
	FixedWidthSmallStillURL      string `json:"fixed_width_small_still_url"`
	FixedWidthSmallWidth         string `json:"fixed_width_small_width"`
	FixedWidthSmallHeight        string `json:"fixed_width_small_height"`
	Caption                      string `json:"caption"`
}

type Images struct {
	DownsizedLarge         The480_WStill  `json:"downsized_large"`
	FixedHeightSmallStill  The480_WStill  `json:"fixed_height_small_still"`
	Original               FixedHeight    `json:"original"`
	FixedHeightDownsampled FixedHeight    `json:"fixed_height_downsampled"`
	DownsizedStill         The480_WStill  `json:"downsized_still"`
	FixedHeightStill       The480_WStill  `json:"fixed_height_still"`
	DownsizedMedium        The480_WStill  `json:"downsized_medium"`
	Downsized              The480_WStill  `json:"downsized"`
	PreviewWebp            The480_WStill  `json:"preview_webp"`
	OriginalMp4            DownsizedSmall `json:"original_mp4"`
	FixedHeightSmall       FixedHeight    `json:"fixed_height_small"`
	FixedHeight            FixedHeight    `json:"fixed_height"`
	DownsizedSmall         DownsizedSmall `json:"downsized_small"`
	Preview                DownsizedSmall `json:"preview"`
	FixedWidthDownsampled  FixedHeight    `json:"fixed_width_downsampled"`
	FixedWidthSmallStill   The480_WStill  `json:"fixed_width_small_still"`
	FixedWidthSmall        FixedHeight    `json:"fixed_width_small"`
	OriginalStill          The480_WStill  `json:"original_still"`
	FixedWidthStill        The480_WStill  `json:"fixed_width_still"`
	Looping                Looping        `json:"looping"`
	FixedWidth             FixedHeight    `json:"fixed_width"`
	PreviewGIF             The480_WStill  `json:"preview_gif"`
	The480WStill           The480_WStill  `json:"480w_still"`
}

type The480_WStill struct {
	URL    string  `json:"url"`
	Width  string  `json:"width"`
	Height string  `json:"height"`
	Size   *string `json:"size,omitempty"`
}

type DownsizedSmall struct {
	Height  string `json:"height"`
	Mp4     string `json:"mp4"`
	Mp4Size string `json:"mp4_size"`
	Width   string `json:"width"`
}

type FixedHeight struct {
	Height   string  `json:"height"`
	Mp4      *string `json:"mp4,omitempty"`
	Mp4Size  *string `json:"mp4_size,omitempty"`
	Size     string  `json:"size"`
	URL      string  `json:"url"`
	Webp     string  `json:"webp"`
	WebpSize string  `json:"webp_size"`
	Width    string  `json:"width"`
	Frames   *string `json:"frames,omitempty"`
	Hash     *string `json:"hash,omitempty"`
}

type Looping struct {
	Mp4     string `json:"mp4"`
	Mp4Size string `json:"mp4_size"`
}

type User struct {
	AvatarURL    string `json:"avatar_url"`
	BannerURL    string `json:"banner_url"`
	Username     string `json:"username"`
	DisplayName  string `json:"display_name"`
	Description  string `json:"description"`
	IsVerified   bool   `json:"is_verified"`
	WebsiteURL   string `json:"website_url"`
	InstagramURL string `json:"instagram_url"`
	BannerImage  string `json:"banner_image"`
	ProfileURL   string `json:"profile_url"`
}

type Meta struct {
	Status     int64  `json:"status"`
	Msg        string `json:"msg"`
	ResponseID string `json:"response_id"`
}

func MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "kakashi" {
		giphyResponseJSON := new(GiphyResponse)
		response, err := http.Get("https://api.giphy.com/v1/gifs/random?api_key=" + os.Getenv("GIPHY_API_KEY") + "&tag=kakashi")

		if err != nil {
			fmt.Println(err)
		}

		defer response.Body.Close()

		body, err2 := ioutil.ReadAll(response.Body)

		if err2 != nil {
			fmt.Println(err2)
		}

		err3 := json.Unmarshal(body, &giphyResponseJSON)
		if err3 != nil {
			fmt.Println(err3)
		}

		_, err = s.ChannelMessageSend(m.ChannelID, giphyResponseJSON.Data.URL)
		if err != nil {
			fmt.Println(err)
		}

	} else if m.Content == "joke" {
		joke, jokeerror := getJoke()
		if jokeerror != nil {
			fmt.Println("Fail fetch joke error")
		}
		_, discord_error := s.ChannelMessageSend(m.ChannelID, joke)
		if discord_error != nil {
			fmt.Println("Fail to send joke error")
		}
	}

}
