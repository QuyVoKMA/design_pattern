package main

import (
	v "observer/visitor/a"
)

func main(){
	square := &v.Square{Side: 2}
	circle := &v.Circle{Radius:40}
	rectange := &v.Rectange{Height:2, Weight: 3}

	areaCal := &v.AreaCal{}
	square.Accept(areaCal)
	circle.Accept(areaCal)
	rectange.Accept(areaCal)

	perimete := &v.AreaCal{}
	square.Accept(perimete)
	circle.Accept(perimete)
	rectange.Accept(perimete)

}



