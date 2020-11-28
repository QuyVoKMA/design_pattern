package a

type IShoe interface {
	setLogo(logo string)
	setSize(size int)
	GetLogo() string
	GetSize() int
}

type iShoe struct {
	logo string
	size int
}

func (i *iShoe)setSize(size int){
	i.size = size
}

func (i *iShoe)setLogo(logo string){
	i.logo = logo
}

func (i *iShoe)GetSize() int{
	return i.size
}

func (i *iShoe)GetLogo() string{
	return i.logo
}
