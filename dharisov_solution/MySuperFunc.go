package dharisov_solution

import (
	"math"
)

// MySuperFuncImpl MySuperFunc - реализация
// 1. `n==0` -> `x1`
// 2. `n==1` -> `x1 * x2`
// 3. `n>1` -> `f(x1, x2, n-2) * f(x1, x2, n-1)`
func MySuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
	// изначально текст полностью повторяет BasicSuperFuncImpl

    var f_number float64 = GetFibonacciNumber(n)

	return math.Pow(x1, GetFibonacciNumber_OverOne(f_number, n)) * math.Pow(x2, f_number)
}

var golden_ratio float64 = (1 + math.Sqrt(5.00000000))/2

func GetFibonacciNumber(n uint8) float64 {
    var n_as_float64 float64 = float64(n)
    var number = (math.Pow(golden_ratio, n_as_float64) - math.Pow(-golden_ratio, -n_as_float64))/(2*golden_ratio-1)
    return math.Round(number)
}

func GetFibonacciNumber_OverOne(f_number float64, n uint8) float64 {
    if (n%2==0){
        return (f_number + math.Sqrt(5*f_number*f_number + 4))/2
    }else{
        return (f_number + math.Sqrt(5*f_number*f_number - 4))/2
    }
}
