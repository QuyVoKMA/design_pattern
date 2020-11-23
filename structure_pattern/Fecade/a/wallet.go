package a

import "fmt"

type wallet struct {
	balance int
}

func newWallet() *wallet {
	return &wallet{
		balance:0,
	}
}

func (w *wallet) creditWallet(money int){
	w.balance +=  money
	fmt.Printf("Wallet balance add successfully")
}

func (w *wallet) debitWallet(money int) error{
	if w.balance < money {
		return fmt.Errorf("Balance is not sufficient")
	}
	w.balance -=money
	fmt.Println("Wallet balance debit successfully")
	return nil
}