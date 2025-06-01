package types

import (
	"math"
)

type (
	Mapping struct {
		Inputs []float64
		Output float64
	}
)

func (m Mapping) AbsoluteOutput() float64 {
	return math.Abs(m.Output)
}
