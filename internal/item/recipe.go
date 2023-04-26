package item

import (
	"github.com/pballok/gw2-crafting-helper/backend/internal/guildwars2"
)

type Recipe struct {
	Id          int
	OutputCount int
	Items       []IngredientItem
	Currencies  []IngredientCurrency
}

type IngredientItem struct {
	Count      int
	Ingredient Item
}

type IngredientCurrency struct {
	Count      int
	Ingredient Currency
}

func NewRecipe(id int) *Recipe {
	return &Recipe{Id: id}
}

func (r *Recipe) FetchAll() error {
	err := r.fetchBase()
	if err != nil {
		return err
	}

	err = r.fetchIngredients()
	if err != nil {
		return err
	}

	return nil
}

func (r *Recipe) fetchBase() error {
	gw2Recipe, err := guildwars2.FetchRecipe(r.Id)
	if err != nil {
		return err
	}

	r.OutputCount = gw2Recipe.OutputCount
	for _, ingredient := range gw2Recipe.Ingredients {
		if ingredient.Type == "Item" {
			item := NewItem(ingredient.Id)
			r.Items = append(r.Items, IngredientItem{
				Count:      ingredient.Count,
				Ingredient: *item,
			})
		} else if ingredient.Type == "Currency" {
			currency := NewCurrency(ingredient.Id)
			r.Currencies = append(r.Currencies, IngredientCurrency{
				Count:      ingredient.Count,
				Ingredient: *currency,
			})
		}
	}

	return nil
}

func (r *Recipe) fetchIngredients() error {
	for i := range r.Items {
		err := r.Items[i].Ingredient.FetchAll()
		if err != nil {
			return err
		}
	}
	for c := range r.Currencies {
		err := r.Currencies[c].Ingredient.FetchAll()
		if err != nil {
			return err
		}
	}

	return nil
}
