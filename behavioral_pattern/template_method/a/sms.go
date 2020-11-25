package a

import "fmt"

type SMS struct {

}

func (s *SMS) genRandomOTP(len int) string{
	randomOTP :="123"
	fmt.Printf("SMS : generating random OTP %s\n", randomOTP)
	return randomOTP
}
func (s *SMS) saveOTPCache(otp string){
	fmt.Printf("SMS: saving OTP: %s to cache\n", otp)
}

func (s *SMS) getMessage(otp string) string{
	return "SMS: saving OTP for login\n"+ otp
}
func (s *SMS) sendNotification(message string) error{
	fmt.Printf("SMS: seding sms: %s \n", message)
	return nil
}