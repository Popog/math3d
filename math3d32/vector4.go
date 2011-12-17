/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "fmt"

type Vector4 [4]float32

func MakeVector4V(v []float32) Vector4 {
	return Vector4{v[0], v[1], v[2], v[3]}
}

func MakeVector4(x, y, z, o float32) Vector4 {
	v := Vector4{x, y, z, o}
	return v
}

// return v1+v2 (won't modify any of them)
func (v1 Vector4) Add(v2 Vector4) Vector4 {
	return Vector4{v1[0] + v2[0], v1[1] + v2[1], v1[2] + v2[2], v1[3] + v2[3]}
}

// return v1-v2 (won't modify any of them)
func (v1 Vector4) Sub(v2 Vector4) Vector4 {
	return Vector4{v1[0] - v2[0], v1[1] - v2[1], v1[2] - v2[2], v1[3] - v2[3]}
}

func (v1 Vector4) Dot(v2 Vector4) float32 {
	return v1[0]*v2[0] + v1[1]*v2[1] + v1[2]*v2[2]
}

func (v1 Vector4) Cross(v2 Vector4) Vector4 {
	return Vector4{v1[1]*v2[2] - v1[2]*v2[1], v1[2]*v2[0] - v1[0]*v2[2], v1[0]*v2[1] - v1[1]*v2[0]}
}

/*
// For those cases when the 4d vector represents just a 3d vector. 4:t axis is ignored
func (v1 Vector4) Dot3d(v2 Vector4) float32 {
	return v1[0]*v2[0]+v1[1]*v2[1]+v1[2]*v2[2]
}

// For those cases when the 4d vector represents just a 3d vector. 4:t axis is ignored
func (v1 Vector4) Cross3d(v2 Vector4) Vector4 {
	return Vector4{v1[1]*v2[2]-v1[2]*v2[1],v1[2]*v2[0]-v1[0]*v2[2],v1[0]*v2[1]-v1[1]*v2[0]} 
}
*/

// If two vectors represents points the distance between them can be calculated
// Forth value is ignored
func (v0 Vector4) Distance3d(v1 Vector4) float32 {
	d0 := v0[0] - v1[0]
	d1 := v0[1] - v1[1]
	d2 := v0[2] - v1[2]
	return Sqrtf(d0*d0 + d1*d1 + d2*d2)
}

func (v Vector4) Length() float32 {
	return Sqrtf(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])
}

// Normalize will modify this vector
func (v Vector4) Normalize() Vector4 {
	l := v.Length()
	v[0] /= l
	v[1] /= l
	v[2] /= l
	return v
}

func (m Vector4) Equal(q Vector4) bool {
	return m[0] == q[0] && m[1] == q[1] && m[2] == q[2] && m[3] == q[3]
}

func (a Vector4) ApproxEquals(b Vector4, ε float32) bool {
	for i := 0; i < 4; i++ {
		if Fabsf(a[i]-b[i]) > ε {
			return false
		}
	}
	return true
}

func (v Vector4) String() string {
	return fmt.Sprintf("[%.5f,%.5f,%.5f,%.5f]", v[0], v[1], v[2], v[3])
}

// Tests to see if the difference between two matrices, element-wise, exceeds ε.

