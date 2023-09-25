package hash

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
)

type HashService struct {
}

func NewHashService() *HashService {
	return &HashService{}
} //end method

func (*HashService) Sha512(text string) (string, error) {
	if text == "" {
		return "", errors.New("text can not be empty")
	}
	hasher := sha512.New()
	_, err := hasher.Write([]byte(text))
	if err != nil {
		return "", err
	} //end if
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString, nil
} //end method

func (*HashService) Sha256(text string) (string, error) {
	hasher := sha256.New()
	if text == "" {
		return "", errors.New("text can not be empty")
	}
	_, err := hasher.Write([]byte(text))
	if err != nil {
		return "", err
	} //end if
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString, nil
} //end method

func (*HashService) Md5(text string) (string, error) {
	if text == "" {
		return "", errors.New("text can not be empty")
	}
	hasher := md5.New()

	_, err := hasher.Write([]byte(text))
	if err != nil {
		return "", err
	} //end if
	hashString := hex.EncodeToString(hasher.Sum(nil))
	return hashString, nil
} //end method
