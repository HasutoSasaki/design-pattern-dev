package main

import "fmt"

type Email struct {
	Otp
}

func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("Email: generating random otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("Email: saving otp %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (s *Email) sendNotification(message string) error {
	fmt.Printf("Email: sending email: %s\n", message)
	return nil
}
