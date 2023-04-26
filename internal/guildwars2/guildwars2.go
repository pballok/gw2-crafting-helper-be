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
	Icon string `json:"icon"`
}

type Prices struct {
	Id   int `json:"id"`
	Buys struct {
		Price int `json:"unit_price"`
	} `json:"buys"`
	Sells struct {
		Price int `json:"unit_price"`
	} `json:"sells"`
}

type Recipe struct {
	Id          int `json:"id"`
	OutputCount int `json:"output_item_count"`
	Ingredients []struct {
		Type  string `json:"type"`
		Id    int    `json:"id"`
		Count int    `json:"count"`
	} `json:"ingredients"`
}

type resultType interface{}

func fetch[R resultType](url string) (*R, error) {
	fmt.Printf("fetching %s\n", url)
	var result R
	resp, err := http.Get(url)
	if err != nil {
		return &result, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return &result, nil
	}

	if resp.StatusCode != http.StatusOK {
		return &result, fmt.Errorf("failed to fetch from guildwars2: API returned with %d", resp.StatusCode)
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return &result, fmt.Errorf("failed to decode from response : %w", err)
	}

	return &result, nil
}

func FetchItem(id int) (*Item, error) {
	requestUrl := fmt.Sprintf("%s/items/%d", api_base, id)
	item, err := fetch[Item](requestUrl)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch item %d : %w", id, err)
	}

	return item, nil
}

func FetchPrices(id int) (*Prices, error) {
	requestUrl := fmt.Sprintf("%s/commerce/prices/%d", api_base, id)
	prices, err := fetch[Prices](requestUrl)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch prices for item %d : %w", id, err)
	}

	return prices, nil
}

func SearchRecipesOutput(id int) ([]int, error) {
	requestUrl := fmt.Sprintf("%s/recipes/search?output=%d", api_base, id)
	recipeIds, err := fetch[[]int](requestUrl)

	if err != nil {
		return nil, fmt.Errorf("failed to search for output recipes of item %d : %w", id, err)
	}

	return *recipeIds, nil
}

func FetchRecipe(id int) (*Recipe, error) {
	requestUrl := fmt.Sprintf("%s/recipes/%d?v=2022-03-09T02:00:00.000Z", api_base, id)
	recipe, err := fetch[Recipe](requestUrl)

	if err != nil {
		return nil, fmt.Errorf("failed to fetch recipe %d : %w", id, err)
	}

	return recipe, nil
}
