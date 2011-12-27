/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "fmt"

// This is a 4x4 matrix of float32, stored in row major
type Matrix4 [4 * 4]float32

// Make a new matrix from the array
func MakeMatrix4(rowMajor bool, v ...float32) (r Matrix4) {
	copy(r[:], v)
	// transform the data to OpenGl format
	if !rowMajor {
		r.TransposeThis()
	}
	return
}

// Set the current matrix to be a zero matrix
func (m *Matrix4) ZeroThis() {
	*m = Matrix4{}
}

// Create a new identity matrix
func MakeMatrix4Identity() (m Matrix4) {
	const size = 3
	for i := 0; i < size; i++ {
		m[i*size+i] = 1
	}
	return
}

// Set the current matrix to be an identity matrix
func (m *Matrix4) IdentityThis() {
	*m = MakeMatrix4Identity()
}

// untested code
func MakeRotationMatrix(look, tmpUp Vector3) Matrix4 {

	look = look.Normalize()
	right := tmpUp.Normalize().Cross(look).Normalize()
	up := look.Cross(right).Normalize()

	return Matrix4{
		right[0], right[1], right[2], 0,
		up[0], up[1], up[2], 0,
		look[0], look[1], look[2], 0,
		0., 0., 0., 1}
}

// d8888b. d8888b.  .d88b.  d8888b. d88888b d8888b. d888888b db    db      d888b  d88888b d888888b d888888b d88888b d8888b. .d8888. 
// 88  `8D 88  `8D .8P  Y8. 88  `8D 88'     88  `8D `~~88~~' `8b  d8'     88' Y8b 88'     `~~88~~' `~~88~~' 88'     88  `8D 88'  YP 
// 88oodD' 88oobY' 88    88 88oodD' 88ooooo 88oobY'    88     `8bd8'      88      88ooooo    88       88    88ooooo 88oobY' `8bo.   
// 88~~~   88`8b   88    88 88~~~   88~~~~~ 88`8b      88       88        88  ooo 88~~~~~    88       88    88~~~~~ 88`8b     `Y8b. 
// 88      88 `88. `8b  d8' 88      88.     88 `88.    88       88        88. ~8~ 88.        88       88    88.     88 `88. db   8D 
// 88      88   YD  `Y88P'  88      Y88888P 88   YD    YP       YP         Y888P  Y88888P    YP       YP    Y88888P 88   YD `8888Y' 

// Returns a row as a vector
func (m Matrix4) GetRow(row int) (r Vector4) {
	for i := 0; i < len(r); i++ {
		r[i] = m.At(row, i)
	}
	return
}

// Returns a column as a vector
func (m Matrix4) GetCol(col int) (r Vector4) {
	for i := 0; i < len(r); i++ {
		r[i] = m.At(i, col)
	}
	return
}

// Returns the element at row,col
func (m Matrix4) At(row, col int) float32 {
	const size = 4
	return m[row*size+col]
}

// Gets the determinant 
func (m Matrix4) Determinant() float32 {
	a0 := m[0]*m[5] - m[4]*m[1]
	a1 := m[0]*m[9] - m[8]*m[1]
	a2 := m[0]*m[13] - m[12]*m[1]
	a3 := m[4]*m[9] - m[8]*m[5]
	a4 := m[4]*m[13] - m[12]*m[5]
	a5 := m[8]*m[13] - m[12]*m[9]
	b0 := m[2]*m[7] - m[6]*m[3]
	b1 := m[2]*m[11] - m[10]*m[3]
	b2 := m[2]*m[15] - m[14]*m[3]
	b3 := m[6]*m[11] - m[10]*m[7]
	b4 := m[6]*m[15] - m[14]*m[7]
	b5 := m[10]*m[15] - m[14]*m[11]
	return a0*b5 - a1*b4 + a2*b3 + a3*b2 - a4*b1 + a5*b0
}

// Todo - fixme
func (m Matrix4) Cofactor() (r Matrix4) {
	r[0] = (m[4]*m[8] - m[5]*m[7])
	r[1] = -(m[3]*m[8] - m[5]*m[6])
	r[2] = (m[3]*m[7] - m[4]*m[6])
	r[3] = -(m[1]*m[8] - m[2]*m[7])
	r[4] = (m[0]*m[8] - m[2]*m[6])
	r[5] = -(m[0]*m[7] - m[1]*m[6])
	r[6] = (m[1]*m[5] - m[2]*m[4])
	r[7] = -(m[0]*m[5] - m[2]*m[3])
	r[8] = (m[0]*m[4] - m[1]*m[3])
	return
}

