package main

import (
	"fmt"
)

type walletFacade struct {
	account      *account
	wallet       *wallet
	securityCode *securityCode
	notification *notification
	ledger       *ledger
}

func newWalletFacade(accountId string, code int) *walletFacade {
	fmt.Println("Staring create account")
	myWalletFacade := &walletFacade{
		account:      newAccount(accountId),
		wallet:       newWallet(),
		securityCode: newSecurityCode(code),
		notification: &notification{},
		ledger:       &ledger{},
	}
	fmt.Println("Account created")
	return myWalletFacade
}

func (w *walletFacade) addMoneyToWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("Starting add money to wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	w.wallet.creditBalance(amount)
	w.notification.sendWalletCreditNotification()
	w.ledger.makeEntry(accountId, "credit", amount)
	return nil
}

func (w *walletFacade) deductMoneyFromWallet(accountId string, securityCode int, amount int) error {
	fmt.Println("Starting debit money to wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return err
	}
	err = w.wallet.debitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.sendWalletDebitNotification()
	w.ledger.makeEntry(accountId, "credit", amount)
	return nil
}
func (w *walletFacade) checkMoneyFromWallet(accountId string, securityCode int) (int, error) {
	fmt.Println("Starting debit money to wallet")
	err := w.account.checkAccount(accountId)
	if err != nil {
		return -1, err
	}
	err = w.securityCode.checkCode(securityCode)
	if err != nil {
		return -1, err
	}
	return w.wallet.getMoney(), nil
}
