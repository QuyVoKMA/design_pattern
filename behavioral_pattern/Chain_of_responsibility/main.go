package main

import c "behavior/Chain_of_responsibility/a"

func main(){
	cashier := &c.Cashier{}

	medical := &c.Medicia{}
	medical.SetNext(cashier)

	docter := &c.Doctor{}
	docter.SetNext(medical)

	reception := &c.Reception{}
	reception.SetNext(docter)

	patient := &c.Patient{Name: "Quy",}
	reception.Execute(patient)

}
