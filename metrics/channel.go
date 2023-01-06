package metrics

func handleChannel() {
	for {
		metric := <- MetricsChannel

		if metric.Data.Tag == "NumberOfRequests" {
			increaseNumberOfRequests(metric.Name)
		}
	}
}