package main
import "os"
import "fmt"
import "math"
import "strconv"

func float(b bool) float64 {
	if b {
		return 1
	}
	return 0
}

func divide(n, d float64) float64 {
	if math.Abs(d) > 1e-9 {
		return n / d
	}
	if math.Abs(n) < 1e-9 {
		return math.NaN()
	}
	if n > 0 {
		return math.Inf(1)
	}
	return math.Inf(-1)
}

var A float64 = -94.033074
var B float64 = -58.953136
var C float64 = -7.262083
var D float64 = 38.998861
var E float64 = -68.250528
var F float64 = -5.966731
var G float64 = 10.523035
var H float64 = -84.791771
var I float64 = 61.754903
var J float64 = -95.718353
var K float64 = -0.434269
var L float64 = -56.628360
var M float64 = -83.967506
var N float64 = -19.120607
var O float64 = 54.137898
var P float64 = 26.635743

func main() {
	for i := 1; i < len(os.Args); i += 3 {

		A, _ = strconv.ParseFloat(os.Args[i+0], 64)
		B, _ = strconv.ParseFloat(os.Args[i+1], 64)
		C, _ = strconv.ParseFloat(os.Args[i+2], 64)

		K = math.Max(L, B)
		K = A * K
		G = N - D
		E = float(I > A)
		F = C * C
		N = F + F
		B = I * I
		A = A - E
		F = N + F
		E = divide(P, M)
		B = math.Min(A, B)
		A = math.Max(J, F)
		P = math.Max(A, E)
		P = P + K
		C = math.Min(D, B)
		P = P + K
		N = P - C
		K = float(K < I)
		C = math.Sin(G)
		P = N - C
		P = P + K
		P = P + K
		P = P + K

	}
	fmt.Println(P)
}