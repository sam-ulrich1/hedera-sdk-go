package hedera

// #include "hedera.h"
import "C"

type TransactionContractUpdate struct {
	transaction
}

func newTransactionContractUpdate(client *Client) TransactionContractUpdate {
	return TransactionContractUpdate{transaction{
		C.hedera_transaction__contract_update__new(client.inner)}}
}

func (tx TransactionContractUpdate) AdminKey(public PublicKey) TransactionContractUpdate {
	C.hedera_transaction__contract_update__set_admin_key(tx.inner, public.inner)
	return tx
}

func (tx TransactionContractUpdate) File(file FileID) TransactionContractUpdate {
	C.hedera_transaction__contract_update__set_file(tx.inner, cFileID(file))
	return tx
}

func (tx TransactionContractUpdate) ProxyAccount(proxyId AccountID) TransactionContractUpdate {
	C.hedera_transaction__contract_update__set_proxy_account(tx.inner, cAccountID(proxyId))
	return tx
}
