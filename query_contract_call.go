package hedera

// #include "hedera.h"
import "C"
import "unsafe"

type QueryContractCall struct {
	query
}

func newQueryContractCall(client *Client, contract ContractID, gas int64, params []byte,
	maxResultSize int64) QueryContractCall {
	cParams := (*C.uint8_t)(unsafe.Pointer(&[]byte{}))
	if len(params) > 0 {
		cParams = (*C.uint8_t)(&params[0])
	}
	return QueryContractCall{
		query{C.hedera_query__contract_call__new(client.inner, cContractID(contract),
			C.uint64_t(gas), cParams, C.size_t(len(params)), C.uint64_t(maxResultSize))},
	}
}

func (query QueryContractCall)Execute() (ContractFunctionResult, error) {
	var funcRes C.HederaContractFunctionResult
	err := C.hedera_query__contract_call__execute(query.inner, &funcRes)
	if err != 0 {
		return ContractFunctionResult{}, hederaLastError()
	}
	return goContractFunctionResult(funcRes), nil
}

