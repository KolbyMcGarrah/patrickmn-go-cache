package cache

import (
	"bytes"
	"context"
	"io/ioutil"
	"runtime"
	"strconv"
	"sync"
	"testing"
	"time"

	pgocache "github.com/patrickmn/go-cache"
)

// tests are from https://github.com/patrickmn/go-cache/blob/master/cache_test.go
// to ensure the wrapping act the same as the original cache

type TestStruct struct {
	Num      int
	Children []*TestStruct
}

func TestCache(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")

	a, found := tc.Get(context.Background(), "a")
	if found || a != nil {
		t.Error("Getting A found value that shouldn't exist:", a)
	}

	b, found := tc.Get(context.Background(), "b")
	if found || b != nil {
		t.Error("Getting B found value that shouldn't exist:", b)
	}

	c, found := tc.Get(context.Background(), "c")
	if found || c != nil {
		t.Error("Getting C found value that shouldn't exist:", c)
	}

	tc.Set(context.Background(), "a", 1, pgocache.DefaultExpiration)
	tc.Set(context.Background(), "b", "b", pgocache.DefaultExpiration)
	tc.Set(context.Background(), "c", 3.5, pgocache.DefaultExpiration)

	x, found := tc.Get(context.Background(), "a")
	if !found {
		t.Error("a was not found while getting a2")
	}
	if x == nil {
		t.Error("x for a is nil")
	} else if a2 := x.(int); a2+2 != 3 {
		t.Error("a2 (which should be 1) plus 2 does not equal 3; value:", a2)
	}

	x, found = tc.Get(context.Background(), "b")
	if !found {
		t.Error("b was not found while getting b2")
	}
	if x == nil {
		t.Error("x for b is nil")
	} else if b2 := x.(string); b2+"B" != "bB" {
		t.Error("b2 (which should be b) plus B does not equal bB; value:", b2)
	}

	x, found = tc.Get(context.Background(), "c")
	if !found {
		t.Error("c was not found while getting c2")
	}
	if x == nil {
		t.Error("x for c is nil")
	} else if c2 := x.(float64); c2+1.2 != 4.7 {
		t.Error("c2 (which should be 3.5) plus 1.2 does not equal 4.7; value:", c2)
	}
}

func TestCacheTimes(t *testing.T) {
	var found bool

	tc := Wrap(pgocache.New(50*time.Millisecond, 1*time.Millisecond), "test")
	tc.Set(context.Background(), "a", 1, pgocache.DefaultExpiration)
	tc.Set(context.Background(), "b", 2, pgocache.NoExpiration)
	tc.Set(context.Background(), "c", 3, 20*time.Millisecond)
	tc.Set(context.Background(), "d", 4, 70*time.Millisecond)

	<-time.After(25 * time.Millisecond)
	_, found = tc.Get(context.Background(), "c")
	if found {
		t.Error("Found c when it should have been automatically deleted")
	}

	<-time.After(30 * time.Millisecond)
	_, found = tc.Get(context.Background(), "a")
	if found {
		t.Error("Found a when it should have been automatically deleted")
	}

	_, found = tc.Get(context.Background(), "b")
	if !found {
		t.Error("Did not find b even though it was set to never expire")
	}

	_, found = tc.Get(context.Background(), "d")
	if !found {
		t.Error("Did not find d even though it was set to expire later than the default")
	}

	<-time.After(20 * time.Millisecond)
	_, found = tc.Get(context.Background(), "d")
	if found {
		t.Error("Found d when it should have been automatically deleted (later than the default)")
	}
}

