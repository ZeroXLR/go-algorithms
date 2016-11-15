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
	if ilast := len(stk) - 1; ilast < 0 {
		isempty = emptypeek
	} else {
		last = stk[ilast]
	}
	return
}

func (que GenericStaque) Peekque() (first Generic, isempty error) {
	if len(que) == 0 {
		isempty = emptypeek
	} else {
		first = que[0]
	}
	return
}

func (stk GenericStaque) Popstk() (modified GenericStaque, last Generic, isempty error) {
	if ilast := len(stk) - 1; ilast < 0 {
		isempty = emptypop // as this generic function will become multi-typed, we have many possible
		return // zero values for last; so, we need named returns to set default values for last
	} else if ilast < cap(stk) / 4 {
		return append(make([]Generic, 0, cap(stk) / 2), stk[:ilast]...), stk[ilast], nil
	} else {
		return stk[:ilast], stk[ilast], nil
	}
}

func (que GenericStaque) Popque() (modified GenericStaque, first Generic, isempty error) {
	if len(que) == 0 {
		isempty = emptypop // as this generic function will become multi-typed, we have many possible
		return // zero values for first; so, we need named returns to set default values for first
	} else if len(que) > cap(que) / 4 {
		return que[1:], que[0], nil
	} else {
		return append(make([]Generic, 0, cap(que) / 2), que[1:]...), que[0], nil
	}
}