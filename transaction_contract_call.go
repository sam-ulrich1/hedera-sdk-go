package hedera

// #include "hedera.h"
import "C"

type TransactionContractCall struct {
	transaction
}

func newTransactionContractCall(client *Client, contract ContractID) TransactionContractCall {
	return TransactionContractCall{transaction{
		C.hedera_transaction__contract_call__new(client.inner, cContractID(contract))}}
}

func (tx TransactionContractCall) Gas(gas uint64) TransactionContractCall {
	C.hedera_transaction__contract_call__set_gas(tx.inner, C.uint64_t(gas))
	return tx
}

func (tx TransactionContractCall) Amount(amount uint64) TransactionContractCall {
	C.hedera_transaction__contract_call__set_amount(tx.inner, C.uint64_t(amount))
	return tx
}

func (tx TransactionContractCall) Parameters(params []byte) TransactionContractCall {
	C.hedera_transaction__contract_call__set_function_parameters(tx.inner, (*C.uint8_t)(&params[0]), C.size_t(len(params)))
	return tx
}
