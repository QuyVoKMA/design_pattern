package a

type ISportFactory interface {
	MakeShoe() IShoe
	MakeShort() IShort
}

func GetSprotFactory(brach string) ISportFactory {
	switch brach {
	case "adidas":
		return &Adidas{}

	case "nike":
		return &Nike{}
	}
	return nil
}
