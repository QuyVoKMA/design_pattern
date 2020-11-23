package main

import (
	d "design_pattern/structure_pattern/decorator/decorate1/a"
	"fmt"
)

func main() {

	houseWood := &d.HouseWood{}
	houseModern :=&d.HouseModern{}
	fmt.Println(houseWood.DoHouse())
	fmt.Println(houseModern.DoHouse())


	yellowDecorator := d.NewYellowDecorator(houseWood)
	blueDecorator :=d.NewBlueDecorator(houseModern)

	fmt.Println(yellowDecorator.DoHouse())
	fmt.Println(blueDecorator.DoHouse())



}

