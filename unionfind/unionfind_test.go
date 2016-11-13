package unionfind

import (
	"testing"
	"math/rand"
)

func TestUFcore(t *testing.T) {
	n := rand.Intn(100)
	uf := core(n)

	if uf.setcount != n {
		t.Errorf("Initially there should be %d components but got uf.setcount == %d", n, uf.setcount)
	}

	for i := 0; i < n; i++ {
		if uf.parent[i] != i {
			t.Errorf("Initially every point is its own parent; so expected uf.parent[i] == i but got uf.parent[%d] == %d", i, uf.parent[i])
		}
	}
}

func TestUFbyrank(t *testing.T) {
	n := rand.Intn(100)
	uf := Byrank(n)

	for i := 0; i < n; i++ {
		if uf.rank[i] != 0 {
			t.Errorf("Initial rank of components should be 0; but got uf.rank[%d] == %d", i, uf.rank[i])
		}
	}
}

func TestUFbysize(t *testing.T) {
	n := rand.Intn(100)
	uf := Bysize(n)

	for i := 0; i < n; i++ {
		if uf.size[i] != 1 {
			t.Errorf("Initial size of components should be 1; but got uf.size[%d] == %d", i, uf.size[i])
		}
	}
}