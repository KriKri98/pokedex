package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/KriKri98/pokedex/internal/pokecache"
)

func (c *Client) Get(url string) (map[string]any, error) {
	var data map[string]any

	storage, ok := c.Cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return data, err
		}

		defer res.Body.Close()
		statusCode := res.StatusCode
		if statusCode >= 400 {
			return data, fmt.Errorf("Status: %v", statusCode)
		}
		storage, err = io.ReadAll(res.Body)
		if err != nil {
			return data, err
		}
		c.Cache.Add(url, storage)

	}
	if err := json.Unmarshal(storage, &data); err != nil {
		return data, err
	}

	return data, nil

}

type Client struct {
	Cache  pokecache.Cache
	Config Config
}

type Config struct {
	Next     string
	Previous string
}

func NewClient(interval time.Duration) Client {
	return Client{
		Cache: *pokecache.NewCache(interval * time.Second),
		Config: Config{
			Next:     "https://pokeapi.co/api/v2/location-area/1",
			Previous: "https://pokeapi.co/api/v2/location-area/1",
		},
	}
}
