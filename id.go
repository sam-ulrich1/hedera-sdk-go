package hedera

// #include <stdlib.h>
// #include "hedera.h"
import "C"
import (
	"encoding/binary"
	"errors"
	"unsafe"
)
import "github.com/markbates/oncer"

type AccountID struct {
	Realm   int64 `json:"realm"`
	Shard   int64 `json:"shard"`
	Account int64 `json:"account"`
}

// Deprecated: Use AccountID{realm, shard, account} instead.
func NewAccountID(realm, shard, account int64) AccountID {
	oncer.Deprecate(0,
		"github.com/launchbadge/hedera-sdk-go#NewAccountID",
		"Use AccountID instead.")

	return AccountID{Realm: realm, Shard: shard, Account: account}
}

type ContractID struct {
	Realm    int64 `json:"realm"`
	Shard    int64 `json:"shard"`
	Contract int64 `json:"contract"`
}

type FileID struct {
	Realm int64 `json:"realm"`
	Shard int64 `json:"shard"`
	File  int64 `json:"file"`
}

func cAccountID(id AccountID) C.HederaAccountId {
	return C.HederaAccountId{
		realm:   C.int64_t(id.Realm),
		shard:   C.int64_t(id.Shard),
		account: C.int64_t(id.Account),
	}
}

func goAccountID(id C.HederaAccountId) AccountID {
	return AccountID{
		Realm:   int64(id.realm),
		Shard:   int64(id.shard),
		Account: int64(id.account),
	}
}

func AccountIDFromString(s string) (AccountID, error) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	var id C.HederaAccountId
	res := C.hedera_account_id_from_str(cStr, &id)
	if res != 0 {
		return AccountID{}, hederaLastError()
	}

	return goAccountID(id), nil
}

func AccountIDFromAddress(b []byte) (AccountID, error) {
	if len(b) != 20 {
		return AccountID{}, errors.New("ILLEGAL ARGUMENT ERROR: A solidity address " +
			"should be exactly 20 bytes long")
	}
	shard := int64(binary.BigEndian.Uint32(b[0:4]))
	realm := int64(binary.BigEndian.Uint64(b[4:12]))
	acct := int64(binary.BigEndian.Uint64(b[12:20]))
	return AccountID{
		Shard: shard,
		Realm: realm,
		Account: acct,
	}, nil
}

func (id AccountID) String() string {
	cID := cAccountID(id)
	return hederaString(C.hedera_account_id_to_str(&cID))
}

func cContractID(id ContractID) C.HederaContractId {
	return C.HederaContractId{
		realm:    C.int64_t(id.Realm),
		shard:    C.int64_t(id.Shard),
		contract: C.int64_t(id.Contract),
	}
}

func goContractID(id C.HederaContractId) ContractID {
	return ContractID{
		Realm:    int64(id.realm),
		Shard:    int64(id.shard),
		Contract: int64(id.contract),
	}
}

func ContractIDFromString(s string) (ContractID, error) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	var id C.HederaContractId
	res := C.hedera_contract_id_from_str(cStr, &id)
	if res != 0 {
		return ContractID{}, hederaLastError()
	}

	return goContractID(id), nil
}

func ContractIDFromAddress(b []byte) (ContractID, error) {
	if len(b) != 20 {
		return ContractID{}, errors.New("ILLEGAL ARGUMENT ERROR: A solidity address " +
			"should be exactly 20 bytes long")
	}
	shard := int64(binary.BigEndian.Uint32(b[0:4]))
	realm := int64(binary.BigEndian.Uint64(b[4:12]))
	acct := int64(binary.BigEndian.Uint64(b[12:20]))
	return ContractID{
		Shard: shard,
		Realm: realm,
		Contract: acct,
	}, nil
}

func (id ContractID) String() string {
	cID := cContractID(id)
	return hederaString(C.hedera_contract_id_to_str(&cID))
}

func cFileID(id FileID) C.HederaFileId {
	return C.HederaFileId{
		realm: C.int64_t(id.Realm),
		shard: C.int64_t(id.Shard),
		file:  C.int64_t(id.File),
	}
}

func goFileID(id C.HederaFileId) FileID {
	return FileID{
		Realm: int64(id.realm),
		Shard: int64(id.shard),
		File:  int64(id.file),
	}
}

func FileIDFromString(s string) (FileID, error) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr))

	var id C.HederaFileId
	res := C.hedera_file_id_from_str(cStr, &id)
	if res != 0 {
		return FileID{}, hederaLastError()
	}

	return goFileID(id), nil
}

func FileIDFromAddress(b []byte) (FileID, error) {
	if len(b) != 20 {
		return FileID{}, errors.New("ILLEGAL ARGUMENT ERROR: A solidity address " +
			"should be exactly 20 bytes long")
	}
	shard := int64(binary.BigEndian.Uint32(b[0:4]))
	realm := int64(binary.BigEndian.Uint64(b[4:12]))
	acct := int64(binary.BigEndian.Uint64(b[12:20]))
	return FileID{
		Shard: shard,
		Realm: realm,
		File: acct,
	}, nil
}

func (id FileID) String() string {
	cID := cFileID(id)
	return hederaString(C.hedera_file_id_to_str(&cID))
}
