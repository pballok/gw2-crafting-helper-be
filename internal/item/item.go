package item

import (
	"github.com/pballok/gw2-crafting-helper/backend/internal/guildwars2"
)

type Item struct {
	Id    int
	Name  string
	Icon  string
	Price struct {
		Buy  int
		Sell int
	}
	Recipes []Recipe
}

func NewItem(id int) *Item {
	return &Item{Id: id}
}

func (i *Item) FetchAll() error {
	err := i.fetchBase()
	if err != nil {
		return err
	}

	err = i.fetchPrices()
	if err != nil {
		return err
	}

	err = i.fetchRecipes()
	if err != nil {
		return err
	}

	return nil
}

func (i *Item) fetchBase() error {
	item, err := guildwars2.FetchItem(i.Id)
	if err != nil {
		return err
	}

	i.Name = item.Name
	i.Icon = item.Icon

	return nil
}

func (i *Item) fetchPrices() error {
	prices, err := guildwars2.FetchPrices(i.Id)
	if err != nil {
		return err
	}

	i.Price.Buy = prices.Buys.Price
	i.Price.Sell = prices.Sells.Price

	return nil
}

func (i *Item) fetchRecipes() error {
	recipeIdList, err := guildwars2.SearchRecipesOutput(i.Id)
	if err != nil {
		return err
	}

	for _, recipeId := range recipeIdList {
		recipe := NewRecipe(recipeId)
		err = recipe.FetchAll()
		if err != nil {
			return err
		}
		i.Recipes = append(i.Recipes, *recipe)
	}

	return nil
}
