package mail

type Driver interface {
	// check up captcha
	Send(email Email, config map[string]string) bool
}
