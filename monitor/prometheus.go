package monitor

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Prometheus struct {
	Host string `yaml:"Host" json:"Host"`
	Port int    `yaml:"Port" json:"Port"`
	Path string `yaml:"Path" json:"Path"`
}

func RecordMetrics() {
	go func() {
		for {
			OpsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	OpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func Start(addr string, port int, path string) {
	RecordMetrics()
	http.Handle(path, promhttp.Handler())
	addr = fmt.Sprintf("%s:%d", addr, port)
	http.ListenAndServe(addr, nil)
}
