package hedera

// #include "hedera.h"
import "C"

type TransactionFileCreate struct {
	transaction
}

func newTransactionFileCreate(client *Client) TransactionFileCreate {
	return TransactionFileCreate{
		transaction{C.hedera_transaction__file_create__new(client.inner)}}
}

func (tx TransactionFileCreate) Key(public PublicKey) TransactionFileCreate {
	C.hedera_transaction__file_create__set_key(tx.inner, public.inner)
	return tx
}

func (tx TransactionFileCreate) Content(content []byte) TransactionFileCreate {
	C.hedera_transaction__file_create__set_contents(tx.inner, (*C.uint8_t)(&content[0]),
		C.size_t(len(content)))
	return tx
}

func (tx TransactionFileCreate) Expiration(exp Timestamp) TransactionFileCreate {
	C.hedera_transaction__file_create__set_expires_at(tx.inner, cTimestamp(exp))
	return tx
}
