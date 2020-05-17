package cache

import (
	"context"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
)

const (
	statusCalled   = "CALLED"
	statusFound    = "FOUND"
	statusNotFound = "NOT_FOUND"
	statusError    = "ERROR"
	statusOK       = "OK"
)

// The following tags are aooplied to stats recorded by this package
var (
	// GoCacheName is the name of the cache instance.
	GoCacheName, _ = tag.NewKey("go_cache_name")

	// GoCacheMethod is the cache method called.
	GoCacheMethod, _ = tag.NewKey("go_cache_method")

	// GoCacheStatus identifies found v.s not found.
	GoCacheStatus, _ = tag.NewKey("go_cache_status")

	DefaultTags = []tag.Key{GoCacheMethod, GoCacheStatus}
)

// The following measures are supported for use in custom views.
var (
	MeasureLatencyMs = stats.Int64("go.cache/latency", "The latency of calls in milliseconds", stats.UnitMilliseconds)
)

// Default distributions used by views in this package
var (
	DefaultMillisecondsDistribution = view.Distribution(
		0.0,
		0.001,
		0.005,
		0.01,
		0.05,
		0.1,
		0.5,
		1.0,
		1.5,
		2.0,
		2.5,
		5.0,
		10.0,
		25.0,
		50.0,
		100.0,
		200.0,
		400.0,
		600.0,
		800.0,
		1000.0,
		1500.0,
		2000.0,
		2500.0,
		5000.0,
		10000.0,
		20000.0,
		40000.0,
		100000.0,
		200000.0,
		500000.0,
	)
)

// Package cache provides some convenience views.
// You still need to register these views for data to actually be collected.
// You can use the RegisterAllViews function for this.
var (
	GoCacheLatencyView = &view.View{
		Name:        "go.cache/client/latency",
		Description: "The distribution of latency of various calls in milliseconds",
		Measure:     MeasureLatencyMs,
		Aggregation: DefaultMillisecondsDistribution,
		TagKeys:     DefaultTags,
	}

	GoCacheCallsView = &view.View{
		Name:        "go.cache/client/calls",
		Description: "The number of various calls of methods",
		Measure:     MeasureLatencyMs,
		Aggregation: view.Count(),
		TagKeys:     DefaultTags,
	}

	DefaultViews = []*view.View{GoCacheLatencyView, GoCacheCallsView}
)

// RegisterAllViews registers all the cache views to enable collection of stats
func RegisterAllViews() error {
	return view.Register(DefaultViews...)
}

func recordCallStats(ctx context.Context, method string, instanceName string) func() {
	var startTime = time.Now()

	return func() {
		var (
			timeSpentMs = time.Since(startTime).Milliseconds()
			tags        = []tag.Mutator{
				tag.Insert(GoCacheName, instanceName),
				tag.Insert(GoCacheMethod, method),
				tag.Insert(GoCacheStatus, statusCalled),
			}
		)

		_ = stats.RecordWithTags(ctx, tags, MeasureLatencyMs.M(timeSpentMs))
	}
}

func recordCallFoundStats(ctx context.Context, method string, instanceName string) func(found bool) {
	var startTime = time.Now()

	return func(found bool) {
		var (
			timeSpentMs = time.Since(startTime).Milliseconds()
			tags        = []tag.Mutator{
				tag.Insert(GoCacheName, instanceName),
				tag.Insert(GoCacheMethod, method),
			}
		)

		if found {
			tags = append(tags, tag.Insert(GoCacheStatus, statusFound))
		} else {
			tags = append(tags, tag.Insert(GoCacheStatus, statusNotFound))
		}

		_ = stats.RecordWithTags(ctx, tags, MeasureLatencyMs.M(timeSpentMs))
	}
}

func recordCallErrorStatus(ctx context.Context, method string, instanceName string) func(err error) {
	var startTime = time.Now()

	return func(err error) {
		var (
			timeSpentMs = time.Since(startTime).Milliseconds()
			tags        = []tag.Mutator{
				tag.Insert(GoCacheName, instanceName),
				tag.Insert(GoCacheMethod, method),
			}
		)

		if err != nil {
			tags = append(tags, tag.Insert(GoCacheStatus, statusError))
		} else {
			tags = append(tags, tag.Insert(GoCacheStatus, statusOK))
		}

		_ = stats.RecordWithTags(ctx, tags, MeasureLatencyMs.M(timeSpentMs))
	}
}
