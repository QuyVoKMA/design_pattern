package a

type BlueDecorator struct {
	house Ihouse
}

func NewBlueDecorator(house Ihouse) *BlueDecorator{
	return &BlueDecorator{
		house: house,
	}
}

func (b *BlueDecorator)DoHouse() string{
	PaintColor := b.house.DoHouse()
	return PaintColor + b.AddColor()
}

func (b *BlueDecorator) AddColor() string{
	return " son mau Blue"
}