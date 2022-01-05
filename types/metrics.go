package types

//Metrics 打点接口
type Metrics interface {
	//With 指定指标的Label名
	With(labels ...string) Series
	//MetricName 获取指标名称
	MetricName() string
}

//Series 是LabelNames/LabelValues组合后产生的唯一线
type Series interface {
	Value(value float64)
}
