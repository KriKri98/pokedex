package pokeapi

import (
	"encoding/json"
	"net/http"
)

func get(url string) (map[string]any, error) {
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
