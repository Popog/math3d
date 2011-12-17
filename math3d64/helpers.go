/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d64

import "math"
import "rand"

// Sed these
//type float64 float64
//type uint64 uint64
//type int64 int64

const mantissaSize = 52
const minInt = math.MinInt32


func floatBits(f float64) int64 { return int64(math.Float64bits(float64(f))) }
func Randf(r *rand.Rand) float64 { return float64(rand.Float32()) }
