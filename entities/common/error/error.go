package error

import "github.com/wildanie12/go-hex-arch-sample/config"

// WebError main struct contains message for both development and production and error code
type WebError struct {
	code int 
	message string
	prodMessage string
	devMessage string
}

// New error object with simple message
func New(code int, message string) error {
	return WebError{
		code: code,
		message: message,
	}
}

// Make error object with full production and development message
func Make(code int, prodMsg string, devMsg string) error {
	return WebError{
		code: code,
		prodMessage: prodMsg,
		devMessage: devMsg,
	}
}

func (err WebError) Error() string {
	env := config.New().App.Env
	message := map[bool]string{ true: err.prodMessage, false: err.devMessage }[env == "production"]
	if message != "" {
		return message
	}
	return err.message
}