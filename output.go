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

var A float64 = 70.561011
var B float64 = 4.513004
var C float64 = 10.418990
var D float64 = 0.630643
var E float64 = 28.396771
var F float64 = 53.670301
var G float64 = 46.649093
var H float64 = 56.034599
var I float64 = 59.182395
var J float64 = 8.374666
var K float64 = -52.772649
var L float64 = 39.913368
var M float64 = 16.349964
var N float64 = -92.659256
var O float64 = 19.787176
var P float64 = 26.108537

func main() {
	A, _ = strconv.ParseFloat(os.Args[1], 64)
	B, _ = strconv.ParseFloat(os.Args[2], 64)
	C, _ = strconv.ParseFloat(os.Args[3], 64)

	L = C * C
	C = K + H
	D = C - A
	J = divide(I, H)
	P = D - J
	P = L + P
	P = L + P
	P = L + P
	C = B * A
	D = C - M
	L = math.Max(C, D)
	P = L + P
	P = L + P
	L = math.Log(M)
	P = L + P

	fmt.Println(P)
}