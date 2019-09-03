package hedera

// #include "hedera.h"
import "C"
import (
	"time"
)

type Timestamp struct {
	Seconds int64
	Nanos uint32
}

func TimeNow() *Timestamp {
	t := time.Now()
	secs := t.Unix()
	nans := t.UnixNano()
	nans -= secs * 1000000000
	return &Timestamp{
		Seconds: secs,
		Nanos:   uint32(nans),
	}
}

func (ts  *Timestamp) AddSeconds(secs int64) {
	ts.Seconds += secs
}

func (ts  *Timestamp) AddNanos(nans uint32) {
	ts.Nanos += nans
}

func cTimestamp(ts Timestamp) C.HederaTimestamp {
	return C.HederaTimestamp{
		seconds: C.int64_t(ts.Seconds),
		nanos: C.uint32_t(ts.Nanos),
	}
}

func goTimestamp(ts C.HederaTimestamp) Timestamp {
	return Timestamp{
		Seconds: int64(ts.seconds),
		Nanos: uint32(ts.nanos),
	}
}
