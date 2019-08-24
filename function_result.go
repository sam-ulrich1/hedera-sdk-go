package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"
import (
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
