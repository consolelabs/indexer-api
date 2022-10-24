package utils

import (
	"encoding/hex"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"golang.org/x/crypto/sha3"
)

func IsEvmAddress(s string) bool {
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")
	return re.MatchString(s)
}

func EqualAddress[T string | common.Hash](x, y T) bool {
	return x == y
}

func calculateKeccak256(addr []byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	hash.Write(addr)
	return hash.Sum(nil)
}

func checksumByte(addr byte, hash byte) string {
	result := strconv.FormatUint(uint64(addr), 16)
	if hash >= 8 {
		return strings.ToUpper(result)
	} else {
		return result
	}
}

func StringToHashAddress(addrStr string) (string, error) {
	if len(addrStr) <= 2 {
		return "", fmt.Errorf("invalid address")
	}
	addrStr = addrStr[2:]
	addr, err := hex.DecodeString(addrStr)
	if err != nil {
		return "", err
	}
	hash := calculateKeccak256([]byte(strings.ToLower(addrStr)))

	result := "0x"

	for i, b := range addr {
		result += checksumByte(b>>4, hash[i]>>4)
		result += checksumByte(b&0xF, hash[i]&0xF)
	}

	return result, nil
}
