package utils

import (
	"encoding/json"
	"net/http"
)

type JokeResponse struct {
	Joke string `json:"joke"`
}

func GetRandomJoke() (string, error) {
	resp, err := http.Get("https://api-ninjas.com/api/jokes")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var jokeResp JokeResponse
	if err := json.NewDecoder(resp.Body).Decode(&jokeResp); err != nil {
		return "", err
	}

	return jokeResp.Joke, nil
}
