package router

import (
	"context"
	"net/http"

	"github.com/devrian/golb/monitor"
	"github.com/devrian/golb/response"
	"github.com/felixge/httpsnoop"
	"github.com/julienschmidt/httprouter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type (
	Options struct {
		Prefix  string
		Timeout int
	}

	HttpRouter struct {
		Httprouter     *httprouter.Router
		WrappedHandler http.Handler
		Options        *Options
		tracer         trace.Tracer
	}

	captureConfig struct {
		captureHandler bool
	}

	httpParamsKey struct{}

	captureHandlerKey struct{}

	Handle func(*http.Request) *response.JSONResponse
)

func GetHttpParam(ctx context.Context, name string) string {
	ps := ctx.Value(httpParamsKey{}).(httprouter.Params)
	return ps.ByName(name)
}

func New(o *Options) *HttpRouter {
	router := &HttpRouter{
		Options: o,
		tracer:  otel.Tracer("router/httprouter"),
	}
	router.Httprouter = httprouter.New()
	return router
}

func (mr *HttpRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	captureConf := captureConfig{
		captureHandler: true,
	}
	ctx := context.WithValue(r.Context(), captureHandlerKey{}, &captureConf)

	m := httpsnoop.CaptureMetrics(mr.Httprouter, w, r.WithContext(ctx))

	if captureConf.captureHandler {
		monitor.FeedHTTPMetrics(m.Code, m.Duration, r.Header.Get("routePath"), r.Method)
	}
}
