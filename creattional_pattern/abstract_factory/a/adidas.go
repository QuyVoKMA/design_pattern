package a

// Day la 1 cai factory

//Muc dich la cai dat cac phuong thuc vao strcut nay
type Adidas struct {}

func (a *Adidas) MakeShoe() IShoe {
	return &adidasShoe{
		iShoe :iShoe{
			logo: "adidas shoe",
			size: 40,
		},
	}
}

func (a *Adidas) MakeShort() IShort{
	return &adidasShort{
		iShort :iShort{
			logo: "adidas short",
			size: 40,
		},
	}
}
