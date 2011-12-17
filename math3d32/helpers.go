/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

import "math"
import "rand"

// Sed these:
//type float32 float32
//type uint32 uint32
//type int32 int32

const mantissaSize = 23
const minInt = math.MinInt32


func floatBits(f float32) int32 { return int32(math.Float32bits(float32(f))) }
func Randf(r *rand.Rand) float32 { return float32(rand.Float32()) }
