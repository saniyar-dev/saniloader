package server

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"saniloader/config"
	"sync"
)

var mu sync.Mutex
var idx int = 0

var cfg config.ConfigType
var s http.Server

func lbHandler(w http.ResponseWriter, r *http.Request) {
    maxLen := len(cfg.Backends)
    mu.Lock()
    currentBackend := cfg.Backends[idx%maxLen]
    fmt.Println(currentBackend.Name, currentBackend.URL)
    targetURL, err := url.Parse(currentBackend.URL)
    if err != nil {
        log.Fatal(err.Error())
    }
    idx++
    mu.Unlock()
    reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
    reverseProxy.ServeHTTP(w, r)
}

func Serve(serveCfg config.ConfigType) {
	cfg = serveCfg
	var err error

    s = http.Server{
        Addr:    ":" + cfg.Proxy.Port,
        Handler: http.HandlerFunc(lbHandler),
    }
    if err = s.ListenAndServe(); err != nil {
        log.Fatal(err.Error())
    }
}

func StopServe() error {
    return s.Close()
}