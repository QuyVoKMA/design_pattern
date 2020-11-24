package a

import "fmt"

/*Cusotmer is struct*/
type Customer struct {
	ID string
}
/*Update is function*/
func (c *Customer) Update(itemName string){
	fmt.Printf("Sending email to customer %s for item %s\n", c.ID, itemName)
}
/*GetID is function*/
func (c *Customer) GetID() string{
	return c.ID
}
