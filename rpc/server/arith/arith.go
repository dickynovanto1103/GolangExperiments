package arith

import (
	"errors"
	"log"
)

type Arith int
type Args struct {
	A, B int
}

type Quotient struct {
	Res, Rem int
}

func (a *Arith) Multiply(args Args, result *int) error {
	res := args.A * args.B
	*result = res
	log.Printf("halo, a: %v b: %v", args.A, args.B)
	return nil
}

func (a *Arith) Divide(args Args, result *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	res := args.A / args.B
	rem := args.A % args.B
	//result.Res = res
	//result.Rem = rem
	*result = Quotient{res, rem}
	return nil
}
