package service

import "github.com/prometheus/client_golang/prometheus"

const (
	serviceName = "userd"
	component   = "service"

	loginHandlerLabel    = "login_handler"
	registerHandlerLabel = "register_handler"
	meHandlerLabel       = "me_handler"

	successResultLabel         = "success"
	userNotFoundResultLabel    = "user_not_found"
	invalidPasswordResultLabel = "invalid_password"
	userExistsResultLabel      = "user_exists"
	invalidTokenResultLabel    = "invalid_token"
	missingTokenResultLabel    = "missing_token"
)

var (
	// Metrics for login API

	loginMetrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: serviceName,
		Subsystem: component,
		Name:      loginHandlerLabel,
	}, []string{"result"})

	registerMetrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: serviceName,
		Subsystem: component,
		Name:      registerHandlerLabel,
	}, []string{"result"})

	meMetrics = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: serviceName,
		Subsystem: component,
		Name:      meHandlerLabel,
	}, []string{"result"})
)
