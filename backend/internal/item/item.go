package item

import "github.com/pballok/gw2-crafting-helper/backend/internal/guildwars2"

type Item struct {
	Id   uint16
	Name string
}

func fromAPI(apiItem *guildwars2.Item) *Item {
	return &Item{
		Id:   apiItem.Id,
		Name: apiItem.Name,
	}
}

func FromID(id uint16) (*Item, error) {
	item, err := guildwars2.FetchItem(id)
	if err != nil {
		return nil, err
	}

	return fromAPI(item), nil
}
