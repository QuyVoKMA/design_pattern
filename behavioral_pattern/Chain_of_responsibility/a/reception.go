package a

import "fmt"

type Reception struct {
	next Department
}

func (r *Reception) Execute(p *Patient){
	if p.isRegister{
		fmt.Println("Patient registion has already done")
		r.next.Execute(p)
		return
	}
	fmt.Println("Reception registing patient")
	p.isRegister= true
	r.next.Execute(p)

}

func (r *Reception) SetNext(next Department){
	r.next=next
}
