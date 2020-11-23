package main

import (
	f "design_pattern/structure_pattern/Fecade/a"
	"fmt"
	"log"
)

func main(){
	fmt.Println()
	walletFacade := f.NewAccountFacade("Quy",123)
	fmt.Println()
	err :=walletFacade.AddMoneyIntoWallet("Quy",123,1000)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println()
	err =walletFacade.DeductMoneyToWallet("Quy", 123, 100)
	if err!=nil{
		log.Fatal(err.Error())
	}


}
