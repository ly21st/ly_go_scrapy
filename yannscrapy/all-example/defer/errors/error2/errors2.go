package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)

	fmt.Println(errors.Unwrap(err3))
	fmt.Println(errors.Unwrap(errors.Unwrap(err3)))
}
