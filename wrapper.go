package cache

import (
	"context"
	"io"
	"time"

	pgocache "github.com/patrickmn/go-cache"
)

// Wrap takes a cache instance and wraps it with OpenCensus instrumentation.
func Wrap(c *pgocache.Cache, instanceName string) *Wrapper {
	return &Wrapper{
		Cache:        c,
		instanceName: instanceName,
	}
}

var _ Cacher = &Wrapper{}

type Wrapper struct {
	Cache        *pgocache.Cache
	instanceName string
}

func (w *Wrapper) Add(ctx context.Context, k string, x interface{}, d time.Duration) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.add", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Add(k, x, d)

	return
}

func (w *Wrapper) Decrement(ctx context.Context, k string, n int64) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrement", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Decrement(k, n)

	return
}

func (w *Wrapper) DecrementFloat(ctx context.Context, k string, n float64) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementfloat", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.DecrementFloat(k, n)

	return
}

func (w *Wrapper) DecrementFloat32(ctx context.Context, k string, n float32) (v float32, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementfloat32", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementFloat32(k, n)

	return
}

func (w *Wrapper) DecrementFloat64(ctx context.Context, k string, n float64) (v float64, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementfloat64", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementFloat64(k, n)

	return
}

func (w *Wrapper) DecrementInt(ctx context.Context, k string, n int) (v int, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt(k, n)

	return
}

func (w *Wrapper) DecrementInt16(ctx context.Context, k string, n int16) (v int16, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint16", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt16(k, n)

	return
}

func (w *Wrapper) DecrementInt32(ctx context.Context, k string, n int32) (v int32, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint32", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt32(k, n)

	return
}

func (w *Wrapper) DecrementInt64(ctx context.Context, k string, n int64) (v int64, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint64", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt64(k, n)

	return
}

func (w *Wrapper) DecrementInt8(ctx context.Context, k string, n int8) (v int8, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementint8", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementInt8(k, n)

	return
}

func (w *Wrapper) DecrementUint(ctx context.Context, k string, n uint) (v uint, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint(k, n)

	return
}

func (w *Wrapper) DecrementUint16(ctx context.Context, k string, n uint16) (v uint16, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint16", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint16(k, n)

	return
}

func (w *Wrapper) DecrementUint32(ctx context.Context, k string, n uint32) (v uint32, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint32", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint32(k, n)

	return
}

func (w *Wrapper) DecrementUint64(ctx context.Context, k string, n uint64) (v uint64, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint64", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint64(k, n)

	return
}

func (w *Wrapper) DecrementUint8(ctx context.Context, k string, n uint8) (v uint8, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuint8", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUint8(k, n)

	return
}

func (w *Wrapper) DecrementUintptr(ctx context.Context, k string, n uintptr) (v uintptr, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.decrementuintptr", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.DecrementUintptr(k, n)

	return
}

