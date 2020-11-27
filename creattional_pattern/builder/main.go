package main

import (
	"crea/builder/a"
	"fmt"
)

func main(){
	nomalBuilder := a.GetBuilder("nomal")
	iglooBuilder :=a.GetBuilder("igloo")

	director := a.NewDirector(nomalBuilder)
	nomalHouse :=director.BuildHouse()


	fmt.Printf("Get Window Type %s\n", nomalHouse.GetWindowType())
	fmt.Printf("Get Door Type %s\n", nomalHouse.GetDoorType())
	fmt.Printf("Get Floor %d\n", nomalHouse.Getfloor())

	director.SetBuilder(iglooBuilder)
	iglooHouse :=director.BuildHouse()

	fmt.Println("-----------------------------------------------------------------")

	fmt.Printf("Get Window Type %s\n", iglooHouse.GetWindowType())
	fmt.Printf("Get Door Type %s\n", iglooHouse.GetDoorType())
	fmt.Printf("Get Floor %d\n", iglooHouse.Getfloor())

}
