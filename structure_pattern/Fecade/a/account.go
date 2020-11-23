package a

import "fmt"

type account struct {
	name string
}

func newAccount(nameAccount string) *account{
	return &account{
		name:nameAccount,
	}
}
func (a *account) checkAccount(accountName string) error{
	if a.name !=accountName {
		return fmt.Errorf("Account name is incorrect")
	}
	fmt.Println("Account verified")
	return nil
}