// Tests to see if the difference between two matrices, element-wise, exceeds ε.
func (m Matrix4) ApproxEquals(q Matrix4, ε float32) bool {
	for i := 0; i < len(m); i++ {
		if ApproxEquals(m[i], q[i], ε) {
			return false
		}
	}
	return true
}

func (m Matrix4) Equal(q Matrix4) bool {
	for i := 0; i < len(m); i++ {
		if m[i] != q[i] {
			return false
		}
	}
	return true

	// return m[0] == q[0] && m[1] == q[1] && m[2] == q[2] && m[3] == q[3] && m[4] == q[4] && m[5] == q[5] &&
	// 	m[6] == q[6] && m[7] == q[7] && m[8] == q[8] && m[9] == q[9] && m[10] == q[10] && m[11] == q[11] &&
	// 	m[12] == q[12] && m[13] == q[13] && m[14] == q[14] && m[15] == q[15]
}

func (m Matrix4) String() string {
	// output in octave format for easy testing
	return fmt.Sprintf("[%.5f,%.5f,%.5f,%.5f;%.5f,%.5f,%.5f,%.5f;%.5f,%.5f,%.5f,%.5f;%.5f,%.5f,%.5f,%.5f]",
		m[0], m[1], m[2], m[3],
		m[4], m[5], m[6], m[7],
		m[8], m[9], m[10], m[11],
		m[12], m[13], m[14], m[15])
}

func (m Matrix4) Inverse() Matrix4 {
	a0 := m[0]*m[5] - m[4]*m[1]
	a1 := m[0]*m[9] - m[8]*m[1]
	a2 := m[0]*m[13] - m[12]*m[1]
	a3 := m[4]*m[9] - m[8]*m[5]
	a4 := m[4]*m[13] - m[12]*m[5]
	a5 := m[8]*m[13] - m[12]*m[9]
	b0 := m[2]*m[7] - m[6]*m[3]
	b1 := m[2]*m[11] - m[10]*m[3]
	b2 := m[2]*m[15] - m[14]*m[3]
	b3 := m[6]*m[11] - m[10]*m[7]
	b4 := m[6]*m[15] - m[14]*m[7]
	b5 := m[10]*m[15] - m[14]*m[11]
	det := a0*b5 - a1*b4 + a2*b3 + a3*b2 - a4*b1 + a5*b0

	id := 1. / det
	return Matrix4{
		id * (+m[5]*b5 - m[9]*b4 + m[13]*b3),
		id * (-m[1]*b5 + m[9]*b2 - m[13]*b1),
		id * (+m[1]*b4 - m[5]*b2 + m[13]*b0),
		id * (-m[1]*b3 + m[5]*b1 - m[9]*b0),
		id * (-m[4]*b5 + m[8]*b4 - m[12]*b3),
		id * (+m[0]*b5 - m[8]*b2 + m[12]*b1),
		id * (-m[0]*b4 + m[4]*b2 - m[12]*b0),
		id * (+m[0]*b3 - m[4]*b1 + m[8]*b0),
		id * (+m[7]*a5 - m[11]*a4 + m[15]*a3),
		id * (-m[3]*a5 + m[11]*a2 - m[15]*a1),
		id * (+m[3]*a4 - m[7]*a2 + m[15]*a0),
		id * (-m[3]*a3 + m[7]*a1 - m[11]*a0),
		id * (-m[6]*a5 + m[10]*a4 - m[14]*a3),
		id * (+m[2]*a5 - m[10]*a2 + m[14]*a1),
		id * (-m[2]*a4 + m[6]*a2 - m[14]*a0),
		id * (+m[2]*a3 - m[6]*a1 + m[10]*a0)}
}

func (m *Matrix4) InverseThis() {
	*m = m.Inverse()
}

// Return the transpose of matrix
func (m Matrix4) Transpose() Matrix4 {
	const size = 4
	for r := 0; r < size; r++ {
		for c := 0; c < r; c++ {
			m[r*size+c], m[c*size+r] = m[c*size+r], m[r*size+c]
		}
	}
	return m
}

