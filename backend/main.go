package main

import (
	"fmt"

	"github.com/pballok/gw2-crafting-helper/backend/internal/item"
)

func main() {
	itemId := uint16(19719)
	item, err := item.FromID(itemId)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Println(item.Name)
}
