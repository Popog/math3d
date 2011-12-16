/*
Note that this code uses row major matrixes
Distributed under the Boost Software License, Version 1.0.
http://www.boost.org/LICENSE_1_0.txt
*/

package math3d32

// import "math"
import "fmt"
import "testing"

var a3Array, aI3Array, b3Array, bI3Array, p3Array, pI3Array, va3Array, v03Array, v13Array [][]float32

const ε3d = 0.0001

func init() {
	// randomly generated test data. Products and inverts generated by numpy
	a3Array = [][]float32{[]float32{4.000000, -1.000000, -3.000000, -6.000000, 8.000000, 8.000000, -1.000000, -8.000000, 4.000000}, []float32{-7.000000, 7.000000, -4.000000, -4.000000, 4.000000, 3.000000, -1.000000, 3.000000, 0.000000}, []float32{-2.000000, 7.000000, 3.000000, 9.000000, 0.000000, -3.000000, 8.000000, 7.000000, 0.000000}, []float32{-2.000000, 8.000000, -9.000000, 9.000000, -1.000000, -8.000000, -7.000000, -9.000000, -8.000000}, []float32{9.000000, -7.000000, -5.000000, 6.000000, -2.000000, 9.000000, 7.000000, -2.000000, -8.000000}, []float32{-2.000000, 9.000000, 1.000000, -5.000000, -8.000000, -3.000000, -9.000000, -1.000000, 8.000000}, []float32{4.000000, -8.000000, -9.000000, -9.000000, -6.000000, 1.000000, -2.000000, 8.000000, 8.000000}, []float32{8.000000, -3.000000, -7.000000, 6.000000, -2.000000, -3.000000, -9.000000, -8.000000, -6.000000}, []float32{-5.000000, 7.000000, 9.000000, 1.000000, 4.000000, 9.000000, 2.000000, 9.000000, 9.000000}, []float32{7.000000, 9.000000, 3.000000, 7.000000, 8.000000, -6.000000, 7.000000, -7.000000, 1.000000}, []float32{-6.000000, 6.000000, -6.000000, -7.000000, 1.000000, 4.000000, 9.000000, -6.000000, 9.000000}, []float32{6.000000, 5.000000, 0.000000, 4.000000, -9.000000, -6.000000, 1.000000, -4.000000, -7.000000}, []float32{-1.000000, 8.000000, 8.000000, 1.000000, 3.000000, -8.000000, -3.000000, 0.000000, 4.000000}, []float32{3.000000, -5.000000, -6.000000, 5.000000, 3.000000, 1.000000, 6.000000, 8.000000, 6.000000}, []float32{-2.000000, -8.000000, -8.000000, -4.000000, -3.000000, -9.000000, -7.000000, 2.000000, 3.000000}, []float32{-8.000000, -1.000000, 1.000000, 0.000000, -7.000000, -1.000000, -5.000000, 4.000000, -8.000000}, []float32{2.000000, 7.000000, 8.000000, 2.000000, -4.000000, 2.000000, 5.000000, 8.000000, -9.000000}, []float32{-9.000000, 4.000000, 2.000000, 0.000000, 4.000000, 3.000000, 5.000000, 0.000000, -9.000000}, []float32{-6.000000, 0.000000, 4.000000, -8.000000, 5.000000, 2.000000, -3.000000, 0.000000, 5.000000}, []float32{3.000000, 8.000000, 2.000000, -2.000000, -5.000000, -8.000000, -4.000000, -5.000000, -8.000000}, []float32{-2.000000, -9.000000, 2.000000, 3.000000, -4.000000, -6.000000, 1.000000, -3.000000, -5.000000}, []float32{-2.000000, -4.000000, -8.000000, -6.000000, -2.000000, 4.000000, -4.000000, -1.000000, 0.000000}, []float32{-1.000000, -4.000000, 9.000000, 3.000000, 9.000000, -3.000000, -5.000000, -1.000000, -5.000000}, []float32{-4.000000, 1.000000, 6.000000, -9.000000, -2.000000, 6.000000, -1.000000, -3.000000, -5.000000}, []float32{2.000000, 3.000000, -1.000000, -5.000000, 0.000000, -7.000000, 7.000000, -1.000000, -8.000000}, []float32{-8.000000, 4.000000, 5.000000, -6.000000, 7.000000, 4.000000, -3.000000, 5.000000, -4.000000}, []float32{-1.000000, 6.000000, -9.000000, 6.000000, 7.000000, -9.000000, -8.000000, 2.000000, -4.000000}, []float32{-2.000000, -8.000000, 6.000000, -9.000000, 8.000000, 7.000000, 7.000000, 0.000000, -4.000000}, []float32{5.000000, -1.000000, -8.000000, 6.000000, 6.000000, 3.000000, 0.000000, 8.000000, 5.000000}, []float32{9.000000, 8.000000, 3.000000, -6.000000, -7.000000, -9.000000, 8.000000, -2.000000, -5.000000}}
	aI3Array = [][]float32{[]float32{0.480000, 0.140000, 0.080000, 0.080000, 0.065000, -0.070000, 0.280000, 0.165000, 0.130000}, []float32{-0.121622, -0.162162, 0.500000, -0.040541, -0.054054, 0.500000, -0.108108, 0.189189, -0.000000}, []float32{-1.000000, -1.000000, 1.000000, 1.142857, 1.142857, -1.000000, -3.000000, -3.333333, 3.000000}, []float32{-0.032922, 0.074588, -0.037551, 0.065844, -0.024177, -0.049897, -0.045267, -0.038066, -0.036008}, []float32{-0.070686, 0.095634, 0.151767, -0.230769, 0.076923, 0.230769, -0.004158, 0.064449, -0.049896}, []float32{-0.100000, -0.108955, -0.028358, 0.100000, -0.010448, -0.016418, -0.100000, -0.123881, 0.091045}, []float32{2.000000, 0.285714, 2.214286, -2.500000, -0.500000, -2.750000, 3.000000, 0.571429, 3.428571}, []float32{-0.067797, 0.214689, -0.028249, 0.355932, -0.627119, -0.101695, -0.372881, 0.514124, 0.011299}, []float32{-0.151515, 0.060606, 0.090909, 0.030303, -0.212121, 0.181818, 0.003367, 0.198653, -0.090909}, []float32{0.034205, 0.030181, 0.078471, 0.049296, 0.014085, -0.063380, 0.105634, -0.112676, 0.007042}, []float32{0.166667, -0.090909, 0.151515, 0.500000, -0.000000, 0.333333, 0.166667, 0.090909, 0.181818}, []float32{0.113372, 0.101744, -0.087209, 0.063953, -0.122093, 0.104651, -0.020349, 0.084302, -0.215116}, []float32{0.054545, -0.145455, -0.400000, 0.090909, 0.090909, -0.000000, 0.040909, -0.109091, -0.050000}, []float32{0.555556, -1.000000, 0.722222, -1.333333, 3.000000, -1.833333, 1.222222, -3.000000, 1.888889}, []float32{-0.023316, -0.020725, -0.124352, -0.194301, 0.160622, -0.036269, 0.075130, -0.155440, 0.067358}, []float32{-0.115385, 0.007692, -0.015385, -0.009615, -0.132692, 0.015385, 0.067308, -0.071154, -0.107692}, []float32{0.038168, 0.242366, 0.087786, 0.053435, -0.110687, 0.022901, 0.068702, 0.036260, -0.041985}, []float32{-0.104651, 0.104651, 0.011628, 0.043605, 0.206395, 0.078488, -0.058140, 0.058140, -0.104651}, []float32{-0.277778, 0.000000, 0.222222, -0.377778, 0.200000, 0.222222, -0.166667, 0.000000, 0.333333}, []float32{0.000000, 0.500000, -0.500000, 0.148148, -0.148148, 0.185185, -0.092593, -0.157407, 0.009259}, []float32{-0.021053, 0.536842, -0.652632, -0.094737, -0.084211, 0.063158, 0.052632, 0.157895, -0.368421}, []float32{0.055556, 0.111111, -0.444444, -0.222222, -0.444444, 0.777778, -0.027778, 0.194444, -0.277778}, []float32{-0.156863, -0.094771, -0.225490, 0.098039, 0.163399, 0.078431, 0.137255, 0.062092, 0.009804}, []float32{-2.153846, 1.000000, -1.384615, 3.923077, -2.000000, 2.307692, -1.923077, 1.000000, -1.307692}, []float32{0.024476, -0.087413, 0.073427, 0.311189, 0.031469, -0.066434, -0.017483, -0.080420, -0.052448}, []float32{-0.246154, 0.210256, -0.097436, -0.184615, 0.241026, 0.010256, -0.046154, 0.143590, -0.164103}, []float32{0.384615, -0.230769, -0.346154, -3.692308, 2.615385, 2.423077, -2.615385, 1.769231, 1.653846}, []float32{0.085106, 0.085106, 0.276596, -0.034574, 0.090426, 0.106383, 0.148936, 0.148936, 0.234043}, []float32{-0.018519, 0.182099, -0.138889, 0.092593, -0.077160, 0.194444, -0.148148, 0.123457, -0.111111}, []float32{-0.037037, -0.074074, 0.111111, 0.222222, 0.150327, -0.137255, -0.148148, -0.178649, 0.032680}}
	b3Array = [][]float32{[]float32{-1.000000, 0.000000, -1.000000, 4.000000, 0.000000, 9.000000, 0.000000, -9.000000, -6.000000}, []float32{6.000000, -5.000000, -6.000000, -5.000000, 7.000000, -1.000000, -7.000000, -7.000000, 4.000000}, []float32{-5.000000, 3.000000, 1.000000, 3.000000, 5.000000, -4.000000, -1.000000, 0.000000, -4.000000}, []float32{-7.000000, -3.000000, -9.000000, 8.000000, -9.000000, 1.000000, -3.000000, -4.000000, -8.000000}, []float32{6.000000, -4.000000, 4.000000, -5.000000, 6.000000, 3.000000, 7.000000, -9.000000, -4.000000}, []float32{-9.000000, 0.000000, 0.000000, 4.000000, -4.000000, -2.000000, 2.000000, -8.000000, -2.000000}, []float32{8.000000, -6.000000, 6.000000, 6.000000, 5.000000, 9.000000, 3.000000, 8.000000, 6.000000}, []float32{-7.000000, -3.000000, 4.000000, 1.000000, 1.000000, 9.000000, 4.000000, 1.000000, -4.000000}, []float32{6.000000, -6.000000, -3.000000, 8.000000, -9.000000, -1.000000, -1.000000, -1.000000, -8.000000}, []float32{-8.000000, 0.000000, -7.000000, 6.000000, -3.000000, 5.000000, 0.000000, -3.000000, -4.000000}, []float32{1.000000, 9.000000, 9.000000, -6.000000, 4.000000, -1.000000, 0.000000, 5.000000, -1.000000}, []float32{-2.000000, -9.000000, -7.000000, -4.000000, 1.000000, -3.000000, 8.000000, -6.000000, -9.000000}, []float32{-6.000000, 9.000000, 7.000000, -3.000000, 3.000000, 8.000000, 8.000000, 9.000000, -7.000000}, []float32{-8.000000, -2.000000, -3.000000, 7.000000, 4.000000, 7.000000, -8.000000, 6.000000, -3.000000}, []float32{-2.000000, 8.000000, 0.000000, -6.000000, 0.000000, -5.000000, 4.000000, -6.000000, -9.000000}, []float32{-6.000000, -5.000000, -3.000000, -7.000000, -8.000000, 1.000000, -2.000000, 3.000000, -6.000000}, []float32{0.000000, 5.000000, 1.000000, -4.000000, 1.000000, 8.000000, 7.000000, -8.000000, 8.000000}, []float32{-2.000000, 4.000000, 3.000000, -3.000000, -5.000000, 1.000000, -1.000000, -9.000000, -9.000000}, []float32{9.000000, -1.000000, 9.000000, -8.000000, -1.000000, 6.000000, 1.000000, 3.000000, 0.000000}, []float32{-4.000000, 4.000000, -2.000000, 9.000000, 7.000000, 2.000000, -7.000000, 0.000000, -2.000000}, []float32{1.000000, -2.000000, 8.000000, -5.000000, -8.000000, -4.000000, -8.000000, 8.000000, 4.000000}, []float32{7.000000, 2.000000, -1.000000, 6.000000, 5.000000, 5.000000, 4.000000, -8.000000, 3.000000}, []float32{4.000000, 0.000000, 2.000000, 9.000000, -2.000000, 5.000000, -7.000000, -1.000000, 3.000000}, []float32{-8.000000, -4.000000, -6.000000, 7.000000, 7.000000, 3.000000, -1.000000, 8.000000, 9.000000}, []float32{3.000000, 3.000000, -5.000000, -9.000000, 8.000000, 6.000000, -9.000000, 8.000000, -5.000000}, []float32{3.000000, -1.000000, 3.000000, -1.000000, 0.000000, 4.000000, -7.000000, 2.000000, -8.000000}, []float32{1.000000, 5.000000, -8.000000, 1.000000, -6.000000, 4.000000, -4.000000, -5.000000, -1.000000}, []float32{-9.000000, 1.000000, -4.000000, 9.000000, 5.000000, -5.000000, 3.000000, 8.000000, 1.000000}, []float32{0.000000, -1.000000, 7.000000, 5.000000, 9.000000, 5.000000, 1.000000, 0.000000, -6.000000}, []float32{-9.000000, -6.000000, 5.000000, -1.000000, 3.000000, -9.000000, 4.000000, 8.000000, 3.000000}}
	bI3Array = [][]float32{[]float32{-1.800000, -0.200000, 0.000000, -0.533333, -0.133333, -0.111111, 0.800000, 0.200000, 0.000000}, []float32{-0.040936, -0.120858, -0.091618, -0.052632, 0.035088, -0.070175, -0.163743, -0.150097, -0.033138}, []float32{-0.130719, 0.078431, -0.111111, 0.104575, 0.137255, -0.111111, 0.032680, -0.019608, -0.222222}, []float32{-0.413043, -0.065217, 0.456522, -0.331522, -0.157609, 0.353261, 0.320652, 0.103261, -0.472826}, []float32{0.115385, -2.000000, -1.384615, 0.038462, -2.000000, -1.461538, 0.115385, 1.000000, 0.615385}, []float32{-0.111111, -0.000000, -0.000000, 0.055556, 0.250000, -0.250000, -0.333333, -1.000000, 0.500000}, []float32{0.500000, -1.000000, 1.000000, 0.107143, -0.357143, 0.428571, -0.392857, 0.976190, -0.904762}, []float32{0.317073, 0.195122, 0.756098, -0.975610, -0.292683, -1.634146, 0.073171, 0.121951, 0.097561}, []float32{0.816092, -0.517241, -0.241379, 0.747126, -0.586207, -0.206897, -0.195402, 0.137931, -0.068966}, []float32{-0.300000, -0.233333, 0.233333, -0.266667, -0.355556, 0.022222, 0.200000, 0.266667, -0.266667}, []float32{-0.003096, -0.167183, 0.139319, 0.018576, 0.003096, 0.164087, 0.092879, 0.015480, -0.179567}, []float32{-0.056017, -0.080913, 0.070539, -0.124481, 0.153527, 0.045643, 0.033195, -0.174274, -0.078838}, []float32{-0.158163, 0.214286, 0.086735, 0.073129, -0.023810, 0.045918, -0.086735, 0.214286, 0.015306}, []float32{-0.192857, -0.085714, -0.007143, -0.125000, 0.000000, 0.125000, 0.264286, 0.228571, -0.064286}, []float32{0.056391, -0.135338, 0.075188, 0.139098, -0.033835, 0.018797, -0.067669, -0.037594, -0.090226}, []float32{0.737705, -0.639344, -0.475410, -0.721311, 0.491803, 0.442623, -0.606557, 0.459016, 0.213115}, []float32{0.154839, -0.103226, 0.083871, 0.189247, -0.015054, -0.008602, 0.053763, 0.075269, 0.043011}, []float32{-0.350649, -0.058442, -0.123377, 0.181818, -0.136364, 0.045455, -0.142857, 0.142857, -0.142857}, []float32{0.048000, -0.072000, -0.008000, -0.016000, 0.024000, 0.336000, 0.061333, 0.074667, 0.045333}, []float32{0.538462, -0.307692, -0.846154, -0.153846, 0.230769, 0.384615, -1.884615, 1.076923, 2.461538}, []float32{-0.000000, -0.076923, -0.076923, -0.055556, -0.072650, 0.038462, 0.111111, -0.008547, 0.019231}, []float32{0.120350, 0.004376, 0.032823, 0.004376, 0.054705, -0.089716, -0.148796, 0.140044, 0.050328}, []float32{0.020000, 0.040000, -0.080000, 1.240000, -0.520000, 0.040000, 0.460000, -0.080000, 0.160000}, []float32{-0.091549, 0.028169, -0.070423, 0.154930, 0.183099, 0.042254, -0.147887, -0.159624, 0.065728}, []float32{0.156863, 0.044563, -0.103387, 0.176471, 0.106952, -0.048128, -0.000000, 0.090909, -0.090909}, []float32{-1.333333, -0.333333, -0.666667, -6.000000, -0.500000, -2.500000, -0.333333, 0.166667, -0.166667}, []float32{0.142077, 0.245902, -0.153005, -0.081967, -0.180328, -0.065574, -0.158470, -0.081967, -0.060109}, []float32{-0.068493, 0.050228, -0.022831, 0.036530, -0.004566, 0.123288, -0.086758, -0.114155, 0.082192}, []float32{0.551020, 0.061224, 0.693878, -0.357143, 0.071429, -0.357143, 0.091837, 0.010204, -0.051020}, []float32{-0.128368, -0.091918, -0.061807, 0.052298, 0.074485, 0.136292, 0.031696, -0.076070, 0.052298}}
	p3Array = [][]float32{[]float32{-8.000000, 27.000000, 5.000000, 38.000000, -72.000000, 30.000000, -31.000000, -36.000000, -95.000000}, []float32{-49.000000, 112.000000, 19.000000, -65.000000, 27.000000, 32.000000, -21.000000, 26.000000, 3.000000}, []float32{28.000000, 29.000000, -42.000000, -42.000000, 27.000000, 21.000000, -19.000000, 59.000000, -20.000000}, []float32{105.000000, -30.000000, 98.000000, -47.000000, 14.000000, -18.000000, 1.000000, 134.000000, 118.000000}, []float32{54.000000, -33.000000, 35.000000, 109.000000, -117.000000, -18.000000, -4.000000, 32.000000, 54.000000}, []float32{56.000000, -44.000000, -20.000000, 7.000000, 56.000000, 22.000000, 93.000000, -60.000000, -14.000000}, []float32{-43.000000, -136.000000, -102.000000, -105.000000, 32.000000, -102.000000, 56.000000, 116.000000, 108.000000}, []float32{-87.000000, -34.000000, 33.000000, -56.000000, -23.000000, 18.000000, 31.000000, 13.000000, -84.000000}, []float32{17.000000, -42.000000, -64.000000, 29.000000, -51.000000, -79.000000, 75.000000, -102.000000, -87.000000}, []float32{-2.000000, -36.000000, -16.000000, -8.000000, -6.000000, 15.000000, -98.000000, 18.000000, -88.000000}, []float32{-42.000000, -60.000000, -54.000000, -13.000000, -39.000000, -68.000000, 45.000000, 102.000000, 78.000000}, []float32{-32.000000, -49.000000, -57.000000, -20.000000, -9.000000, 53.000000, -42.000000, 29.000000, 68.000000}, []float32{46.000000, 87.000000, 1.000000, -79.000000, -54.000000, 87.000000, 50.000000, 9.000000, -49.000000}, []float32{-11.000000, -62.000000, -26.000000, -27.000000, 8.000000, 3.000000, -40.000000, 56.000000, 20.000000}, []float32{20.000000, 32.000000, 112.000000, -10.000000, 22.000000, 96.000000, 14.000000, -74.000000, -37.000000}, []float32{53.000000, 51.000000, 17.000000, 51.000000, 53.000000, -1.000000, 18.000000, -31.000000, 67.000000}, []float32{28.000000, -47.000000, 122.000000, 30.000000, -10.000000, -14.000000, -95.000000, 105.000000, -3.000000}, []float32{4.000000, -74.000000, -41.000000, -15.000000, -47.000000, -23.000000, -1.000000, 101.000000, 96.000000}, []float32{-50.000000, 18.000000, -54.000000, -110.000000, 9.000000, -42.000000, -22.000000, 18.000000, -27.000000}, []float32{46.000000, 68.000000, 6.000000, 19.000000, -43.000000, 10.000000, 27.000000, -51.000000, 14.000000}, []float32{27.000000, 92.000000, 28.000000, 71.000000, -22.000000, 16.000000, 56.000000, -18.000000, 0.000000}, []float32{-70.000000, 40.000000, -42.000000, -38.000000, -54.000000, 8.000000, -34.000000, -13.000000, -1.000000}, []float32{-103.000000, -1.000000, 5.000000, 114.000000, -15.000000, 42.000000, 6.000000, 7.000000, -30.000000}, []float32{33.000000, 71.000000, 81.000000, 52.000000, 70.000000, 102.000000, -8.000000, -57.000000, -48.000000}, []float32{-12.000000, 22.000000, 13.000000, 48.000000, -71.000000, 60.000000, 102.000000, -51.000000, -1.000000}, []float32{-63.000000, 18.000000, -48.000000, -53.000000, 14.000000, -22.000000, 14.000000, -5.000000, 43.000000}, []float32{41.000000, 4.000000, 41.000000, 49.000000, 33.000000, -11.000000, 10.000000, -32.000000, 76.000000}, []float32{-36.000000, 6.000000, 54.000000, 174.000000, 87.000000, 3.000000, -75.000000, -25.000000, -32.000000}, []float32{-13.000000, -14.000000, 78.000000, 33.000000, 48.000000, 54.000000, 45.000000, 72.000000, 10.000000}, []float32{-77.000000, -6.000000, -18.000000, 25.000000, -57.000000, 6.000000, -90.000000, -94.000000, 43.000000}}
	pI3Array = [][]float32{[]float32{-0.880000, -0.265000, -0.130000, -0.297778, -0.101667, -0.047778, 0.400000, 0.125000, 0.050000}, []float32{0.019783, -0.004162, -0.080897, 0.012565, -0.006638, -0.008772, 0.029582, 0.028397, -0.156920}, []float32{0.553688, 0.590725, -0.542484, 0.385621, 0.422658, -0.366013, 0.611578, 0.685652, -0.614379}, []float32{-0.011362, -0.046609, 0.002326, -0.015454, -0.034364, 0.007593, 0.017646, 0.039419, -0.000168}, []float32{0.459140, -0.232049, -0.374940, 0.464897, -0.244363, -0.382776, -0.241484, 0.127619, 0.217576}, []float32{0.011111, 0.012106, 0.003151, 0.044444, 0.022305, -0.028441, -0.116667, -0.015174, 0.071393}, []float32{6.500000, 1.214286, 7.285714, 2.392857, 0.454082, 2.688776, -5.940476, -1.117347, -6.656463}, []float32{-0.233981, 0.334436, -0.020256, 0.571310, -0.866060, 0.038859, 0.002067, -0.010610, -0.013366}, []float32{-0.140137, 0.111227, 0.002090, -0.131661, 0.128527, -0.019854, 0.033554, -0.054801, 0.013584}, []float32{0.002884, -0.038632, -0.007109, -0.024301, -0.015560, 0.001766, -0.008182, 0.039839, -0.003085}, []float32{-0.060888, 0.012947, -0.030866, 0.031992, 0.013228, 0.033680, -0.006708, -0.024768, -0.013416}, []float32{-0.012961, 0.010126, -0.018757, -0.005223, -0.027562, 0.017104, -0.005778, 0.018009, -0.004174}, []float32{0.014402, 0.033024, 0.058929, 0.003703, -0.017811, -0.031548, 0.015376, 0.030427, 0.033929}, []float32{-0.001587, -0.042857, 0.004365, 0.083333, -0.250000, 0.145833, -0.236508, 0.614286, -0.349603}, []float32{0.030630, -0.034594, 0.002961, 0.004743, -0.011239, -0.014804, 0.002104, 0.009389, 0.003701}, []float32{-0.110971, 0.124338, 0.030013, 0.108291, -0.102301, -0.029004, 0.079918, -0.080738, -0.006557}, []float32{0.006156, 0.051995, 0.007707, 0.005828, 0.047222, 0.016630, 0.009029, 0.006259, 0.004638}, []float32{0.041321, -0.055931, 0.004247, -0.027616, -0.006475, -0.013346, 0.029485, 0.006229, 0.024502}, []float32{0.015200, -0.014400, -0.008000, -0.060622, 0.004800, 0.113778, -0.052800, 0.014933, 0.045333}, []float32{0.032764, 0.448006, -0.334046, -0.001425, -0.171652, 0.123219, -0.068376, -1.489316, 1.164530}, []float32{0.003239, -0.005668, 0.023482, 0.010076, -0.017634, 0.017499, -0.000517, 0.063405, -0.080139}, []float32{0.004802, 0.017809, -0.059203, -0.009421, -0.041272, 0.065524, -0.040785, -0.068989, 0.161075}, []float32{-0.010196, -0.000327, -0.002157, -0.240000, -0.200000, -0.320000, -0.058039, -0.046732, -0.108431}, []float32{0.443120, -0.218310, 0.283857, 0.303359, -0.169014, 0.152763, -0.434092, 0.237089, -0.249549}, []float32{0.019514, -0.003995, 0.013980, 0.038443, -0.008190, 0.008377, 0.029879, 0.010172, -0.001271}, []float32{0.420513, -0.456410, 0.235897, 1.684615, -1.741026, 0.989744, 0.058974, -0.053846, 0.061538}, []float32{-0.453132, 0.339639, 0.293611, 0.805801, -0.568726, -0.517024, 0.398907, -0.284153, -0.243169}, []float32{-0.010966, -0.004688, -0.018945, 0.021629, 0.021058, 0.038473, 0.008805, -0.005465, -0.016905}, []float32{-0.107332, 0.181280, -0.141723, 0.066138, -0.114638, 0.103175, 0.006803, 0.009637, -0.005102}, []float32{-0.006515, 0.006733, -0.003667, -0.005576, -0.017025, 0.000041, -0.025826, -0.023126, 0.015672}}

	va3Array = [][]float32{[]float32{7, -9, -2, 5, 6, 6, -9, 7, 8}, []float32{-7, -5, 8, 2, -6, 5, 9, 1, -8}, []float32{5, 3, -8, -8, 1, 1, -7, 0, -8}, []float32{-9, -9, -1, -9, -2, -4, 5, 0, 8}, []float32{-5, -1, -3, -2, 6, -7, -8, -5, 8}, []float32{9, 9, -4, -3, 7, 6, 5, -3, 5}, []float32{0, -6, 9, -2, -5, 3, -8, -5, 9}, []float32{4, -1, 6, -2, -5, -3, 9, -3, -4}, []float32{-3, 1, 8, 6, 7, -5, -7, 6, -9}, []float32{-6, 2, -2, 5, -7, 8, 0, 7, -8}, []float32{-6, 2, 4, -4, 6, -9, 0, -4, -4}, []float32{9, 8, 3, 5, 2, 3, 3, -5, 7}, []float32{4, 0, -7, -8, 1, -5, 1, -9, 1}, []float32{5, 4, -6, 2, -1, 6, 1, -2, -2}, []float32{-1, -1, -1, 2, 2, -9, 5, 8, -2}, []float32{-8, 9, 1, -1, -6, -3, -7, -4, 3}, []float32{-8, -4, -7, 6, 2, -2, 8, -5, 9}, []float32{-4, 5, 5, 5, 3, 8, -3, 7, 4}, []float32{-3, -9, 5, -5, -4, -1, 0, 1, -7}, []float32{0, -6, 8, 8, -8, -2, -8, -4, -5}, []float32{-9, -7, -1, 4, 6, -3, 2, -2, 8}, []float32{-6, -9, 2, -6, 7, -9, 5, 9, 9}, []float32{-3, 6, -7, 8, 6, -4, -5, -5, -6}, []float32{-7, -4, -4, -7, -4, 1, 0, 8, 7}, []float32{-1, -6, -1, -4, 6, 7, -8, -6, 2}, []float32{8, -3, 7, 0, 0, -3, 3, -3, -7}, []float32{-5, -6, 3, 3, -2, -7, 2, -9, -2}, []float32{-6, 3, -5, 2, 7, 6, 8, -3, 1}, []float32{-8, 3, 4, -7, 8, 2, 8, 4, 6}, []float32{9, -1, -7, -9, 0, -5, -7, -1, 6}}
	v03Array = [][]float32{[]float32{6, -9, 6}, []float32{1, 6, 5}, []float32{4, 6, -8}, []float32{2, 8, 6}, []float32{5, -7, 8}, []float32{4, 1, -1}, []float32{-9, 1, 9}, []float32{0, 4, 7}, []float32{-6, 3, 2}, []float32{-1, -6, 1}, []float32{-3, 3, -3}, []float32{9, -7, 6}, []float32{-6, 4, 3}, []float32{2, 5, -7}, []float32{-2, -6, 1}, []float32{-5, 4, 5}, []float32{1, -6, -4}, []float32{4, 5, -2}, []float32{-5, 7, -4}, []float32{8, -1, 0}, []float32{-6, -2, 3}, []float32{-3, -1, -6}, []float32{4, -6, 0}, []float32{-7, 4, -3}, []float32{-1, -9, 1}, []float32{-3, -9, -5}, []float32{-6, -1, -3}, []float32{3, 5, -6}, []float32{9, -1, -3}, []float32{-6, -8, -5}}
	v13Array = [][]float32{[]float32{-57, -66, -18}, []float32{50, -36, -2}, []float32{28, 18, 38}, []float32{-60, -34, 14}, []float32{-75, -87, 98}, []float32{28, 46, -15}, []float32{-74, 4, 3}, []float32{55, -41, -40}, []float32{22, 27, -81}, []float32{-24, 47, -54}, []float32{6, 24, -27}, []float32{64, 28, 48}, []float32{-53, -23, 25}, []float32{13, 17, 32}, []float32{-5, -2, 54}, []float32{1, -89, -2}, []float32{-76, 4, -31}, []float32{15, 21, 52}, []float32{-20, 13, -4}, []float32{-8, -40, 66}, []float32{52, 24, 36}, []float32{-6, -34, -51}, []float32{-60, -12, -4}, []float32{21, -12, 11}, []float32{29, -54, -60}, []float32{-39, 24, 41}, []float32{21, 65, -5}, []float32{-56, 62, 9}, []float32{-89, 7, 16}, []float32{53, 11, 52}}
}

