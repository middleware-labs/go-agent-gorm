package gorm

import (
	"github.com/middleware-labs/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/attribute"
	"gorm.io/gorm"
)

func NewPlugin(opts ...otelgorm.Option) gorm.Plugin {
	return otelgorm.NewPlugin(opts...)
}

func WithAttributes(attrs ...attribute.KeyValue) otelgorm.Option {
	return otelgorm.WithAttributes(attrs...)
}

func WithDBName(name string) otelgorm.Option {
	return otelgorm.WithDBName(name)
}

