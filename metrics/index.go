package metrics

import (
	"net/http"
	"log"
)

type MetricsType struct {
	Tag string
	value int
}

type MetricsChannelType struct {
	Name string
	Data MetricsType
}

func ServeMetrics(metricsChannel chan MetricsChannelType) {
	s := http.Server{
			Addr: ":" + "9191",
			Handler: http.HandlerFunc(metricsHttpHandler),
		}

		if err := s.ListenAndServe(); err != nil {
			log.Fatal(err.Error())
		}
}