// 
func makeA3() Matrix3 {
	return MakeMatrix3([]float32{1, 2, 3, 4, 5, -6, 7, 8, 9}, true)
}

// Symmetrical B.Transpose()==B
func makeB3() Matrix3 {
	return MakeMatrix3([]float32{1, 0, 1, 0, 0, 1, 1, 1, 0}, true)
}

func makeC3() Matrix3 {
	return MakeMatrix3([]float32{0, 1, 1, 1, 0, 0, 0, 0, 0}, true)
}

// Symmetrical D.Transpose()==D
func makeD3() Matrix3 {
	return MakeMatrix3([]float32{1, 1, 1, 1, 1, 1, 1, 1, 1}, true)
}

func makeM2() Matrix3 {
	return MakeMatrix3([]float32{1, 4, 7, 2, 5, -8, 3, 6, 9}, true)
}

func makeM4() Matrix3 {
	return MakeMatrix3([]float32{10, 4, 8, 1, 5, 9, 2, 6, 10}, true)
}

func makeM2M4() Matrix3 {
	return MakeMatrix3([]float32{18, 32, 56, 57, 77, 137, 80, 42, 74}, true)
}

func makeM4M2() Matrix3 {
	return MakeMatrix3([]float32{82, -24, 126, 84, -45, 114, 96, -46, 132}, true)
}