func TestNewFrom(t *testing.T) {
	m := map[string]pgocache.Item{
		"a": pgocache.Item{
			Object:     1,
			Expiration: 0,
		},
		"b": pgocache.Item{
			Object:     2,
			Expiration: 0,
		},
	}
	tc := Wrap(pgocache.NewFrom(pgocache.DefaultExpiration, 0, m), "test")
	a, found := tc.Get(context.Background(), "a")
	if !found {
		t.Fatal("Did not find a")
	}
	if a.(int) != 1 {
		t.Fatal("a is not 1")
	}
	b, found := tc.Get(context.Background(), "b")
	if !found {
		t.Fatal("Did not find b")
	}
	if b.(int) != 2 {
		t.Fatal("b is not 2")
	}
}

func TestStorePointerToStruct(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "foo", &TestStruct{Num: 1}, pgocache.DefaultExpiration)
	x, found := tc.Get(context.Background(), "foo")
	if !found {
		t.Fatal("*TestStruct was not found for foo")
	}
	foo := x.(*TestStruct)
	foo.Num++

	y, found := tc.Get(context.Background(), "foo")
	if !found {
		t.Fatal("*TestStruct was not found for foo (second time)")
	}
	bar := y.(*TestStruct)
	if bar.Num != 2 {
		t.Fatal("TestStruct.Num is not 2")
	}
}

func TestIncrementWithInt(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint", 1, pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "tint")
	if !found {
		t.Error("tint was not found")
	}
	if x.(int) != 3 {
		t.Error("tint is not 3:", x)
	}
}

func TestIncrementWithInt8(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint8", int8(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "tint8")
	if !found {
		t.Error("tint8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("tint8 is not 3:", x)
	}
}

func TestIncrementWithInt16(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint16", int16(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "tint16")
	if !found {
		t.Error("tint16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("tint16 is not 3:", x)
	}
}

func TestIncrementWithInt32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint32", int32(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "tint32")
	if !found {
		t.Error("tint32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("tint32 is not 3:", x)
	}
}

func TestIncrementWithInt64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint64", int64(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "tint64")
	if !found {
		t.Error("tint64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("tint64 is not 3:", x)
	}
}

func TestIncrementWithUint(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint", uint(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tuint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "tuint")
	if !found {
		t.Error("tuint was not found")
	}
	if x.(uint) != 3 {
		t.Error("tuint is not 3:", x)
	}
}

func TestIncrementWithUintptr(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuintptr", uintptr(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tuintptr", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get(context.Background(), "tuintptr")
	if !found {
		t.Error("tuintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("tuintptr is not 3:", x)
	}
}

func TestIncrementWithUint8(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint8", uint8(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tuint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "tuint8")
	if !found {
		t.Error("tuint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("tuint8 is not 3:", x)
	}
}

func TestIncrementWithUint16(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint16", uint16(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tuint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get(context.Background(), "tuint16")
	if !found {
		t.Error("tuint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("tuint16 is not 3:", x)
	}
}

func TestIncrementWithUint32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint32", uint32(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tuint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "tuint32")
	if !found {
		t.Error("tuint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("tuint32 is not 3:", x)
	}
}

func TestIncrementWithUint64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint64", uint64(1), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "tuint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}

	x, found := tc.Get(context.Background(), "tuint64")
	if !found {
		t.Error("tuint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("tuint64 is not 3:", x)
	}
}

func TestIncrementWithFloat32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float32", float32(1.5), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "float32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}
}

func TestIncrementWithFloat64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float64", float64(1.5), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "float64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	x, found := tc.Get(context.Background(), "float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}
}

func TestIncrementFloatWithFloat32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float32", float32(1.5), pgocache.DefaultExpiration)
	err := tc.IncrementFloat(context.Background(), "float32", 2)
	if err != nil {
		t.Error("Error incrementfloating:", err)
	}
	x, found := tc.Get(context.Background(), "float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}
}

func TestIncrementFloatWithFloat64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float64", float64(1.5), pgocache.DefaultExpiration)
	err := tc.IncrementFloat(context.Background(), "float64", 2)
	if err != nil {
		t.Error("Error incrementfloating:", err)
	}
	x, found := tc.Get(context.Background(), "float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}
}

func TestDecrementWithInt(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int", int(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "int", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "int")
	if !found {
		t.Error("int was not found")
	}
	if x.(int) != 3 {
		t.Error("int is not 3:", x)
	}
}

func TestDecrementWithInt8(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int8", int8(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "int8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "int8")
	if !found {
		t.Error("int8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("int8 is not 3:", x)
	}
}

func TestDecrementWithInt16(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int16", int16(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "int16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "int16")
	if !found {
		t.Error("int16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("int16 is not 3:", x)
	}
}

func TestDecrementWithInt32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int32", int32(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "int32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "int32")
	if !found {
		t.Error("int32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("int32 is not 3:", x)
	}
}

func TestDecrementWithInt64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int64", int64(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "int64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "int64")
	if !found {
		t.Error("int64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("int64 is not 3:", x)
	}
}

