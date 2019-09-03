package hedera

type Status uint8

const (
	StatusOk                             Status = 0
	StatusInvalidTransaction             Status = 1
	StatusPayerAccountNotFound           Status = 2
	StatusInvalidNodeAccount             Status = 3
	StatusTransactionExpired             Status = 4
	StatusInvalidTransactionStart        Status = 5
	StatusInvalidTransactionDuration     Status = 6
	StatusInvalidSignature               Status = 7
	StatusMemoTooLong                    Status = 8
	StatusInsufficientTxFee              Status = 9
	StatusInsufficientPayerBalance       Status = 10
	StatusDuplicateTransaction           Status = 11
	StatusBusy                           Status = 12
	StatusNotSupported                   Status = 13
	StatusInvalidFileId                  Status = 14
	StatusInvalidAccountId               Status = 15
	StatusInvalidContractId              Status = 16
	StatusInvalidTransactionId           Status = 17
	StatusReceiptNotFound                Status = 18
	StatusRecordNotFound                 Status = 19
	StatusInvalidSolidityId              Status = 20
	StatusUnknown                        Status = 21
	StatusSuccess                        Status = 22
	StatusFailInvalid                    Status = 23
	StatusFailFee                        Status = 24
	StatusFailBalance                    Status = 25
	StatusKeyRequired                    Status = 26
	StatusBadEncoding                    Status = 27
	StatusInsufficientAccountBalance     Status = 28
	StatusInvalidSolidityAddress         Status = 29
	StatusInsufficientGas                Status = 30
	StatusContractSizeLimitExceeded      Status = 31
	StatusLocalCallModificationException Status = 32
	StatusContractRevertExecuted         Status = 33
	StatusContractExecutionException     Status = 34
	StatusInvalidReceivingNodeAccount    Status = 35
	StatusMissingQueryHeader             Status = 36
	StatusAccountUpdateFailed            Status = 37
	StatusInvalidKeyEncoding             Status = 38
	StatusNullSolidityAddress            Status = 39
	StatusContractUpdateFailed           Status = 40
	StatusInvalidQueryHeader             Status = 41
	StatusInvalidFeeSubmitted            Status = 42
	StatusInvalidPayerSignature          Status = 43
	StatusKeyNotProvided                 Status = 44
	StatusInvalidExpirationTime          Status = 45
	StatusNoWaclKey                      Status = 46
	StatusFileContentEmpty               Status = 47
	StatusInvalidAccountAmounts          Status = 48
	StatusEmptyTransactionBody           Status = 49
	StatusInvalidTransactionBody         Status = 50
	InvalidSignatureTypeMismatch         Status = 51
	InvalidSignatureCountMismatch        Status = 52
	EmptyClaimBody                       Status = 53
	EmptyClaimHash                       Status = 54
	EmptyClaimKeys                       Status = 55
	InvalidClaimHashSize                 Status = 56
	EmptyQueryBody                       Status = 57
	EmptyClaimQuery                      Status = 58
	ClaimNotFound                        Status = 59
	AccountIdDoesNotExist                Status = 60
	ClaimAlreadyExists                   Status = 61
	InvalidFileWACL                      Status = 62
	SerializationFailed                  Status = 63
	TransactionOversize                  Status = 64
	TransactionTooManyLayers             Status = 65
	ContractDeleted                      Status = 66
	PlatformNotActive                    Status = 67
	KeyPrefixMismatch                    Status = 68
	TransactionNotCreated                Status = 69
	InvalidRenewalPeriod                 Status = 70
	InvalidPayerAccount                  Status = 71
	AccountDeleted                       Status = 72
	FileDeleted                          Status = 73
	AccountRepeatedInAccountAmounts      Status = 74
	SettingNegativeAccountBalance        Status = 75
	ObtainerRequired                     Status = 76
	ObtainerSameContractId               Status = 77
	ObtainerDoesNotExist                 Status = 78
	ModifyingImmutableContract           Status = 79
	FileSystemException                  Status = 80
	AutorenewDurationNotInRange          Status = 81
	ErrorDecodingBytestring              Status = 82
	ContractFileEmpty                    Status = 83
	ContractBytecodeEmpty                Status = 84
	InvalidInitialBalance                Status = 85
	InvalidReceiveRecordThreshold        Status = 86
	InvalidSendRecordThreshold           Status = 87
	AccountIsNotGenesisAccount           Status = 88
	PayerAccountUnauthorized             Status = 89
	InvalidFreezeTransactionBody         Status = 90
	FreezeTransactionBodyNotFound        Status = 91
	TransferListSizeLimitExceeded        Status = 92
	ResultSizeLimitExceeded              Status = 93
	NotSpecialAccount                    Status = 94
	ContractNegativeGas                  Status = 95
	ContractNegativeValue                Status = 96
	InvalidFeeFile						 Status = 97
	InvalidExchangeRateFile				 Status = 98
	InsufficientLocalCallGas			 Status = 99
	EntityNotAllowedToDelete			 Status = 100
	AuthorizationFailed					 Status = 101
	FileUploadedProtoInvalid			 Status = 102
	FileUploadedProtoNotSavedToDisk		 Status = 103
	FeeScheduleFilePartUploaded			 Status = 104
	ExchangeRateChangeLimitExceeded		 Status = 105
)