func Test_3Approximates1(t *testing.T) {
	A := makeA3()
	B := makeB3()
	C := makeC3()
	if A.ApproxEquals(B, ε3d) {
		t.Fail()
	}
	if A.ApproxEquals(C, ε3d) {
		t.Fail()
	}
	if C.ApproxEquals(B, ε3d) {
		t.Fail()
	}
}

func Test_3Approximates2(t *testing.T) {
	B := makeB3()
	b := makeB3().ScalarMultiply(1.0 + ε3d/2)
	C := makeC3()
	c := makeC3().ScalarMultiply(1.0 + ε3d/2)
	D := makeD3()
	d := makeD3().ScalarMultiply(1.0 + ε3d/2)

	if !B.ApproxEquals(b, ε3d) {
		fmt.Println(B)
		fmt.Println(b)
		t.Fail()
	}
	if !C.ApproxEquals(c, ε3d) {
		fmt.Println(C)
		fmt.Println(c)
		t.Fail()
	}
	if !D.ApproxEquals(d, ε3d) {
		fmt.Println(D)
		fmt.Println(d)
		t.Fail()
	}
}

func Test_3Transpose1(t *testing.T) {
	// Testing Transpose(), not TransposeThis()
	// Transpose does not change the object (const operation)
	A := makeA3() // symmetrical
	a := A.Transpose()
	B := makeB3()
	b := B.Transpose()
	C := makeC3() // Symmetrical
	c := C.Transpose()
	D := makeD3()
	d := D.Transpose()

	if A.Equals(a) {
		fmt.Println("A", A)
		fmt.Println("a", a)
		t.Fail()
	}
	if !B.Equals(b) {
		// Symmetrical
		fmt.Println("B", B)
		fmt.Println("b", b)
		t.Fail()
	}
	if C.Equals(c) {
		fmt.Println("C", C)
		fmt.Println("c", c)
		t.Fail()
	}
	if !D.Equals(d) {
		// Symmetrical
		fmt.Println("D", D)
		fmt.Println("d", d)
		t.Fail()
	}
}

