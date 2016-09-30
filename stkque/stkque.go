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
	lastindex := len(*stk) - 1
	if lastindex < 0 {
		return nil, Empty(*stk)
	}
	return (*stk)[lastindex], nil
}

func (que *Slice) Peekque() (Element, error) {
	if len(*que) == 0 {
		return nil, Empty(*que)
	}
	return (*que)[0], nil
}

func shrink(slc Slice, fromincl, toexcl int) Slice {
	/*
	  As per Dr. Sedgewick's suggestion, we shrink dynamic arrays to half only
	  when their lens reach a quarter of their caps. This prevents thrashing (
	  assuming append doubles the cap on resizing)
	*/
	if len(slc) <= cap(slc) / 4 {
		newslc := make(Slice, cap(slc) / 2)
		copy(newslc, slc[fromincl:toexcl])
		return newslc[fromincl:toexcl]
	}
	return slc[fromincl:toexcl]
}

func (stk *Slice) Popstk() (Element, error) {
	lastindex := len(*stk) - 1
	if lastindex < 0 {
		return nil, Empty(*stk)
	}

	last := (*stk)[lastindex] // save last value; it won't be available after shrink
	*stk = shrink(*stk, 0, lastindex)
	return last, nil
}

func (que *Slice) Popque() (Element, error) {
	len := len(*que)
	if len == 0 {
		return nil, Empty(*que)
	}

	first := (*que)[0] // save first value; it won't be available after shrink
	*que = shrink(*que, 1, len)
	return first, nil
}