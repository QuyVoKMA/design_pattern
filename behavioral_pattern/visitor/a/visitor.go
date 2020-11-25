package a

type Visitor interface {
	visitForSquare(s *Square)
	visitForCirCle(c *Circle)
	visitForRectange(r *Rectange)
}
