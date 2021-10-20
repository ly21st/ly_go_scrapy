package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func fn1() error {

	e1 := errors.New("error")
	// e1 = errors.Wrap(e1, "e1")
	return e1
}

func fn2() error {
	e1 := fn1()
	// e2 := errors.Wrap(e1, "e2:[e1]")
	return e1
}

func fn3() error {
	e2 := fn2()
	// e3 := errors.Wrap(e2, "e3:[e2]")
	// e3 := fmt.Errorf("e3:[%w]", e2)
	return e2
}

func main() {

	e3 := fn3()
	e4 := errors.WithMessage(e3, "main:e4")

	fmt.Printf("%T:%v\n", errors.Cause(e4), errors.Cause(e4))
	fmt.Println()
	fmt.Printf("%+v\n", e4)

	// type stackTracer interface {
	// 	StackTrace() errors.StackTrace
	// }

	// err, ok := errors.Cause(e4).(stackTracer)
	// if !ok {
	// 	panic("oops, err does not implement stackTracer")
	// }

	// st := err.StackTrace()
	// // fmt.Printf("%+v", st[0:2]) // top two frames
	// fmt.Printf("%+v", st[:])

}
