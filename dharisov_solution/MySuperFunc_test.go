package dharisov_solution

import "testing"
import "math"
import "github.com/comdiv/task_func_optimize_base_go/basis"

//BenchmarkBasicSuperFuncImpl оставлен для наглядного сравнения производительности
func BenchmarkBasicSuperFuncImpl(b *testing.B) {
	basis.SuperFuncBenchmark(basis.BasicSuperFuncImpl, b)
}

func TestMySuperFuncImpl(t *testing.T) {
	basis.SuperFuncTestCase(MySuperFuncImpl, t)
}

func BenchmarkMySuperFuncImpl(b *testing.B) {
	basis.SuperFuncBenchmark(MySuperFuncImpl, b)
}

func TestPower(t *testing.T){
        var actual float64
        var expected float64
    	for i := 2; i < 15; i++ {
    	    actual = power(float64(i), uint64(i))
    	    expected = math.Pow(float64(i),float64(i))

    		      if actual != expected {
                     t.Errorf("For i: %d, power does not work. Actual: %f. Expected: %f.", i, actual,  expected)
                  }
    	}
}