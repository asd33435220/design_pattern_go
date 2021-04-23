package main

import "fmt"

func main() {
	fmt.Println()
	walletFacade := newWalletFacade("zhuyuchen", 1234)

	err := walletFacade.addMoneyToWallet("zhuyuchen", 1234, 10)

	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Println()


	err = walletFacade.deductMoneyFromWallet("zhuyuchen", 1234, 5)
	if err != nil {
		fmt.Println(err.Error())
	}
	money, err := walletFacade.checkMoneyFromWallet("zhuyuchen", 1234)
	if err != nil {

		fmt.Println(err.Error())
	}
	fmt.Printf("%d $ leaf\n", money)
	err = walletFacade.deductMoneyFromWallet("zhuyuchen", 1234, 6)
	if err != nil {
		fmt.Println(err.Error())
	}
}
