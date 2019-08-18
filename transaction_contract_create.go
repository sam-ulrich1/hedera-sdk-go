package hedera

// #include "hedera.h"
import "C"
import "unsafe"

type TransactionContractCreate struct {
	transaction
}

func newTransactionContractCreate(client *Client) TransactionContractCreate {
	return TransactionContractCreate{transaction{
		C.hedera_transaction__contract_create__new(client.inner)}}
}

func (tx TransactionContractCreate) File(file FileID) TransactionContractCreate {
	C.hedera_transaction__contract_create__set_file(tx.inner, cFileID(file))
	return tx
}

func (tx TransactionContractCreate) Gas(gas int64) TransactionContractCreate {
	C.hedera_transaction__contract_create__set_gas(tx.inner, C.uint64_t(gas))
	return tx
}

func (tx TransactionContractCreate) AdminKey(public PublicKey) TransactionContractCreate {
	C.hedera_transaction__contract_create__set_admin_key(tx.inner, public.inner)
	return tx
}

func (tx TransactionContractCreate) InitialBalance(balance int64) TransactionContractCreate {
	C.hedera_transaction__contract_create__set_initial_balance(tx.inner, C.uint64_t(balance))
	return tx
}

func (tx TransactionContractCreate) ProxyAccount(proxyId AccountID) TransactionContractCreate {
	C.hedera_transaction__contract_create__set_proxy_account(tx.inner, cAccountID(proxyId))
	return tx
}

func (tx TransactionContractCreate) ProxyFraction(fraction int32) TransactionContractCreate {
	C.hedera_transaction__contract_create__set_proxy_fraction(tx.inner, C.int32_t(fraction))
	return tx
}

func (tx TransactionContractCreate) AutoRenew(period Duration) TransactionContractCreate {
	C.hedera_transaction__contract_create__set_auto_renew_period(tx.inner, cDuration(period))
	return tx
}

func (tx TransactionContractCreate) ConstructorParams(params []byte) TransactionContractCreate {
	C.hedera_transaction__contract_create__set_constructor_parameters(tx.inner, (*C.uint8_t)(unsafe.Pointer(&params)),
		C.size_t(len(params)))
	return tx
}
