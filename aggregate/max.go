package aggregate

import (
	"math"
)

func NewMaxBuilder() *MaxBuilder {
	return &MaxBuilder{}
}

type MaxBuilder struct {
}

func (builder *MaxBuilder) Build() Function {
	return &MaxFunction{value: math.MinInt64}
}

type MaxFunction struct {
	value float64
}

func (f *MaxFunction) Accumulate(value float64) {
	if f.value < value {
		f.value = value
	}
}

func (f *MaxFunction) Aggregate() float64 {
	return f.value
}

func (f *MaxFunction) Values() []float64 {
	return []float64{float64(f.value)}
}
