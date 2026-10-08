package main

import "fmt"

// HomologyGroup describes H_k = ℤ^freeRank ⊕ (⊕ ℤ/torsion[i])
type HomologyGroup struct {
	Dim      int
	FreeRank int
	Torsion  []int
}

func (h HomologyGroup) String() string {
	if h.FreeRank == 0 && len(h.Torsion) == 0 {
		return "0"
	}
	s := ""
	if h.FreeRank > 0 {
		if h.FreeRank == 1 {
			s += "Z"
		} else {
			s += fmt.Sprintf("Z^%d", h.FreeRank)
		}
	}
	for _, t := range h.Torsion {
		if t <= 1 {
			continue
		}
		if s != "" {
			s += " + "
		}
		s += fmt.Sprintf("Z/%d", t)
	}
	return s
}

// BoundaryMatrix builds the matrix of ∂_k : C_k → C_{k-1}
// with respect to the ordered bases of the chain groups.
func BoundaryMatrix(c *SimplicialComplex, k int) *MatrixInt {
	ck := c.ChainGroup(k)
	ckm1 := c.ChainGroup(k - 1)

	m := NewMatrixInt(len(ckm1), len(ck))

	for j, s := range ck {
		for _, term := range s.Faces() {
			i := c.Index(term.Simplex)
			if i < 0 {
				panic(fmt.Sprintf("face %v não encontrada no complexo", term.Simplex))
			}
			m.Data[i][j] += term.Coeff
		}
	}
	return m
}

// ComputeHomology computes H_k for k = 0..maxDim.
func ComputeHomology(c *SimplicialComplex) []HomologyGroup {
	maxDim := c.MaxDim()
	result := make([]HomologyGroup, maxDim+1)

	for k := 0; k <= maxDim; k++ {
		var rankK, rankKm1 int
		var torsion []int

		// rank of the chain group C_k
		rankK = len(c.ChainGroup(k))

		// ∂_{k+1}: C_{k+1} → C_k
		var rankBoundaryNext int
		if k+1 <= maxDim {
			bNext := BoundaryMatrix(c, k+1)
			inv := bNext.SmithNormalForm()
			rankBoundaryNext = len(inv)
		}

		// ∂_k: C_k → C_{k-1}
		if k > 0 {
			bk := BoundaryMatrix(c, k)
			inv := bk.SmithNormalForm()
			rankKm1 = len(inv)
			for _, v := range inv {
				if v > 1 {
					torsion = append(torsion, v)
				}
			}
		} else {
			rankKm1 = 0
		}

		// H_k = ker(∂_k) / im(∂_{k+1})
		// rank(ker ∂_k) = dim C_k - rank(∂_k)
		kerRank := rankK - rankKm1
		freeRank := kerRank - rankBoundaryNext
		if freeRank < 0 {
			freeRank = 0
		}

		result[k] = HomologyGroup{
			Dim:      k,
			FreeRank: freeRank,
			Torsion:  torsion,
		}
	}
	return result
}
