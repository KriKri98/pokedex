package pokeapi

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/KriKri98/pokedex/internal/pokecache"
)

func (c *Client) Get(url string) ([]byte, error) {
	var data []byte

	storage, ok := c.Cache.Get(url)
	if !ok {
		res, err := http.Get(url)
		if err != nil {
			return data, err
		}

		defer res.Body.Close()
		statusCode := res.StatusCode
		if statusCode >= 400 {
			return data, fmt.Errorf("status: %v", statusCode)
		}
		storage, err = io.ReadAll(res.Body)
		if err != nil {
			return data, err
		}
		c.Cache.Add(url, storage)
	}

	return storage, nil

}

type Client struct {
	Cache   *pokecache.Cache
	Config  Config
	Pokedex map[string]Pokemon
}

type Config struct {
	Next     string
	Previous string
}

func NewClient(interval time.Duration) Client {
	return Client{
		Cache: pokecache.NewCache(interval),
		Config: Config{
			Next: "https://pokeapi.co/api/v2/location-area",
		},
		Pokedex: make(map[string]Pokemon),
	}
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}
