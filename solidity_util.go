package hedera

import (
	"encoding/hex"
	"errors"
)

const (
	ADDRESS_LEN = 20
	ADDRESS_LEN_HEX = ADDRESS_LEN * 2
)

func addressForEntity(shard int64, realm int64, entity int64) string {
	// TODO: Add check to ensure 'shard' is within 32 bit range
	var buf []byte
	buf = append(buf, byte(shard))
	buf = append(buf, byte(realm))
	buf = append(buf, byte(entity))
	return hex.EncodeToString(buf)
}

func addressForAccount(accountID AccountID) string {
	return addressForEntity(
		accountID.Shard,
		accountID.Realm,
		accountID.Account)
}

func addressForContract(contractID ContractID) string {
	return addressForEntity(
		contractID.Shard,
		contractID.Realm,
		contractID.Contract)
}

func addressForFile(fileID FileID) string {
	return addressForEntity(
		fileID.Shard,
		fileID.Realm,
		fileID.File)
}

func CheckAddressLen(address []byte) error {
	if len(address) != ADDRESS_LEN  {
		return errors.New("ILLEGAL ARGUMENT ERROR: Solidity addresses must be 20 bytes or 40 hex chars")
	}
	return nil
}

func DecodeAddress(address string) ([]byte, error) {
	if len(address) != ADDRESS_LEN_HEX {
		return nil, errors.New("ILLEGAL ARGUMENT ERROR: Solidity addresses must be 20 bytes or 40 hex chars")
	}

	b, err := hex.DecodeString(address)
	if err != nil {
		return nil, errors.New("ILLEGAL ARGUMENT ERROR: failed to decode Solidity address as hex; " + err.Error())
	}
	return b, nil
}
