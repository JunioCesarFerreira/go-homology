package main

import "fmt"

func main() {
	S1()    // Circle S¹ → H_0=Z, H_1=Z
	D2()    // Solid disk → H_0=Z
	S2()    // Sphere S² → H_0=Z, H_2=Z
	T2()    // Torus → H_0=Z, H_1=Z², H_2=Z
	RP2()   // Projective plane RP² → H_0=Z, H_1=Z/2
	Klein() // Klein bottle → H_0=Z, H_1=Z⊕Z/2
}

func S1() {
	fmt.Println("=== Boundary of a triangle (S¹) ===")
	c := NewComplex()
	c.Add(NewSimplex(0, 1))
	c.Add(NewSimplex(1, 2))
	c.Add(NewSimplex(2, 0))
	mostrarHomologia(c)
}

func D2() {
	fmt.Println("=== Filled triangle (disk D²) ===")
	c := NewComplex()
	c.Add(NewSimplex(0, 1, 2))
	mostrarHomologia(c)
}

func S2() {
	fmt.Println("=== Boundary of a tetrahedron (sphere S²) ===")
	c := NewComplex()
	// the 4 faces of the tetrahedron
	c.Add(NewSimplex(0, 1, 2))
	c.Add(NewSimplex(0, 1, 3))
	c.Add(NewSimplex(0, 2, 3))
	c.Add(NewSimplex(1, 2, 3))
	mostrarHomologia(c)
}

func T2() {
	fmt.Println("=== Torus T² (minimal 7-vertex triangulation) ===")
	c := NewComplex()
	// Minimal 7-vertex triangulation of the torus (Möbius torus, realized by
	// the Császár polyhedron): triangles {i, i+1, i+3} and {i, i+2, i+3} mod 7.
	for i := 0; i < 7; i++ {
		c.Add(NewSimplex(i, (i+1)%7, (i+3)%7))
		c.Add(NewSimplex(i, (i+2)%7, (i+3)%7))
	}
	mostrarHomologia(c)
}

func RP2() {
	fmt.Println("=== Real projective plane RP² ===")
	c := NewComplex()
	// RP² = quotient of the sphere by the antipodal map.
	// 6-vertex triangulation (hemi-icosahedron model).
	faces := [][3]int{
		{0, 1, 2}, {0, 1, 5}, {0, 2, 3}, {0, 3, 4}, {0, 4, 5},
		{1, 2, 4}, {1, 3, 4}, {1, 3, 5}, {2, 3, 5}, {2, 4, 5},
	}
	for _, f := range faces {
		c.Add(NewSimplex(f[0], f[1], f[2]))
	}
	mostrarHomologia(c)
}

func Klein() {
	fmt.Println("=== Klein bottle (9-vertex triangulation) ===")
	c := NewComplex()
	// The Klein bottle is the square [0,3]×[0,3] with the bottom and top sides
	// glued directly, (i, 3) ~ (i, 0), and the left and right sides glued with
	// a flip, (3, j) ~ (0, 3-j). Splitting each unit square of the 3×3 grid
	// into two triangles gives a 9-vertex triangulation.
	// v maps the grid point (i, j), with 0 ≤ i, j ≤ 3, to its vertex label.
	v := func(i, j int) int {
		if i == 3 {
			i, j = 0, 3-j
		}
		return 3*i + j%3
	}
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			c.Add(NewSimplex(v(i, j), v(i+1, j), v(i+1, j+1)))
			c.Add(NewSimplex(v(i, j), v(i, j+1), v(i+1, j+1)))
		}
	}
	mostrarHomologia(c)
}

func mostrarHomologia(c *SimplicialComplex) {
	fmt.Println("Complex:")
	fmt.Print(c.String())

	hom := ComputeHomology(c)
	fmt.Println("Homology groups:")
	betti := make([]int, len(hom))
	for k, h := range hom {
		fmt.Printf("  H_%d = %s\n", h.Dim, h)
		betti[k] = h.FreeRank
	}
	// β_k = rank H_k, indexed by k
	fmt.Println("Betti numbers:", betti)
	fmt.Println()
}
