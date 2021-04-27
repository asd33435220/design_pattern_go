package main

func main() {
	shirtItem := newItem("Nike Shirt")

	zhuyuchen := &customer{id: "zhuyuchen"}
	liangyilang := &customer{id: "liangyilang"}
	chenjiashen := &customer{id: "chenjiashen"}

	shirtItem.register(zhuyuchen)
	shirtItem.register(liangyilang)
	shirtItem.register(chenjiashen)

	shirtItem.updateAvailability()

	shirtItem.deregister(zhuyuchen)
	shirtItem.updateAvailability()


}
