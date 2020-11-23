package a

import "fmt"

type notification struct {

}

func (n *notification) sendWalletCreditNotification(){
	fmt.Println("Sending wallet credit notification")
}

func (n *notification) sendWalletDeditNotification(){
	fmt.Println("Sending wallet dedit notification")
}

