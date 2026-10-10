package main

import "fmt"

// SimplicialComplex is a collection of simplices closed under taking faces.
type SimplicialComplex struct {
	// Simplices grouped by dimension
	byDim map[int]map[string]Simplex
	// Canonical order of each dimension (used to index matrices)
	order  map[int][]Simplex
	index  map[int]map[string]int
	maxDim int
}

func NewComplex() *SimplicialComplex {
	return &SimplicialComplex{
		byDim:  make(map[int]map[string]Simplex),
		order:  make(map[int][]Simplex),
		index:  make(map[int]map[string]int),
		maxDim: -1,
	}
}

// Add inserts a simplex and all of its faces (closure).
func (c *SimplicialComplex) Add(s Simplex) {
	if _, ok := c.byDim[s.Dim][s.Key()]; ok {
		return
	}
	if c.byDim[s.Dim] == nil {
		c.byDim[s.Dim] = make(map[string]Simplex)
		c.index[s.Dim] = make(map[string]int)
	}

	c.byDim[s.Dim][s.Key()] = s
	c.index[s.Dim][s.Key()] = len(c.order[s.Dim])
	c.order[s.Dim] = append(c.order[s.Dim], s)

	if s.Dim > c.maxDim {
		c.maxDim = s.Dim
	}

	// recursively add the faces
	for i := range s.Vertices {
		faceVerts := append([]int{}, s.Vertices[:i]...)
		faceVerts = append(faceVerts, s.Vertices[i+1:]...)
		c.Add(NewSimplex(faceVerts...))
	}
}

// ChainGroup returns the simplices of dimension k in order.
func (c *SimplicialComplex) ChainGroup(k int) []Simplex {
	return c.order[k]
}

// Index returns the index of a simplex in the basis of its chain group.
func (c *SimplicialComplex) Index(s Simplex) int {
	m, ok := c.index[s.Dim]
	if !ok {
		return -1
	}
	i, ok := m[s.Key()]
	if !ok {
		return -1
	}
	return i
}

func (c *SimplicialComplex) MaxDim() int { return c.maxDim }

func (c *SimplicialComplex) String() string {
	s := ""
	for k := 0; k <= c.maxDim; k++ {
		s += fmt.Sprintf("Dimension %d: %v\n", k, c.order[k])
	}
	return s
}
