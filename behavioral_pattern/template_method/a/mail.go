package a

import "fmt"

type Email struct {

}

func (s *Email) genRandomOTP(len int) string{
	randomOTP :="456"
	fmt.Printf("EMAIL : generating random OTP %s\n", randomOTP)
	return randomOTP
}
func (s *Email) saveOTPCache(otp string){
	fmt.Printf("EMAIL: saving OTP: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string{
	return "EMAIL: saving OTP for to login\n"+ otp
}
func (s *Email) sendNotification(message string) error{
	fmt.Printf("EMAIL: seding EMAIL: %s \n", message)
	return nil
}
