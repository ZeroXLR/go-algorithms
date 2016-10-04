package unionfind

type UFcore struct {
	setcount int
	parent []int
}

func (uf *UFcore) Count() int {
	return uf.setcount
}

type ElementOutOfBounds string
func (eoob ElementOutOfBounds) Error() string {
	return string(eoob)
}

func (uf *UFcore) Find(p int) (int, error) {
	if p < 0 || p >= len(uf.parent) {
		return 0, ElementOutOfBounds("Element is Out of Bounds")
	}
	for p != uf.parent[p] {
		uf.parent[p], p = uf.parent[uf.parent[p]], uf.parent[p]
	}
	return p, nil
}

func (uf *UFcore) Connected(p0, p1 int) (bool, error) {
	root0, err := uf.Find(p0)
	var root1 int; if err == nil {
		root1, err = uf.Find(p1)
	}
	return err == nil && root0 == root1, err
}

func Core(n int) *UFcore {
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &UFcore{n, parent}
}

type UFbyrank struct {
	*UFcore
	rank []byte
}

func (uf *UFbyrank) Union(p0, p1 int) error {
	root0, err := uf.Find(p0)
	if err == nil {
		root1, err := uf.Find(p1)
		if err == nil && root0 != root1 {
			if w0, w1 := uf.rank[root0], uf.rank[root1]; w1 > w0 {
				uf.parent[root0] = root1
			} else if w0 > w1 {
				uf.parent[root1] = root0
			} else {
				uf.parent[root1] = root0
				uf.rank[root0]++
			}
			uf.setcount--
		}
	}
	return err
}

type UFbysize struct {
	*UFcore
	size []int
}

func (uf *UFbysize) Union(p0, p1 int) error {
	root0, err := uf.Find(p0)
	if err == nil {
		root1, err := uf.Find(p1)
		if err == nil && root0 != root1 {
			if w0, w1 := uf.size[root0], uf.size[root1]; w1 > w0 {
				uf.parent[root0] = root1
				uf.size[root1] += w0
			} else {
				uf.parent[root1] = root0
				uf.size[root0] += w1
			}
			uf.setcount--
		}
	}
	return err
}

func Byrank(n int) *UFbyrank {
	return &UFbyrank{Core(n), make([]byte, n)}
}

func New(n int) *UFbyrank {
	return Byrank(n)
}

func Bysize(n int) *UFbysize {
	size := make([]int, n)
	for i := 0; i < n; i++ {
		size[i] = 1
	}
	return &UFbysize{Core(n), size}
}