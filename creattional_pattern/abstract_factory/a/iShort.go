package a

type IShort interface {
	setSize(size int)
	setLogo(logo string)
	GetSize() int
	GetLogo() string
}

type iShort struct {
	logo string
	size int
}

func (i *iShort) setSize(size int){
	i.size = size
}

func (i *iShort) setLogo(logo string){
	i.logo = logo
}

func (i *iShort) GetSize() int{
	return i.size
}
func (i *iShort) GetLogo() string{
	return i.logo
}