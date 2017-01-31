package main
import "math"


func CalcLog(x float64) float64 {
	return math.Log2(x)
}
func CalcLogArg(y string, x float64) float64 {
	if y == "2" {
	return math.Log2(x)
	}
	
	if y == "10" {
	return math.Log10(x)
	}
	return 0
}