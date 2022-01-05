package aggregate

type Function interface {
	Accumulate(value float64)
	Aggregate() float64
	Values() []float64
}

type Builder interface {
	Build() Function
}
