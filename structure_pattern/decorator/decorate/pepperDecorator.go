package decorate

type PapperDecorator struct {
	pizza Ipizze
}

func NewPapperDecorator(pizza Ipizze) *PapperDecorator{
	return &PapperDecorator{
		pizza: pizza,
	}
}

func (pp *PapperDecorator) Dopizza() string{
	pizzaType :=pp.pizza.Dopizza()
	return pizzaType + pp.addFlavor()
}

func (p *PapperDecorator)addFlavor() string{
	return "  papper"
}