package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	// err := errors.Errorf("whoops: %s", "foo")
	// fmt.Printf("%+v", err)

	cause := errors.New("whoops")
	err := errors.WithStack(cause)
	fmt.Printf("%+v", err)
}
