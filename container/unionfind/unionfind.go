package unionfind

import (
	"sync"
)

// UnionFind implements Find, Union GetSet functions in thread-safe and find with path compression.
// The element in union is exists or not exists is guaranteed by user/caller
// TODO: add rank to  https://oi-wiki.org/ds/dsu/
type UnionFind interface {
	Add(i int)
	Find(i int) int
	Union(i, j int)
	GetSet(i int) []int
	DeleteSet(i int)
}

type uf struct {
	sync.RWMutex

	parent map[int]int
	next   map[int]int
}

func NewUnionFind() UnionFind {
	return &uf{
		parent: make(map[int]int),
		next:   make(map[int]int),
	}
}

// Add inserts a new element into the UnionFind structure that initially is in its own partition.
func (x *uf) Add(i int) {
	x.Lock()
	defer x.Unlock()

	x.parent[i] = i
	x.next[i] = i
}

// Find returns the root of the disjoint set in which i belongs, return 0 if i is being deleted
func (x *uf) Find(i int) int {
	x.Lock()
	defer x.Unlock()

	_, ok := x.parent[i]
	if !ok {
		return int(0)
	}

	root := x.find(i, false)
	return root
}

// Union lets j's root connect to i's
func (x *uf) Union(i, j int) {
	x.Lock()
	defer x.Unlock()

	pr := x.find(i, true)
	qr := x.find(j, true)

	if pr != qr {
		x.parent[qr] = pr
		// concatenate the two circulars into one
		x.next[pr], x.next[qr] = x.next[qr], x.next[pr]
	}
}

// GetSet returns all the members of a set with same parent,
// it guarantees the first element of returned slice is current argument.
func (x *uf) GetSet(i int) []int {
	x.RLock()
	defer x.RUnlock()

	return x.getSet(i)
}

func (x *uf) DeleteSet(i int) {
	x.Lock()
	defer x.Unlock()

	ids := x.getSet(i)
	for _, id := range ids {
		delete(x.parent, id)
		delete(x.next, id)
	}
}

func (x *uf) getSet(i int) []int {
	cluster := []int{i}
	// traversal the circular
	for r := i; x.next[i] != r; i = x.next[i] {
		cluster = append(cluster, x.next[i])
	}

	return cluster
}

func (x *uf) find(i int, compression bool) (r int) {
	// find the root by traversing up tree
	for r = i; x.parent[r] != r; r = x.parent[r] {
	}

	if !compression {
		return
	}

	for x.parent[i] != r {
		// set all elements along path to point directly to root
		i, x.parent[i] = x.parent[i], r
	}

	return
}