func TestDecrementWithUint(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint", uint(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "uint", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "uint")
	if !found {
		t.Error("uint was not found")
	}
	if x.(uint) != 3 {
		t.Error("uint is not 3:", x)
	}
}

func TestDecrementWithUintptr(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uintptr", uintptr(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "uintptr", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "uintptr")
	if !found {
		t.Error("uintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("uintptr is not 3:", x)
	}
}

func TestDecrementWithUint8(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint8", uint8(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "uint8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "uint8")
	if !found {
		t.Error("uint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("uint8 is not 3:", x)
	}
}

func TestDecrementWithUint16(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint16", uint16(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "uint16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "uint16")
	if !found {
		t.Error("uint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("uint16 is not 3:", x)
	}
}

func TestDecrementWithUint32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint32", uint32(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "uint32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "uint32")
	if !found {
		t.Error("uint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("uint32 is not 3:", x)
	}
}

func TestDecrementWithUint64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint64", uint64(5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "uint64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "uint64")
	if !found {
		t.Error("uint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("uint64 is not 3:", x)
	}
}

func TestDecrementWithFloat32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float32", float32(5.5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3:", x)
	}
}

func TestDecrementWithFloat64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float64", float64(5.5), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3:", x)
	}
}

func TestDecrementFloatWithFloat32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float32", float32(5.5), pgocache.DefaultExpiration)
	err := tc.DecrementFloat(context.Background(), "float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3:", x)
	}
}

func TestDecrementFloatWithFloat64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float64", float64(5.5), pgocache.DefaultExpiration)
	err := tc.DecrementFloat(context.Background(), "float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	x, found := tc.Get(context.Background(), "float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3:", x)
	}
}

func TestIncrementInt(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint", 1, pgocache.DefaultExpiration)
	n, err := tc.IncrementInt(context.Background(), "tint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tint")
	if !found {
		t.Error("tint was not found")
	}
	if x.(int) != 3 {
		t.Error("tint is not 3:", x)
	}
}

