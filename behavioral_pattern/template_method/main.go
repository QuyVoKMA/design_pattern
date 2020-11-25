package main

import (
	t "observer/template_method/a"
)

func main() {
	smsOTP := &t.SMS{}
	o :=&t.OTP{
		IOTP: smsOTP,
	}
	o.GenAndSendOTP(4)
	emailOTP :=&t.Email{}
	o =&t.OTP{
		IOTP:emailOTP,
	}
	o.GenAndSendOTP(4)
}
