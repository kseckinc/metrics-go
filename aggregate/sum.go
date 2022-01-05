package aggregate

func NewSumBuilder() *SumBuilder {
	return &SumBuilder{}
}

type SumBuilder struct {
}

func (builder *SumBuilder) Build() Function {
	return &SumFunction{}
}

type SumFunction struct {
	value float64
}

func (f *SumFunction) Accumulate(value float64) {
	f.value += value
}

func (f *SumFunction) Aggregate() float64 {
	return f.value
}

func (f *SumFunction) Values() []float64 {
	return []float64{float64(f.value)}
}
