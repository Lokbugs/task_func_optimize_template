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

	return math.Pow(x1, GetFibonacciNumber_OverOne(f_number, n) - f_number) * math.Pow(x1 * x2, f_number)
}

var golden_ratio float64 = (1 + math.Sqrt(5.00000000))/2

func GetFibonacciNumber(n uint8) uint32 {
    var gr_in_power = math.Pow(golden_ratio, float64(n))
        if (n % 2 == 0){
            return uint32(math.Round((gr_in_power - 1/gr_in_power)/(2*golden_ratio-1)))
        } else {
            return uint32(math.Round((gr_in_power + 1/gr_in_power)/(2*golden_ratio-1)))
        }
}

func GetFibonacciNumber_OverOne(f_number float64, n uint8) uint32 {
    if (n%2==0){
        return uint32((f_number + math.Sqrt(5*f_number*f_number + 4))/2)
    }else{
        return uint32((f_number + math.Sqrt(5*f_number*f_number - 4))/2)
    }
}

func power(number float64, power uint64) float64 {
    var count uint8 = 31
    var not_zero_beat_number uint4 = 32
    for power >> count < 0 {
        not_zero_beat_number -= 1
        count -= 1
    }
}

func get_older_beat(number uint8) uint8 {
    var count uint8 = 31

    for (number >> count) == 0 {
            count -= 1
        }

    return count
}