func statusText(s Status) string {
	switch s {
	default:
		return "UNKNOWN STATUS CODE: " + s.String()
	case 0:
		return "OK"
	case 1:
		return "INVALID TRANSACTION"
	case 2:
		return "PAYER ACCOUNT NOT FOUND"
	case 3:
		return "INVALID NODE ACCOUNT"
	case 4:
		return "TRANSACTION EXPIRED"
	case 5:
		return "INVALID TRANSACTION START"
	case 6:
		return "INVALID TRANSACTION DURATION"
	case 7:
		return "INVALID SIGNATURE"
	case 8:
		return "MEMO TOO LONG"
	case 9:
		return "INSUFFICIENT TX FEE"
	case 10:
		return "INSUFFICIENT PAYER BALANCE"
	case 11:
		return "DUPLICATE TRANSACTION"
	case 12:
		return "BUSY"
	case 13:
		return "NOT SUPPORTED"
	case 14:
		return "INVALID FILE ID"
	case 15:
		return "INVALID ACCOUNT ID"
	case 16:
		return "INVALID CONTRACT ID"
	case 17:
		return "INVALID TRANSACTION ID"
	case 18:
		return "RECEIPT NOT FOUND"
	case 19:
		return "RECORD NOT FOUND"
	case 20:
		return "INVALID SOLIDITY ID"
	case 21:
		return "UNKNOWN"
	case 22:
		return "SUCCESS"
	case 23:
		return "FAIL INVALID"
	case 24:
		return "FAIL FEE"
	case 25:
		return "FAIL BALANCE"
	case 26:
		return "KEY REQUIRED"
	case 27:
		return "BAD ENCODING"
	case 28:
		return "INSUFFICIENT ACCOUNT BALANCE"
	case 29:
		return "INVALID SOLIDITY ADDRESS"
	case 30:
		return "INSUFFICIENT GAS"
	case 31:
		return "CONTRACT SIZE LIMIT EXCEEDED"
	case 32:
		return "LOCAL CALL MODIFICATION EXCEPTION"
	case 33:
		return "CONTRACT EXECUTION EXCEPTION"
	case 34:
		return "INVALID RECEIVING NODE ACCOUNT"
	case 35:
		return "INVALID RECEIVING NODE ACCOUNT"
	case 36:
		return "MISSING QUERY HEADER"
	case 37:
		return "ACCOUNT UPDATE FAILED"
	case 38:
		return "INVALID KEY ENCODING"
	case 39:
		return "NULL SOLIDITY ADDRESS"
	case 40:
		return "CONTRACT UPDATE FAILED"
	case 41:
		return "INVALID QUERY HEADER"
	case 42:
		return "INVALID FEE SUBMITTED"
	case 43:
		return "INVALID PAYER SIGNATURE"
	case 44:
		return "KEY NOT PROVIDED"
	case 45:
		return "INVALID EXPIRATION TIME"
	case 46:
		return "NO WACL KEY"
	case 47:
		return "FILE CONTENT EMPTY"
	case 48:
		return "INVALID ACCOUNT AMOUNTS"
	case 49:
		return "EMPTY TRANSACTION BODY"
	case 50:
		return "INVALID TRANSACTION BODY"
	case 51:
		return "INVALID SIGNATURE TYPE MISMATCH"
	case 52:
		return "INVALID SIGNATURE COUNT MISMATCH"
	case 53:
		return "EMPTY CLAIM BODY"
	case 54:
		return "EMPTY CLAIM HASH"
	case 55:
		return "EMPTY CLAIM KEYS"
	case 56:
		return "INVALID CLAIM HASH SIZE"
	case 57:
		return "EMPTY QUERY BODY"
	case 58:
		return "EMPTY CLAIM QUERY"
	case 59:
		return "CLAIM NOT FOUND"
	case 60:
		return "ACCOUNT ID DOES NOT EXIST"
	case 61:
		return "CLAIM ALREADY EXISTS"
	case 62:
		return "INVALID FILE WACL"
	case 63:
		return "SERIALIZATION FAILED"
	case 64:
		return "TRANSACTION OVERSIZE"
	case 65:
		return "TRANSACTION TOO MANY LAYERS"
	case 66:
		return "CONTRACT DELETED"
	case 67:
		return "PLATFORM NOT ACTIVE"
	case 68:
		return "KEY PREFIX MISMATCH"
	case 69:
		return "TRANSACTION NOT CREATED"
	case 70:
		return "INVALID RENEWAL PERIOD"
	case 71:
		return "INVALID PAYER ACCOUNT"
	case 72:
		return "ACCOUNT DELETED"
	case 73:
		return "FILE DELETED"
	case 74:
		return "ACCOUNT REPEATED IN ACCOUNT AMOUNTS"
	case 75:
		return "SETTING NEGATIVE ACCOUNT BALANCE"
	case 76:
		return "OBTAINER REQUIRED"
	case 78:
		return "OBTAINER SAME CONTRACT ID"
	case 79:
		return "MODIFYING IMMUTABLE CONTRACT"
	case 80:
		return "FILE SYSTEM EXCEPTION"
	case 81:
		return "AUTO RENEW DURATION NOT IN RANGE"
	case 82:
		return "ERROR DECODING BYTESTRING"
	case 83:
		return "CONTRACT FILE EMPTY"
	case 84:
		return "CONTRACT BYTECODE EMPTY"
	case 85:
		return "INVALID INITIAL BALANCE"
	case 86:
		return "INVALID RECEIVE RECORD THRESHOLD"
	case 87:
		return "INVALID SEND RECORD THRESHOLD"
	case 88:
		return "ACCOUNT IS NOT GENESIS ACCOUNT"
	case 89:
		return "PAYER ACCOUNT UNAUTHORIZED"
	case 90:
		return "INVALID FREEZE TRANSACTION"
	case 91:
		return "FREEZE TRANSACTION BODY NOT FOUND"
	case 92:
		return "TRANSFER LIST SIZE LIMIT EXCEEDED"
	case 93:
		return "RESULT SIZE LIMIT EXCEEDED"
	case 94:
		return "NOT SPECIAL ACCOUNT"
	case 95:
		return "CONTRACT NEGATIVE GAS"
	case 96:
		return "CONTRACT NEGATIVE VALUE"
	case 97:
		return "INVALID FEE FILE"
	case 98:
		return "INVALID EXCHANGE RATE FILE"
	case 99:
		return "INSUFFICIENT LOCAL CALL GAS"
	case 100:
		return "ENTITY NOT ALLOWED TO DELETE"
	case 101:
		return "AUTHORIZATION FAILED"
	case 102:
		return "FILE UPLOADED PROTO INVALID"
	case 103:
		return "FILE UPLOADED PROTO NOT SAVED TO DISK"
	case 104:
		return "FEE SCHEDULE FILE PART UPLOADED"
	case 105:
		return "EXCHANGE RATE CHANGE LIMIT EXCEEDED"
	}
}
