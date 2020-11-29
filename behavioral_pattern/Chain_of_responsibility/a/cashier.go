package a

import "fmt"

type Cashier struct {
	next Department
}
func (c *Cashier) Execute(p *Patient){
	if p.isPaid{
		fmt.Println("Patient already paid for hostipal")
		c.next.Execute(p)
		return
	}
	fmt.Println("Patient is paying for hospital")
	p.isPaid=true
	c.next.Execute(p)
}

func (c *Cashier) SetNext(next Department){
	c.next=next
}
