package a

type Square struct {
	Side int
}


func (s *Square) Accept(v *AreaCal) {
	v.visitForSquare(s)
}