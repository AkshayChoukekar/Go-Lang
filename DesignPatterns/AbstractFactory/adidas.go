package main

type adidas struct{}

func (a *adidas) makeShoe() iShoe {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: 7,
		},
	}
}

func (s *adidas) makeShort() iShort {
	return &adidasShort{
		short: short{
			logo: "adidas",
			size: 7,
		},
	}

}
