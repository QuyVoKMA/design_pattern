package a

type OTP struct {
	IOTP IOtp
	//Khai báo cái biên(trong golang ko có khái niệm thừa kế) nó sẽ có được tất cả các phương thức cài đặt trong propeties này.
}
func (o *OTP) GenAndSendOTP(otpLenght int) error {
	otp :=o.IOTP.genRandomOTP(otpLenght)
	o.IOTP.saveOTPCache(otp)
	message :=o.IOTP.getMessage(otp)
	err := o.IOTP.sendNotification(message)
	if err != nil{
		return err
	}
	return nil
}

