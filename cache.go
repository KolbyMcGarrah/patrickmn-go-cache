// Package cache instruments patrickmn/go-cache interactions with Open Census
package cache

import (
	"context"
	"io"
	"time"

	pgocache "github.com/patrickmn/go-cache"
)

// Cacher defines a context aware implementation go-cache
type Cacher interface {
	Add(c context.Context, k string, x interface{}, d time.Duration) error
	Decrement(c context.Context, k string, n int64) error
	DecrementFloat(c context.Context, k string, n float64) error
	DecrementFloat32(c context.Context, k string, n float32) (float32, error)
	DecrementFloat64(c context.Context, k string, n float64) (float64, error)
	DecrementInt(c context.Context, k string, n int) (int, error)
	DecrementInt16(c context.Context, k string, n int16) (int16, error)
	DecrementInt32(c context.Context, k string, n int32) (int32, error)
	DecrementInt64(c context.Context, k string, n int64) (int64, error)
	DecrementInt8(c context.Context, k string, n int8) (int8, error)
	DecrementUint(c context.Context, k string, n uint) (uint, error)
	DecrementUint16(c context.Context, k string, n uint16) (uint16, error)
	DecrementUint32(c context.Context, k string, n uint32) (uint32, error)
	DecrementUint64(c context.Context, k string, n uint64) (uint64, error)
	DecrementUint8(c context.Context, k string, n uint8) (uint8, error)
	DecrementUintptr(c context.Context, k string, n uintptr) (uintptr, error)
	Delete(c context.Context, k string)
	DeleteExpired(c context.Context)
	Flush(c context.Context)
	Get(c context.Context, k string) (interface{}, bool)
	GetWithExpiration(c context.Context, k string) (interface{}, time.Time, bool)
	Increment(c context.Context, k string, n int64) error
	IncrementFloat(c context.Context, k string, n float64) error
	IncrementFloat32(c context.Context, k string, n float32) (float32, error)
	IncrementFloat64(c context.Context, k string, n float64) (float64, error)
	IncrementInt(c context.Context, k string, n int) (int, error)
	IncrementInt16(c context.Context, k string, n int16) (int16, error)
	IncrementInt32(c context.Context, k string, n int32) (int32, error)
	IncrementInt64(c context.Context, k string, n int64) (int64, error)
	IncrementInt8(c context.Context, k string, n int8) (int8, error)
	IncrementUint(c context.Context, k string, n uint) (uint, error)
	IncrementUint16(c context.Context, k string, n uint16) (uint16, error)
	IncrementUint32(c context.Context, k string, n uint32) (uint32, error)
	IncrementUint64(c context.Context, k string, n uint64) (uint64, error)
	IncrementUint8(c context.Context, k string, n uint8) (uint8, error)
	IncrementUintptr(c context.Context, k string, n uintptr) (uintptr, error)
	ItemCount(c context.Context) int
	Items(c context.Context) map[string]pgocache.Item
	Load(c context.Context, r io.Reader) error
	LoadFile(c context.Context, fname string) error
	OnEvicted(c context.Context, f func(string, interface{}))
	Replace(c context.Context, k string, x interface{}, d time.Duration) error
	Save(c context.Context, w io.Writer) (err error)
	SaveFile(c context.Context, fname string) error
	Set(c context.Context, k string, x interface{}, d time.Duration)
	SetDefault(c context.Context, k string, x interface{})
}
