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

var A float64 = -4.563879
var B float64 = -45.027600
var C float64 = 55.134696
var D float64 = 55.084808
var E float64 = 81.762317
var F float64 = 44.051770
var G float64 = 37.212538
var H float64 = 84.682005
var I float64 = -33.626764
var J float64 = 88.412611
var K float64 = -81.315664
var L float64 = 65.917293
var M float64 = -27.110642
var N float64 = -0.922462
var O float64 = 16.447717
var P float64 = 36.550981

func main() {
	A, _ = strconv.ParseFloat(os.Args[1], 64)
	B, _ = strconv.ParseFloat(os.Args[2], 64)
	C, _ = strconv.ParseFloat(os.Args[3], 64)

	H = math.Min(C, E)
	N = C * H
	F = A * B
	B = math.Cos(M)
	G = P - A
	J = G - I
	O = D * E
	B = math.Min(C, C)
	J = J + N
	E = B + C
	A = E * B
	P = math.Cos(O)
	E = math.Sin(P)
	N = E + F
	F = O - N
	J = J + N
	J = J + N
	M = L - J
	P = A - M

	fmt.Println(P)
}