func Test_3Transpose2(t *testing.T) {

	// Testing TransposeThis()
	// the operation should change the object, making A and a identical
	A := makeA3()
	a := A.Transpose()
	B := makeB3()
	b := B.Transpose()
	C := makeC3()
	c := C.Transpose()
	D := makeD3()
	d := D.Transpose()
	
	A.TransposeThis()
	B.TransposeThis()
	C.TransposeThis()
	D.TransposeThis()

	if !A.Equals(a) {
		fmt.Println("A", A)
		fmt.Println("a", a)
		t.Fail()
	}
	if !B.Equals(b) {
		// Symmetrical
		fmt.Println("B", B)
		fmt.Println("b", b)
		t.Fail()
	}
	if !C.Equals(c) {
		fmt.Println("C", C)
		fmt.Println("c", c)
		t.Fail()
	}
	if !D.Equals(d) {
		fmt.Println("D", D)
		fmt.Println("d", d)
		t.Fail()
	}
}

func Test_3Inv1(t *testing.T) {
	for i := 0; i < len(a3Array); i++ {
		A := MakeMatrix3(a3Array[i], true)
		B := MakeMatrix3(b3Array[i], true)
		AI := MakeMatrix3(aI3Array[i], true)
		BI := MakeMatrix3(bI3Array[i], true)
		ai := A.Inverse()
		bi := B.Inverse()
		if !AI.ApproxEquals(ai, ε3d) {
			fmt.Println("failed to inverse 'A'. Expected result as 'AI', got 'ai'")
			fmt.Println("A	=", A)
			fmt.Println("AI	=", AI)
			fmt.Println("ai	=", ai)
			fmt.Println()
			t.Fail()
		}
		if !BI.ApproxEquals(bi, ε3d) {
			fmt.Println("failed to inverse 'B'. Expected result as 'BI', got 'bi'")
			fmt.Println("B	=", B)
			fmt.Println("BI	=", BI)
			fmt.Println("bi	=", bi)
			fmt.Println()
			t.Fail()
		}
	}
}

