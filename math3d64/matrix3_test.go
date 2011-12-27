/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d64

import "testing"
import "rand"
import "time"

func TestMakeMatrix3(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}

		m1 := MakeMatrix3(true, data...)
		m2 := MakeMatrix3(false, data...)

		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				if m1.At(r, c) != data[r*size+c] {
					t.Errorf("m[%d][%d] (%f) != data[%d] (%f)", r, c, m1.At(r, c), r*size+c, data[r*size+c])
				}
				if m2.At(r, c) != data[c*size+r] {
					t.Errorf("m[%d][%d] (%f) != data[%d] (%f)", r, c, m1.At(r, c), r*size+c, data[r*size+c])
				}
			}
		}
	}
}

func TestMatrix3_ZeroThis(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}

		m1 := MakeMatrix3(true, data...)
		m1.ZeroThis()

		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				if m1.At(r, c) != 0 {
					t.Errorf("m[%d][%d] (%f) != 0", r, c, m1.At(r, c))
				}
			}
		}
	}
}

func TestMakeMatrix3Identity(t *testing.T) {
	const size = 3
	for iterations := 0; iterations < 1000; iterations++ {
		m1 := MakeMatrix3Identity()

		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				if r == c {
					if m1.At(r, c) != 1 {
						t.Errorf("m[%d][%d] (%f) != 0", r, c, m1.At(r, c))
					}
				} else if m1.At(r, c) != 0 {
					t.Errorf("m[%d][%d] (%f) != 0", r, c, m1.At(r, c))
				}
			}
		}
	}
}

func TestMatrix3_IdentityThis(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}

		m1 := MakeMatrix3(true, data...)
		m1.IdentityThis()

		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				if r == c {
					if m1.At(r, c) != 1 {
						t.Errorf("m[%d][%d] (%f) != 0", r, c, m1.At(r, c))
					}
				} else if m1.At(r, c) != 0 {
					t.Errorf("m[%d][%d] (%f) != 0", r, c, m1.At(r, c))
				}
			}
		}
	}
}

func TestMatrix3_GetRow(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}

		m1 := MakeMatrix3(true, data...)

		for r := 0; r < size; r++ {
			row := m1.GetRow(r)
			for c := 0; c < size; c++ {
				if m1.At(r, c) != row[c] {
					t.Errorf("m[%d][%d] (%f) != row[%d] (%f)", r, c, m1.At(r, c), c, row[c])
				}
			}
		}
	}
}

func TestMatrix3_GetCol(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}

		m1 := MakeMatrix3(true, data...)

		for c := 0; c < size; c++ {
			col := m1.GetCol(c)
			for r := 0; r < size; r++ {
				if m1.At(r, c) != col[r] {
					t.Errorf("m[%d][%d] (%f) != col[%d] (%f)", r, c, m1.At(r, c), r, col[r])
				}
			}
		}
	}
}

func TestMatrix3_Determinant(t *testing.T) {
	// TODO: Get a bunch of test data to do this
}

func TestMatrix3_Cofactor(t *testing.T) {
	// TODO: Get a bunch of test data to do this
}

func TestMatrix3_Equals(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data1 := make([]float64, size*size)
	data2 := make([]float64, size*size)

	// Test Equals
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data1[i] = (Randf(r) - 0.5) * 1000
		}

		m1 := MakeMatrix3(true, data1...)
		m2 := MakeMatrix3(true, data1...)

		if !m1.Equals(m2) {
			t.Error("m1 != m2\n\tm1:", m1, "\n\tm2:", m2)
		}
		if !m2.Equals(m1) {
			t.Error("m2 != m1\n\tm1:", m1, "\n\tm2:", m2)
		}
	}

	// Test !Equals
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data1[i] = (Randf(r) - 0.5) * 1000
		}

		m1 := MakeMatrix3(true, data1...)

		for element := 0; element < size*size; element++ {
			for i := 0; i < size*size; i++ {
				data2[i] = data1[i]
			}
			data2[element] += 1
			m2 := MakeMatrix3(true, data2...)

			if m1.Equals(m2) {
				t.Error("m1 == m2\n\tm1:", m1, "\n\tm2:", m2)
			}
			if m2.Equals(m1) {
				t.Error("m2 == m1\n\tm1:", m1, "\n\tm2:", m2)
			}
		}
	}
}

