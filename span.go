package cache

import (
	"context"

	"go.opencensus.io/trace"
)

// SpanWrapper holds a pointer to a span that allows us to call a function to close the span
type SpanWrapper struct {
	span *trace.Span
}

// AllowTrace checks to see if we should start a trace on the given function call
func AllowTrace(ctx context.Context, allow, root bool) bool {
	return allow && (root || trace.FromContext(ctx) != nil)
}

// StartSpan creates a span on the given call and returns a SpanWrapper. SpanWrapper will be nil if no parentSpan exists and creating new spans is disabled
func StartSpan(ctx context.Context, spanName string, options TraceOptions) *SpanWrapper {
	parentSpan := trace.FromContext(ctx)
	if !options.AllowRoot && parentSpan == nil {
		return nil
	}
	var span *trace.Span
	_, span = trace.StartSpan(ctx, spanName,
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithSampler(options.Sampler),
	)
	if len(options.DefaultAttributes) > 0 {
		span.AddAttributes(options.DefaultAttributes...)
	}
	return &SpanWrapper{
		span: span,
	}
}

// EndSpanWithErr sets the status of the span based on the supplied error and then ends the span
func (s *SpanWrapper) EndSpanWithErr(err error) {
	s.setSpanStatus(err)
	s.span.End()
}

// EndSpan sets the status of the span and then ends the span
func (s *SpanWrapper) EndSpan() {
	s.span.End()
}

func (s *SpanWrapper) setSpanStatus(err error) {
	var status trace.Status
	if err == nil {
		status.Code = trace.StatusCodeOK
	} else {
		status.Code = trace.StatusCodeUnknown
		status.Message = err.Error()
	}
	s.span.SetStatus(status)
}
