package errors_test

import (
	"fmt"
	"testing"

	errors "github.com/walteh/go-errors/pkg-errors"
)

func TestFrameMarshalText(t *testing.T) {
	frame := errors.Frame(0)
	dt, err := frame.MarshalText()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(dt))
}