func TestMatrix3_ApproxEquals(t *testing.T) {
	// TODO: Get a bunch of test data to do this
}

func TestMatrix3_Inverse(t *testing.T) {
	// TODO: Get a bunch of test data to do this
}

func TestMatrix3_InverseThis(t *testing.T) {
	// TODO: Get a bunch of test data to do this
}

func TestMatrix3_Transpose(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)

	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}

		m1 := MakeMatrix3(true, data...)
		m2 := m1.Transpose()

		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				if m1.At(r, c) != m2.At(c, r) {
					t.Errorf("m[%d][%d] (%f) != m[%d][%d] (%f)", r, c, m1.At(r, c), c, r, m2.At(c, r))
				}
			}
		}
	}
}

func TestMatrix3_TransposeThis(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)

	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}

		m1 := MakeMatrix3(true, data...)
		m2 := MakeMatrix3(true, data...)
		m2.TransposeThis()

		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				if m1.At(r, c) != m2.At(c, r) {
					t.Errorf("m[%d][%d] (%f) != m[%d][%d] (%f)", r, c, m1.At(r, c), c, r, m2.At(c, r))
				}
			}
		}
	}
}

func TestMatrix3_ScalarMultiply(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}

		scale := (Randf(r) - 0.5) * 1000
		m1 := MakeMatrix3(true, data...)
		m2 := m1.ScalarMultiply(scale)

		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				if scale*m1.At(r, c) != m2.At(r, c) {
					t.Errorf("scale*m[%d][%d] (%f) != m[%d][%d] (%f)", r, c, scale*m1.At(r, c), r, c, m2.At(r, c))
				}
			}
		}
	}
}

func TestMatrix3_ScalarMultiplyThis(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}

		scale := (Randf(r) - 0.5) * 1000
		m1 := MakeMatrix3(true, data...)
		m2 := MakeMatrix3(true, data...)
		m2.ScalarMultiplyThis(scale)

		for r := 0; r < size; r++ {
			for c := 0; c < size; c++ {
				if scale*m1.At(r, c) != m2.At(r, c) {
					t.Errorf("scale*m[%d][%d] (%f) != m[%d][%d] (%f)", r, c, scale*m1.At(r, c), r, c, m2.At(r, c))
				}
			}
		}
	}
}

func TestMatrix3_MultiplyV(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))

	// Simple math test
	for iterations := 0; iterations < 1000; iterations++ {
		data1 := (Randf(r) - 0.5) * 1000
		data2 := (Randf(r) - 0.5) * 1000

		data3 := data1 * data2
		for element := 0; element < size*size; element++ {

			var m Matrix3
			m[element] = data1

			row_index := element / size
			col_index := element % size

			for v_element := 0; v_element < size; v_element++ {
				var v Vector3
				v[v_element] = data2

				r1 := m.MultiplyV(v)
				if col_index != v_element {
					var r2 Vector3
					if !r1.Equals(r2) {
						t.Error("r1 != r2\n\tm:", m, "\n\r1:", r1, "\n\r2:", r2)
					}
				} else {
					var r2 Vector3
					r2[row_index] = data3
					if !r1.Equals(r2) {
						t.Error("r1 != r2\n\tm:", m, "\n\r1:", r1, "\n\r2:", r2)
					}
				}
			}
		}
	}

	// Test identity matrix
	{
		data := make([]float64, size)
		m := MakeMatrix3Identity()
		for iterations := 0; iterations < 1000; iterations++ {
			for i := 0; i < size; i++ {
				data[i] = (Randf(r) - 0.5) * 1000
			}

			v1 := MakeVector3(data...)
			v2 := m.MultiplyV(v1)
			if !v1.Equals(v2) {
				t.Error("v1 != v2\n\tm:", m, "\n\v1:", v1, "\n\v2:", v2)
			}
		}
	}

	// TODO: Get a bunch of test data to do this
	input_matrix_data := [...][size * size]float64{}
	_ = input_matrix_data
}

