package metrics

import (
	"net/http"
	"time"
)

type HTTPMiddleware struct {
	instrumentationType InstrumentationType
	instrumentation     *Instrumentation
}

func NewHTTPMiddleware(instrumentationType InstrumentationType, instrumentation *Instrumentation) HTTPMiddleware {
	return HTTPMiddleware{
		instrumentationType: instrumentationType,
		instrumentation:     instrumentation,
	}
}

func (m HTTPMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.instrumentation.CounterVecs[m.instrumentationType].
			WithLabelValues(r.Method, r.URL.Path).Inc()

		start := time.Now()
		next.ServeHTTP(w, r)
		elapsed := time.Since(start).Seconds()

		m.instrumentation.HistogramVecs[m.instrumentationType].
			WithLabelValues(r.Method, r.URL.Path).Observe(elapsed)
	})
}
