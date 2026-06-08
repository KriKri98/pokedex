package pokeapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/KriKri98/pokedex/internal/pokecache"
)

func Get(url string) (map[string]any, error) {
	var data map[string]any

	res, err := http.Get(url)
	if err != nil {
		return data, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&data); err != nil {
		return data, err
	}
	return data, nil

}

type Client struct {
	cache  pokecache.Cache
	config Config
}

type Config struct {
	Next     string
	Previous string
}

func NewClient(interval time.Duration) Client {
	return Client{
		cache: *pokecache.NewCache(interval * time.Second),
		config: Config{
			Next:     "https://pokeapi.co/api/v2/location-area/1",
			Previous: "https://pokeapi.co/api/v2/location-area/1",
		},
	}
}
