package main

import "fmt"

func main() {
	S1()  // Círculo S¹ → H_0=Z, H_1=Z
	D2()  // Disco sólido → H_0=Z
	S2()  // Esfera S² → H_0=Z, H_2=Z
	T2()  // Toro → H_0=Z, H_1=Z², H_2=Z
	RP2() // Plano projetivo RP² → H_0=Z, H_1=Z/2
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
	// 4 faces do tetraedro
	c.Add(NewSimplex(0, 1, 2))
	c.Add(NewSimplex(0, 1, 3))
	c.Add(NewSimplex(0, 2, 3))
	c.Add(NewSimplex(1, 2, 3))
	mostrarHomologia(c)
}

func T2() {
	fmt.Println("=== Toro T² (triangulação mínima com 7 vértices) ===")
	c := NewComplex()
	// Triangulação do toro com 7 vértices (Császár / Möbius–Kantor)
	faces := [][3]int{
		{0, 1, 2}, {0, 2, 3}, {0, 3, 4}, {0, 4, 5}, {0, 5, 6}, {0, 6, 1},
		{1, 3, 6}, {1, 3, 4}, {1, 2, 4}, {2, 4, 6}, {2, 5, 6}, {2, 3, 5},
		{3, 5, 6}, {1, 5, 4},
	}
	for _, f := range faces {
		c.Add(NewSimplex(f[0], f[1], f[2]))
	}
	mostrarHomologia(c)
}

func RP2() {
	fmt.Println("=== Plano projetivo real RP² ===")
	c := NewComplex()
	// RP² = quociente da esfera por antipodal.
	// Triangulação com 6 vértices (modelo icosaédrico reduzido).
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
