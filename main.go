package main

import (
	"github.com/pehringer/fungen/internal/ga"
)

// Definition:
//   f(x, y) = 0.26 × (x² + y²) − 0.48 × x × y
// Domain:
//   x, y ∈ [−10, 10]
// Global Minimum:
//   f(0, 0) = 0
// Characteristics:
//   Convex, symmetric, and unimodal.
var matyasFunction []ga.Test = []ga.Test{
	{Inputs: []float32{0, 0},	Expected: []float32{0.0}},
	{Inputs: []float32{1, 1},	Expected: []float32{0.04}},
	{Inputs: []float32{-1, -1},	Expected: []float32{0.04}},
	{Inputs: []float32{1, -1},	Expected: []float32{0.52}},
	{Inputs: []float32{-1, 1},	Expected: []float32{0.52}},
	{Inputs: []float32{2, 2},	Expected: []float32{0.16}},
	{Inputs: []float32{-2, -2},	Expected: []float32{0.16}},
	{Inputs: []float32{2, -2},	Expected: []float32{2.08}},
	{Inputs: []float32{-2, 2},	Expected: []float32{2.08}},
	{Inputs: []float32{3, 3},	Expected: []float32{0.36}},
	{Inputs: []float32{-3, -3},	Expected: []float32{0.36}},
	{Inputs: []float32{3, -3},	Expected: []float32{4.68}},
	{Inputs: []float32{-3, 3},	Expected: []float32{4.68}},
	{Inputs: []float32{4, 4},	Expected: []float32{0.64}},
	{Inputs: []float32{-4, -4},	Expected: []float32{0.64}},
	{Inputs: []float32{4, -4},	Expected: []float32{8.32}},
	{Inputs: []float32{-4, 4},	Expected: []float32{8.32}},
	{Inputs: []float32{5, 5},	Expected: []float32{1.0}},
	{Inputs: []float32{-5, -5},	Expected: []float32{1.0}},
	{Inputs: []float32{5, -5},	Expected: []float32{12.0}},
	{Inputs: []float32{-5, 5},	Expected: []float32{12.0}},
	{Inputs: []float32{6, 6},	Expected: []float32{1.44}},
	{Inputs: []float32{-6, -6},	Expected: []float32{1.44}},
	{Inputs: []float32{6, -6},	Expected: []float32{16.64}},
	{Inputs: []float32{-6, 6},	Expected: []float32{16.64}},
	{Inputs: []float32{7, 7},	Expected: []float32{1.96}},
	{Inputs: []float32{-7, -7},	Expected: []float32{1.96}},
	{Inputs: []float32{7, -7},	Expected: []float32{21.92}},
	{Inputs: []float32{-7, 7},	Expected: []float32{21.92}},
	{Inputs: []float32{8, 8},	Expected: []float32{2.56}},
	{Inputs: []float32{-8, -8},	Expected: []float32{2.56}},
	{Inputs: []float32{8, -8},	Expected: []float32{27.84}},
	{Inputs: []float32{-8, 8},	Expected: []float32{27.84}},
}

// Definition:
//   f(x, y) = (x² + y − 11)² + (x + y² − 7)²
// Domain:
//   x, y ∈ [−5, 5]
// Global Minima:
//   Four global minima at:
//     (3.0, 2.0)
//     (−2.805118, 3.131312)
//     (−3.779310, −3.283186)
//     (3.584428, −1.848126)
// Characteristics:
//   Multimodal with multiple global minima.
var himmelblauFunction []ga.Test = []ga.Test{
	{Inputs: []float32{0.0, 0.0},			Expected: []float32{170.0}},
	{Inputs: []float32{1.0, 1.0},			Expected: []float32{106.0}},
	{Inputs: []float32{-1.0, -1.0},			Expected: []float32{218.0}},
	{Inputs: []float32{2.0, 2.0},			Expected: []float32{36.0}},
	{Inputs: []float32{-2.0, -2.0},			Expected: []float32{388.0}},
	{Inputs: []float32{3.0, 2.0},			Expected: []float32{0.0}},
	{Inputs: []float32{-2.805118, 3.131312},	Expected: []float32{0.0}},
	{Inputs: []float32{-3.779310, -3.283186}, 	Expected: []float32{0.0}},
	{Inputs: []float32{3.584428, -1.848126}, 	Expected: []float32{0.0}},
	{Inputs: []float32{5.0, 5.0},			Expected: []float32{250.0}},
	{Inputs: []float32{-5.0, -5.0},			Expected: []float32{890.0}},
	{Inputs: []float32{4.0, 4.0},			Expected: []float32{98.0}},
	{Inputs: []float32{-4.0, -4.0},			Expected: []float32{530.0}},
	{Inputs: []float32{2.0, -2.0},			Expected: []float32{196.0}},
	{Inputs: []float32{-2.0, 2.0},			Expected: []float32{10.0}},
	{Inputs: []float32{1.0, -1.0},			Expected: []float32{122.0}},
	{Inputs: []float32{-1.0, 1.0},			Expected: []float32{122.0}},
	{Inputs: []float32{0.5, 0.5},			Expected: []float32{116.5}},
	{Inputs: []float32{-0.5, -0.5},			Expected: []float32{194.5}},
	{Inputs: []float32{1.5, 1.5},			Expected: []float32{82.25}},
	{Inputs: []float32{-1.5, -1.5},			Expected: []float32{194.25}},
	{Inputs: []float32{2.5, 2.5},			Expected: []float32{20.25}},
	{Inputs: []float32{-2.5, -2.5},			Expected: []float32{324.25}},
	{Inputs: []float32{3.5, 3.5},			Expected: []float32{42.25}},
	{Inputs: []float32{-3.5, -3.5},			Expected: []float32{626.25}},
	{Inputs: []float32{4.5, 4.5},			Expected: []float32{162.25}},
	{Inputs: []float32{-4.5, -4.5},			Expected: []float32{882.25}},
	{Inputs: []float32{5.5, 5.5},			Expected: []float32{338.25}},
	{Inputs: []float32{-5.5, -5.5},			Expected: []float32{1194.25}},
	{Inputs: []float32{6.0, 6.0},			Expected: []float32{450.0}},
	{Inputs: []float32{-6.0, -6.0},			Expected: []float32{1450.0}},
	{Inputs: []float32{3.0, -2.0},			Expected: []float32{10.0}},
	{Inputs: []float32{-3.0, 2.0},			Expected: []float32{10.0}},
}

