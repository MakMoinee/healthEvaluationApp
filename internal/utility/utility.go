package utility

import (
	log "github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	errorCode = "error_code"
	state     = "state"
)

var AppFailures = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "app_custom_err_code",
		Help: "Application error codes, specifying the failure responses",
	},
	[]string{errorCode},
)

var AppMetric = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ez_shop_service",
		Help: "Custom Metrics",
	},
	[]string{state},
)

func LogPrometheusErrorCustom(errCode string) {
	log.Debug("")
	AppFailures.With(prometheus.Labels{
		errorCode: errCode}).Inc()
}

func LogPrometheusErr(k string, v string) {
	log.Debug(log.WithFields(log.Fields{
		k: v,
	}))
	LogPrometheusErrorCustom(k + v)
}
func LogPrometheusCustom(stateStr string) {
	AppMetric.With(prometheus.Labels{
		state: stateStr}).Inc()
}
func Debug(k string, v string) {
	LogPrometheusCustom(k + v)
}
