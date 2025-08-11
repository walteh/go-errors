package main

import (
	"github.com/walteh/go-errors"
)

func main() {
	panic(errors.WithDetails(errors.Base("panic error"), "key", "value"))
}
