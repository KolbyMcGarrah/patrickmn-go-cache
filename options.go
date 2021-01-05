package cache

import "go.opencensus.io/trace"

const defaultInstanceName = "default"

// TraceOption allows for managing cache trace configurations using funcitonal options
type TraceOption func(o *TraceOptions)

// TraceOptions holds configurations of the go-cache tracing middleware
// by default all options are initialized to false.
type TraceOptions struct {
	// AllowRoot, if set to true, will allow go-cache to create root spans in
	// absence of existing spans or even context.
	// Default is to not trace go-cache calls if no existing parent span is found
	// in context or when using methods not taking context.
	AllowRoot bool

	// InstanceName identifies the cache
	InstanceName string

	// DefaultAttributes will set to each span as default
	DefaultAttributes []trace.Attribute

	// Sampler to use when creating spans
	Sampler trace.Sampler

	// Setting the below options will control whether or not spans are created
	// on their call.
	Add               bool
	Decrement         bool
	DecrementFloat    bool
	DecrementFloat32  bool
	DecrementFloat64  bool
	DecrementInt      bool
	DecrementInt16    bool
	DecrementInt32    bool
	DecrementInt64    bool
	DecrementInt8     bool
	DecrementUint     bool
	DecrementUint16   bool
	DecrementUint32   bool
	DecrementUint64   bool
	DecrementUint8    bool
	DecrementUintptr  bool
	Delete            bool
	DeleteExpired     bool
	Flush             bool
	Get               bool
	GetWithExpiration bool
	Increment         bool
	IncrementFloat    bool
	IncrementFloat32  bool
	IncrementFloat64  bool
	IncrementInt      bool
	IncrementInt16    bool
	IncrementInt32    bool
	IncrementInt64    bool
	IncrementInt8     bool
	IncrementUint     bool
	IncrementUint16   bool
	IncrementUint32   bool
	IncrementUint64   bool
	IncrementUint8    bool
	IncrementUintptr  bool
	ItemCount         bool
	Items             bool
	Load              bool
	LoadFile          bool
	OnEvicted         bool
	Replace           bool
	Save              bool
	SaveFile          bool
	Set               bool
	SetDefault        bool
}

// WithAllTraceOptions enables all available traceoptions
func WithAllTraceOptions() TraceOption {
	return func(o *TraceOptions) {
		*o = AllTraceOptions
	}
}

// AllTraceOptions has all tracing options enabled
var AllTraceOptions = TraceOptions{
	Add:               true,
	Decrement:         true,
	DecrementFloat:    true,
	DecrementFloat32:  true,
	DecrementFloat64:  true,
	DecrementInt:      true,
	DecrementInt16:    true,
	DecrementInt32:    true,
	DecrementInt64:    true,
	DecrementInt8:     true,
	DecrementUint:     true,
	DecrementUint16:   true,
	DecrementUint32:   true,
	DecrementUint64:   true,
	DecrementUint8:    true,
	DecrementUintptr:  true,
	Delete:            true,
	DeleteExpired:     true,
	Flush:             true,
	Get:               true,
	GetWithExpiration: true,
	Increment:         true,
	IncrementFloat:    true,
	IncrementFloat32:  true,
	IncrementFloat64:  true,
	IncrementInt:      true,
	IncrementInt16:    true,
	IncrementInt32:    true,
	IncrementInt64:    true,
	IncrementInt8:     true,
	IncrementUint:     true,
	IncrementUint16:   true,
	IncrementUint32:   true,
	IncrementUint64:   true,
	IncrementUint8:    true,
	IncrementUintptr:  true,
	ItemCount:         true,
	Items:             true,
	Load:              true,
	LoadFile:          true,
	OnEvicted:         true,
	Replace:           true,
	Save:              true,
	SaveFile:          true,
	Set:               true,
	SetDefault:        true,
}

// WithOptions sets the go-cache tracing options with a single TraceOptions object
func WithOptions(options TraceOptions) TraceOption {
	return func(o *TraceOptions) {
		*o = options
		o.DefaultAttributes = append(
			[]trace.Attribute(nil), options.DefaultAttributes...,
		)
	}
}

// WithAllowRoot if set to true, will allow go-cache to create root spans in
// absence of exisiting spans or even context.
// Default is to not trace go-cache calls if no existing parent span is found
// in context or when using methods not taking context.
func WithAllowRoot(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.AllowRoot = b
	}
}

// WithInstanceName sets cache instance name.
func WithInstanceName(instanceName string) TraceOption {
	return func(o *TraceOptions) {
		o.InstanceName = instanceName
	}
}

// WithAdd if set to true, will allow spans on Add
func WithAdd(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Add = b
	}
}

// WithDecrement if set to true, will allow spans on Decrement
func WithDecrement(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Decrement = b
	}
}

// WithDecrementFloat if set to true, will allow spans on DecrementFloat
func WithDecrementFloat(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementFloat = b
	}
}

// WithDecrementFloat32 if set to true, will allow spans on DecrementFloat32
func WithDecrementFloat32(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementFloat32 = b
	}
}

// WithDecrementFloat64 if set to true, will allow spans on DecrementFloat64
func WithDecrementFloat64(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementFloat64 = b
	}
}

