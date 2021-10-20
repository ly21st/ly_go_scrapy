package main

import (
	"errors"
	"fmt"
)

func main() {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)

	fmt.Println(errors.Is(err3, err2))
	fmt.Println(errors.Is(err3, err1))
}
