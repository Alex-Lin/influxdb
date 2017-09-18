package tsdb

import "github.com/influxdata/influxdb/query"

// EOF represents a "not found" key returned by a Cursor.
const EOF = query.ZeroTime

// Cursor represents an iterator over a series.
type Cursor interface {
	Close()
	SeriesKey() string
	Err() error
}

type IntegerCursor interface {
	Cursor
	Next() (key int64, value int64)
}

type IntegerBatchCursor interface {
	Cursor
	Next() (keys []int64, values []int64)
}

type FloatCursor interface {
	Cursor
	Next() (key int64, value float64)
}

type FloatBatchCursor interface {
	Cursor
	Next() (keys []int64, values []float64)
}

type UnsignedCursor interface {
	Cursor
	Next() (key int64, value uint64)
}

type UnsignedBatchCursor interface {
	Cursor
	Next() (keys []int64, values []uint64)
}

type StringCursor interface {
	Cursor
	Next() (key int64, value string)
}

type StringBatchCursor interface {
	Cursor
	Next() (keys []int64, values []string)
}

type BooleanCursor interface {
	Cursor
	Next() (key int64, value bool)
}

type BooleanBatchCursor interface {
	Cursor
	Next() (keys []int64, values []bool)
}

type CursorRequest struct {
	Measurement string
	Series      string
	Field       string
	Ascending   bool
	StartTime   int64
	EndTime     int64
}
