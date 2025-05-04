package aesutils

import (
	"crypto/aes"
	"errors"
)

func CifrarAES_OFB(c0 []byte, key []byte, pixels []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(c0) != aes.BlockSize {
		return nil, errors.New("el vector de inicialización (IV) debe tener 16 bytes")
	}

	pixels = PadPKCS7(pixels, aes.BlockSize)

	ciphertext := make([]byte, len(pixels))
	stream := make([]byte, aes.BlockSize)
	copy(stream, c0)

	for i := 0; i < len(pixels); i += aes.BlockSize {
		blockOut := make([]byte, aes.BlockSize)
		block.Encrypt(blockOut, stream)

		chunkSize := aes.BlockSize
		if i+chunkSize > len(pixels) {
			chunkSize = len(pixels) - i
		}

		for j := 0; j < chunkSize; j++ {
			ciphertext[i+j] = pixels[i+j] ^ blockOut[j]
		}

		copy(stream, blockOut)
	}

	return ciphertext, nil
}

func DecifrarAES_OFB(c0 []byte, key []byte, pixels []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(c0) != aes.BlockSize {
		return nil, errors.New("el vector de inicialización (IV) debe tener 16 bytes")
	}

	plaintext := make([]byte, len(pixels))
	stream := make([]byte, aes.BlockSize)
	copy(stream, c0)

	for i := 0; i < len(pixels); i += aes.BlockSize {
		blockOut := make([]byte, aes.BlockSize)
		block.Encrypt(blockOut, stream)

		chunkSize := aes.BlockSize
		if i+chunkSize > len(pixels) {
			chunkSize = len(pixels) - i
		}

		for j := 0; j < chunkSize; j++ {
			plaintext[i+j] = pixels[i+j] ^ blockOut[j]
		}

		copy(stream, blockOut)
	}

	return UnpadPKCS7(plaintext)
}