// Transposes the matrix in-place
func (m *Matrix4) TransposeThis() {
	*m = m.Transpose()
}

func (m Matrix4) ScalarMultiply(scalar float32) Matrix4 {
	for i := 0; i < len(m); i++ {
		m[i] *= scalar
	}
	return m
}

func (m *Matrix4) ScalarMultiplyThis(scalar float32) {
	*m = m.ScalarMultiply(scalar)
}

// Mutiply this matrix with a column vector v, resulting in another column vector
func (m Matrix4) MultiplyV(v Vector4) (r Vector4) {
	for i := 0; i < len(r); i++ {
		r[i] = m.GetRow(i).Dot(v)
	}
	return
}

// Returns m * q
func (m Matrix4) RightMultiply(q Matrix4) (result Matrix4) {
	const size = 4
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			result[r*size+c] = m.GetRow(r).Dot(q.GetCol(c))
		}
	}
	return
}

// return q * m (if for some reason you wanted that)
func (m Matrix4) LeftMultiply(q Matrix4) (result Matrix4) {
	const size = 4
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			result[r*size+c] = q.GetRow(r).Dot(m.GetCol(c))
		}
	}
	return
}

func (m Matrix4) Add(q Matrix4) (r Matrix4) {
	for i := 0; i < len(m); i++ {
		r[i] = m[i] + q[i]
	}
	return
}

func (m *Matrix4) AddThis(q Matrix4) {
	*m = m.Add(q)
}

func (m Matrix4) Sub(q Matrix4) (r Matrix4) {
	for i := 0; i < len(m); i++ {
		r[i] = m[i] - q[i]
	}
	return
}

func (m *Matrix4) SubThis(q Matrix4) {
	*m = m.Sub(q)
}

/*
// Orthogonalize will modify this matrix (fixme)
func (m Matrix4) Orthogonalize(){
	i := MakeVf(m[0],m[1],m[2])
	j := MakeVf(m[3],m[4],m[5]) 
	k := MakeVf(m[6],m[7],m[8]).Normalize()
	i = j.Cross(k).Normalize()
	j=k.Cross(i)
	m[0]=i[0]; m[3]=j[0]; m[6]=k[0]
	m[1]=i[3]; m[4]=j[3]; m[7]=k[3]
	m[2]=i[6]; m[5]=j[6]; m[8]=k[6]
}

// Orthogonalize will not modify this matrix (fixme)
func (m1 Matrix4) Orthogonalized() Matrix4{
	m := m1.Copy()
	m.Orthogonalize()
	return m
}

*/

func (m Matrix4) ToQuaternion() Quaternion {
	// Algorithm in Ken Shoemake's article in 1987 SIGGRAPH course notes
	// article "HQuaternion Calculus and Fast Animation".
	toQuaternionNext := []int{1, 2, 0}

	q := Quaternion{0, 0, 0, 0}
	//fmt.Println("q = ", q)
	trace := m[0] + m[5] + m[10]
	var root float32
	//fmt.Printf("trace = %f\n", trace)
	if trace > 0. {
		// |w| > 1/2, may as well choose w > 1/2
		root = Sqrtf(trace + 1.0) // 2w
		q[0] = 0.5 * root
		root = 0.5 / root // 1/(4w)
		q[1] = (m.At(2, 1) - m.At(1, 2)) * root
		q[2] = (m.At(0, 2) - m.At(2, 0)) * root
		q[3] = (m.At(1, 0) - m.At(0, 1)) * root
	} else {
		// |w| <= 1/2
		i := 0
		if m.At(1, 1) > m.At(0, 0) {
			i = 1
		}
		if m.At(2, 2) > m.At(i, i) {
			i = 2
		}
		j := toQuaternionNext[i]
		k := toQuaternionNext[j]

		root = Sqrtf(m.At(i, i) - m.At(j, j) - m.At(k, k) + 1.)
		quat := q[1:]
		//fmt.Printf("Quat = [%f,%f,%f]\n", quat[0],quat[1],quat[2])
		quat[i] = 0.5 * root
		root = 0.5 / root
		q[0] = (m.At(k, j) - m.At(j, k)) * root
		quat[j] = (m.At(j, i) + m.At(i, j)) * root
		quat[k] = (m.At(k, i) + m.At(i, k)) * root
	}
	return q
}