func Test_3Transpose3(t *testing.T) {
	for i := 0; i < len(a3Array); i++ {
		A := MakeMatrix3(a3Array[i], true)
		at := A.Transpose()
		B := MakeMatrix3(b3Array[i], true)
		bt := B.Transpose()
		btat := bt.RightMultiply(at)
		P := MakeMatrix3(p3Array[i], true)
		pt := P.Transpose()
		if !pt.ApproxEquals(btat, ε3d) {
			fmt.Println("failed to transpose and multiply 'Bt*At'. Expected result as 'Pt', got 'btat'")
			fmt.Println("A	=", A)
			fmt.Println("B	=", B)
			fmt.Println("Pt	=", pt)
			fmt.Println("btat=", btat)
			fmt.Println()
			t.Fail()
		}
	}
}

func Test_3Mutiply1(t *testing.T) {
	for i := 0; i < len(a3Array); i++ {
		A := MakeMatrix3(a3Array[i], true)
		B := MakeMatrix3(b3Array[i], true)
		P := MakeMatrix3(p3Array[i], true)
		p := A.RightMultiply(B)
		if !P.ApproxEquals(p, ε3d) {
			fmt.Println("failed to multiply 'A*B'. Expected result as 'P', got 'p'")
			fmt.Println("A	=", A)
			fmt.Println("B	=", B)
			fmt.Println("P	=", P)
			fmt.Println("p	=", p)
			fmt.Println()
			t.Fail()
		}
	}
}

