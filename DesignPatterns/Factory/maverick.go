package main

type maverick struct {
	gun
}

func newMaverick() iGun {
	return &maverick{
		gun: gun{
			name:  "MAVERICK GUN",
			power: 2000,
		},
	}
}
