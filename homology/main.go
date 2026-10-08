package main

import "fmt"

func main() {
	S1()  // Circle S¹ → H_0=Z, H_1=Z
	D2()  // Solid disk → H_0=Z
	S2()  // Sphere S² → H_0=Z, H_2=Z
	T2()  // Torus → H_0=Z, H_1=Z², H_2=Z
	RP2() // Projective plane RP² → H_0=Z, H_1=Z/2
}

func S1() {
	fmt.Println("=== Borda de um triângulo (S¹) ===")
	c := NewComplex()
	c.Add(NewSimplex(0, 1))
	c.Add(NewSimplex(1, 2))
	c.Add(NewSimplex(2, 0))
	mostrarHomologia(c)
}

func D2() {
	fmt.Println("=== Triângulo cheio (disco D²) ===")
	c := NewComplex()
	c.Add(NewSimplex(0, 1, 2))
	mostrarHomologia(c)
}

func S2() {
	fmt.Println("=== Borda de um tetraedro (esfera S²) ===")
	c := NewComplex()
	// the 4 faces of the tetrahedron
	c.Add(NewSimplex(0, 1, 2))
	c.Add(NewSimplex(0, 1, 3))
	c.Add(NewSimplex(0, 2, 3))
	c.Add(NewSimplex(1, 2, 3))
	mostrarHomologia(c)
}

func T2() {
	fmt.Println("=== Toro T² (triangulação mínima com 7 vértices) ===")
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
	fmt.Println("=== Plano projetivo real RP² ===")
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

func mostrarHomologia(c *SimplicialComplex) {
	fmt.Println("Complexo:")
	fmt.Print(c.String())

	hom := ComputeHomology(c)
	fmt.Println("Grupos de homologia:")
	for _, h := range hom {
		fmt.Printf("  H_%d = %s\n", h.Dim, h)
	}
	fmt.Println()
}
