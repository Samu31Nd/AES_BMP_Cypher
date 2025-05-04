package aesutils

import (
	"crypto/aes"
	"errors"
	"labAES28_04/utils"
)

func CifrarAES_CTR(c0 []byte, key []byte, pixels []byte) ([]byte, error) {
	return aesCTRInternal(c0, key, pixels)
}

func DecifrarAES_CTR(c0 []byte, key []byte, pixels []byte) ([]byte, error) {
	return aesCTRInternal(c0, key, pixels)
}

func aesCTRInternal(counter []byte, key []byte, data []byte) ([]byte, error) {
	if len(counter) != aes.BlockSize {
		return nil, errors.New("el contador debe tener 16 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	result := make([]byte, len(data))
	streamBlock := make([]byte, aes.BlockSize)
	ctr := make([]byte, aes.BlockSize)
	copy(ctr, counter)

	for i := 0; i < len(data); i += aes.BlockSize {
		block.Encrypt(streamBlock, ctr)

		chunkSize := aes.BlockSize
		if i+chunkSize > len(data) {
			chunkSize = len(data) - i
		}

		for j := 0; j < chunkSize; j++ {
			result[i+j] = data[i+j] ^ streamBlock[j]
		}

		utils.IncrementCounter(ctr)
	}

	return result, nil
}