func TestIncrementInt8(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint8", int8(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementInt8(context.Background(), "tint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tint8")
	if !found {
		t.Error("tint8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("tint8 is not 3:", x)
	}
}

func TestIncrementInt16(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint16", int16(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementInt16(context.Background(), "tint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tint16")
	if !found {
		t.Error("tint16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("tint16 is not 3:", x)
	}
}

func TestIncrementInt32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint32", int32(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementInt32(context.Background(), "tint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tint32")
	if !found {
		t.Error("tint32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("tint32 is not 3:", x)
	}
}

func TestIncrementInt64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tint64", int64(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementInt64(context.Background(), "tint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tint64")
	if !found {
		t.Error("tint64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("tint64 is not 3:", x)
	}
}

func TestIncrementUint(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint", uint(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementUint(context.Background(), "tuint", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tuint")
	if !found {
		t.Error("tuint was not found")
	}
	if x.(uint) != 3 {
		t.Error("tuint is not 3:", x)
	}
}

func TestIncrementUintptr(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuintptr", uintptr(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementUintptr(context.Background(), "tuintptr", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tuintptr")
	if !found {
		t.Error("tuintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("tuintptr is not 3:", x)
	}
}

func TestIncrementUint8(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint8", uint8(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementUint8(context.Background(), "tuint8", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tuint8")
	if !found {
		t.Error("tuint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("tuint8 is not 3:", x)
	}
}

func TestIncrementUint16(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint16", uint16(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementUint16(context.Background(), "tuint16", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tuint16")
	if !found {
		t.Error("tuint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("tuint16 is not 3:", x)
	}
}

func TestIncrementUint32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint32", uint32(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementUint32(context.Background(), "tuint32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tuint32")
	if !found {
		t.Error("tuint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("tuint32 is not 3:", x)
	}
}

func TestIncrementUint64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "tuint64", uint64(1), pgocache.DefaultExpiration)
	n, err := tc.IncrementUint64(context.Background(), "tuint64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "tuint64")
	if !found {
		t.Error("tuint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("tuint64 is not 3:", x)
	}
}

func TestIncrementFloat32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float32", float32(1.5), pgocache.DefaultExpiration)
	n, err := tc.IncrementFloat32(context.Background(), "float32", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3.5 {
		t.Error("Returned number is not 3.5:", n)
	}
	x, found := tc.Get(context.Background(), "float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3.5 {
		t.Error("float32 is not 3.5:", x)
	}
}

func TestIncrementFloat64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float64", float64(1.5), pgocache.DefaultExpiration)
	n, err := tc.IncrementFloat64(context.Background(), "float64", 2)
	if err != nil {
		t.Error("Error incrementing:", err)
	}
	if n != 3.5 {
		t.Error("Returned number is not 3.5:", n)
	}
	x, found := tc.Get(context.Background(), "float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3.5 {
		t.Error("float64 is not 3.5:", x)
	}
}

func TestDecrementInt8(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int8", int8(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementInt8(context.Background(), "int8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "int8")
	if !found {
		t.Error("int8 was not found")
	}
	if x.(int8) != 3 {
		t.Error("int8 is not 3:", x)
	}
}

func TestDecrementInt16(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int16", int16(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementInt16(context.Background(), "int16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "int16")
	if !found {
		t.Error("int16 was not found")
	}
	if x.(int16) != 3 {
		t.Error("int16 is not 3:", x)
	}
}

func TestDecrementInt32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int32", int32(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementInt32(context.Background(), "int32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "int32")
	if !found {
		t.Error("int32 was not found")
	}
	if x.(int32) != 3 {
		t.Error("int32 is not 3:", x)
	}
}

func TestDecrementInt64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int64", int64(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementInt64(context.Background(), "int64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "int64")
	if !found {
		t.Error("int64 was not found")
	}
	if x.(int64) != 3 {
		t.Error("int64 is not 3:", x)
	}
}

func TestDecrementUint(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint", uint(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementUint(context.Background(), "uint", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "uint")
	if !found {
		t.Error("uint was not found")
	}
	if x.(uint) != 3 {
		t.Error("uint is not 3:", x)
	}
}

func TestDecrementUintptr(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uintptr", uintptr(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementUintptr(context.Background(), "uintptr", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "uintptr")
	if !found {
		t.Error("uintptr was not found")
	}
	if x.(uintptr) != 3 {
		t.Error("uintptr is not 3:", x)
	}
}

func TestDecrementUint8(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint8", uint8(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementUint8(context.Background(), "uint8", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "uint8")
	if !found {
		t.Error("uint8 was not found")
	}
	if x.(uint8) != 3 {
		t.Error("uint8 is not 3:", x)
	}
}

func TestDecrementUint16(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint16", uint16(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementUint16(context.Background(), "uint16", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "uint16")
	if !found {
		t.Error("uint16 was not found")
	}
	if x.(uint16) != 3 {
		t.Error("uint16 is not 3:", x)
	}
}

