package server

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"sync"
	"saniloader/config"
)

var mu sync.Mutex
var idx int = 0

var cfg config.ConfigType

func lbHandler(w http.ResponseWriter, r *http.Request) {
    maxLen := len(cfg.Backends)
    mu.Lock()
    currentBackend := cfg.Backends[idx%maxLen]
    targetURL, err := url.Parse(currentBackend.URL)
    if err != nil {
        log.Fatal(err.Error())
    }
    idx++
    mu.Unlock()
    reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
    reverseProxy.ServeHTTP(w, r)
}

func Serve() {
	var err error

	cfg, err = config.ReadConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

    s := http.Server{
        Addr:    ":" + cfg.Proxy.Port,
        Handler: http.HandlerFunc(lbHandler),
    }
    if err = s.ListenAndServe(); err != nil {
        log.Fatal(err.Error())
    }
}