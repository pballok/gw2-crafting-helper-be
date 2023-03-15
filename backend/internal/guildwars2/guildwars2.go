package guildwars2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const api_base = "https://api.guildwars2.com/v2"

type Item struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func FetchItem(id int) (*Item, error) {
	resp, err := http.Get(fmt.Sprintf("%s/items/%d", api_base, id))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch item %d : %w", id, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return &Item{}, nil
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch item %d from guildwars2: API returned with %d", id, resp.StatusCode)
	}

	item := Item{}
	err = json.NewDecoder(resp.Body).Decode(&item)
	if err != nil {
		return nil, fmt.Errorf("failed to decode item %d from response : %w", id, err)
	}

	return &item, nil
}
