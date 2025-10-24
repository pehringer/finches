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

var A float64 = -12.567277
var B float64 = 81.606477
var C float64 = 45.307862
var D float64 = 86.146217
var E float64 = 64.306770
var F float64 = -82.968181
var G float64 = -32.260088
var H float64 = 96.811005
var I float64 = 10.356790
var J float64 = 3.491086
var K float64 = -58.116895
var L float64 = -16.986303
var M float64 = 92.462796
var N float64 = 17.830892
var O float64 = 31.559845
var P float64 = -0.749817

func main() {
	for i := 1; i < len(os.Args); i += 1 {
		A, _ = strconv.ParseFloat(os.Args[i+0], 64)

		E = E + A
		M = math.Sin(D)
		O = E - D
		B = math.Cos(N)
		O = math.Sqrt(A)
		D = math.Min(D, A)
		G = float(A < O)
		A = A - G
		G = math.Sin(G)
		I = D * A
		A = A - G
		M = I + E
		G = C * P
		P = divide(A, G)
		G = math.Sqrt(M)
		E = G * P
		D = E + P
		K = L * P
		P = P + D
		D = math.Min(P, A)
		P = math.Sin(D)
		P = P + K

		fmt.Println(P)
	}
}