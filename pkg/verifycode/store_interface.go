package verifycode

type Store interface {
	// save captcha
	Set(id string, value string) bool

	// get captcha
	Get(id string, clear bool) string

	// verify captcha
	Verify(id, answer string, clear bool) bool
}