package main

type ak47 struct {
	gun
}

func newAk47() iGun {
	return &ak47{
		gun: gun{
			name:  "AK47 Gun",
			power: 1000,
		},
	}
}
