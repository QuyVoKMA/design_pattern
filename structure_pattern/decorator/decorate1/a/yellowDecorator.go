package a

type YellowDecorator struct {
	house Ihouse
}

func NewYellowDecorator(house Ihouse) *YellowDecorator{
	return &YellowDecorator{
		house:house,
	}
}

func (y *YellowDecorator) DoHouse() string{
	PaintColor :=y.house.DoHouse()
	return PaintColor + y.AddColor()
}
func (y *YellowDecorator)AddColor() string{
 	return "  Yellow"
}
