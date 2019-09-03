package hedera

import (
	"encoding/hex"
	"errors"
	"math/big"
	"strconv"
)

type CallParams struct {
	funcSelector *FunctionSelector
	Args []Argument
}

func NewCallParams(fs *FunctionSelector) *CallParams {
	var args []Argument
	return &CallParams{
		funcSelector: fs,
		Args:         args,
	}
}

func NewConstructorCallParams() *CallParams {
	var fs *FunctionSelector
	return NewCallParams(fs)
}

func (cp *CallParams) addParamType(paramType string) {
	if cp.funcSelector != nil {
		cp.funcSelector.AddParamType(paramType)
	}
}

func (cp *CallParams) AddString(param string) error {
	encString, err := encodeString(param)
	if err != nil { return err }

	arg, err := NewArgument(encString, true)
	if err != nil { return err }

	cp.addParamType("string")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddStringArray(param []string) error {
	var list [][]byte
	for index, _ := range param {
		b, err := encodeString(param[index])
		if err != nil { return err }
		list = append(list, b)
	}

	argBytes, err := encodeByteArray(list, true)
	if err != nil { return err }
	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("string[]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFixedStringArray(param []string, fixedLen int) error {
	err := checkFixedArrayLen(param, fixedLen)
	if err != nil { return err }

	var list [][]byte
	for index, _ := range param {
		b, err := encodeString(param[index])
		if err != nil { return err }
		list = append(list, b)
	}

	argBytes, err := encodeByteArray(list, true)
	if err != nil { return err }
	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("string[" + strconv.Itoa(fixedLen) + "]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddBytes(param []byte) error {
	encBytes, err := encodeBytes(param)
	if err != nil { return err }

	arg, err := NewArgument(encBytes, true)
	if err != nil { return err }

	cp.addParamType("bytes")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFixedBytes(param []byte, fixedLen int) error {
	err := checkFixedArrayLen(param, fixedLen)
	if err != nil { return err }

	if fixedLen > 32 {
		return errors.New("ILLEGAL ARGUMENT ERROR: bytesN cannot have a length greater than 32; given length: " +
			strconv.Itoa(fixedLen))
	}

	encBytes := encodeFixedBytes(param)

	arg, err := NewArgument(encBytes, false)
	if err != nil { return err }

	cp.addParamType("bytes" + strconv.Itoa(fixedLen))
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddByteArray(param [][]byte) error {
	var list [][]byte
	for index, _ := range param {
		b, err := encodeBytes(param[index])
		if err != nil { return err }
		list = append(list, b)
	}

	argBytes, err := encodeByteArray(list, true)
	if err != nil { return err }
	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("bytes[]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFixedByteArray(param [][]byte, byteLen int) error {
	for _, v := range param {
		err := checkFixedArrayLen(v, byteLen)
		if err != nil { return err }
	}

	if byteLen > 32 {
		return errors.New("ILLEGAL ARGUMENT ERROR: bytesN cannot have a length greater than 32; given length: " +
			strconv.Itoa(byteLen))
	}

	var list [][]byte
	for _, v := range param {
		b := encodeFixedBytes(v)
		list = append(list, b)
	}

	argBytes, err := encodeByteArray(list, true)
	if err != nil { return err }
	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("bytes" + strconv.Itoa(byteLen)  + "[]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddByteFixedArray(param [][]byte, fixedLen int) error {
	err := checkFixedArrayLen(param, fixedLen)
	if err != nil { return err }

	var list [][]byte
	for _, v := range param {
		b, err := encodeBytes(v)
		if err != nil { return err }
		list = append(list, b)
	}

	argBytes, err := encodeByteArray(list, true)
	if err != nil { return err }
	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("bytes[" + strconv.Itoa(fixedLen) + "]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFixedByteFixedArray(param [][]byte, fixedByteLen int, fixedLen int) error {
	err := checkFixedArrayLen(param, fixedLen)
	if err != nil { return err }

	for _, v := range param {
		err := checkFixedArrayLen(v, fixedByteLen)
		if err != nil { return err }
	}

	var list [][]byte
	for _, v := range param {
		b := encodeFixedBytes(v)
		list = append(list, b)
	}

	argBytes, err := encodeByteArray(list, true)
	if err != nil { return err }
	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("bytes" + strconv.Itoa(fixedByteLen) + "[" + strconv.Itoa(fixedLen) + "]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddBool(b bool) error {
	val := 0
	if b == true { val = 1 }

	encBool, err := int256(int64(val), 8)
	if err != nil { return err }

	arg, err := NewArgument(encBool, false)
	if err != nil { return err }

	cp.addParamType("bool")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddInt(value int64, width int) error {
	err := checkIntWidth(width)
	if err != nil { return err }

	encInt, err := int256(value, width)
	if err != nil { return err }

	arg, err := NewArgument(encInt, false)
	if err != nil { return err }

	cp.addParamType("int" + strconv.Itoa(width))
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddBigInt(value big.Int, width int) error {
	err := checkBigInt(value, width, true)
	if err != nil { return err }

	encInt, err := bigInt256(value)
	if err != nil { return err }
	arg, err := NewArgument(encInt, false)
	if err != nil { return err }

	cp.addParamType("int" + strconv.Itoa(width))
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddIntArray(value []int, width int) error {
	argBytes, err := encodeIntArray(value, width, true)
	if err != nil { return err }

	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("int" + strconv.Itoa(width) + "[]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFixedIntArray(value []int, width int, fixedLen int) error {
	err := checkFixedArrayLen(value, fixedLen)
	if err != nil { return err }

	argBytes, err := encodeIntArray(value, width, true)
	if err != nil { return err }

	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("int" + strconv.Itoa(width) + "[" + strconv.Itoa(fixedLen) + "]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddBitIntArray(value []big.Int, width int) error {
	argBytes, err := encodeBigIntArray(value, width, true)
	if err != nil { return err }

	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("int" + strconv.Itoa(width) + "[]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFixedBigIntArray(value []big.Int, width int, fixedLen int) error {
	err := checkFixedArrayLen(value, fixedLen)
	if err != nil { return err }

	argBytes, err := encodeBigIntArray(value, width, true)
	if err != nil { return err }

	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("int" + strconv.Itoa(width) + "[" + strconv.Itoa(fixedLen) + "]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddUint(value uint64, width int) error {
	err := checkIntWidth(width)
	if err != nil { return err }

	encVal, err := uint256(value)
	if err != nil { return err }

	arg, err := NewArgument(encVal, false)
	if err != nil { return err }

	cp.addParamType("uint" + strconv.Itoa(width))
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddUintArray(value []uint, width int) error {
	argBytes, err := encodeUintArray(value, width, true)
	if err != nil { return err }

	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("uint" + strconv.Itoa(width) + "[]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFixedUintArray(value []uint, width int, fixedLen int) error {
	err := checkFixedArrayLen(value, fixedLen)
	if err != nil { return err }

	argBytes, err := encodeUintArray(value, width, true)
	if err != nil { return err }

	arg, err := NewArgument(argBytes, true)
	if err != nil { return err }

	cp.addParamType("uint" + strconv.Itoa(width) + "[" + strconv.Itoa(fixedLen) + "]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddAddress(address []byte) error {
	err := CheckAddressLen(address)
	if err != nil { return err }

	arg, err := NewArgument(leftPad(address, false), false)
	if err != nil { return err }

	cp.addParamType("address")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddAddressString(address string) error {
	addr, err := DecodeAddress(address)
	if err != nil { return err }

	return cp.AddAddress(addr)
}

func (cp *CallParams) AddAddressArray(addresses [][]byte) error {
	var buf [][]byte
	for _, v := range addresses {
		err := CheckAddressLen(v)
		if err != nil { return err }
		buf = append(buf, leftPad(v, false))
	}

	encAddrs, err := encodeByteArray(buf, true)
	if err != nil { return err }

	arg, err := NewArgument(encAddrs, true)
	if err != nil { return err }

	cp.addParamType("address[]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFixedAddressArray(addresses [][]byte, fixedLen int) error {
	err := checkFixedArrayLen(addresses, fixedLen)
	if err != nil { return err }

	var buf [][]byte
	for _, v := range addresses {
		err := CheckAddressLen(v)
		if err != nil { return err }
		buf = append(buf, leftPad(v, false))
	}

	encAddrs, err := encodeByteArray(buf, true)
	if err != nil { return err }

	arg, err := NewArgument(encAddrs, true)
	if err != nil { return err }

	cp.addParamType("address[" + strconv.Itoa(fixedLen) + "]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddAddressStringArray(addresses []string) error {
	var buf [][]byte
	for _, v := range addresses {
		addr, err := DecodeAddress(v)
		if err != nil { return err }
		buf = append(buf, leftPad(addr, false))
	}

	encAddrs, err := encodeByteArray(buf, true)
	if err != nil { return err }

	arg, err := NewArgument(encAddrs, true)
	if err != nil { return err }

	cp.addParamType("address[]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFixedAddressStringArray(addresses []string, fixedLen int) error {
	err := checkFixedArrayLen(addresses, fixedLen)
	if err != nil { return err }

	var buf [][]byte
	for _, v := range addresses {
		addr, err := DecodeAddress(v)
		if err != nil { return err }
		buf = append(buf, leftPad(addr, false))
	}

	encAddrs, err := encodeByteArray(buf, true)
	if err != nil { return err }

	arg, err := NewArgument(encAddrs, true)
	if err != nil { return err }

	cp.addParamType("address[" + strconv.Itoa(fixedLen) + "]")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFunction(address []byte, selector []byte) error {
	err := CheckAddressLen(address)
	if err != nil { return err }

	if len(selector) != SELECTOR_LEN {
		return errors.New("ILLEGAL ARGUMENT ERROR: function selectors must be 4 bytes or 8 hex chars")
	}

	var output []byte
	output = append(output, address...)
	output = append(output, selector...)

	arg, err := NewArgument(rightPad(output), false)
	if err != nil { return err }

	cp.addParamType("function")
	cp.Args = append(cp.Args, *arg)
	return nil
}

func (cp *CallParams) AddFunctionString(address string, selector string) error {
	if len(selector) != SELECTOR_LEN_HEX {
		return errors.New("ILLEGAL ARGUMENT ERROR: function selectors must be 4 bytes or 8 hex chars")
	}

	selectorBytes, err := hex.DecodeString(selector)
	if err != nil {
		return errors.New("failed to decode Solidity function selector as hex; " + err.Error())
	}

	addressBytes, err := DecodeAddress(address)
	if err != nil { return err }

	return cp.AddFunction(addressBytes, selectorBytes)
}

func (cp *CallParams) AddFunctionFS(address string, selector FunctionSelector) error {
	addr, err := DecodeAddress(address)
	if err != nil { return err }

	fs := selector.FinishIntermediate()

	return cp.AddFunction(addr, fs[0:4])
}

func (cp *CallParams) ToProto() ([]byte, error) {
	dynamicOffset := len(cp.Args) * 32

	var paramBytes [][]byte

	if cp.funcSelector != nil {
		fs := cp.funcSelector.FinishIntermediate()
		paramBytes = append(paramBytes, fs[0:4])
	}

	var dynamicArgs [][]byte

	for _, arg := range cp.Args {
		if arg.Dynamic == true {
			offset, err := int256(int64(dynamicOffset), 256)
			if err != nil { return nil, err }

			paramBytes = append(paramBytes, offset)
			dynamicArgs = append(dynamicArgs, arg.Value)
			dynamicOffset += len(arg.Value)
		} else {
			paramBytes = append(paramBytes, arg.Value)
		}
	}

	paramBytes = append(paramBytes, dynamicArgs...)

	var out []byte
	for _, v := range paramBytes {
		out = append(out, v...)
	}

	return out, nil
}
