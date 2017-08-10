package abcgo

import (
	"fmt"
	"math"
	"testing"

	"github.com/atgjack/prob"
)

func TestRejectionSampler(t *testing.T) {
	data := []float64{ // true value: 10, 0.3 size: 20
		-0.97282116, 0.49710874, -0.08851066, 0.64683817, 0.16645932,
		-0.98981693, -0.40585841, -0.70108671, 0.25486693, -0.06730565,
		-0.5142657, -0.10126306, 0.67930431, 0.8177273, 0.54188112,
		-0.59839544, 0.27969346, -0.69082318, -0.93346952, 0.68355628,
	}

	generatorF := func(params ...float64) []float64 {
		b, _ := prob.NewBinomial(params[0], params[1])
		simData := make([]float64, 20)
		for i := range simData {
			simData[i] = b.Random()
		}
		return simData
	}

	filterF := func(y, x []float64, epsilon float64) bool {
		var distance, d float64
		for i := range y {
			d = math.Abs(y[i] - x[i])
			distance += math.Pow(d, 2)
		}
		if distance < epsilon {
			return true
		}
		return false
	}

	proposers := Proposers{
		NewNormalProposer(9, 3),
		NewNormalProposer(0, 1),
	}

	params := RejectionSamplerN(100000, data, 10, proposers, generatorF, filterF)

	fmt.Println(len(params))
}
