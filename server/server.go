package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"saniloader/config"
	"saniloader/metrics"
	"sync"
)

var mu sync.Mutex
var idx int = 0

var ServerConfig config.ConfigType

func addMetrics(currentBackend config.BackendType) {
    metrics.MetricsChannel <- metrics.MetricsChannelType{Name: currentBackend.Name, Data: metrics.MetricsType{Tag: "NumberOfRequests", Value: 1}}
}

func lbHandler(w http.ResponseWriter, r *http.Request) {
    maxLen := len(ServerConfig.Backends)
    mu.Lock()
    currentBackend := ServerConfig.Backends[idx%maxLen]
    fmt.Println(currentBackend.Name, currentBackend.URL)
    targetURL, err := url.Parse(currentBackend.URL)
    if err != nil {
        log.Fatal(err.Error())
    }
    idx++
    mu.Unlock()
    reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
    reverseProxy.ServeHTTP(w, r)
    addMetrics(currentBackend)
}

func Serve() {
	var err error
	ServerConfig, err = config.GetCfg()
    if err != nil {
        log.Fatal(err.Error())
    }

    s := http.Server{
        Addr:    ":" + ServerConfig.Proxy.Port,
        Handler: http.HandlerFunc(lbHandler),
    }
    if err = s.ListenAndServe(); err != nil {
        log.Fatal(err.Error())
    }
}
