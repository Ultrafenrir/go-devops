package handlers

type MetricValue struct {
	val       [8]byte
	isCounter bool
}

type Metric struct {
	MetricType string
	Value      string
}
