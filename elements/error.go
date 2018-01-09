package elements

import (
	"fmt"
)

// Error is the error type.
type Error string

// Error returns the error message.
func (e Error) Error() string {
	return string(e)
}

// FormatError is an error that creates a unique message from the state at the time of creation.
type FormatError struct {
	message string
	items   []interface{}
}

// Error returns the error message.
func (e FormatError) Error() string {
	return fmt.Sprintf(e.message, e.items...)
}

// CumulativeError defines a collection of errors.
type CumulativeError struct {
	items []error
}

// Append the error to this error.
func (cumErr *CumulativeError) Append(err ...error) {
	for _, e := range err {
		if e != nil {
			switch v := e.(type) {
			case *CumulativeError:
				cumErr.Append(v.items...)
			default:
				cumErr.items = append(cumErr.items, e)
			}
		}
	}
}

// Error returns the error message.
func (cumErr *CumulativeError) Error() string {
	message := ""
	for index, part := range cumErr.items {
		message += fmt.Sprintf("%d: %s\n", index, part)
	}

	return message
}

// ErrorList returns the error collection.
func (cumErr *CumulativeError) ErrorList() []error {
	return cumErr.items
}

// NewError creates a new error.
func NewError(message string, contents ...interface{}) (err error) {

	if len(contents) == 0 {
		err = Error(message)
	} else {
		err = &FormatError{
			message: message,
			items:   contents,
		}
	}

	return err
}

// AppendError will take the original error and append the new one.
func AppendError(errors ...error) (err error) {
	switch len(errors) {

	// handle the simple cases where one or both are nil.
	case 0:
		break
	case 1:
		err = errors[0]

	default:
		for i := 0; i < len(errors); i++ {
			if errors[i] != nil {
				switch v := errors[i].(type) {
				case *CumulativeError:
					v.Append(errors[i+1:]...)
					err = v
				default:
					if len(errors)-i > 1 {
						cumErr := &CumulativeError{}
						cumErr.Append(errors[i:]...)
						err = cumErr
					} else {
						err = errors[i]
					}
				}
				break
			}
		}
	}

	return err
}
