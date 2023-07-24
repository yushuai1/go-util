package commons

import "errors"

func ToErr(msg string) error {
	return errors.New(msg)
}
