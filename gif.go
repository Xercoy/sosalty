package yousosalty

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// APIKey is a string variable representing the Giphy API key used when sending a
// request.
var APIKey = os.Getenv("YOUSOSALTYKEY")

func init() {
	if APIKey == "" {
		log.Println("API Key not found in environment variable, switching to default.")

		// Default to public beta key
		APIKey = "dc6zaTOxFJmzC"
	} else {
		log.Println("API Key in environment variable found.")
	}
}

// GiphyResponse represents the information received by sending a request to GIPHY.
type GiphyResponse struct {
	Data Gif `json:"data"`
}

// Gif represents information about the URL of the particular GIF in a response.
type Gif struct {
	ImageURL string `json:"image_url"`
}

func GetGifURL() (string, error) {

	url := `http://api.giphy.com/v1/gifs/random?api_key=` + APIKey + `&tag="salty"`

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	var r GiphyResponse

	err = json.Unmarshal(responseBytes, &r)

	return r.Data.ImageURL, err

	/*	Old code, we got a private key!
		    var urls []string

			urls = append(urls, "https://media4.giphy.com/media/10sU6K4B9SIIBW/giphy.gif")
			urls = append(urls, "https://media4.giphy.com/media/7BY2M4A5RxyV2/giphy.gif")
			urls = append(urls, "https://media3.giphy.com/media/uyFbFcWuKcaeQ/giphy.gif")
			urls = append(urls, "https://media1.giphy.com/media/Mjn2djt0EFVle/giphy.gif")
			urls = append(urls, "https://media.giphy.com/media/Z3gSAkTpSnB0A/giphy.gif")
			urls = append(urls, "https://media.giphy.com/media/PjpEeI11T7qSI/giphy.gif")
			urls = append(urls, "http://media2.giphy.com/media/JWi3xPG4hfsac/giphy.gif")
			rand.Seed(time.Now().UnixNano())
			randomIndex := rand.Intn(len(urls))

			return urls[randomIndex], nil*/

	//Melee
	// http://media4.giphy.com/media/3WEhRB0yfxlmM/giphy.gif

	// General
	// http://media0.giphy.com/media/e2ZPwmGerfGdW/giphy.gif
	// http://media2.giphy.com/media/3oEdv9HsJoI7rwoDIs/giphy.gif
	// http://media4.giphy.com/media/3dx4cmZftWOk/giphy.gif
	// https://media3.giphy.com/media/qivbcSpS2mdAA/giphy.gif
}
