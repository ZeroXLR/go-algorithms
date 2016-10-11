package staque

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=staque_specialized.go gen "Generic=BUILTINS"

type Generic generic.Type
type GenericStaque []Generic

func NewGeneric() GenericStaque {
	return GenericStaque{}
}

func (staque *GenericStaque) Push(xs ...Generic) {
	*staque = append(*staque, xs...)
}

func (stk *GenericStaque) Peekstk() (last Generic, isempty error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return last, Empty("Cannot Peek() on empty staque")
	}
	return (*stk)[ilast], nil
}

func (que *GenericStaque) Peekque() (first Generic, isempty error) {
	if len(*que) == 0 {
		return first, Empty("Cannot Peek() on empty staque")
	}
	return (*que)[0], nil
}

func (stk *GenericStaque) Popstk() (last Generic, isempty error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return last, Empty("Cannot Pop() on empty staque")
	}

	last = (*stk)[ilast]
	if ilast < cap(*stk) / 4 {
		*stk = append(make([]Generic, 0, cap(*stk) / 2), (*stk)[:ilast]...)
	} else {
		*stk = (*stk)[:ilast]
	}
	return last, nil
}

func (que *GenericStaque) Popque() (first Generic, isempty error) {
	len := len(*que)
	if len == 0 {
		return first, Empty("Cannot Pop() on empty staque")
	}

	first = (*que)[0]
	if len > cap(*que) / 4 {
		*que = (*que)[1:]
	} else {
		*que = append(make([]Generic, 0, cap(*que) / 2), (*que)[1:]...)
	}
	return first, nil
}