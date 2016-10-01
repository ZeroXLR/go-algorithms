package stkque

/*
  Unfortunately without true generics, I have to allow arbitrary Elements into the stkque. The alternatives to this are:
  1) Provide a stkque that allows only one type of Element (like int) which is very limiting.
  2) Hand-code stkques for every type: int, rune, byte, string, etc. Well, this is exactly the kind of meaningless effort that generics should solve. Also, I cannot predict every type that is ever going to come into existence!
*/
type Element interface{}
type Slice []Element

func New() Slice {
	return Slice{}
}

func (stk *Slice) Len() int {
	return len(*stk)
}

func (stk *Slice) Cap() int {
	return cap(*stk)
}

func (stk *Slice) Push(xs ...Element) {
	*stk = append(*stk, xs...)
}

type Empty Slice
func (empty Empty) Error() string {
	return "Cannot Peek() or Pop() on empty stkque"
}

func (stk *Slice) Peekstk() (Element, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return nil, Empty(*stk)
	}
	return (*stk)[ilast], nil
}

func (que *Slice) Peekque() (Element, error) {
	if len(*que) == 0 {
		return nil, Empty(*que)
	}
	return (*que)[0], nil
}

func (stk *Slice) Popstk() (Element, error) {
	ilast := len(*stk) - 1
	if ilast < 0 {
		return nil, Empty(*stk)
	}

	last := (*stk)[ilast] // save last value; it won't be available afterwards
	if ilast < cap(*stk) / 4 {
		*stk = append(make(Slice, 0, cap(*stk) / 2), (*stk)[:ilast]...)
	} else {
		*stk = (*stk)[:ilast]
	}
	return last, nil
}

func (que *Slice) Popque() (Element, error) {
	len := len(*que)
	if len == 0 {
		return nil, Empty(*que)
	}

	first := (*que)[0] // save first value; it won't be available afterwards
	if len > cap(*que) / 4 {
		*que = (*que)[1:]
	} else {
		*que = append(make(Slice, 0, cap(*que) / 2), (*que)[1:]...)
	}
	return first, nil
}