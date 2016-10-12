package staque

import (
	"testing"
	"math/rand"
)

func TestPush(t *testing.T) {
	slc := NewGeneric()

	/* Push a random number of times */
	randvals := make([]Generic, rand.Intn(100))
	for i := 0; i < len(randvals); i++ {
		randvals[i] = float64(i)
	}
	slc = slc.Push(randvals...)
	/* Stack must be populated by only the pushed values */
	for i, v := range slc {
		if v != float64(i) {
			t.Errorf("slc[%d] == %f, want %f", i, v, float64(i))
			break
		}
	}
}

func TestPeek(t *testing.T) {
 	slc := NewGeneric()

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
	vals := [4]Generic{byte(42), '界', "gopher", false}
	for _, val := range vals {
		slc = slc.Push(val)
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
	slc := NewGeneric()

	/* Popstking an empty array should throw a non-nil error*/
	slc, v, err := slc.Popstk()
	if err == nil || v != nil {
		t.Errorf("slc.Popstk() == %v, %v, want %v, !%v: Popstking an empty array should throw non-nil error", v, err, nil, nil)
	}

	/* Push some power of 2 times */
	vals := make([]Generic, 128)
	for i := 0; i < 128; i++ {
		vals[i] = i
	}

	slc = slc.Push(vals...)
	/* Start Popstking; check return values AND len(slc) */
	for i := 127; i > -1; i-- {
		slc, v, err = slc.Popstk()
		if !(v == i && err == nil && len(slc) == i) {
			t.Errorf("slc.Popstk() == %d, %v and len(slc) == %d, want %d, %v and %d", v, err, len(slc), i, nil, i)
			break
		}
	}

	slc = slc.Push(vals...)
	/* Start Popqueing; check return values AND len(slc) */
	for i, j := 0, 127; i < 128; i, j = i + 1, j - 1 {
		slc, v, err = slc.Popque()
		if !(v == i && err == nil && len(slc) == j) {
			t.Errorf("slc.Popque() == %d, %v and len(slc) == %d, want %d, %v and %d", v, err, len(slc), i, nil, j)
			break
		}
	}
}

func BenchmarkPush1By1(b *testing.B) {
	slc := NewGeneric()
	for i := 0; i < b.N; i++ {
		slc = slc.Push(i)
	}
}

func BenchmarkPushAll(b *testing.B) {
	slc, vals := NewGeneric(), make([]Generic, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = i
	}
	slc = slc.Push(vals...)
}

func BenchmarkPush1By1String(b *testing.B) {
	slc := NewString()
	for i := 0; i < b.N; i++ {
		slc = slc.Push("Meow Meow Meow Meow")
	}
}

func BenchmarkPushAllString(b *testing.B) {
	slc, vals := NewString(), make([]string, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = "Meow Meow Meow Meow"
	}
	slc = slc.Push(vals...)
}

func BenchmarkPush1By1Int(b *testing.B) {
	slc := NewInt()
	for i := 0; i < b.N; i++ {
		slc = slc.Push(i)
	}
}

func BenchmarkPushAllInt(b *testing.B) {
	slc, vals := NewInt(), make([]int, b.N)
	for i := 0; i < b.N; i++ {
		vals[i] = i
	}
	slc = slc.Push(vals...)
}