package metrics

import (
	"net/http"
	"log"
)

func metricsHttpHandler(w http.ResponseWriter, r *http.Request) {
	
}

func runMetricsServer() {
	s := http.Server{
		Addr: ":" + "9191",
		Handler: http.HandlerFunc(metricsHttpHandler),
	}

	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}