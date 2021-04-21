package main

import "fmt"

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")
	printGunDetail(ak47)
	printGunDetail(musket)
}

func printGunDetail(gun iGun) {
	fmt.Println("gun's name :", gun.getName())
	fmt.Println("gun's power :", gun.getPower())

}
