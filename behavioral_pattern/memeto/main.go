package main

import (
	"behavior/memeto/a"
	"fmt"
)

func main(){
 caretaker := &a.CareTaker{
 	MementoArry: make([]*a.Memento,0),
 }
 originator := &a.Originator{
 	State: "A",
 }
 fmt.Printf(" Orig current state: %s \n", originator.GetState())
 caretaker.AddMemento(originator.CreateMemeto())

	originator.SetState("B")
	fmt.Printf(" Orig current state: %s \n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemeto())
	originator.SetState("C")
	fmt.Printf(" Orig current state: %s \n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemeto())

 	originator.RestoreMemeto(caretaker.GetMemento(1))
	fmt.Printf(" Orig current state: %s \n", originator.GetState())
	originator.RestoreMemeto(caretaker.GetMemento(0))
	fmt.Printf(" Orig current state: %s \n", originator.GetState())
	originator.RestoreMemeto(caretaker.GetMemento(2))
	fmt.Printf(" Orig current state: %s \n", originator.GetState())
	originator.RestoreMemeto(caretaker.GetMemento(0))
	fmt.Printf(" Orig current state: %s \n", originator.GetState())


}
