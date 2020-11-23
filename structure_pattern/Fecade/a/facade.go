package a

import "fmt"

type WalletFacade struct {
	account *account
	securityCode *securityCode
	wallet  *wallet
	notification *notification
}
func NewAccountFacade(accountID string, code int) *WalletFacade{
	fmt.Println("Starting create account")
	walletFacade:= &WalletFacade{
		account:newAccount(accountID),
		securityCode: newSecurityCode(code),
		wallet: newWallet(),
		notification: &notification{},
	}
	return walletFacade
}

func (w *WalletFacade) AddMoneyIntoWallet(accountID string, securityCode int, money int) error{
	fmt.Println("Starting add money into Wallet")
	err := w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkSecurityCode(securityCode)
	if err != nil{
		return err
	}
	w.wallet.creditWallet(money)
	w.notification.sendWalletCreditNotification()
	return nil
}

func (w *WalletFacade) DeductMoneyToWallet(accountID string, securityCode int, money int) error{
	fmt.Println("Starting debit money to wallet")
	err :=w.account.checkAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.checkSecurityCode(securityCode)
	if err != nil {
		return err
	}
	err =w.wallet.debitWallet(money)
	if err != nil{
		return err
	}
	w.notification.sendWalletDeditNotification()
	return nil
}