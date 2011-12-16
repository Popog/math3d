/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "fmt"

type Vector3 [3]Float

func MakeVector3(v []Float) (r Vector3) {
	for i := 0; i < len(r); i++ { r[i] = v[i] }
	return
}

// Entrywise addition
func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2]}
}
// In-place entrywise addition
func (v1 * Vector3) AddThis(v2 Vector3) {
	*v1 = v1.Add(v2)
}

// Entrywise subtraction
func (v1 Vector3) Sub(v2 Vector3) Vector3 {
	return Vector3{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2]}
}
// In-place entrywise subtraction
func (v1 * Vector3) SubThis(v2 Vector3) {
	*v1 = v1.Sub(v2)
}

// Entrywise product (Hadamard product?)
func (v1 Vector3) Mul(v2 Vector3) (r Vector3) {
	for i := 0; i < len(v1); i++ { r[i] = v1[i]*v2[i] }
	return
}
// In-place entrywise product (Hadamard product?)
func (v1 *Vector3) MulThis(v2 Vector3) {
	*v1 = v1.Mul(v2)
}
// Entrywise quotient (Hadamard quotient?)
func (v1 Vector3) Div(v2 Vector3) (r Vector3) {
	for i := 0; i < len(v1); i++ { r[i] = v1[i]/v2[i] }
	return
}
// In-place entrywise quotient (Hadamard quotient?)
func (v1 *Vector3) DivThis(v2 Vector3) {
	*v1 = v1.Div(v2)
}



// Scalar multiplication
func (v Vector3) ScalarMultiply(scalar Float) (r Vector3) {
	for i := 0; i < len(v); i++ { r[i] = v[i]*scalar }
	return
}
// In place scalar multiplication
func (v *Vector3) ScalarMultiplyThis(scalar Float) {
	*v = v.ScalarMultiply(scalar)
}

func (v1 Vector3) Dot(v2 Vector3) (r Float) {
	for i := 0; i < len(v1); i++ { r += v1[i]*v2[i] }
	return
}

func (v1 Vector3) Cross(v2 Vector3) Vector3 {
	return Vector3{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

// The magnitude squared of a vector
func (v Vector3) LengthSq() (m Float) {
	return v.Dot(v)
}
// The magnitude of a vector
func (v Vector3) Length() Float {
	return Sqrtf(v.LengthSq())
}

// If two vectors represents points the distance squared between them can be calculated
func (v0 Vector3) DistanceSq(v1 Vector3) Float {
	return v0.Sub(v1).LengthSq()
}
// If two vectors represents points the distance between them can be calculated
func (v0 Vector3) Distance(v1 Vector3) Float {
	return Sqrtf(v0.DistanceSq(v1))
}


// Normalize will modify this vector
func (v Vector3) Normalize() Vector3 {
	return v.ScalarMultiply(1.0 / v.Length())
}
// In place normalize
func (v *Vector3) NormalizeThis() {
	*v = v.Normalize()
}


func (m1 Vector3) Equals(q Vector3) bool {
	return m1[0] == q[0] && m1[1] == q[1] && m1[2] == q[2]
}

func (a Vector3) ApproxEquals(b Vector3, ε Float) bool {
	for i := 0; i < 3; i++ {
		if Fabsf(a[i]-b[i]) > ε {
			return false
		}
	}
	return true
}

// untested
func (v Vector3) Yaw() Float {
	return -Atan2f(v[0], v[2])
}

// untested
func (v Vector3) Pitch() Float {
	return -Atan2f(v[1], Sqrtf(v[0]*v[0]+v[2]*v[2]))
}

func (v Vector3) String() string {
	return fmt.Sprintf("[%.5f,%.5f,%.5f]", v[0], v[1], v[2])
}

// p1,p2,p3 represents points
func SurfaceNormal(p1, p2, p3 Vector3) Vector3 {
	u := Vector3{p2[0] - p1[0], p2[1] - p1[2], p2[2] - p1[2]}
	v := Vector3{p3[0] - p1[0], p3[1] - p1[2], p3[2] - p1[2]}
	return Vector3{u[1]*v[2] - u[2]*v[1], u[2]*v[0] - u[0]*v[2], u[0]*v[1] - u[1]*v[0]}
}