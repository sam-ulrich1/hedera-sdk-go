package hedera

type TransactionReceipt struct {
	Status     Status
	AccountID  *AccountID
	ContractID *ContractID
	FileID     *FileID
}

func (status Status) String() string {
	return statusText(status)
}
