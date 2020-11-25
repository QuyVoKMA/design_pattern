package a

type Rectange struct {
	Height int
	Weight int
}

func (r *Rectange) Accept(v Visitor){
	v.visitForRectange(r)
}
