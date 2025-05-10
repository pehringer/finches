package types

import (
	"math"
)

type (
	Mapping struct {
		Input  float64
		Output float64
	}
)

func (m Mapping) AbsoluteOutput() float64 {
	return math.Abs(m.Output)
}
