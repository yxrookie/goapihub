package sms

type Driver interface {
	// send message
	Send(phone string, message Message, config map[string]string) bool
}