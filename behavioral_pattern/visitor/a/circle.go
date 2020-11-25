package a

type Circle struct {
	Radius int
}

func (c *Circle) Accept(v Visitor){
	v.visitForCirCle(c)
}
