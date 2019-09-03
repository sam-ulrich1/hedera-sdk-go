package hedera

// #include "hedera.h"
import "C"

type TransactionFileDelete struct {
	transaction
}

func newTransactionFileDelete(client *Client, fileID FileID) TransactionFileDelete {
	return TransactionFileDelete{
		transaction{
			C.hedera_transaction__file_delete__new(
				client.inner,
				cFileID(fileID), )}}
}
