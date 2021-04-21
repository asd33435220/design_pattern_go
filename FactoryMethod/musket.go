package main

type musket struct {
	gun
}

func newMusket() iGun {
	return &musket{gun{
		name:  "musket gun",
		power: 1,
	}}
}
