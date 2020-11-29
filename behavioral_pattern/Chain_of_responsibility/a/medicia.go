package a

import "fmt"

type Medicia struct {
	next Department
}

func (r *Medicia) Execute(p *Patient){
	if p.isMedicineProvide{
		fmt.Println("Medicia Department already provided Medicine for patient")
		r.next.Execute(p)
		return
	}
	fmt.Println("Medicia Department is providing medicine for patient")
	p.isMedicineProvide=true
	r.next.Execute(p)
}

func (r *Medicia) SetNext(next Department){
	r.next=next
}
