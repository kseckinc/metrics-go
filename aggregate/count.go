package aggregate

func NewCountBuilder() *CountBuilder {
	return &CountBuilder{}
}

type CountBuilder struct {
}

func (builder *CountBuilder) Build() Function {
	return &CountFunction{}
}

type CountFunction struct {
	value int64
}

func (f *CountFunction) Accumulate(value float64) {
	f.value += 1
}

func (f *CountFunction) Aggregate() float64 {
	return float64(f.value)
}
func (f *CountFunction) Values() []float64 {
	return []float64{float64(f.value)}
}
