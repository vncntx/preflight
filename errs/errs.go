package errs

import "strings"

// CombinedError wraps multiple errors
type CombinedError struct {
	e string
}

// Combine returns a new CombinedError
func Combine(errors ...error) error {
	e := []string{}
	for _, err := range errors {
		if err != nil {
			e = append(e, err.Error())
		}
	}

	if len(e) < 1 {
		return nil
	}

	return &CombinedError{
		e: strings.Join(e, "; "),
	}
}

func (c *CombinedError) Error() string {
	return c.e
}
