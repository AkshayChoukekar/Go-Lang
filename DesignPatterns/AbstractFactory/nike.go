package main

type nike struct{}

func (a *nike) makeShoe() iShoe {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: 7,
		},
	}
}

func (s *nike) makeShort() iShort {
	return &nikeShort{
		short: short{
			logo: "nike",
			size: 7,
		},
	}

}
