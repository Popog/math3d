/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "math"
import "rand"

// Set these based on the type of Float
type floatType float32
type uintType uint32
type intType int32
const mantissaSize = 23
const minInt = math.MinInt32


func floatBits(f floatType) intType { return intType(math.Float32bits(float32(f))) }
func Randf(r *rand.Rand) floatType { return floatType(rand.Float32()) }
