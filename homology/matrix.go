package main

// MatrixInt é uma matriz com entradas inteiras.
type MatrixInt struct {
	Rows, Cols int
	Data       [][]int
}

func NewMatrixInt(rows, cols int) *MatrixInt {
	d := make([][]int, rows)
	for i := range d {
		d[i] = make([]int, cols)
	}
	return &MatrixInt{Rows: rows, Cols: cols, Data: d}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return abs(a)
}

// SmithNormalForm calcula a forma normal de Smith da matriz.
// Retorna a lista de divisores elementares (invariantes) não-nulos.
// Realiza operações de linha e coluna sobre ℤ.
func (m *MatrixInt) SmithNormalForm() []int {
	if m.Rows == 0 || m.Cols == 0 {
		return nil
	}

	t := 0 // índice do pivô atual
	for t < m.Rows && t < m.Cols {
		// Encontra pivô não-nulo de menor valor absoluto no sub-bloco.
		pivot := -1
		pivotRow, pivotCol := -1, -1
		best := 0
		for i := t; i < m.Rows; i++ {
			for j := t; j < m.Cols; j++ {
				v := abs(m.Data[i][j])
				if v != 0 && (best == 0 || v < best) {
					best = v
					pivotRow, pivotCol = i, j
				}
			}
		}
		if best == 0 {
			break // sub-bloco todo zero
		}
		_ = pivot

		// Move pivô para (t,t)
		m.Data[t], m.Data[pivotRow] = m.Data[pivotRow], m.Data[t]
		for i := 0; i < m.Rows; i++ {
			m.Data[i][t], m.Data[i][pivotCol] = m.Data[i][pivotCol], m.Data[i][t]
		}

		// Limpa linha e coluna usando operações elementares.
		done := false
		for !done {
			done = true
			p := m.Data[t][t]
			// Zera a coluna t
			for i := t + 1; i < m.Rows; i++ {
				if m.Data[i][t] != 0 {
					q := m.Data[i][t] / p
					for j := t; j < m.Cols; j++ {
						m.Data[i][j] -= q * m.Data[t][j]
					}
					if m.Data[i][t] != 0 {
						// resto não-zero: troca e continua
						m.Data[t], m.Data[i] = m.Data[i], m.Data[t]
						done = false
						break
					}
				}
			}
			if !done {
				continue
			}
			// Zera a linha t
			for j := t + 1; j < m.Cols; j++ {
				if m.Data[t][j] != 0 {
					q := m.Data[t][j] / m.Data[t][t]
					for i := t; i < m.Rows; i++ {
						m.Data[i][j] -= q * m.Data[i][t]
					}
					if m.Data[t][j] != 0 {
						for i := 0; i < m.Rows; i++ {
							m.Data[i][t], m.Data[i][j] = m.Data[i][j], m.Data[i][t]
						}
						done = false
						break
					}
				}
			}
		}

		// Garante divisibilidade: se algum elemento do sub-bloco não é
		// divisível pelo pivô, soma a linha correspondente à linha t.
		p := m.Data[t][t]
		divisible := true
		for i := t + 1; i < m.Rows && divisible; i++ {
			for j := t + 1; j < m.Cols; j++ {
				if m.Data[i][j]%p != 0 {
					for k := t; k < m.Cols; k++ {
						m.Data[t][k] += m.Data[i][k]
					}
					divisible = false
					done = false
					break
				}
			}
		}
		if !divisible {
			continue
		}

		t++
	}

	// Coleta invariantes (diagonal)
	var invariants []int
	for i := 0; i < m.Rows && i < m.Cols; i++ {
		if m.Data[i][i] != 0 {
			invariants = append(invariants, abs(m.Data[i][i]))
		}
	}
	return invariants
}
