package unionfind

import "errors"

type ufcore struct {
	setcount int
	parent []int
}

func (uf ufcore) Count() int {
	return uf.setcount
}

var eoob = errors.New("Element is Out of Bounds")

func (uf ufcore) Find(p int) (int, error) {
	if p < 0 || p >= len(uf.parent) {
		return 0, eoob
	}
	root := p
	for root != uf.parent[root] {
		root = uf.parent[root]
	}
	for p != root {
		uf.parent[p], p = root, uf.parent[p]
	}
	return root, nil
}

func (uf ufcore) Connected(p0, p1 int) (bool, error) {
	root0, err := uf.Find(p0)
	var root1 int; if err == nil {
		root1, err = uf.Find(p1)
	}
	return err == nil && root0 == root1, err
}

func core(n int) *ufcore {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &ufcore{n, parent}
}

type UFbyrank struct {
	*ufcore
	rank []byte
}

func (uf UFbyrank) Union(p0, p1 int) (err error) {
	root0, err := uf.Find(p0)
	if err != nil {
		return
	}
	root1, err := uf.Find(p1)
	if err != nil {
		return
	}
	if root0 == root1 {
		return
	}
	if w0, w1 := uf.rank[root0], uf.rank[root1]; w1 > w0 {
		uf.parent[root0] = root1
	} else if w0 > w1 {
		uf.parent[root1] = root0
	} else {
		uf.parent[root1] = root0
		uf.rank[root0]++
	}
	uf.setcount--
	return
}

type UFbysize struct {
	*ufcore
	size []int
}

func (uf UFbysize) Union(p0, p1 int) (err error) {
	root0, err := uf.Find(p0)
	if err != nil {
		return
	}
	root1, err := uf.Find(p1)
	if err != nil {
		return
	}
	if root0 == root1 {
		return
	}
	if w0, w1 := uf.size[root0], uf.size[root1]; w1 > w0 {
		uf.parent[root0] = root1
		uf.size[root1] += w0
	} else {
		uf.parent[root1] = root0
		uf.size[root0] += w1
	}
	uf.setcount--
	return
}

func Byrank(n int) UFbyrank {
	return UFbyrank{core(n), make([]byte, n)}
}

func New(n int) UFbyrank {
	return Byrank(n)
}

func Bysize(n int) UFbysize {
	size := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1
	}
	return UFbysize{core(n), size}
}