func TestDecrementUint32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint32", uint32(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementUint32(context.Background(), "uint32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "uint32")
	if !found {
		t.Error("uint32 was not found")
	}
	if x.(uint32) != 3 {
		t.Error("uint32 is not 3:", x)
	}
}

func TestDecrementUint64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint64", uint64(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementUint64(context.Background(), "uint64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "uint64")
	if !found {
		t.Error("uint64 was not found")
	}
	if x.(uint64) != 3 {
		t.Error("uint64 is not 3:", x)
	}
}

func TestDecrementFloat32(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float32", float32(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementFloat32(context.Background(), "float32", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "float32")
	if !found {
		t.Error("float32 was not found")
	}
	if x.(float32) != 3 {
		t.Error("float32 is not 3:", x)
	}
}

func TestDecrementFloat64(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "float64", float64(5), pgocache.DefaultExpiration)
	n, err := tc.DecrementFloat64(context.Background(), "float64", 2)
	if err != nil {
		t.Error("Error decrementing:", err)
	}
	if n != 3 {
		t.Error("Returned number is not 3:", n)
	}
	x, found := tc.Get(context.Background(), "float64")
	if !found {
		t.Error("float64 was not found")
	}
	if x.(float64) != 3 {
		t.Error("float64 is not 3:", x)
	}
}

func TestAdd(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	err := tc.Add(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
	if err != nil {
		t.Error("Couldn't add foo even though it shouldn't exist")
	}
	err = tc.Add(context.Background(), "foo", "baz", pgocache.DefaultExpiration)
	if err == nil {
		t.Error("Successfully added another foo when it should have returned an error")
	}
}

func TestReplace(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	err := tc.Replace(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
	if err == nil {
		t.Error("Replaced foo when it shouldn't exist")
	}
	tc.Set(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
	err = tc.Replace(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
	if err != nil {
		t.Error("Couldn't replace existing key foo")
	}
}

func TestDelete(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
	tc.Delete(context.Background(), "foo")
	x, found := tc.Get(context.Background(), "foo")
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != nil {
		t.Error("x is not nil:", x)
	}
}

func TestItemCount(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "foo", "1", pgocache.DefaultExpiration)
	tc.Set(context.Background(), "bar", "2", pgocache.DefaultExpiration)
	tc.Set(context.Background(), "baz", "3", pgocache.DefaultExpiration)
	if n := tc.ItemCount(context.Background()); n != 3 {
		t.Errorf("pgocache.Item count is not 3: %d", n)
	}
}

func TestFlush(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
	tc.Set(context.Background(), "baz", "yes", pgocache.DefaultExpiration)
	tc.Flush(context.Background())
	x, found := tc.Get(context.Background(), "foo")
	if found {
		t.Error("foo was found, but it should have been deleted")
	}
	if x != nil {
		t.Error("x is not nil:", x)
	}
	x, found = tc.Get(context.Background(), "baz")
	if found {
		t.Error("baz was found, but it should have been deleted")
	}
	if x != nil {
		t.Error("x is not nil:", x)
	}
}

func TestIncrementOverflowInt(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "int8", int8(127), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "int8", 1)
	if err != nil {
		t.Error("Error incrementing int8:", err)
	}
	x, _ := tc.Get(context.Background(), "int8")
	int8 := x.(int8)
	if int8 != -128 {
		t.Error("int8 did not overflow as expected; value:", int8)
	}

}

func TestIncrementOverflowUint(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint8", uint8(255), pgocache.DefaultExpiration)
	err := tc.Increment(context.Background(), "uint8", 1)
	if err != nil {
		t.Error("Error incrementing int8:", err)
	}
	x, _ := tc.Get(context.Background(), "uint8")
	uint8 := x.(uint8)
	if uint8 != 0 {
		t.Error("uint8 did not overflow as expected; value:", uint8)
	}
}

func TestDecrementUnderflowUint(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "uint8", uint8(0), pgocache.DefaultExpiration)
	err := tc.Decrement(context.Background(), "uint8", 1)
	if err != nil {
		t.Error("Error decrementing int8:", err)
	}
	x, _ := tc.Get(context.Background(), "uint8")
	uint8 := x.(uint8)
	if uint8 != 255 {
		t.Error("uint8 did not underflow as expected; value:", uint8)
	}
}

func TestOnEvicted(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "foo", 3, pgocache.DefaultExpiration)
	works := false
	tc.OnEvicted(context.Background(), func(k string, v interface{}) {
		if k == "foo" && v.(int) == 3 {
			works = true
		}
		tc.Set(context.Background(), "bar", 4, pgocache.DefaultExpiration)
	})
	tc.Delete(context.Background(), "foo")
	x, _ := tc.Get(context.Background(), "bar")
	if !works {
		t.Error("works bool not true")
	}
	if x.(int) != 4 {
		t.Error("bar was not 4")
	}
}

