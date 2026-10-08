package main

import (
	"fmt"
	"sort"
)

// Simplex representa um k-simplex como um conjunto ordenado de vértices.
// Ex: {0,1,2} é um 2-simplex (triângulo).
type Simplex struct {
	Vertices []int
	Dim      int
}

// NewSimplex cria um simplex normalizado (vértices ordenados, sem repetição).
func NewSimplex(vertices ...int) Simplex {
	// copia e ordena
	vs := make([]int, len(vertices))
	copy(vs, vertices)
	sort.Ints(vs)

	// remove duplicatas
	uniq := vs[:0]
	for i, v := range vs {
		if i == 0 || v != vs[i-1] {
			uniq = append(uniq, v)
		}
	}

	return Simplex{Vertices: uniq, Dim: len(uniq) - 1}
}

// Key retorna uma string canônica para usar como chave de mapa.
func (s Simplex) Key() string {
	return fmt.Sprint(s.Vertices)
}

// Faces retorna as (k-1)-faces do simplex, com sinal (+/-) alternado.
// A i-ésima face é obtida removendo o i-ésimo vértice, com sinal (-1)^i.
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

// BoundaryTerm é um simplex com coeficiente inteiro.
type BoundaryTerm struct {
	Simplex Simplex
	Coeff   int
}

func (s Simplex) String() string {
	return fmt.Sprintf("%v", s.Vertices)
}
