package main

import (
	"abs/a"
	"fmt"
)

func main(){
	adidasFactory := a.GetSprotFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	printShoeDetails(adidasShoe)
	adidasShort := adidasFactory.MakeShort()
	printShortDetails(adidasShort)

	nikeFactory := a.GetSprotFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	printShoeDetails(nikeShoe)
	nikeShort := nikeFactory.MakeShort()
	printShortDetails(nikeShort)
}

func printShoeDetails(s  a.IShoe){
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n",s.GetSize())
}

func printShortDetails(s  a.IShort){
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n",s.GetSize())
}
