package main

import (
	d "design_pattern/structure_pattern/decorator/decorate"
	"fmt"
)

func main() {

	c :=&d.TomatoPizza{}
	chicken :=&d.ChickenPizza{}

	fmt.Println(c.Dopizza())
	fmt.Println(chicken.Dopizza())

	papperDecorator :=d.NewPapperDecorator(chicken)
	cheeseDecorator :=d.NewCheeseDecorator(chicken)

	fmt.Println(papperDecorator.Dopizza())
	fmt.Println(cheeseDecorator.DoPizza())

}
