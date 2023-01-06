package metrics

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
) 

func metricsHttpHandler(w http.ResponseWriter, r *http.Request) {
	name := strings.Trim(r.URL.Path, "/")
	metrics, err := getMetrics(name)
	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := json.Marshal(metrics)
	if err != nil {
		log.Fatal(err.Error())
	}

	w.Header().Set("Content-type", "application/json")
	w.Write(res)
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