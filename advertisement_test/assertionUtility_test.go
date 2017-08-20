package advertisement_test

import (
	"testing"
	"fmt"
)

func AssertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Assert Failed: %v != %v", a, b)
	}
	t.Fatal(message)
}


func AssertNotEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a != b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Assert Failed: %v == %v", a, b)
	}
	t.Fatal(message)
}

func AssertNull(t *testing.T, a interface{}, message string) {
	if a == nil {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Assert Failed: %v != null", a)
	}
	t.Fatal(message)
}

func AssertNotNull(t *testing.T, a interface{}, message string) {
	if a != nil {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("Assert Failed: %v == null", a)
	}
	t.Fatal(message)
}