func TestMatrix3_RightMultiply(t *testing.T) {
	// TODO: Get a bunch of test data to do this
}

func TestMatrix3_LeftMultiply(t *testing.T) {
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data := make([]float64, size*size)
	for iterations := 0; iterations < 1000; iterations++ {
		// Initialize the data
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}
		m1 := MakeMatrix3(true, data...)
		for i := 0; i < size*size; i++ {
			data[i] = (Randf(r) - 0.5) * 1000
		}
		m2 := MakeMatrix3(true, data...)

		m3 := m1.RightMultiply(m2)
		m4 := m2.LeftMultiply(m1)

		if !m3.Equals(m4) {
			t.Error("m3 != m4\n\tm1:", m1, "\n\tm2:", m2, "\n\tm3:", m3, "\n\tm4:", m4)
		}
	}
}

func BenchmarkMatrix3_ScalarMultiply(b *testing.B) {
	b.StopTimer()
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data1 := make([]float64, size*size)
	data2 := (Randf(r) - 0.5) * 1000
	for i := 0; i < size*size; i++ {
		data1[i] = (Randf(r) - 0.5) * 1000
	}
	m1 := MakeMatrix3(true, data1...)

	b.StartTimer()
	for iterations := 0; iterations < b.N; iterations++ {
		m1.ScalarMultiply(data2)
	}
}

func BenchmarkMatrix3_RightMultiply1(b *testing.B) {
	b.StopTimer()
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data1 := make([]float64, size*size)
	data2 := make([]float64, size*size)
	for i := 0; i < size*size; i++ {
		data1[i] = (Randf(r) - 0.5) * 1000
	}
	for i := 0; i < size*size; i++ {
		data2[i] = (Randf(r) - 0.5) * 1000
	}
	m1 := MakeMatrix3(true, data1...)
	m2 := MakeMatrix3(true, data2...)

	b.StartTimer()
	for iterations := 0; iterations < b.N; iterations++ {
		m1.RightMultiply(m2)
	}
}
func BenchmarkMatrix3_RightMultiply2(b *testing.B) {
	b.StopTimer()
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data1 := make([]float64, size*size)
	data2 := make([]float64, size*size)
	for i := 0; i < size*size; i++ {
		data1[i] = (Randf(r) - 0.5) * 1000
	}
	for i := 0; i < size*size; i++ {
		data2[i] = (Randf(r) - 0.5) * 1000
	}
	m1 := MakeMatrix3(true, data1...)
	m2 := MakeMatrix3(true, data2...)

	b.StartTimer()
	for iterations := 0; iterations < b.N; iterations++ {
		m1.RightMultiply2(m2)
	}
}
func BenchmarkMatrix3_RightMultiply3(b *testing.B) {
	b.StopTimer()
	const size = 3
	r := rand.New(rand.NewSource(time.Nanoseconds()))
	data1 := make([]float64, size*size)
	data2 := make([]float64, size*size)
	for i := 0; i < size*size; i++ {
		data1[i] = (Randf(r) - 0.5) * 1000
	}
	for i := 0; i < size*size; i++ {
		data2[i] = (Randf(r) - 0.5) * 1000
	}
	m1 := MakeMatrix3(true, data1...)
	m2 := MakeMatrix3(true, data2...)

	b.StartTimer()
	for iterations := 0; iterations < b.N; iterations++ {
		m1.RightMultiply3(m2)
	}
}
