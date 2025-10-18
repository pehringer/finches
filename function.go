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

var A float64 = 72.272701
var B float64 = 32.765586
var C float64 = 72.773138
var D float64 = -71.790308
var E float64 = 55.892720
var F float64 = -43.801348
var G float64 = -7.714681
var H float64 = 40.435154
var I float64 = 7.663220
var J float64 = -2.946256
var K float64 = 66.987884
var L float64 = -75.982768
var M float64 = 42.344338
var N float64 = 21.162335
var O float64 = 0.384904
var P float64 = 25.689506

func main() {
	A, _ = strconv.ParseFloat(os.Args[1], 64)
	B, _ = strconv.ParseFloat(os.Args[2], 64)
	C, _ = strconv.ParseFloat(os.Args[3], 64)

	L = C * C
	K = math.Exp(N)
	G = divide(H, O)
	O = A * B
	P = C * C
	P = P + O
	E = I - G
	P = P - A
	C = math.Cos(A)
	B = math.Cos(E)
	N = math.Asin(J)
	A = math.Min(B, C)
	P = P + O
	P = P - A
	A = A + A
	A = A - L
	P = P - A
	J = math.Log(D)
	P = P - A

	fmt.Println(P)
}