// Definition:
//   f(x, y) = 2x² − 1.05x⁴ + (x⁶)/6 + x × y + y²
// Domain:
//   x, y ∈ [−5, 5]
// Global Minimum:
//   f(0, 0) = 0
// Characteristics:
//   Non-convex with multiple local minima.
var camelFunction []ga.Test = []ga.Test{
	{Inputs: []float32{ 0.0,  0.0}, Expected: []float32{   0.000000}},
	{Inputs: []float32{ 1.0,  1.0}, Expected: []float32{   3.116667}},
	{Inputs: []float32{ 1.0, -1.0}, Expected: []float32{   1.116667}},
	{Inputs: []float32{-1.0,  1.0}, Expected: []float32{   1.116667}},
	{Inputs: []float32{-1.0, -1.0}, Expected: []float32{   3.116667}},
	{Inputs: []float32{ 2.0,  2.0}, Expected: []float32{   9.866667}},
	{Inputs: []float32{ 2.0, -2.0}, Expected: []float32{   1.866667}},
	{Inputs: []float32{-2.0,  2.0}, Expected: []float32{   1.866667}},
	{Inputs: []float32{-2.0, -2.0}, Expected: []float32{   9.866667}},
	{Inputs: []float32{ 3.0,  3.0}, Expected: []float32{  72.450000}},
	{Inputs: []float32{ 3.0, -3.0}, Expected: []float32{  54.450000}},
	{Inputs: []float32{-3.0,  3.0}, Expected: []float32{  54.450000}},
	{Inputs: []float32{-3.0, -3.0}, Expected: []float32{  72.450000}},
	{Inputs: []float32{ 4.0,  1.0}, Expected: []float32{ 450.866669}},
	{Inputs: []float32{ 1.0,  4.0}, Expected: []float32{  21.116667}},
	{Inputs: []float32{-4.0, -1.0}, Expected: []float32{ 450.866669}},
}

// Definition:
//   f(x, y) = [1 + (x + y + 1)² × (19 − 14x + 3x² − 14y + 6xy + 3y²)]
//   × [30 + (2x − 3y)² × (18 − 32x + 12x² + 48y − 36xy + 27y²)]
// Domain:
//   x, y ∈ [−2, 2]
// Global Minimum:
//   f(0, −1) = 3
// Characteristics:
//   Complex landscape with multiple local minima
var goldsteinPriceFunction []ga.Test = []ga.Test{
	{Inputs: []float32{0.0, -1.0}, Expected: []float32{3.0}},
	{Inputs: []float32{0.0,  2.0}, Expected: []float32{3.0}},
	{Inputs: []float32{2.0, -1.0}, Expected: []float32{3.0}},
	{Inputs: []float32{2.0,  2.0}, Expected: []float32{3.0}},
	{Inputs: []float32{0.0,  0.0}, Expected: []float32{600.0}},
	{Inputs: []float32{1.0,  1.0}, Expected: []float32{1876.0}},
	{Inputs: []float32{-1.0, -1.0}, Expected: []float32{2100.0}},
	{Inputs: []float32{1.0, -1.0}, Expected: []float32{7100.0}},
	{Inputs: []float32{-2.0, -2.0}, Expected: []float32{24376.0}},
	{Inputs: []float32{-2.0,  2.0}, Expected: []float32{956600.0}},
	{Inputs: []float32{2.0, -2.0}, Expected: []float32{316600.0}},
	{Inputs: []float32{2.0,  2.0}, Expected: []float32{76728.0}},
	{Inputs: []float32{2.0,  0.0}, Expected: []float32{1736.0}},
	{Inputs: []float32{0.0,  2.0}, Expected: []float32{224616.0}},
	{Inputs: []float32{-2.0,  0.0}, Expected: []float32{126600.0}},
	{Inputs: []float32{0.0, -2.0}, Expected: []float32{66600.0}},
}


func main() {
	ga.Evolution(himmelblauFunction, 0.5, 12, 48, 4096)
}
