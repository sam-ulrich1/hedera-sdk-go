package hedera

import (
	"bytes"
	"encoding/binary"
	"errors"
	"math"
	"math/big"
	"reflect"
	"strconv"
)

func getBigIntBitLength(i big.Int) (int, error) {
	actualBitLen := i.BitLen()
	if actualBitLen < 0 {
		return -1, errors.New("ILLEGAL ARGUMENT ERROR: Big Int must have a positive bit length")
	} else if actualBitLen > 256 {
		return -1, errors.New("ILLEGAL ARGUMENT ERROR: " +
			"Solidity integer width must be a multiple of 8, in the closed range [8, 256]")
	} else if 0 <= actualBitLen && actualBitLen <= 8 {
		return 8, nil
	} else if 9 <= actualBitLen && actualBitLen <= 16 {
		return 16, nil
	} else if 17 <= actualBitLen && actualBitLen <= 32 {
		return 32, nil
	} else if 33 <= actualBitLen && actualBitLen <= 64 {
		return 64, nil
	} else if 65 <= actualBitLen && actualBitLen <= 128 {
		return 128, nil
	} else if 129 <= actualBitLen && actualBitLen <= 256 {
		return 256, nil
	}
	return -1, errors.New("ILLEGAL ARGUMENT ERROR: Big Int bit length was not identifiable")
}

func checkIntWidth(width int) error {
	if width % 8 != 0 || width < 8 || width > 256 {
		return errors.New("ILLEGAL ARGUMENT ERROR: " +
			"Solidity integer width must be a multiple of 8, in the closed range [8, 256]")
	}
	return nil
}

func checkBigInt(val big.Int, width int, signed bool) error {
	err := checkIntWidth(width)
	if err != nil { return err }

	actualBitLen := val.BitLen()
	if signed == true { actualBitLen++ }

	if actualBitLen > 256 {
		return errors.New("ILLEGAL ARGUMENT ERROR: Big Int out of range for Solidity integers")
	}

	if width < actualBitLen {
		return errors.New("ILLEGAL ARGUMENT ERROR: Big Int bit length is greater than the nominal parameter width")
	}

	return nil
}

func checkFixedArrayLen(arr interface{}, fixedLen int) error {
	arrVal := reflect.ValueOf(arr)
	if arrVal.Len() != fixedLen {
		return errors.New("ILLEGAL ARGUMENT ERROR: fixedLen (" + strconv.Itoa(fixedLen) +
			") does not match array length (" + strconv.Itoa(arrVal.Len()) + ")")
	}
	return nil
}

/////////////////////////////////////////////////////////////
// Modified version of go-ethereum math library functions found at:
// https://github.com/ethereum/go-ethereum/blob/dbb03fe9893dd19f6b1de1ee3b768317f22fd135/common/math/big.go

// Originally ReadBits
func readBigIntBits(bigint *big.Int, buf []byte) {
	wordBits := 32 << (uint64(^big.Word(0)) >> 63)
	wordBytes := wordBits / 8

	i := len(buf)
	for _, d := range bigint.Bits() {
		for j := 0; j < wordBytes && i > 0; j++ {
			i--
			buf[i] = byte(d)
			d >>= 8
		}
	}
}

// Originally PaddedBytes
func padBigBytes(bigint *big.Int, n int) []byte {
	if bigint.BitLen()/8 >= n {
		return bigint.Bytes()
	}
	ret := make([]byte, n)
	readBigIntBits(bigint, ret)
	return ret
}
/////////////////////////////////////////////////////////////

func createPadding() []byte {
	var nPad []byte
	for i := 0; i < 31; i++ {
		nPad = append(nPad, 0)
	}
	return nPad
}

func createNegativePadding() []byte {
	var nPad []byte
	for i := 0; i < 31; i++ {
		nPad = append(nPad, byte(0xFF))
	}
	return nPad
}

func leftPad(input []byte, negative bool) []byte {
	rem := 32 - len(input) % 32
	if rem == 32 { return input }

	var padding []byte
	if negative == true {
		padding = createNegativePadding()
	} else {
		padding = createPadding()
	}

	padding = padding[0:rem]
	return append(padding, input...)
}

func rightPad(input []byte) []byte {
	rem := 32 - len(input) % 32
	if rem == 32 { return input }

	return append(input, createPadding()[0:rem]...)
}

func int256(val int64, bitWidth int) ([]byte, error) {
	bitWidth = int(math.Min(float64(bitWidth), 64))
	buf := bytes.Buffer{}

	err := binary.Write(&buf, binary.BigEndian, val)
	if err != nil { return nil, err }
	output := buf.Bytes()

	return leftPad(output, val < 0), nil
}

func uint256(val uint64) ([]byte, error) {
	buf := bytes.Buffer{}
	err := binary.Write(&buf, binary.BigEndian, val)
	if err != nil { return nil, err }

	return leftPad(buf.Bytes(), false), nil
}

func bigInt256(val big.Int) ([]byte, error) {
	bitLen, err := getBigIntBitLength(val)
	if err != nil { return nil, err }

	cmp := val.Cmp(big.NewInt(0))
	return leftPad(padBigBytes(&val, bitLen), cmp > -1), nil
}

func encodeBytes(b []byte) ([]byte, error) {
	ib, err := int256(int64(len(b)), 32)
	if err != nil { return nil, err }
	return append(ib, rightPad(b)...), nil
}

func encodeString(str string) ([]byte, error) {
	b := []byte(str)
	return encodeBytes(b)
}

func encodeFixedBytes(b []byte) []byte {
	return leftPad(b, false)
}

func encodeByteArray(byteArray [][]byte, prependLen bool) ([]byte, error) {
	var list []byte
	for _, v := range byteArray {
		list = append(list, v...)
	}

	if prependLen == true {
		encBytes, err := int256(int64(len(byteArray)), 32)
		if err != nil { return nil, err }
		return append(encBytes, list...), nil
	}

	return list, nil
}

func encodeIntArray(intArray []int, intWidth int, prependLen bool) ([]byte, error) {
	err := checkIntWidth(intWidth)
	if err != nil { return nil, err }

	var b []byte
	for _, v := range intArray {
		i, err := int256(int64(v), intWidth)
		if err != nil { return nil, err }
		b = append(b, i...)
	}

	if prependLen == true {
		encBytes, err := int256(int64(len(intArray)), 32)
		if err != nil { return nil, err }
		return append(encBytes, b...), nil
	}
	return b, nil
}

func encodeBigIntArray(intArray []big.Int, intWidth int, prependLen bool) ([]byte, error) {
	err := checkIntWidth(intWidth)
	if err != nil { return nil, err }

	var b []byte
	for _, v := range intArray {
		i, err := bigInt256(v)
		if err != nil { return nil, err }
		b = append(b, i...)
	}

	if prependLen == true {
		encBytes, err := int256(int64(len(intArray)), 32)
		if err != nil { return nil, err }
		return append(encBytes, b...), nil
	}
	return b, nil
}

func encodeUintArray(intArray []uint, intWidth int, prependLen bool) ([]byte, error) {
	err := checkIntWidth(intWidth)
	if err != nil { return nil, err }

	var b []byte
	for _, v := range intArray {
		encV, err := uint256(uint64(v))
		if err != nil { return nil, err }

		b = append(b, encV...)
	}

	if prependLen == true {
		encBytes, err := int256(int64(len(intArray)), 32)
		if err != nil { return nil, err }
		return append(encBytes, b...), nil
	}
	return b, nil
}
