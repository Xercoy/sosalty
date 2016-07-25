package sosalty

import (
	"math/rand"
	"time"
)

type Request struct {
	Data GifResponse `json:"data"`
}

type GifResponse struct {
	ImageURL string `json:"image_url"`
}

func GetGifURL() (string, error) {
	var urls []string

	urls = append(urls, "https://media4.giphy.com/media/10sU6K4B9SIIBW/giphy.gif")
	urls = append(urls, "https://media4.giphy.com/media/7BY2M4A5RxyV2/giphy.gif")
	urls = append(urls, "https://media3.giphy.com/media/uyFbFcWuKcaeQ/giphy.gif")
	urls = append(urls, "https://media1.giphy.com/media/Mjn2djt0EFVle/giphy.gif")
	urls = append(urls, "https://media.giphy.com/media/Z3gSAkTpSnB0A/giphy.gif")
	urls = append(urls, "https://media.giphy.com/media/PjpEeI11T7qSI/giphy.gif")

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(urls))

	return urls[randomIndex], nil

	/* Tabled for now until we get a private API key
	url := `http://api.giphy.com/v1/gifs/random?api_key=dc6zaTOxFJmzC&tag="salty"`

	response, err := http.Get(url)
	if err != nil {
		return Request{}, err
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Request{}, err
	}

	var r Request

	err = json.Unmarshal(responseBytes, &r)

	return r.ImageURL, err
	*/
}