func TestFileSerialization(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Add(context.Background(), "a", "a", pgocache.DefaultExpiration)
	tc.Add(context.Background(), "b", "b", pgocache.DefaultExpiration)
	f, err := ioutil.TempFile("", "go-cache-cache.dat")
	if err != nil {
		t.Fatal("Couldn't create cache file:", err)
	}
	fname := f.Name()
	f.Close()
	tc.SaveFile(context.Background(), fname)

	oc := pgocache.New(pgocache.DefaultExpiration, 0)
	oc.Add("a", "aa", 0) // this should not be overwritten
	err = oc.LoadFile(fname)
	if err != nil {
		t.Error(err)
	}
	a, found := oc.Get("a")
	if !found {
		t.Error("a was not found")
	}
	astr := a.(string)
	if astr != "aa" {
		if astr == "a" {
			t.Error("a was overwritten")
		} else {
			t.Error("a is not aa")
		}
	}
	b, found := oc.Get("b")
	if !found {
		t.Error("b was not found")
	}
	if b.(string) != "b" {
		t.Error("b is not b")
	}
}

func TestSerializeUnserializable(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	ch := make(chan bool, 1)
	ch <- true
	tc.Set(context.Background(), "chan", ch, pgocache.DefaultExpiration)
	fp := &bytes.Buffer{}
	err := tc.Save(context.Background(), fp) // this should fail gracefully
	if err.Error() != "gob NewTypeObject can't handle type: chan bool" {
		t.Error("Error from Save was not gob NewTypeObject can't handle type chan bool:", err)
	}
}

func BenchmarkCacheGetExpiring(b *testing.B) {
	benchmarkCacheGet(b, 5*time.Minute)
}

func BenchmarkCacheGetNotExpiring(b *testing.B) {
	benchmarkCacheGet(b, pgocache.NoExpiration)
}

func benchmarkCacheGet(b *testing.B, exp time.Duration) {
	b.StopTimer()
	tc := Wrap(pgocache.New(exp, 0), "test")
	tc.Set(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Get(context.Background(), "foo")
	}
}

func BenchmarkRWMutexMapGet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m["foo"]
		mu.RUnlock()
	}
}

func BenchmarkRWMutexInterfaceMapGetStruct(b *testing.B) {
	b.StopTimer()
	s := struct{ name string }{name: "foo"}
	m := map[interface{}]string{
		s: "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m[s]
		mu.RUnlock()
	}
}

func BenchmarkRWMutexInterfaceMapGetString(b *testing.B) {
	b.StopTimer()
	m := map[interface{}]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.RLock()
		_, _ = m["foo"]
		mu.RUnlock()
	}
}

func BenchmarkCacheGetConcurrentExpiring(b *testing.B) {
	benchmarkCacheGetConcurrent(b, 5*time.Minute)
}

func BenchmarkCacheGetConcurrentNotExpiring(b *testing.B) {
	benchmarkCacheGetConcurrent(b, pgocache.NoExpiration)
}

