package staque

import "github.com/cheekybits/genny/generic"

//go:generate genny -in=$GOFILE -out=staque_specialized.go gen "Generic=BUILTINS"

type Generic generic.Type
type GenericStaque []Generic

func NewGeneric() GenericStaque {
	return GenericStaque{}
}

func (staque GenericStaque) Push(xs ...Generic) GenericStaque {
	return append(staque, xs...)
}

func (stk GenericStaque) Peekstk() (last Generic, isempty error) {
	ilast := len(stk) - 1
	if ilast < 0 {
		return last, &empty{"Cannot Peek() on empty staque"}
	}
	return stk[ilast], nil
}

func (que GenericStaque) Peekque() (first Generic, isempty error) {
	if len(que) == 0 {
		return first, &empty{"Cannot Peek() on empty staque"}
	}
	return que[0], nil
}

func (stk GenericStaque) Popstk() (modified GenericStaque, last Generic, isempty error) {
	ilast := len(stk) - 1
	if ilast < 0 {
		return nil, last, &empty{"Cannot Pop() on empty staque"}
	}

	if ilast < cap(stk) / 4 {
		return append(make([]Generic, 0, cap(stk) / 2), stk[:ilast]...), stk[ilast], nil
	} else {
		return stk[:ilast], stk[ilast], nil
	}
}

func (que GenericStaque) Popque() (modified GenericStaque, first Generic, isempty error) {
	len := len(que)
	if len == 0 {
		return nil, first, &empty{"Cannot Pop() on empty staque"}
	}

	if len > cap(que) / 4 {
		return que[1:], que[0], nil
	} else {
		return append(make([]Generic, 0, cap(que) / 2), que[1:]...), que[0], nil
	}
}