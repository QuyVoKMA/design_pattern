package a

import "fmt"

type Doctor struct {
	next Department
}

func (r *Doctor) Execute(p *Patient){
	if p.isDoctorCheck{
		fmt.Println("doctor already cheked by Patient")
		r.next.Execute(p)
		return
	}
	fmt.Println("Doctor is checking for patient")
	p.isDoctorCheck=true
	r.next.Execute(p)
}

func (r *Doctor) SetNext(next Department){
	r.next=next
}