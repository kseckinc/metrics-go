package aggregate

import (
	"math"
)

func NewMinBuilder() *MinBuilder {
	return &MinBuilder{}
}

type MinBuilder struct {
}

func (builder *MinBuilder) Build() Function {
	return &MinFunction{value: math.MaxFloat64}
}

type MinFunction struct {
	value float64
}

func (f *MinFunction) Accumulate(value float64) {
	if f.value > value {
		f.value = value
	}
}

func (f *MinFunction) Aggregate() float64 {
	return f.value
}

func (f *MinFunction) Values() []float64 {
	return []float64{f.value}
}