func benchmarkCacheGetConcurrent(b *testing.B, exp time.Duration) {
	b.StopTimer()
	tc := Wrap(pgocache.New(exp, 0), "test")
	tc.Set(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
	wg := new(sync.WaitGroup)
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers)
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < each; j++ {
				tc.Get(context.Background(), "foo")
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkRWMutexMapGetConcurrent(b *testing.B) {
	b.StopTimer()
	m := map[string]string{
		"foo": "bar",
	}
	mu := sync.RWMutex{}
	wg := new(sync.WaitGroup)
	workers := runtime.NumCPU()
	each := b.N / workers
	wg.Add(workers)
	b.StartTimer()
	for i := 0; i < workers; i++ {
		go func() {
			for j := 0; j < each; j++ {
				mu.RLock()
				_, _ = m["foo"]
				mu.RUnlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func BenchmarkCacheGetManyConcurrentExpiring(b *testing.B) {
	benchmarkCacheGetManyConcurrent(b, 5*time.Minute)
}

func BenchmarkCacheGetManyConcurrentNotExpiring(b *testing.B) {
	benchmarkCacheGetManyConcurrent(b, pgocache.NoExpiration)
}

func benchmarkCacheGetManyConcurrent(b *testing.B, exp time.Duration) {
	// This is the same as BenchmarkCacheGetConcurrent, but its result
	// can be compared against BenchmarkShardedCacheGetManyConcurrent
	// in sharded_test.go.
	b.StopTimer()
	n := 10000
	tc := Wrap(pgocache.New(exp, 0), "test")
	keys := make([]string, n)
	for i := 0; i < n; i++ {
		k := "foo" + strconv.Itoa(i)
		keys[i] = k
		tc.Set(context.Background(), k, "bar", pgocache.DefaultExpiration)
	}
	each := b.N / n
	wg := new(sync.WaitGroup)
	wg.Add(n)
	for _, v := range keys {
		go func(k string) {
			for j := 0; j < each; j++ {
				tc.Get(context.Background(), k)
			}
			wg.Done()
		}(v)
	}
	b.StartTimer()
	wg.Wait()
}

func BenchmarkCacheSetExpiring(b *testing.B) {
	benchmarkCacheSet(b, 5*time.Minute)
}

func BenchmarkCacheSetNotExpiring(b *testing.B) {
	benchmarkCacheSet(b, pgocache.NoExpiration)
}

func benchmarkCacheSet(b *testing.B, exp time.Duration) {
	b.StopTimer()
	tc := Wrap(pgocache.New(exp, 0), "test")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
	}
}

func BenchmarkRWMutexMapSet(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		mu.Unlock()
	}
}

func BenchmarkCacheSetDelete(b *testing.B) {
	b.StopTimer()
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
		tc.Delete(context.Background(), "foo")
	}
}

func BenchmarkRWMutexMapSetDelete(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		mu.Unlock()
		mu.Lock()
		delete(m, "foo")
		mu.Unlock()
	}
}

func BenchmarkCacheSetDeleteSingleLock(b *testing.B) {
	b.StopTimer()
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.Set(context.Background(), "foo", "bar", pgocache.DefaultExpiration)
		tc.Delete(context.Background(), "foo")
	}
}

func BenchmarkRWMutexMapSetDeleteSingleLock(b *testing.B) {
	b.StopTimer()
	m := map[string]string{}
	mu := sync.RWMutex{}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		mu.Lock()
		m["foo"] = "bar"
		delete(m, "foo")
		mu.Unlock()
	}
}

func BenchmarkIncrementInt(b *testing.B) {
	b.StopTimer()
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")
	tc.Set(context.Background(), "foo", 0, pgocache.DefaultExpiration)
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.IncrementInt(context.Background(), "foo", 1)
	}
}

