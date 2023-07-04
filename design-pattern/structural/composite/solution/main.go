package main

import "fmt"

type Item interface {
	Cost() float32
}

type RealItem struct {
	Name  string
	Price float32
}

func (item RealItem) Cost() float32 {
	return item.Price
}

type Box struct {
	Children []Item
}

func (box Box) Cost() float32 {
	var cost float32 = 0.0
	for _, child := range box.Children {
		cost += child.Cost()
	}
	return cost
}

func CreatePackage() Item {
	return Box{
		Children: []Item{
			RealItem{
				Name:  "Mouse",
				Price: 20.5,
			},
			Box{
				Children: []Item{
					RealItem{
						Name:  "Keyboard",
						Price: 60,
					},
					RealItem{
						Name:  "Charger",
						Price: 15,
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
