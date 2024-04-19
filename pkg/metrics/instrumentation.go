package metrics

import "github.com/prometheus/client_golang/prometheus"

// InstrumentationType is the type of instrumentation the metric is capturing.
// Use this to define your own instrumentation types. e.g.:
//
// const (
//
//	InstrumentationTypeUndefined InstrumentationType = iota
//	InstrumentationTypeRequests
//	InstrumentationTypeInfo
//
// )
type InstrumentationType uint64

// Instrumentation is a collection of metric instrumentations within a namespace
// that can be registered with Prometheus.
type Instrumentation struct {
	namespace     string
	Counters      map[InstrumentationType]prometheus.Counter
	CounterVecs   map[InstrumentationType]*prometheus.CounterVec
	Gauges        map[InstrumentationType]prometheus.Gauge
	GaugeVecs     map[InstrumentationType]*prometheus.GaugeVec
	Histograms    map[InstrumentationType]prometheus.Histogram
	HistogramVecs map[InstrumentationType]*prometheus.HistogramVec
	Summaries     map[InstrumentationType]prometheus.Summary
	SummaryVecs   map[InstrumentationType]*prometheus.SummaryVec
}

// NewInstrumentation creates a new Instrumentation instance with the given namespace.
func NewInstrumentation(namespace string) *Instrumentation {
	return &Instrumentation{
		namespace:     namespace,
		Counters:      make(map[InstrumentationType]prometheus.Counter),
		CounterVecs:   make(map[InstrumentationType]*prometheus.CounterVec),
		Gauges:        make(map[InstrumentationType]prometheus.Gauge),
		GaugeVecs:     make(map[InstrumentationType]*prometheus.GaugeVec),
		Histograms:    make(map[InstrumentationType]prometheus.Histogram),
		HistogramVecs: make(map[InstrumentationType]*prometheus.HistogramVec),
		Summaries:     make(map[InstrumentationType]prometheus.Summary),
		SummaryVecs:   make(map[InstrumentationType]*prometheus.SummaryVec),
	}
}

// WithCounter creates a new Prometheus Counter collector and adds it to the Instrumentation instance.
func (i *Instrumentation) WithCounter(t InstrumentationType, name, help string) *Instrumentation {
	counter := prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: i.namespace,
		Name:      name,
		Help:      help,
	})
	i.Counters[t] = counter
	return i
}

// WithCounterVec creates a new Prometheus CounterVec collector and adds it to the Instrumentation instance.
func (i *Instrumentation) WithCounterVec(t InstrumentationType, name string, help string, labels ...string) *Instrumentation {
	counterVec := prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: i.namespace,
		Name:      name,
		Help:      help,
	}, labels)
	i.CounterVecs[t] = counterVec
	return i
}

// WithGauge creates a new Prometheus Gauge collector and adds it to the Instrumentation instance.
func (i *Instrumentation) WithGauge(t InstrumentationType, name string, help string) *Instrumentation {
	gauge := prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: i.namespace,
		Name:      name,
		Help:      help,
	})
	i.Gauges[t] = gauge
	return i
}

// WithGaugeVec creates a new Prometheus GaugeVec collector and adds it to the Instrumentation instance.
func (i *Instrumentation) WithGaugeVec(t InstrumentationType, name string, help string, labels ...string) *Instrumentation {
	gaugeVec := prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Namespace: i.namespace,
		Name:      name,
		Help:      help,
	}, labels)
	i.GaugeVecs[t] = gaugeVec
	return i
}

// WithHistogram creates a new Prometheus Histogram collector and adds it to the Instrumentation instance.
func (i *Instrumentation) WithHistogram(t InstrumentationType, name string, help string, buckets []float64) *Instrumentation {
	histogram := prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: i.namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	})
	i.Histograms[t] = histogram
	return i
}

// WithHistogramVec creates a new Prometheus HistogramVec collector and adds it to the Instrumentation instance.
func (i *Instrumentation) WithHistogramVec(t InstrumentationType, name string, help string,
	labels []string, buckets []float64) *Instrumentation {
	histogramVec := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: i.namespace,
		Name:      name,
		Help:      help,
		Buckets:   buckets,
	}, labels)
	i.HistogramVecs[t] = histogramVec
	return i
}

// WithSummary creates a new Prometheus Summary collector and adds it to the Instrumentation instance.
func (i *Instrumentation) WithSummary(t InstrumentationType, name, help string, objectives map[float64]float64) *Instrumentation {
	summary := prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace:  i.namespace,
		Name:       name,
		Help:       help,
		Objectives: objectives,
	})
	i.Summaries[t] = summary
	return i
}

// WithSummaryVec creates a new Prometheus SummaryVec collector and adds it to the Instrumentation instance.
func (i *Instrumentation) WithSummaryVec(t InstrumentationType, name, help string,
	objectives map[float64]float64, labels []string) *Instrumentation {
	summaryVec := prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Namespace:  i.namespace,
		Name:       name,
		Help:       help,
		Objectives: objectives,
	}, labels)
	i.SummaryVecs[t] = summaryVec
	return i
}

// Collectors returns all Prometheus collectors that have been added to the Instrumentation instance.
func (i *Instrumentation) Collectors() []prometheus.Collector {
	var collectors []prometheus.Collector
	for _, counter := range i.Counters {
		collectors = append(collectors, counter)
	}
	for _, counterVec := range i.CounterVecs {
		collectors = append(collectors, counterVec)
	}
	for _, gauge := range i.Gauges {
		collectors = append(collectors, gauge)
	}
	for _, gaugeVec := range i.GaugeVecs {
		collectors = append(collectors, gaugeVec)
	}
	for _, histogram := range i.Histograms {
		collectors = append(collectors, histogram)
	}
	for _, histogramVec := range i.HistogramVecs {
		collectors = append(collectors, histogramVec)
	}
	for _, summary := range i.Summaries {
		collectors = append(collectors, summary)
	}
	for _, summaryVec := range i.SummaryVecs {
		collectors = append(collectors, summaryVec)
	}
	return collectors
}
