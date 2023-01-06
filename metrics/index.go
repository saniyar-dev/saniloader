package metrics

type MetricsType struct {
	Tag string
	Value int
}

type MetricsChannelType struct {
	Name string
	Data MetricsType
}

var MetricsChannel = make(chan MetricsChannelType)

func RunMetrics() {
	go handleChannel()
	go runMetricsServer()
}