func BenchmarkDeleteExpiredLoop(b *testing.B) {
	b.StopTimer()
	tc := Wrap(pgocache.New(5*time.Minute, 0), "test")
	for i := 0; i < 100000; i++ {
		tc.Set(context.Background(), strconv.Itoa(i), "bar", pgocache.DefaultExpiration)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tc.DeleteExpired(context.Background())
	}
}

func TestGetWithExpiration(t *testing.T) {
	tc := Wrap(pgocache.New(pgocache.DefaultExpiration, 0), "test")

	a, expiration, found := tc.GetWithExpiration(context.Background(), "a")
	if found || a != nil || !expiration.IsZero() {
		t.Error("Getting A found value that shouldn't exist:", a)
	}

	b, expiration, found := tc.GetWithExpiration(context.Background(), "b")
	if found || b != nil || !expiration.IsZero() {
		t.Error("Getting B found value that shouldn't exist:", b)
	}

	c, expiration, found := tc.GetWithExpiration(context.Background(), "c")
	if found || c != nil || !expiration.IsZero() {
		t.Error("Getting C found value that shouldn't exist:", c)
	}

	tc.Set(context.Background(), "a", 1, pgocache.DefaultExpiration)
	tc.Set(context.Background(), "b", "b", pgocache.DefaultExpiration)
	tc.Set(context.Background(), "c", 3.5, pgocache.DefaultExpiration)
	tc.Set(context.Background(), "d", 1, pgocache.NoExpiration)
	tc.Set(context.Background(), "e", 1, 50*time.Millisecond)

	x, expiration, found := tc.GetWithExpiration(context.Background(), "a")
	if !found {
		t.Error("a was not found while getting a2")
	}
	if x == nil {
		t.Error("x for a is nil")
	} else if a2 := x.(int); a2+2 != 3 {
		t.Error("a2 (which should be 1) plus 2 does not equal 3; value:", a2)
	}
	if !expiration.IsZero() {
		t.Error("expiration for a is not a zeroed time")
	}

	x, expiration, found = tc.GetWithExpiration(context.Background(), "b")
	if !found {
		t.Error("b was not found while getting b2")
	}
	if x == nil {
		t.Error("x for b is nil")
	} else if b2 := x.(string); b2+"B" != "bB" {
		t.Error("b2 (which should be b) plus B does not equal bB; value:", b2)
	}
	if !expiration.IsZero() {
		t.Error("expiration for b is not a zeroed time")
	}

	x, expiration, found = tc.GetWithExpiration(context.Background(), "c")
	if !found {
		t.Error("c was not found while getting c2")
	}
	if x == nil {
		t.Error("x for c is nil")
	} else if c2 := x.(float64); c2+1.2 != 4.7 {
		t.Error("c2 (which should be 3.5) plus 1.2 does not equal 4.7; value:", c2)
	}
	if !expiration.IsZero() {
		t.Error("expiration for c is not a zeroed time")
	}

	x, expiration, found = tc.GetWithExpiration(context.Background(), "d")
	if !found {
		t.Error("d was not found while getting d2")
	}
	if x == nil {
		t.Error("x for d is nil")
	} else if d2 := x.(int); d2+2 != 3 {
		t.Error("d (which should be 1) plus 2 does not equal 3; value:", d2)
	}
	if !expiration.IsZero() {
		t.Error("expiration for d is not a zeroed time")
	}

	x, expiration, found = tc.GetWithExpiration(context.Background(), "e")
	if !found {
		t.Error("e was not found while getting e2")
	}
	if x == nil {
		t.Error("x for e is nil")
	} else if e2 := x.(int); e2+2 != 3 {
		t.Error("e (which should be 1) plus 2 does not equal 3; value:", e2)
	}
	if expiration.UnixNano() != tc.Items(context.Background())["e"].Expiration {
		t.Error("expiration for e is not the correct time")
	}
	if expiration.UnixNano() < time.Now().UnixNano() {
		t.Error("expiration for e is in the past")
	}
}
