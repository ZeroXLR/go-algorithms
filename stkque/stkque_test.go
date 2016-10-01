package stkque

import (
	"testing"
	"math/rand"
)

func TestLen(t *testing.T) {
	slc := New()

	/* Newly created stkque must have zero length */
	if slc.Len() != 0 {
		t.Errorf("slc.Len() == %d, want %d: New stkque must be empty", slc.Len(), 0)
	}

	/* Push a random number of times */
	randvals := make([]Element, rand.Intn(100))
	slc.Push(randvals...)
	/* Length must equal number of times pushed */
	if len, checklen := slc.Len(), len(randvals); len != checklen {
		t.Errorf("slc.Len() == %d, want %d", len, checklen)
	}
}

func TestPush(t *testing.T) {
	slc := New()

	/* Push a random number of times */
	randvals := make([]Element, rand.Intn(100))
	for i := 0; i < len(randvals); i++ {
		randvals[i] = float64(i)
	}
	slc.Push(randvals...)
	/* Stack must be populated by only the pushed values */
	for i, v := range slc {
		if v != float64(i) {
			t.Errorf("slc[%d] == %f, want %f", i, v, float64(i))
			break
		}
	}
}

func TestPeek(t *testing.T) {
 	slc := New()

	/* Peekstking an empty array should throw a non-nil error*/
	v, err := slc.Peekstk()
	if err == nil || v != nil {
		t.Errorf("slc.Peekstk() == %v, %v, want %v, !%v: Peekstking an empty array should throw non-nil error", v, err, nil, nil)
	}

	/* Peekqueing an empty array should throw a non-nil error*/
	v, err = slc.Peekque()
	if err == nil || v != nil {
		t.Errorf("slc.Peekque() == %v, %v, want %v, !%v: Peekqueing an empty array should throw non-nil error", v, err, nil, nil)
	}

	/* Push various values and check both Peeks */
	vals := [4]Element{byte(42), '界', "gopher", false}
	for _, val := range vals {
		slc.Push(val)
		v, err = slc.Peekstk()
		if !(v == val && err == nil) {
			t.Errorf("slc.Peekstk() == %d, %v want %d, %v", v, err, val, nil)
		}
		v, err = slc.Peekque()
		if !(v == vals[0] && err == nil) {
			t.Errorf("slc.Peekque() == %d, %v want %d, %v", v, err, vals[0], nil)
		}
	}
}

func TestPop(t *testing.T) {
	slc := New()

	/* Popstking an empty array should throw a non-nil error*/
	v, err := slc.Popstk()
	if err == nil || v != nil {
		t.Errorf("slc.Popstk() == %v, %v, want %v, !%v: Popstking an empty array should throw non-nil error", v, err, nil, nil)
	}

	/* Push some power of 2 times */
	vals := make([]Element, 128)
	for i := 0; i < 128; i++ {
		vals[i] = i
	}

	slc.Push(vals...)
	/* Start Popstking; check return values AND len(slc) */
	for i := 127; i > -1; i-- {
		v, err = slc.Popstk()
		if !(v == i && err == nil && len(slc) == i) {
			t.Errorf("slc.Popstk() == %d, %v and len(slc) == %d, want %d, %v and %d", v, err, len(slc), i, nil, i)
			break
		}
	}

	slc.Push(vals...)
	/* Start Popqueing; check return values AND len(slc) */
	for i, j := 0, 127; i < 128; i, j = i + 1, j - 1 {
		v, err = slc.Popque()
		if !(v == i && err == nil && len(slc) == j) {
			t.Errorf("slc.Popque() == %d, %v and len(slc) == %d, want %d, %v and %d", v, err, len(slc), i, nil, j)
			break
		}
	}
}

func BenchmarkPushOneByOne(b *testing.B) {
	slc := New()
	for i := 0; i < b.N; i++ {
		slc.Push(i)
	}
}

func BenchmarkPushSimultaneously(b *testing.B) {
	slc, vals := New(), make([]Element, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = i
	}
	slc.Push(vals...)
}