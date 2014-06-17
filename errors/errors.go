package errors

import kstr "github.com/koanhealth/gotools/strings"

type ErrorSlice []error

func (errs ErrorSlice) Error() string {
	message := kstr.NewStringBuilder()
	message.Print("Errors: ")
	for _, err := range errs {
		message.Printf("\n  %s", err.Error())
	}
	return message.String()
}
