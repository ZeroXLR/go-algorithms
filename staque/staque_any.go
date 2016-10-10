package staque

/*
  Unfortunately without true generics, I have to allow arbitrary elements into the staque. The alternatives to this are:
  1) Provide a staque that allows only one type of element (like int) which is very limiting.
  2) Hand-code staques for every type: int, rune, byte, string, etc. Well, this is exactly the kind of meaningless effort that generics should solve. Also, I cannot predict every type that is ever going to come into existence!
*/
type any interface{}
type AnyStaque []any

func New() AnyStaque {
	return AnyStaque{}
}

func (stk *AnyStaque) Push(xs ...any) {
	*stk = append(*stk, xs...)
}

type Empty string
func (empty Empty) Error() string {
	return string(empty)
}

func (stk *AnyStaque) Peekstk() (any, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return nil, Empty("Cannot Peek() on empty staque")
	}
	return (*stk)[ilast], nil
}

func (que *AnyStaque) Peekque() (any, error) {
	if len(*que) == 0 {
		return nil, Empty("Cannot Peek() on empty staque")
	}
	return (*que)[0], nil
}

func (stk *AnyStaque) Popstk() (any, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return nil, Empty("Cannot Pop() on empty staque")
	}

	last := (*stk)[ilast] // save last value; it won't be available afterwards
	if ilast < cap(*stk) / 4 {
		*stk = append(make(AnyStaque, 0, cap(*stk) / 2), (*stk)[:ilast]...)
	} else {
		*stk = (*stk)[:ilast]
	}
	return last, nil
}

func (que *AnyStaque) Popque() (any, error) {
	len := len(*que)
	if len == 0 {
		return nil, Empty("Cannot Pop() on empty staque")
	}

	first := (*que)[0] // save first value; it won't be available afterwards
	if len > cap(*que) / 4 {
		*que = (*que)[1:]
	} else {
		*que = append(make(AnyStaque, 0, cap(*que) / 2), (*que)[1:]...)
	}
	return first, nil
}