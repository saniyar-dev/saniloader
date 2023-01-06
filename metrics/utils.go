package metrics

func handleChannel() {
	for {
		metric := <- MetricsChannel

		if metric.Data.Tag == "NumberOfRequests" {
			increaseNumberOfRequests(metric.Name)
		}
	}
}

func getMetrics(name string) ([]MetricsType, error) {
	var ans []MetricsType
	norData, err := getNumberOfRequests(name)
	if err != nil {
		return []MetricsType{}, err
	}
	ans = append(ans, MetricsType{Tag: "NumberOfRequests", Value: norData})


	return ans, nil
}