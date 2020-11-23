package decorate

type CheeseDecorator struct {
	pizza Ipizze
}

func NewCheeseDecorator(pizza Ipizze) *CheeseDecorator{
	return &CheeseDecorator{
		pizza:pizza, // gan thuoc tinh vs tham so truyen vao
	}
}

func (c *CheeseDecorator) DoPizza() string{
	// Giua lai hanh vi cu
	pizzaType :=c.pizza.Dopizza() // goi ra de giu lai hanh vi cu
	return pizzaType +  c.addFlavor() // Them cai mui vi moi
}

func (c *CheeseDecorator)addFlavor() string{
	return "   cheese"
}