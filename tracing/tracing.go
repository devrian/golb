package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

// InitFromEnv returns an instance of Jaeger Tracer that read from env
// Env example in GKE deployment
//   - name: JAEGER_AGENT_HOST
//     valueFrom:
//     fieldRef:
//     fieldPath: status.hostIP
//   - name: JAEGER_SERVICE_NAME
//     value: service_name
//   - name: JAEGER_SAMPLER_PARAM
//     value: "1"
//   - name: JAEGER_SAMPLER_TYPE
//     value: const
func InitFromEnv(service string) (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	tracer, closer, err := cfg.NewTracer(config.Logger(jaeger.StdLogger))
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}

	return tracer, closer
}
