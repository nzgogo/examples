package main

import (
	"examples/prometheus_custom_metrics/globals"
	"github.com/nzgogo/micro/codec"
	"github.com/nzgogo/micro/constant"
	"github.com/nzgogo/micro/router"
)

func ContextMetricWrapper(handler router.Handler) router.Handler {
	return func(msg *codec.Message, reply string) *router.Error {
		if msg.Type == constant.REQUEST {
			globals.ContextCnt.Inc()
		} else if msg.Type == constant.RESPONSE {
			globals.ContextCnt.Desc()
		}
		err := handler(msg, reply)
		return err
	}
}