func Test_3Mutiply2(t *testing.T) {

	I := MakeMatrix3Identity()

	A := makeA3()
	ai := A.Inverse()
	da := A.Determinant()
	B := makeB3()
	bi := makeB3().Inverse()
	C := makeC3()
	ci := makeC3().Inverse()

	a := A.RightMultiply(ai)
	if !I.ApproxEquals(a, ε3d) {
		fmt.Println("a=", A)
		fmt.Println("det(a)=", da)
		fmt.Println("ai=", A)
		fmt.Println("I=", a)
		t.Fail()
	}
	b := B.RightMultiply(bi)
	if !I.ApproxEquals(b, ε3d) {
		fmt.Println("b", b)
		fmt.Println("I", I)
		t.Fail()
	}
	c := C.RightMultiply(ci)
	if !I.ApproxEquals(c, ε3d) {
		fmt.Println("c", c)
		fmt.Println("i", I)
		t.Fail()
	}
}

func Test_3Mutiply3(t *testing.T) {

	A := makeA3()
	da := A.Determinant()
	B := makeB3()
	bi := makeB3().Inverse()
	C := makeC3()
	ci := makeC3().Inverse()

	a := A.RightMultiply(B).RightMultiply(C).RightMultiply(ci).RightMultiply(bi)
	if !A.ApproxEquals(a, ε3d) {
		fmt.Println("a=", A)
		fmt.Println("det(a)=", da)
		fmt.Println("ai=", A)
		fmt.Println("I=", a)
		t.Fail()
	}
}

