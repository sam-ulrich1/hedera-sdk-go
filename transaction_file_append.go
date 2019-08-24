package hedera

// #include "hedera.h"
import "C"

type TransactionFileAppend struct {
	transaction
}

func newTransactionFileAppend(client *Client, fileID FileID, content []byte) TransactionFileAppend {
	return TransactionFileAppend{
		transaction{
			C.hedera_transaction__file_append__new(
				client.inner,
				cFileID(fileID),
				(*C.uint8_t)(&content[0]),
				C.size_t(len(content)))},
	}
}
