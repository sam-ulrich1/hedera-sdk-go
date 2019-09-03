package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"
import (
	"encoding/binary"
	"encoding/hex"
	"math/big"
	"unsafe"
)

type ContractLogInfo struct {
	Contract ContractID
	Bloom []byte
	Topic [][]byte
	Data []byte
}

func cContractLogInfo(cli ContractLogInfo) C.HederaContractLogInfo {
	return C.HederaContractLogInfo{
		contract_id:   cContractID(cli.Contract),
		bloom:   C.HederaArray{
			ptr: unsafe.Pointer(&cli.Bloom),
			len: C.size_t(len(cli.Bloom)),
		},
		topic: C.HederaArray{
			ptr: unsafe.Pointer(&cli.Topic),
			len: C.size_t(len(cli.Topic)),
		},
		data: C.HederaArray{
			ptr: unsafe.Pointer(&cli.Data),
			len: C.size_t(len(cli.Data)),
		},
	}
}

func goContractLogInfo(cli C.HederaContractLogInfo) ContractLogInfo {
	return ContractLogInfo{
		Contract:   goContractID(cli.contract_id),
		Bloom: *(*[]byte)(unsafe.Pointer(&cli.bloom)),
		Topic: *(*[][]byte)(unsafe.Pointer(&cli.topic)),
		Data: *(*[]byte)(unsafe.Pointer(&cli.data)),
	}
}

type ContractFunctionResult struct {
	Contract ContractID
	Result []byte
	Error string
	Bloom []byte
	GasUsed uint64
	LogInfo []ContractLogInfo
}

func cContractFunctionResult(fs ContractFunctionResult) C.HederaContractFunctionResult {
	return C.HederaContractFunctionResult{
		contract_id:   cContractID(fs.Contract),
		contract_call_result:   C.HederaArray{
			ptr: unsafe.Pointer(&fs.Result),
			len: C.size_t(len(fs.Result)),
		},
		error_message: C.CString(fs.Error),
		bloom: C.HederaArray{
			ptr: unsafe.Pointer(&fs.Bloom),
			len: C.size_t(len(fs.Bloom)),
		},
		gas_used: C.uint64_t(fs.GasUsed),
		log_info: C.HederaArray{
			ptr: unsafe.Pointer(&fs.LogInfo),
			len: C.size_t(len(fs.LogInfo)),
		},
	}
}

func goContractFunctionResult(fs C.HederaContractFunctionResult) ContractFunctionResult {
	cLogInfo := *(*[]C.HederaContractLogInfo)(unsafe.Pointer(&fs.log_info))

	var logInfo []ContractLogInfo
	for _, cli := range cLogInfo {
		logInfo = append(logInfo, goContractLogInfo(cli))
	}

	return ContractFunctionResult{
		Contract:   goContractID(fs.contract_id),
		Result: *(*[]byte)(unsafe.Pointer(&fs.contract_call_result)),
		Error: C.GoString(fs.error_message),
		Bloom: *(*[]byte)(unsafe.Pointer(&fs.bloom)),
		GasUsed: uint64(fs.gas_used),
		LogInfo: logInfo,
	}
}

func (fr *ContractFunctionResult) getByteBuffer(offset int) byte {
	return fr.Result[offset]
}

func (fr *ContractFunctionResult) getIntValueAt(valueOffset int) int {
	return int(fr.getByteBuffer(valueOffset+28))
}

func (fr *ContractFunctionResult) getInt256(valIndex int) []byte {
	return fr.Result[valIndex * 32 : (valIndex + 1) * 32]
}

func (fr *ContractFunctionResult) getArrayLength(offset int) int64 {
	lengthStart := offset + 24
	lengthEnd := lengthStart + 8
	lengthBytes := fr.Result[lengthStart:lengthEnd]
	return int64(binary.BigEndian.Uint64(lengthBytes))
}

func (fr *ContractFunctionResult) GetInt(valIndex int) int {
	return fr.getIntValueAt(valIndex * 32)
}

func (fr *ContractFunctionResult) GetLong(valIndex int) int64 {
	return int64(fr.getIntValueAt(valIndex * 32 + 24))
}

func (fr *ContractFunctionResult) GetBigInt(valIndex int) big.Int {
	return *big.NewInt(int64(fr.getIntValueAt(valIndex * 32)))
}

func (fr *ContractFunctionResult) GetBytes(valIndex int) ([]byte, error) {
	offset := fr.GetInt(valIndex)
	l := fr.getIntValueAt(int(offset))
	return fr.Result[offset + 32 : offset + 32 + l], nil
}

func (fr *ContractFunctionResult) GetByteArray(valIndex int) [][]byte {
	offset := (valIndex * 32) + 32
	start := offset + 32
	length := fr.getArrayLength(offset)
	var bytes [][]byte
	for i := 0; i < int(length); i++ {
		bOffset := (i * 32) + start
		byteChunk := fr.Result[bOffset : bOffset + 32]
		bytes = append(bytes, byteChunk)
	}
	return bytes
}

func (fr *ContractFunctionResult) GetString(valIndex int) (string, error) {
	s, err := fr.GetBytes(valIndex)
	return string(s), err
}

func (fr *ContractFunctionResult) GetBool(valIndex int) bool {
	return fr.getByteBuffer(valIndex * 32 + 31) != 0
}

func (fr *ContractFunctionResult) GetAddress(valIndex int) []byte {
	offset := valIndex * 32
	return fr.Result[offset + 12 : offset + 32]
}

func (fr *ContractFunctionResult) GetAddressArray(valIndex int) []string {
	offset := (valIndex * 32) + 32
	start := offset + 32
	length := fr.getArrayLength(offset)
	var addrs []string
	for i := 0; i < int(length); i++ {
		addrOffset := (i * 32) + start
		byteChunk := fr.Result[addrOffset + 12 : addrOffset + 32]
		addrs = append(addrs, hex.EncodeToString(byteChunk))
	}
	return addrs
}
