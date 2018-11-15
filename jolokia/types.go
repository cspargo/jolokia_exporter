package jolokia

import (
	"encoding/json"
	"fmt"
)

// Config is holding a list of metrics that should be exported
type Config struct {
	Metrics []MetricMapping `json:"metrics"`
}

// A MetricMapping is the assignment of a JMX source path to a target prom key name
type MetricMapping struct {
	Source MetricSource `json:"source"`
	Target string       `json:"target"`
}

type MetricTarget struct {
	Url string `json:"url,omitempty"`
}

// MetricSource defines what path the metric should be load from
type MetricSource struct {
	Mbean        string        `json:"mbean"`
	Attribute    string        `json:"attribute"`
	Path         string        `json:"path"`
	MetricTarget *MetricTarget `json:"target",omitempty`
}

// RequestMetric holds the info for jolokia what to export
type RequestMetric struct {
	Type         string        `json:"type"`
	Attribute    string        `json:"attribute,omitempty"`
	Mbean        string        `json:"mbean"`
	Path         string        `json:"path,omitempty"`
	MetricTarget *MetricTarget `json:"target",omitempty`
}

func (m RequestMetric) String() string {
	return sanitize(fmt.Sprintf("%s:%s:%s", m.Mbean, m.Attribute, m.Path))
}

// Request is a jolokia request holding a slice of RequestMetrics
type Request []RequestMetric

// Response is a jolokia response with metrics
type Response []struct {
	Request   RequestMetric   `json:"request"`
	Value     json.RawMessage `json:"value"`
	Error     string          `json:"error"`
	ErrorType string          `json:"error_type"`
	Status    uint            `json:"status"`
}

// SimpleValue holds a numeric value returned by jolokia
type SimpleValue interface{}

// NestedValue is holding a struct of information returned by jolokia
type NestedValue map[string]json.RawMessage
