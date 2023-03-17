package item

import "github.com/pballok/gw2-crafting-helper/backend/internal/guildwars2"

type Item struct {
	Id    int
	Name  string
	Icon  string
	Price struct {
		Buy  int
		Sell int
	}
}

func NewItem(id int) (*Item, error) {
	item := Item{Id: id}

	err := item.fetchBase()
	if err != nil {
		return nil, err
	}

	err = item.fetchPrices()
	if err != nil {
		return nil, err
	}

	return &item, nil
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
