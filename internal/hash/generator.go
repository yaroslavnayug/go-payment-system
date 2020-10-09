package hash

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

const (
	hashCustomerKey = "customer"
)

func GenerateUniqueCustomerID(firstName string, passportNumber string, timestamp int64) (string, error) {
	baseString := fmt.Sprintf("%s%s%s%d", firstName, passportNumber, hashCustomerKey, timestamp)
	return getHashForString(baseString)
}

func getHashForString(baseString string) (string, error) {
	hasher := md5.New()
	_, err := hasher.Write([]byte(baseString))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
