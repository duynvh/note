package main

import "fmt"

type Item struct {
	Name string
	Price float32

	// Ops, an item always has children!! Why?
	Children []Item
}

func (item Item) Cost() float32 {
	cost := item.Price

	for _, child := range item.Children {
		cost += child.Cost()
	}

	return cost
}

func CreatePackage() Item {
	return Item{
		// It's a box, not an item
		Name: "root box",
		// so it have no price
		Price: 0,
		// but is has children
		Children: []Item{
			// Here it is!! My real item here.
			{
				Name: "Mouse",
				Price: 20.5,
				Children: nil,
			},
			// Another box contains item
			{
				Name: "sub box",
				Price: 0,
				Children: []Item{
					{
						Name: "Keyboard",
						Price: 60,
						Children: nil,
					},
					{
						Name: "Charger",
						Price: 15,
						Children: nil,
					},
				},
			},
		},
	}
}

func main() {
	pkg := CreatePackage()
	fmt.Println(pkg.Cost())
}