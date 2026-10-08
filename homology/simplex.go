package main

import (
	"fmt"
	"sort"
)

// Simplex represents a k-simplex as an ordered set of vertices.
// E.g. {0,1,2} is a 2-simplex (triangle).
type Simplex struct {
	Vertices []int
	Dim      int
}

// NewSimplex creates a normalized simplex (sorted vertices, no repetitions).
func NewSimplex(vertices ...int) Simplex {
	// copy and sort
	vs := make([]int, len(vertices))
	copy(vs, vertices)
	sort.Ints(vs)

	// remove duplicates
	uniq := vs[:0]
	for i, v := range vs {
		if i == 0 || v != vs[i-1] {
			uniq = append(uniq, v)
		}
	}

	return Simplex{Vertices: uniq, Dim: len(uniq) - 1}
}

// Key returns a canonical string to be used as a map key.
func (s Simplex) Key() string {
	return fmt.Sprint(s.Vertices)
}

// Faces returns the (k-1)-faces of the simplex with alternating signs (+/-).
// The i-th face is obtained by removing the i-th vertex, with sign (-1)^i.
func (s Simplex) Faces() []BoundaryTerm {
	faces := make([]BoundaryTerm, 0, len(s.Vertices))
	for i := range s.Vertices {
		faceVerts := make([]int, 0, len(s.Vertices)-1)
		faceVerts = append(faceVerts, s.Vertices[:i]...)
		faceVerts = append(faceVerts, s.Vertices[i+1:]...)

		sign := 1
		if i%2 == 1 {
			sign = -1
		}

		faces = append(faces, BoundaryTerm{
			Simplex: NewSimplex(faceVerts...),
			Coeff:   sign,
		})
	}
	return faces
}

// BoundaryTerm is a simplex with an integer coefficient.
type BoundaryTerm struct {
	Simplex Simplex
	Coeff   int
}

func (s Simplex) String() string {
	return fmt.Sprintf("%v", s.Vertices)
}
