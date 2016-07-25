package sosalty

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Data GifResponse `json:"data"`
}

type GifResponse struct {
	ImageURL string `json:"image_url"`
}

func GetGifURL() (string, error) {
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
}
