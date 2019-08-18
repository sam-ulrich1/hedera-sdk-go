package hedera

// #include "hedera.h"
import "C"

type Duration struct {
	Seconds uint64
	Nanos uint32
}

func NewDuration(seconds uint64, nanos uint32) *Duration {
	return &Duration{
		Seconds: seconds,
		Nanos:   nanos,
	}
}

func cDuration(duration Duration) C.HederaDuration {
	return C.HederaDuration{
		seconds: C.uint64_t(duration.Seconds),
		nanos: C.uint32_t(duration.Nanos),
	}
}

func goDuration(duration C.HederaDuration) Duration {
	return Duration{
		Seconds: uint64(duration.seconds),
		Nanos: uint32(duration.nanos),
	}
}