func Test_3Transpose4(t *testing.T) {
	for i := 0; i < len(a3Array); i++ {
		A := MakeMatrix3(a3Array[i], true)
		att := A.Transpose().Transpose()
		if !att.Equals(A) {
			fmt.Println("failed to double transpose 'A'. Expected result as 'A', got 'att' (!Equal)")
			fmt.Println("A	=", A)
			fmt.Println("att=", att)
			fmt.Println()
			t.Fail()
		}
	}
}

func Test_3VectorMul_1(t *testing.T) {
	for i := 0; i < len(v03Array); i++ {
		V0 := MakeVector3(v03Array[i])
		M := MakeMatrix3(va3Array[i], false)
		V1 := MakeVector3(v13Array[i])
		v1 := M.MultiplyV(V0)

		if !v1.ApproxEquals(V1, ε3d) {
			fmt.Println("failed to multiply 'V0' with 'M', Expected result as 'V1', got 'v1'")
			fmt.Println("V0	=", V0)
			fmt.Println("M	=", M)
			fmt.Println("V1	=", V1)
			fmt.Println("v1	=", v1)
			fmt.Println()
			t.Fail()
		}
	}
}

func Test_3VectorMul_2(t *testing.T) {
	for i := 0; i < len(v03Array); i++ {
		V0 := MakeVector3(v03Array[i])
		v0 := V0.Add(V0).Add(V0).Sub(V0)
		M := MakeMatrix3(va3Array[i], false)
		V1 := MakeVector3(v13Array[i])
		V1 = V1.Add(V1)

		v1 := M.MultiplyV(v0)

		if !v1.ApproxEquals(V1, ε3d) {
			fmt.Println("failed to multiply 'V0' with 'M', Expected result as 'V1', got 'v1'")
			fmt.Println("V0	=", V0)
			fmt.Println("M	=", M)
			fmt.Println("V1	=", V1)
			fmt.Println("v1	=", v1)
			fmt.Println()
			t.Fail()
		}
	}
}

