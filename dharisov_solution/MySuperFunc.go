package dharisov_solution

import (
	"math"
)

var golden_ratio float64 = (1 + sqrt5)/2
var sqrt5 = math.Sqrt(5)

func MySuperFuncImpl(x1 float64, x2 float64, n uint8) float64 {
	// изначально текст полностью повторяет BasicSuperFuncImpl
    if n == 0 {return x1}
    if n == 1 {return x1 * x2}

    var f_number = GetFibonacciNumber(n)
    var f float64 = float64(f_number)
    var f_number_next uint64
    if n > 2 {
        f_number_next = uint64(f*golden_ratio)
    }else {
        if n%2 == 0 {
            f_number_next = uint64(math.Round((f + math.Sqrt(5*f*f + 4))/2))
        } else {
            f_number_next = uint64(math.Round((f + math.Sqrt(5*f*f - 4))/2))
        }
    }

    return power(x1, (f_number_next - f_number)) * power(x1*x2, f_number)
}

func GetFibonacciNumber(n uint8) uint64 {
    var gr_in_power = power(golden_ratio, uint64(n))

    return uint64(math.Round((gr_in_power)/sqrt5))
}

func power(number float64, power uint64) float64 {
	var result float64 = 1
	var base_coef = number
	var current_power_value = power

	for current_power_value > 0 {
		if current_power_value&1 == 1 {
			result = result * base_coef
		}
		base_coef = base_coef * base_coef

		current_power_value = current_power_value >> 1
	}

	return result
}