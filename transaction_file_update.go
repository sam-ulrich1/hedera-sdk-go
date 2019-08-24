package hedera

// #include "hedera.h"
import "C"

type TransactionFileUpdate struct {
	transaction
}

func newTransactionFileUpdate(client *Client, fileID FileID) TransactionFileUpdate {
	return TransactionFileUpdate{
		transaction{
			C.hedera_transaction__file_update__new(
				client.inner,
				cFileID(fileID))}}
}

func (tx TransactionFileUpdate) Key(public PublicKey) TransactionFileUpdate {
	C.hedera_transaction__file_update__set_key(tx.inner, public.inner)
	return tx
}

//func (tx TransactionFileUpdate) Content(content []byte) TransactionFileUpdate {
//	C.hedera_transaction__file_update__set_contents(tx.inner, (*C.uint8_t)(&content[0]),
//		C.size_t(len(content)))
//	return tx
//}