func (w *Wrapper) Delete(ctx context.Context, k string) {
	var statsFunc = recordCallStats(ctx, "go.cache.delete", w.instanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.Delete(k)

}

func (w *Wrapper) DeleteExpired(ctx context.Context) {
	var statsFunc = recordCallStats(ctx, "go.cache.deleteexpired", w.instanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.DeleteExpired()

}

func (w *Wrapper) Flush(ctx context.Context) {
	var statsFunc = recordCallStats(ctx, "go.cache.flush", w.instanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.Flush()

}

func (w *Wrapper) Get(ctx context.Context, k string) (v interface{}, found bool) {
	var statsFunc = recordCallFoundStats(ctx, "go.cache.get", w.instanceName)
	defer func() {
		statsFunc(found)
	}()

	v, found = w.Cache.Get(k)

	return
}

func (w *Wrapper) GetWithExpiration(ctx context.Context, k string) (v interface{}, exp time.Time, found bool) {
	var statsFunc = recordCallFoundStats(ctx, "go.cache.getwithexpiration", w.instanceName)
	defer func() {
		statsFunc(found)
	}()

	v, exp, found = w.Cache.GetWithExpiration(k)

	return
}

func (w *Wrapper) Increment(ctx context.Context, k string, n int64) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.increment", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Increment(k, n)

	return
}

func (w *Wrapper) IncrementFloat(ctx context.Context, k string, n float64) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementfloat", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.IncrementFloat(k, n)

	return
}

func (w *Wrapper) IncrementFloat32(ctx context.Context, k string, n float32) (v float32, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementfloat32", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementFloat32(k, n)

	return
}

func (w *Wrapper) IncrementFloat64(ctx context.Context, k string, n float64) (v float64, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementfloat64", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementFloat64(k, n)

	return
}

func (w *Wrapper) IncrementInt(ctx context.Context, k string, n int) (v int, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt(k, n)

	return
}

func (w *Wrapper) IncrementInt16(ctx context.Context, k string, n int16) (v int16, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint16", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt16(k, n)

	return
}

func (w *Wrapper) IncrementInt32(ctx context.Context, k string, n int32) (v int32, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint32", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt32(k, n)

	return
}

func (w *Wrapper) IncrementInt64(ctx context.Context, k string, n int64) (v int64, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint64", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt64(k, n)

	return
}

func (w *Wrapper) IncrementInt8(ctx context.Context, k string, n int8) (v int8, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementint8", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementInt8(k, n)

	return
}

func (w *Wrapper) IncrementUint(ctx context.Context, k string, n uint) (v uint, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint(k, n)

	return
}

func (w *Wrapper) IncrementUint16(ctx context.Context, k string, n uint16) (v uint16, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint16", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint16(k, n)

	return
}

func (w *Wrapper) IncrementUint32(ctx context.Context, k string, n uint32) (v uint32, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint32", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint32(k, n)

	return
}

func (w *Wrapper) IncrementUint64(ctx context.Context, k string, n uint64) (v uint64, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint64", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint64(k, n)

	return
}

func (w *Wrapper) IncrementUint8(ctx context.Context, k string, n uint8) (v uint8, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuint8", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUint8(k, n)

	return
}

func (w *Wrapper) IncrementUintptr(ctx context.Context, k string, n uintptr) (v uintptr, err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.incrementuintptr", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	v, err = w.Cache.IncrementUintptr(k, n)

	return
}

func (w *Wrapper) ItemCount(ctx context.Context) (c int) {
	var statsFunc = recordCallStats(ctx, "go.cache.itemcount", w.instanceName)
	defer func() {
		statsFunc()
	}()

	c = w.Cache.ItemCount()

	return
}

func (w *Wrapper) Items(ctx context.Context) (items map[string]pgocache.Item) {
	var statsFunc = recordCallStats(ctx, "go.cache.items", w.instanceName)
	defer func() {
		statsFunc()
	}()

	items = w.Cache.Items()

	return
}

func (w *Wrapper) Load(ctx context.Context, r io.Reader) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.load", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Load(r)

	return
}

func (w *Wrapper) LoadFile(ctx context.Context, fname string) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.loadfile", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.LoadFile(fname)

	return
}

func (w *Wrapper) OnEvicted(ctx context.Context, f func(string, interface{})) {
	var statsFunc = recordCallStats(ctx, "go.cache.onevicted", w.instanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.OnEvicted(f)
}

func (w *Wrapper) Replace(ctx context.Context, k string, x interface{}, d time.Duration) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.replace", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Replace(k, x, d)

	return
}

func (w *Wrapper) Save(ctx context.Context, wr io.Writer) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.save", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.Save(wr)

	return
}

func (w *Wrapper) SaveFile(ctx context.Context, fname string) (err error) {
	var statsFunc = recordCallErrorStatus(ctx, "go.cache.savefile", w.instanceName)
	defer func() {
		statsFunc(err)
	}()

	err = w.Cache.SaveFile(fname)
	return
}

func (w *Wrapper) Set(ctx context.Context, k string, x interface{}, d time.Duration) {
	var statsFunc = recordCallStats(ctx, "go.cache.set", w.instanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.Set(k, x, d)
}

func (w *Wrapper) SetDefault(ctx context.Context, k string, x interface{}) {
	var statsFunc = recordCallStats(ctx, "go.cache.setdefault", w.instanceName)
	defer func() {
		statsFunc()
	}()

	w.Cache.SetDefault(k, x)
}
