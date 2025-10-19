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

var A float64 = 26.531112
var B float64 = -80.993321
var C float64 = 89.356283
var D float64 = 4.848743
var E float64 = -23.970255
var F float64 = 34.824811
var G float64 = 28.174037
var H float64 = -41.478763
var I float64 = -57.840115
var J float64 = 43.051414
var K float64 = -61.234842
var L float64 = 99.687047
var M float64 = 79.082284
var N float64 = -41.485818
var O float64 = 72.751574
var P float64 = -3.864708

func main() {
	A, _ = strconv.ParseFloat(os.Args[1], 64)

	I = math.Min(G, E)
	J = A - O
	N = divide(K, A)
	E = math.Acos(K)
	L = I - N
	L = L + D
	B = math.Sin(M)
	F = math.Log(G)
	F = math.Pow(F, L)
	E = math.Min(B, A)
	D = P * I
	L = D + A
	C = float(N > K)
	B = math.Sin(E)
	L = G - L
	N = N + P
	M = math.Sin(L)
	M = math.Asin(M)
	L = F + M
	O = math.Pow(B, C)
	A = B + A
	O = A + O
	E = math.Sin(J)
	J = O * A
	M = C + D
	J = math.Max(A, J)
	E = math.Asin(E)
	D = D + K
	O = math.Cos(O)
	F = math.Max(E, L)
	F = math.Max(F, A)
	L = L - E
	L = L - E
	A = B + A
	H = A + O
	G = math.Sin(O)
	C = K - M
	D = A + D
	E = J + H
	N = math.Min(D, N)
	H = math.Sin(F)
	N = N - F
	F = math.Sin(J)
	O = math.Cos(D)
	K = math.Max(B, L)
	I = math.Sin(C)
	N = N - F
	J = math.Max(G, J)
	B = math.Sin(G)
	D = L + J
	G = math.Asin(O)
	B = divide(I, J)
	O = divide(I, E)
	N = N - H
	A = B + A
	B = math.Sin(M)
	C = divide(D, G)
	A = B + A
	F = C + J
	M = math.Sin(K)
	J = math.Max(C, N)
	G = math.Cos(O)
	B = divide(I, J)
	A = B + A
	I = math.Exp(I)
	A = B + A
	I = math.Max(I, G)
	P = divide(A, I)

	fmt.Println(P)
}