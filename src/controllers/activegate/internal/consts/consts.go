package consts

const (
	ActiveGateContainerName = "dynatrace-operator" // TODO Does it make sense as the container name? Do we need a link to dynatracev1beta1.OperatorName?

	HttpsServicePortName   = "https"
	HttpsServicePort       = 443
	HttpsServiceTargetPort = "ag-https"
	HttpServicePortName    = "http"
	HttpServicePort        = 80
	HttpServiceTargetPort  = "ag-http"

	EecContainerName = ActiveGateContainerName + "-eec"

	StatsDContainerName    = ActiveGateContainerName + "-statsd"
	StatsDIngestPortName   = "statsd"
	StatsDIngestPort       = 18125
	StatsDIngestTargetPort = "statsd-port"
)
