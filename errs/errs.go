package errs

import "strings"

// CombinedError wraps multiple errors
type CombinedError struct {
	errors []error
}

// Combine returns a new CombinedError
func Combine(errors ...error) error {
	return &CombinedError{errors}
}

func (c *CombinedError) Error() string {
	e := []string{}
	for _, err := range c.errors {
		e = append(e, err.Error())
	}

	return strings.Join(e, "; ")
}
