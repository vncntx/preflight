package preflight

import "testing"

// Unit returns a new unit test
func Unit(t *testing.T) *Test {
	return &Test{t}
}
