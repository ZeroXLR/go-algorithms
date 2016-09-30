package stack

import (
	"testing"
	"math/rand"
)

func TestLen(t *testing.T) {
	stk := New()

	/* Newly created stack must have zero length */
	if stk.Len() != 0 {
		t.Errorf("stk.Len() == %d, want %d: New stack must be empty", stk.Len(), 0)
	}

	/* Push a random number of times */
	randvals := make([]Element, rand.Intn(100))
	stk.Push(randvals...)
	/* Length must equal number of times pushed */
	if len, checklen := stk.Len(), len(randvals); len != checklen {
		t.Errorf("stk.Len() == %d, want %d", len, checklen)
	}
}

func TestPush(t *testing.T) {
	stk := New()

	/* Push a random number of times */
	randvals := make([]Element, rand.Intn(100))
	for i := 0; i < len(randvals); i++ {
		randvals[i] = float64(i)
	}
	stk.Push(randvals...)
	/* Stack must be populated by only the pushed values */
	for i, v := range stk {
		if v != float64(i) {
			t.Errorf("stk[%d] == %f, want %f", i, v, float64(i))
			break
		}
	}
}

func TestPeek(t *testing.T) {
 	stk := New()

	/* Peeking an empty array should throw a non-nil error*/
	v, err := stk.Peek()
	if err == nil || v != nil {
		t.Errorf("stk.Peek() == %v, %v, want %v, !%v: Peeking an empty array should throw non-nil error", v, err, nil, nil)
	}

	/* Push various values and check Peek */
	stk.Push(byte(42))
	v, err = stk.Peek()
	if !(v == byte(42) && err == nil) {
		t.Errorf("stk.Peek() == %d, %v want %d, %v", v, err, byte(42), nil)
	}
	stk.Push("gopher")
	v, err = stk.Peek()
	if !(v == "gopher" && err == nil) {
		t.Errorf("stk.Peek() == %s, %v want %s, %v", v, err, "gopher", nil)
	}
	stk.Push(false)
	v, err = stk.Peek()
	if !(v == false && err == nil) {
		t.Errorf("stk.Peek() == %t, %v want %t, %v", v, err, true, nil)
	}
}

func TestPop(t *testing.T) {
	stk := New()

	/* Popping an empty array should throw a non-nil error*/
	v, err := stk.Pop()
	if err == nil || v != nil {
		t.Errorf("stk.Pop() == %v, %v, want %v, !%v: Popping an empty array should throw non-nil error", v, err, nil, nil)
	}

	/* Push some power of 2 times */
	vals := make([]Element, 128)
	for i := 0; i < 128; i++ {
		vals[i] = i
	}
	stk.Push(vals...)

	/* Start Popping; check return values AND len(stk) */
	for i := 127; i > -1; i-- {
		v, err = stk.Pop()
		if !(v == i && err == nil && len(stk) == i) {
			t.Errorf("stk.Pop() == %d, %v and len(stk) == %d want %d, %v and %d", v, err, i, i, nil, i)
			break
		}
	}
}

func BenchmarkPushOneByOne(b *testing.B) {
	stk := New()
	for i := 0; i < b.N; i++ {
		stk.Push(i)
	}
}

func BenchmarkPushSimultaneously(b *testing.B) {
	stk, vals := New(), make([]Element, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = i
	}
	stk.Push(vals...)
}