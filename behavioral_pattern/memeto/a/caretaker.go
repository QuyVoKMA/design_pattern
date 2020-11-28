package a

type CareTaker struct {
	MementoArry []*Memento
}

func (c *CareTaker)AddMemento(m *Memento){
	c.MementoArry = append(c.MementoArry,m)
}

func(c *CareTaker)GetMemento(index int) *Memento{
	return c.MementoArry[index]
}