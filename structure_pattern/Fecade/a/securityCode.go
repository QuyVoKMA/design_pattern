package a

import "fmt"

type securityCode struct {
	code int

}

func newSecurityCode(code int) *securityCode {
	return &securityCode{
		code : code,
	}
}

func (s *securityCode) checkSecurityCode(code int) error {
	if s.code !=code {
		return fmt.Errorf("SecurityCode is incorrect")
	}
	fmt.Println("SecurityCode verified")
	return nil
}
