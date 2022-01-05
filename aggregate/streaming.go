package aggregate

func NewStreamingBuilder() *StreamingBuilder {
	return &StreamingBuilder{}
}

type StreamingBuilder struct {
}

func (builder *StreamingBuilder) Build() Function {
	return &StreamingFunction{}
}

type StreamingFunction struct {
	values []float64
}

func (f *StreamingFunction) Accumulate(value float64) {
	f.values = append(f.values, value)
}

func (f *StreamingFunction) Aggregate() float64 {
	sum := 0.0
	for _, value := range f.values {
		sum += value
	}
	return sum
}

func (f *StreamingFunction) Values() []float64 {
	return f.values
}
