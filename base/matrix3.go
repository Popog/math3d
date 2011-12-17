/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "fmt"

type Matrix3 [3*3]floatType


// Constructors
func MakeMatrix3(v []floatType, rowMajor bool) (r Matrix3) {
	for i := 0; i < len(r); i++ { r[i] = v[i] }
	// transform the data to OpenGl format
	if !rowMajor { r.TransposeThis() }
	return
}
func (m * Matrix3) ZeroThis() {
	*m = Matrix3{}
}
func MakeMatrix3Identity() (m Matrix3) {
	const size = 3
	for i := 0; i < size; i++ { m[i*size + i] = 1 }
	return
}
func (m *Matrix3) IdentityThis() {
	*m = MakeMatrix3Identity()
}


// d8888b. d8888b.  .d88b.  d8888b. d88888b d8888b. d888888b db    db      d888b  d88888b d888888b d888888b d88888b d8888b. .d8888. 
// 88  `8D 88  `8D .8P  Y8. 88  `8D 88'     88  `8D `~~88~~' `8b  d8'     88' Y8b 88'     `~~88~~' `~~88~~' 88'     88  `8D 88'  YP 
// 88oodD' 88oobY' 88    88 88oodD' 88ooooo 88oobY'    88     `8bd8'      88      88ooooo    88       88    88ooooo 88oobY' `8bo.   
// 88~~~   88`8b   88    88 88~~~   88~~~~~ 88`8b      88       88        88  ooo 88~~~~~    88       88    88~~~~~ 88`8b     `Y8b. 
// 88      88 `88. `8b  d8' 88      88.     88 `88.    88       88        88. ~8~ 88.        88       88    88.     88 `88. db   8D 
// 88      88   YD  `Y88P'  88      Y88888P 88   YD    YP       YP         Y888P  Y88888P    YP       YP    Y88888P 88   YD `8888Y' 

// Returns a row as a vector
func (m Matrix3) GetRow(row int) (r Vector3) {
	for i := 0; i < len(r); i++ { r[i] = m.At(row, i) }
	return
}

// Returns a column as a vector
func (m Matrix3) GetCol(col int) (r Vector3) {
	for i := 0; i < len(r); i++ { r[i] = m.At(i, col) }
	return
}

// Returns the element at row,col
func (m Matrix3) At(row, col int) floatType {
	const size = 3
	return m[row*size+col]
}

func (m Matrix3) Determinant() floatType {
	return m[0]*(m[4]*m[8]-m[5]*m[7]) - m[1]*(m[3]*m[8]-m[5]*m[6]) + m[2]*(m[3]*m[7]-m[4]*m[6])
}

func (m Matrix3) Cofactor() (r Matrix3) {
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
func (m Matrix3) ApproxEquals(q Matrix3, ε floatType) bool {
	for i := 0; i < len(m); i++ {
		if ApproxEquals(m[i], q[i], ε) {
			return false
		}
	}
	return true
}

func (m Matrix3) Equals(q Matrix3) bool {
	for i := 0; i < len(m); i++ {
		if(m[i] != q[i]) { return false }
	}
	return true
	//return m[0] == q[0] && m[3] == q[3] && m[6] == q[6] && m[1] == q[1] && m[4] == q[4] && m[7] == q[7] && m[2] == q[2] && m[5] == q[5] && m[8] == q[8]
}

func (m Matrix3) String() string {
	// output in octave format for easy testing
	return fmt.Sprintf("[%.5f,%.5f,%.5f;%.5f,%.5f,%.5f;%.5f,%.5f,%.5f]", m[0], m[1], m[2], m[3], m[4], m[5], m[6], m[7], m[8])
}



func (m Matrix3) Inverse() (r Matrix3) {
	d := 1.0 / m.Determinant()
	r[0] = d * (m[4]*m[8] - m[5]*m[7])
	r[1] = -d * (m[1]*m[8] - m[2]*m[7])
	r[2] = d * (m[1]*m[5] - m[2]*m[4])
	r[3] = -d * (m[3]*m[8] - m[5]*m[6])
	r[4] = d * (m[0]*m[8] - m[2]*m[6])
	r[5] = -d * (m[0]*m[5] - m[2]*m[3])
	r[6] = d * (m[3]*m[7] - m[4]*m[6])
	r[7] = -d * (m[0]*m[7] - m[1]*m[6])
	r[8] = d * (m[0]*m[4] - m[1]*m[3])
	return
}

func (m *Matrix3) InverseThis() {
	*m = m.Inverse()
}

// Return the transpose of matrix
func (m Matrix3) Transpose() Matrix3 {
	const size = 3
	for r := 0; r < size; r++ {
		for c := 0; c < r; c++ {
			m[r*size + c], m[c*size + r] = m[c*size + r], m[r*size + c]
		}
	}
	return m
}

// Transposes the matrix in-place
func (m * Matrix3) TransposeThis() {
	*m = m.Transpose()
}

func (m Matrix3) ScalarMultiply(scalar floatType) Matrix3 {
	for i := 0; i < len(m); i++ { m[i] *= scalar }
	return m
}

func (m *Matrix3) ScalarMultiplyThis(scalar floatType) {
	*m = m.ScalarMultiply(scalar)
}

// Mutiply this matrix with a column vector v, resulting in another column vector
func (m Matrix3) MultiplyV(v Vector3) (r Vector3) {
	for i := 0; i < len(r); i++ { r[i] = m.GetRow(i).Dot(v) }
	return
}

// Returns m * q
func (m Matrix3) RightMultiply(q Matrix3) (result Matrix3) {
	const size = 3
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			result[r*size + c] = m.GetRow(r).Dot(q.GetCol(c))
		}
	}
	return
}

// Return q * m (if for some reason you wanted that)
func (m Matrix3) LeftMultiply(q Matrix3) (result Matrix3) {
	const size = 3
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			result[r*size + c] = q.GetRow(r).Dot(m.GetCol(c))
		}
	}
	return
}

/*
// Orthogonalize will modify this matrix
func (m Matrix3) Orthogonalize(){
	i := MakeVector3(m[0],m[1],m[2])
	j := MakeVector3(m[3],m[4],m[5]) 
	k := MakeVector3(m[6],m[7],m[8]).Normalize()
	i = j.Cross(k).Normalize()
	j=k.Cross(i)
	m[0]=i[0]; m[3]=j[0]; m[6]=k[0]
	m[1]=i[3]; m[4]=j[3]; m[7]=k[3]
	m[2]=i[6]; m[5]=j[6]; m[8]=k[6]
}

func (m1 Matrix3) Orthogonalized() Matrix3{
	m := m1.Copy()
	m.Orthogonalize()
	return m
}
*/