func Test_MinAngle_1(t *testing.T) {
	for i := 0.; i <= 360.; i++ {
		a0 := float32(i) * Deg2Rad
		a1 := (float32(i) + 1) * Deg2Rad
		a2 := (float32(i) - 1) * Deg2Rad

		r1 := MinAngleBetween(a1, a0) * Rad2Deg
		r1e := float32(1.)
		r2 := MinAngleBetween(a2, a0) * Rad2Deg
		r2e := float32(-1.)

		if !ApproxEquals(r1, r1e, ε3d) {
			fmt.Printf("Failed to calculate min angle between %f and %f\n", a0*Rad2Deg, a1*Rad2Deg)
			fmt.Printf(" Expected result: %f , got %f\n\n", r1e, r1)
			t.Fail()
		}
		if !ApproxEquals(r2, r2e, ε3d) {
			fmt.Printf("Failed to calculate min angle between %f and %f\n", a0*Rad2Deg, a2*Rad2Deg)
			fmt.Printf(" Expected result: %f , got %f\n\n", r2e, r2)
			t.Fail()
		}
	}
}

func Test_MinAngle_2(t *testing.T) {
	a1 := float32(180.) * Deg2Rad
	a2 := float32(-180.) * Deg2Rad
	r := MinAngleBetween(a1, a2)
	if !ApproxEquals(r, 0., ε3d) {
		fmt.Printf("Failed to calculate min angle between %f and %f\n", a1*Rad2Deg, a2*Rad2Deg)
		fmt.Printf(" Expected result: %f , got %f\n\n", 0., r)
		t.Fail()
	}
}
