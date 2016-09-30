package stack

/*
  Unfortunately without true generics, I have to allow arbitrary elements into
  the stack. The alternatives to this are:
  1) Provide a stack that allows only one type of element (like int) which is
     very limiting.
  2) Hand-code stacks for every type: int, rune, byte, string, etc. Well, this
     is exactly the kind of meaningless effort that generics should solve. Also,
     I cannot predict every type that is ever going to come into existence!
*/
type element interface{}
type SliceStack []element

func New() SliceStack {
	return SliceStack{}
}

func (stk *SliceStack) Len() int {
	return len(*stk)
}

func (stk *SliceStack) Cap() int {
	return cap(*stk)
}

func (stk *SliceStack) Push(e element) {
	*stk = append(*stk, e)
}

type EmptyStack SliceStack
func (empty EmptyStack) Error() string {
	return "Cannot Peek() or Pop() on empty stack"
}

func (stk *SliceStack) Peek() (element, error) {
	lastindex := len(*stk) - 1
	if lastindex < 0 {
		return nil, EmptyStack(*stk)
	}
	return (*stk)[lastindex], nil
}

func shrink(stk SliceStack, upto int) SliceStack {
	/*
	  As per Dr. Sedgewick's suggestion, we shrink dynamic arrays to half only
	  when their lens reach a quarter of their caps. This prevents thrashing (
	  assuming append doubles the cap upon resizing)
	*/
	if len(stk) <= cap(stk) / 4 {
		newstk := make(SliceStack, cap(stk) / 2)
		copy(newstk, stk[:upto])
		return newstk[:upto]
	}
	return stk[:upto]
}

func (stk *SliceStack) Pop() (element, error) {
	lastindex := len(*stk) - 1
	if lastindex < 0 {
		return nil, EmptyStack(*stk)
	}

	last := (*stk)[lastindex] // save last value; it won't be available after shrink
	*stk = shrink(*stk, lastindex)
	return last, nil
}