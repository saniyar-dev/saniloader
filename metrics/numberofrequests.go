package metrics

var numberOfRequestsMap = make(map[string]int)

func getNumberOfRequests(name string) (int, error) {
	return numberOfRequestsMap[name], nil
}

func increaseNumberOfRequests(name string) error {
	numberOfRequestsMap[name] = numberOfRequestsMap[name] +1
	return nil
}