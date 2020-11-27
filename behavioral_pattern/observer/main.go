package main

import (
	o "behavior/observer/a"
)

func main(){
	//Item moi
	shirtItem := o.NewItem("New Nike Shirt")

	// Các email được đăng ký
	customerObserver1 :=&o.Customer{ID:"abc@gmail.com"}
	customerObserver2 :=&o.Customer{ID:"xyz@gmail.com"}

	// Register cho customer
	shirtItem.Register(customerObserver1)
	shirtItem.Register(customerObserver2)
	//DeRegister cho customer
	shirtItem.DeRegister(customerObserver2)
	shirtItem.UpdateAvailability()
}