// WithDecrementInt if set to true, will allow spans on DecrementInt
func WithDecrementInt(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementInt = b
	}
}

// WithDecrementInt16 if set to true, will allow spans on DecrementInt16
func WithDecrementInt16(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementInt16 = b
	}
}

// WithDecrementInt64 if set to true, will allow spans on DecrementInt64
func WithDecrementInt64(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementInt64 = b
	}
}

// WithDecrementInt8 if set to true, will allow spans on DecrementInt8
func WithDecrementInt8(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementInt8 = b
	}
}

// WithDecrementUint if set to true, will allow spans on DecrementUint
func WithDecrementUint(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementUint = b
	}
}

// WithDecrementUint16 if set to true, will allow spans on DecrementUint16
func WithDecrementUint16(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementUint16 = b
	}
}

// WithDecrementUint32 if set to true, will allow spans on DecrementUint32
func WithDecrementUint32(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementUint32 = b
	}
}

// WithDecrementUint64 if set to true, will allow spans on DecrementUint64
func WithDecrementUint64(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementUint64 = b
	}
}

// WithDecrementUint8 if set to true, will allow spans on DecrementUint8
func WithDecrementUint8(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementUint8 = b
	}
}

// WithDecrementUintptr if set to true, will allow spans on DecrementUintptr
func WithDecrementUintptr(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DecrementUintptr = b
	}
}

// WithDelete if set to true, will allow spans on Delete
func WithDelete(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Delete = b
	}
}

// WithDeleteExpired if set to true, will allow spans on DeleteExpired
func WithDeleteExpired(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.DeleteExpired = b
	}
}

// WithFlush if set to true, will allow spans on Flush
func WithFlush(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Flush = b
	}
}

// WithGet if set to true, will allow spans on Get
func WithGet(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Get = b
	}
}

// WithGetWithExpiration if set to true, will allow spans on GetWithExpiration
func WithGetWithExpiration(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.GetWithExpiration = b
	}
}

// WithIncrement if set to true, will allow spans on Increment
func WithIncrement(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Increment = b
	}
}

// WithIncrementFloat if set to true, will allow spans on IncrementFloat
func WithIncrementFloat(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementFloat = b
	}
}

// WithIncrementFloat32 if set to true, will allow spans on IncrementFloat32
func WithIncrementFloat32(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementFloat32 = b
	}
}

// WithIncrementFloat64 if set to true, will allow spans on IncrementFloat64
func WithIncrementFloat64(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementFloat64 = b
	}
}

// WithIncrementInt if set to true, will allow spans on IncrementInt
func WithIncrementInt(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementInt = b
	}
}

// WithIncrementInt16 if set to true, will allow spans on IncrementInt16
func WithIncrementInt16(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementInt16 = b
	}
}

// WithIncrementInt32 if set to true, will allow spans on IncrementInt32
func WithIncrementInt32(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementInt32 = b
	}
}

// WithIncrementInt64 if set to true, will allow spans on IncrementInt64
func WithIncrementInt64(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementInt64 = b
	}
}

// WithIncrementInt8 if set to true, will allow spans on IncrementInt8
func WithIncrementInt8(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementInt8 = b
	}
}

// WithIncrementUint if set to true, will allow spans on IncrementUint
func WithIncrementUint(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementUint = b
	}
}

// WithIncrementUint16 if set to true, will allow spans on IncrementUint16
func WithIncrementUint16(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementUint16 = b
	}
}

// WithIncrementUint32 if set to true, will allow spans on IncrementUint32
func WithIncrementUint32(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementUint32 = b
	}
}

// WithIncrementUint64 if set to true, will allow spans on IncrementUint64
func WithIncrementUint64(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementUint64 = b
	}
}

// WithIncrementUint8 if set to true, will allow spans on IncrementUint8
func WithIncrementUint8(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementUint8 = b
	}
}

// WithIncrementUintptr if set to true, will allow spans on IncrementUintptr
func WithIncrementUintptr(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.IncrementUintptr = b
	}
}

// WithItemCount if set to true, will allow spans on ItemCount
func WithItemCount(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.ItemCount = b
	}
}

// WithItems if set to true, will allow spans on Items
func WithItems(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Items = b
	}
}

// WithLoad if set to true, will allow spans on Load
func WithLoad(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Load = b
	}
}

// WithLoadFile if set to true, will allow spans on LoadFile
func WithLoadFile(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.LoadFile = b
	}
}

// WithOnEvicted if set to true, will allow spans on OnEvicted
func WithOnEvicted(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.OnEvicted = b
	}
}

// WithReplace if set to true, will allow spans on Replace
func WithReplace(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Replace = b
	}
}

// WithSave if set to true, will allow spans on Save
func WithSave(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Save = b
	}
}

// WithSaveFile if set to true, will allow spans on SaveFile
func WithSaveFile(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.SaveFile = b
	}
}

// WithSet if set to true, will allow spans on Set
func WithSet(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.Set = b
	}
}

// WithSetDefault if set to true, will allow spans on SetDefault
func WithSetDefault(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.SetDefault = b
	}
}
