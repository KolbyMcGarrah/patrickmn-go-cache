package cache

import (
	"context"
	"io"
	"time"

	pgocache "github.com/patrickmn/go-cache"
	"go.opencensus.io/trace"
)

// Wrap takes a cache instance and wraps it with OpenCensus instrumentation.
func Wrap(c *pgocache.Cache, options ...TraceOption) *Wrapper {
	o := TraceOptions{}
	for _, option := range options {
		option(&o)
	}
	if o.InstanceName == "" {
		o.InstanceName = defaultInstanceName
	} else {
		o.DefaultAttributes = append(o.DefaultAttributes, trace.StringAttribute("cache.instance", o.InstanceName))
	}
	return &Wrapper{
		Cache:   c,
		options: o,
	}
}

var _ Cacher = &Wrapper{}

// Wrapper wraps a pgocache Cache instance with an instance name that can be used to record metrics while preforming cache methods.
type Wrapper struct {
	Cache   *pgocache.Cache
	options TraceOptions
}

// Add implementes the pggocache add method with metrics
func (w *Wrapper) Add(ctx context.Context, k string, x interface{}, d time.Duration) (err error) {
	if AllowTrace(ctx, w.options.Add, w.options.AllowRoot) {
		span := StartSpan(ctx, "Add", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.add", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Add(k, x, d)

	return
}

// Decrement implementes the pggocache decrement method with metrics
func (w *Wrapper) Decrement(ctx context.Context, k string, n int64) (err error) {
	if AllowTrace(ctx, w.options.Decrement, w.options.AllowRoot) {
		span := StartSpan(ctx, "Decrement", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrement", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Decrement(k, n)

	return
}

// DecrementFloat implements the pggocache decrementfloat method with metrics
func (w *Wrapper) DecrementFloat(ctx context.Context, k string, n float64) (err error) {
	if AllowTrace(ctx, w.options.DecrementFloat, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementFloat", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementfloat", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.DecrementFloat(k, n)

	return
}

// DecrementFloat32 implments pggocache decremnetfloat32 method with metrics
func (w *Wrapper) DecrementFloat32(ctx context.Context, k string, n float32) (v float32, err error) {
	if AllowTrace(ctx, w.options.DecrementFloat32, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementFloat32", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementfloat32", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementFloat32(k, n)

	return
}

// DecrementFloat64 implments pggocache decremnetfloat64 method with metrics
func (w *Wrapper) DecrementFloat64(ctx context.Context, k string, n float64) (v float64, err error) {
	if AllowTrace(ctx, w.options.DecrementFloat64, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementFloat64", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementfloat64", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementFloat64(k, n)

	return
}

// DecrementInt implments pggocache decremnetint method with metrics
func (w *Wrapper) DecrementInt(ctx context.Context, k string, n int) (v int, err error) {
	if AllowTrace(ctx, w.options.DecrementInt, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementInt", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt(k, n)

	return
}

// DecrementInt16 implments pggocache decremnetint16 method with metrics
func (w *Wrapper) DecrementInt16(ctx context.Context, k string, n int16) (v int16, err error) {
	if AllowTrace(ctx, w.options.DecrementInt16, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementInt16", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint16", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt16(k, n)

	return
}

// DecrementInt32 implments pggocache decremnetint32 method with metrics
func (w *Wrapper) DecrementInt32(ctx context.Context, k string, n int32) (v int32, err error) {
	if AllowTrace(ctx, w.options.DecrementInt32, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementInt32", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint32", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt32(k, n)

	return
}

// DecrementInt64 implments pggocache decremnetint64 method with metrics
func (w *Wrapper) DecrementInt64(ctx context.Context, k string, n int64) (v int64, err error) {
	if AllowTrace(ctx, w.options.DecrementInt64, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementInt64", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint64", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt64(k, n)

	return
}

// DecrementInt8 implments pggocache decremnetint8 method with metrics
func (w *Wrapper) DecrementInt8(ctx context.Context, k string, n int8) (v int8, err error) {
	if AllowTrace(ctx, w.options.DecrementInt8, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementInt8", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint8", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt8(k, n)

	return
}

// DecrementUint implments pggocache decremnetuint method with metrics
func (w *Wrapper) DecrementUint(ctx context.Context, k string, n uint) (v uint, err error) {
	if AllowTrace(ctx, w.options.DecrementUint, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementUint", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint(k, n)

	return
}

// DecrementUint16 implments pggocache decremnetuint16 method with metrics
func (w *Wrapper) DecrementUint16(ctx context.Context, k string, n uint16) (v uint16, err error) {
	if AllowTrace(ctx, w.options.DecrementUint16, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementUint16", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint16", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint16(k, n)

	return
}

// DecrementUint32 implments pggocache decremnetuint32 method with metrics
func (w *Wrapper) DecrementUint32(ctx context.Context, k string, n uint32) (v uint32, err error) {
	if AllowTrace(ctx, w.options.DecrementUint32, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementUint32", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint32", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint32(k, n)

	return
}

// DecrementUint64 implments pggocache decremnetuint64 method with metrics
func (w *Wrapper) DecrementUint64(ctx context.Context, k string, n uint64) (v uint64, err error) {
	if AllowTrace(ctx, w.options.DecrementUint64, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementUint64", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint64", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint64(k, n)

	return
}

// DecrementUint8 implments pggocache decremnetUint8 method with metrics
func (w *Wrapper) DecrementUint8(ctx context.Context, k string, n uint8) (v uint8, err error) {
	if AllowTrace(ctx, w.options.DecrementUint8, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementUint8", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint8", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint8(k, n)

	return
}

// DecrementUintptr implments pggocache decremnetuintptr method with metrics
func (w *Wrapper) DecrementUintptr(ctx context.Context, k string, n uintptr) (v uintptr, err error) {
	if AllowTrace(ctx, w.options.DecrementUintptr, w.options.AllowRoot) {
		span := StartSpan(ctx, "DecrementUintptr", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuintptr", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUintptr(k, n)

	return
}

// Delete implments pggocache delete method with metrics
func (w *Wrapper) Delete(ctx context.Context, k string) {
	if AllowTrace(ctx, w.options.Delete, w.options.AllowRoot) {
		span := StartSpan(ctx, "Delete", w.options)
		if span != nil {
			defer func() {
				span.EndSpan()
			}()
		}
	}
	var statsFunc = recordCallStats(ctx, "go.cache.delete", w.options.InstanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.Delete(k)

}

// DeleteExpired implments pggocache deleteexpired method with metrics
func (w *Wrapper) DeleteExpired(ctx context.Context) {
	if AllowTrace(ctx, w.options.DeleteExpired, w.options.AllowRoot) {
		span := StartSpan(ctx, "DeleteExpired", w.options)
		if span != nil {
			defer func() {
				span.EndSpan()
			}()
		}
	}
	var statsFunc = recordCallStats(ctx, "go.cache.deleteexpired", w.options.InstanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.DeleteExpired()

}

// Flush implments pggocache flush method with metrics
func (w *Wrapper) Flush(ctx context.Context) {
	if AllowTrace(ctx, w.options.Flush, w.options.AllowRoot) {
		span := StartSpan(ctx, "Flush", w.options)
		if span != nil {
			defer func() {
				span.EndSpan()
			}()
		}
	}
	var statsFunc = recordCallStats(ctx, "go.cache.flush", w.options.InstanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.Flush()

}

// Get implments pggocache get method with metrics
func (w *Wrapper) Get(ctx context.Context, k string) (v interface{}, found bool) {
	if AllowTrace(ctx, w.options.Get, w.options.AllowRoot) {
		span := StartSpan(ctx, "Get", w.options)
		if span != nil {
			defer func() {
				span.EndSpan()
			}()
		}
	}
	var statsFunc = recordCallFoundStats(ctx, "go.cache.get", w.options.InstanceName)
	defer func() {
		statsFunc(found)
	}()

	v, found = w.Cache.Get(k)

	return
}

// GetWithExpiration implments pggocache getwithexpiration method with metrics
func (w *Wrapper) GetWithExpiration(ctx context.Context, k string) (v interface{}, exp time.Time, found bool) {
	if AllowTrace(ctx, w.options.GetWithExpiration, w.options.AllowRoot) {
		span := StartSpan(ctx, "GetWithExpiration", w.options)
		if span != nil {
			defer func() {
				span.EndSpan()
			}()
		}
	}
	var statsFunc = recordCallFoundStats(ctx, "go.cache.getwithexpiration", w.options.InstanceName)
	defer func() {
		statsFunc(found)
	}()

	v, exp, found = w.Cache.GetWithExpiration(k)

	return
}

// Increment implments pggocache increment method with metrics
func (w *Wrapper) Increment(ctx context.Context, k string, n int64) (err error) {
	if AllowTrace(ctx, w.options.Increment, w.options.AllowRoot) {
		span := StartSpan(ctx, "Increment", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.increment", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Increment(k, n)

	return
}

// IncrementFloat implments pggocache incrementfloat method with metrics
func (w *Wrapper) IncrementFloat(ctx context.Context, k string, n float64) (err error) {
	if AllowTrace(ctx, w.options.IncrementFloat, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementFloat", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementfloat", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.IncrementFloat(k, n)

	return
}

// IncrementFloat32 implments pggocache incrementfloat32 method with metrics
func (w *Wrapper) IncrementFloat32(ctx context.Context, k string, n float32) (v float32, err error) {
	if AllowTrace(ctx, w.options.IncrementFloat32, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementFloat32", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementfloat32", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementFloat32(k, n)

	return
}

// IncrementFloat64 implments pggocache incrementfloat64 method with metrics
func (w *Wrapper) IncrementFloat64(ctx context.Context, k string, n float64) (v float64, err error) {
	if AllowTrace(ctx, w.options.IncrementFloat64, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementFloat64", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementfloat64", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementFloat64(k, n)

	return
}

// IncrementInt implments pggocache incrementint method with metrics
func (w *Wrapper) IncrementInt(ctx context.Context, k string, n int) (v int, err error) {
	if AllowTrace(ctx, w.options.IncrementInt, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementInt", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt(k, n)

	return
}

// IncrementInt16 implments pggocache incrementint16 method with metrics
func (w *Wrapper) IncrementInt16(ctx context.Context, k string, n int16) (v int16, err error) {
	if AllowTrace(ctx, w.options.IncrementInt16, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementInt16", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint16", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt16(k, n)

	return
}

// IncrementInt32 implments pggocache incrementint32 method with metrics
func (w *Wrapper) IncrementInt32(ctx context.Context, k string, n int32) (v int32, err error) {
	if AllowTrace(ctx, w.options.IncrementInt32, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementInt32", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint32", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt32(k, n)

	return
}

// IncrementInt64 implments pggocache incrementint64 method with metrics
func (w *Wrapper) IncrementInt64(ctx context.Context, k string, n int64) (v int64, err error) {
	if AllowTrace(ctx, w.options.IncrementInt64, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementInt64", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint64", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt64(k, n)

	return
}

// IncrementInt8 implments pggocache incrementint8 method with metrics
func (w *Wrapper) IncrementInt8(ctx context.Context, k string, n int8) (v int8, err error) {
	if AllowTrace(ctx, w.options.IncrementInt8, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementInt8", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint8", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt8(k, n)

	return
}

// IncrementUint implments pggocache incrementuint method with metrics
func (w *Wrapper) IncrementUint(ctx context.Context, k string, n uint) (v uint, err error) {
	if AllowTrace(ctx, w.options.IncrementUint, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementUint", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint(k, n)

	return
}

// IncrementUint16 implments pggocache incrementuint16 method with metrics
func (w *Wrapper) IncrementUint16(ctx context.Context, k string, n uint16) (v uint16, err error) {
	if AllowTrace(ctx, w.options.IncrementUint16, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementUint16", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint16", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint16(k, n)

	return
}

// IncrementUint32 implments pggocache incrementuint32 method with metrics
func (w *Wrapper) IncrementUint32(ctx context.Context, k string, n uint32) (v uint32, err error) {
	if AllowTrace(ctx, w.options.IncrementUint32, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementUint32", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint32", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint32(k, n)

	return
}

// IncrementUint64 implments pggocache incrementuint64 method with metrics
func (w *Wrapper) IncrementUint64(ctx context.Context, k string, n uint64) (v uint64, err error) {
	if AllowTrace(ctx, w.options.IncrementUint64, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementUint64", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint64", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint64(k, n)

	return
}

// IncrementUint8 implments pggocache incrementuint8 method with metrics
func (w *Wrapper) IncrementUint8(ctx context.Context, k string, n uint8) (v uint8, err error) {
	if AllowTrace(ctx, w.options.IncrementUint8, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementUint8", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint8", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint8(k, n)

	return
}

// IncrementUintptr implments pggocache incrementuintptr method with metrics
func (w *Wrapper) IncrementUintptr(ctx context.Context, k string, n uintptr) (v uintptr, err error) {
	if AllowTrace(ctx, w.options.IncrementUintptr, w.options.AllowRoot) {
		span := StartSpan(ctx, "IncrementUintptr", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuintptr", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUintptr(k, n)

	return
}

// ItemCount implments pggocache itemcount method with metrics
func (w *Wrapper) ItemCount(ctx context.Context) (c int) {
	if AllowTrace(ctx, w.options.ItemCount, w.options.AllowRoot) {
		span := StartSpan(ctx, "ItemCount", w.options)
		if span != nil {
			defer func() {
				span.EndSpan()
			}()
		}
	}
	var statsFunc = recordCallStats(ctx, "go.cache.itemcount", w.options.InstanceName)
	defer func() {
		statsFunc()
	}()

	c = w.Cache.ItemCount()

	return
}

// Items implments pggocache items method with metrics
func (w *Wrapper) Items(ctx context.Context) (items map[string]pgocache.Item) {
	if AllowTrace(ctx, w.options.Items, w.options.AllowRoot) {
		span := StartSpan(ctx, "Items", w.options)
		if span != nil {
			defer func() {
				span.EndSpan()
			}()
		}
	}
	var statsFunc = recordCallStats(ctx, "go.cache.items", w.options.InstanceName)
	defer func() {
		statsFunc()
	}()

	items = w.Cache.Items()

	return
}

// Load implments pggocache load method with metrics
func (w *Wrapper) Load(ctx context.Context, r io.Reader) (err error) {
	if AllowTrace(ctx, w.options.Load, w.options.AllowRoot) {
		span := StartSpan(ctx, "Load", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.load", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Load(r)

	return
}

// LoadFile implments pggocache loadfile method with metrics
func (w *Wrapper) LoadFile(ctx context.Context, fname string) (err error) {
	if AllowTrace(ctx, w.options.LoadFile, w.options.AllowRoot) {
		span := StartSpan(ctx, "LoadFile", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.loadfile", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.LoadFile(fname)

	return
}

// OnEvicted implments pggocache onevicted method with metrics
func (w *Wrapper) OnEvicted(ctx context.Context, f func(string, interface{})) {
	if AllowTrace(ctx, w.options.OnEvicted, w.options.AllowRoot) {
		span := StartSpan(ctx, "OnEvicted", w.options)
		if span != nil {
			defer func() {
				span.EndSpan()
			}()
		}
	}
	var statsFunc = recordCallStats(ctx, "go.cache.onevicted", w.options.InstanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.OnEvicted(f)
}

// Replace implments pggocache replace method with metrics
func (w *Wrapper) Replace(ctx context.Context, k string, x interface{}, d time.Duration) (err error) {
	if AllowTrace(ctx, w.options.Replace, w.options.AllowRoot) {
		span := StartSpan(ctx, "Replace", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.replace", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Replace(k, x, d)

	return
}

// Save implments pggocache save method with metrics
func (w *Wrapper) Save(ctx context.Context, wr io.Writer) (err error) {
	if AllowTrace(ctx, w.options.Save, w.options.AllowRoot) {
		span := StartSpan(ctx, "Save", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.save", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Save(wr)

	return
}

// SaveFile implments pggocache savefile method with metrics
func (w *Wrapper) SaveFile(ctx context.Context, fname string) (err error) {
	if AllowTrace(ctx, w.options.SaveFile, w.options.AllowRoot) {
		span := StartSpan(ctx, "SaveFile", w.options)
		if span != nil {
			defer func() {
				span.EndSpanWithErr(err)
			}()
		}
	}
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.savefile", w.options.InstanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.SaveFile(fname)
	return
}

// Set implments pggocache set method with metrics
func (w *Wrapper) Set(ctx context.Context, k string, x interface{}, d time.Duration) {
	if AllowTrace(ctx, w.options.Set, w.options.AllowRoot) {
		span := StartSpan(ctx, "Set", w.options)
		if span != nil {
			defer func() {
				span.EndSpan()
			}()
		}
	}
	var statsFunc = recordCallStats(ctx, "go.cache.set", w.options.InstanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.Set(k, x, d)
}

// SetDefault implments pggocache setdefault method with metrics
func (w *Wrapper) SetDefault(ctx context.Context, k string, x interface{}) {
	if AllowTrace(ctx, w.options.SetDefault, w.options.AllowRoot) {
		span := StartSpan(ctx, "SaveDefault", w.options)
		if span != nil {
			defer span.EndSpan()
		}
	}
	var statsFunc = recordCallStats(ctx, "go.cache.setdefault", w.options.InstanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.SetDefault(k, x)
}
