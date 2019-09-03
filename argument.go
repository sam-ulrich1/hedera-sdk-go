package hedera

import (
	"errors"
	"strconv"
)

type Argument struct {
	Value []byte
	Dynamic bool
}

func NewArgument(value []byte, dynamic bool) (*Argument, error) {
	if dynamic == false && len(value) != 32 {
		return nil, errors.New("ILLEGAL ARGUMENT ERROR: value argument that was not 32 bytes; value was " +
			strconv.Itoa(len(value)) + " bytes")
	}
	return &Argument{
		Value:   value,
		Dynamic: dynamic,
	}, nil
}
