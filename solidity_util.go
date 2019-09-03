package hedera

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"errors"
)

const (
	ADDRESS_LEN = 20
	ADDRESS_LEN_HEX = ADDRESS_LEN * 2
)

func addressForEntity(shard int64, realm int64, entity int64) (string, error) {
	// TODO: Add check to ensure 'shard' is within 32 bit range
	buf := new(bytes.Buffer)

	err := binary.Write(buf, binary.BigEndian, int32(shard))
	if err != nil { return "", err }

	err = binary.Write(buf, binary.BigEndian, realm)
	if err != nil { return "", err }

	err = binary.Write(buf, binary.BigEndian, entity)
	if err != nil { return "", err }

	return hex.EncodeToString(buf.Bytes()), nil
}

func AddressForAccount(accountID AccountID) (string, error) {
	return addressForEntity(
		accountID.Shard,
		accountID.Realm,
		accountID.Account)
}

func AddressForContract(contractID ContractID) (string, error) {
	return addressForEntity(
		contractID.Shard,
		contractID.Realm,
		contractID.Contract)
}

func AddressForFile(fileID FileID) (string, error) {
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
