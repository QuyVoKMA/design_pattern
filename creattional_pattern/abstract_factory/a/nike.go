package a

type Nike struct {

}
func (n *Nike) MakeShoe() IShoe{
	return &nikeShoe{
		iShoe:iShoe{
			logo: "make nike",
			size: 15,
		},
	}
}

func (n *Nike) MakeShort() IShort{
	return &nikeShort{
		iShort :iShort{
			logo: "make nike",
			size: 15,
		},
	}
}



