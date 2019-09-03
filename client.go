package hedera

// #include <stdlib.h>
// #include "hedera.h"
//
// int operatorSecretCallback(void*, HederaSecretKey*);
//
// static void _client_set_operator(HederaClient* client, HederaAccountId id, void* user_data) {
//   hedera_client_set_operator(client, id, &operatorSecretCallback, user_data);
// }
import "C"
import (
	"unsafe"

	"github.com/markbates/oncer"
	"github.com/mattn/go-pointer"
)

type Client struct {
	inner *C.HederaClient
}

func Dial(address string) (Client, error) {
	cAddress := C.CString(address)
	defer C.free(unsafe.Pointer(cAddress))

	var inner *C.HederaClient
	res := C.hedera_client_new(cAddress, &inner)
	if res != 0 {
		return Client{}, hederaLastError()
	}

	return Client{inner}, nil
}

func (client Client) SetNode(node AccountID) {
	C.hedera_client_set_node(client.inner, cAccountID(node))
}

//export operatorSecretCallback
func operatorSecretCallback(v unsafe.Pointer, out *C.HederaSecretKey) C.int {
	cb := pointer.Restore(v).(*func() SecretKey)
	key := (*cb)()
	*out = key.inner

	return 0
}

func (client Client) SetOperator(operator AccountID, secretCallback func() SecretKey) {
	userData := pointer.Save(&secretCallback)
	C._client_set_operator(client.inner, cAccountID(operator), userData)
}

func (client Client) Close() {
	C.hedera_client_close(client.inner)
}

// Deprecated: Use Client.TransferCrypto() instead
func (client Client) CryptoTransfer() TransactionCryptoTransfer {
	oncer.Deprecate(0,
		"github.com/launchbadge/hedera-sdk-go#Client.CryptoTransfer()",
		"Use Client.TransferCrypto() instead.")

	return client.TransferCrypto()
}

// Deprecated: Use Client.Account(id).Balance() instead
func (client Client) GetAccountBalance(id AccountID) QueryCryptoGetAccountBalance {
	oncer.Deprecate(0,
		"github.com/launchbadge/hedera-sdk-go#Client.GetAccountBalance(id)",
		"Use Client.Account(id).Balance() instead.")

	return newQueryCryptoGetAccountBalance(client, id)
}

// Deprecated: Use Client.Transaction(id).Receipt() instead
func (client Client) GetTransactionReceipt(id *TransactionID) QueryTransactionGetReceipt {
	oncer.Deprecate(0,
		"github.com/launchbadge/hedera-sdk-go#Client.GetTransactionReceipt(id)",
		"Use Client.Transaction(id).Receipt() instead.")

	return newQueryTransactionGetReceipt(client, *id)
}

func (client Client) TransferCrypto() TransactionCryptoTransfer {
	return newTransactionCryptoTransfer(&client)
}

func (client Client) CreateAccount() TransactionCryptoCreate {
	return newTransactionCryptoCreate(&client)
}

func (client Client) CreateContract() TransactionContractCreate {
	return newTransactionContractCreate(&client)
}

func (client Client) UpdateContract() TransactionContractUpdate {
	return newTransactionContractUpdate(&client)
}

func (client Client) CallContract(contract ContractID) TransactionContractCall {
	return newTransactionContractCall(&client, contract)
}

func (client Client) CallContractQuery(contract ContractID, gas int64, params []byte,
	maxResultSize int64) QueryContractCall {
	return newQueryContractCall(&client, contract, gas, params, maxResultSize)
}

func (client Client) CreateFile() TransactionFileCreate {
	return newTransactionFileCreate(&client)
}

func (client Client) UpdateFile(fileID FileID) TransactionFileUpdate {
	return newTransactionFileUpdate(&client, fileID)
}

func (client Client) AppendFile(fileID FileID, content []byte) TransactionFileAppend {
	return newTransactionFileAppend(&client, fileID, content)
}

func (client Client) DeleteFile(fileID FileID) TransactionFileDelete {
	return newTransactionFileDelete(&client, fileID)
}

func (client Client) Account(id AccountID) PartialAccountMessage {
	return PartialAccountMessage{client, id}
}

func (client Client) Transaction(id TransactionID) PartialTransactionMessage {
	return PartialTransactionMessage{client, id}
}

func (client Client) FileContents(fileID FileID) QueryFileGetContents {
	return newQueryFileGetContents(client, fileID)
}

type PartialAccountMessage struct {
	client  Client
	account AccountID
}

func (m PartialAccountMessage) Balance() QueryCryptoGetAccountBalance {
	return newQueryCryptoGetAccountBalance(m.client, m.account)
}

type PartialTransactionMessage struct {
	client      Client
	transaction TransactionID
}

func (m PartialTransactionMessage) Receipt() QueryTransactionGetReceipt {
	return newQueryTransactionGetReceipt(m.client, m.transaction)
}
