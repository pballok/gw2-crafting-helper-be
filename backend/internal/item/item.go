package item

import "github.com/pballok/gw2-crafting-helper/backend/internal/guildwars2"

type Item struct {
	Id   int
	Name string
}

func NewItem(id int) (*Item, error) {
	item, err := guildwars2.FetchItem(id)
	if err != nil {
		return nil, err
	}

	return &Item{
		Id:   item.Id,
		Name: item.Name,
	}, nil
}
