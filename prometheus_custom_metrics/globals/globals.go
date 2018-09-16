package globals

import (
	"github.com/nzgogo/micro"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

const (
	SrvName    = "gogo-core-user"
	SrvVersion = "v1"
)

var (
	Service    = gogo.NewService(SrvName, SrvVersion)
	ContextCnt = promauto.NewCounter(prometheus.CounterOpts{
		Name: "user_service_context_total",
		Help: "The total number of contexts in pool",
	})
)
