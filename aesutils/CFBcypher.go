package aesutils

import (
	"crypto/aes"
	"errors"
)

func CifrarAES_CFB(c0 []byte, key []byte, pixels []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(c0) != aes.BlockSize {
		return nil, errors.New("el vector de inicialización (IV) debe tener 16 bytes")
	}

	pixels = PadPKCS7(pixels, aes.BlockSize)

	ciphertext := make([]byte, len(pixels))
	feedback := make([]byte, aes.BlockSize)
	copy(feedback, c0)

	for i := 0; i < len(pixels); i += aes.BlockSize {
		blockOut := make([]byte, aes.BlockSize)
		block.Encrypt(blockOut, feedback)

		chunkSize := aes.BlockSize
		if i+chunkSize > len(pixels) {
			chunkSize = len(pixels) - i
		}

		for j := 0; j < chunkSize; j++ {
			ciphertext[i+j] = blockOut[j] ^ pixels[i+j]
		}

		copy(feedback, ciphertext[i:i+chunkSize])
	}

	return ciphertext, nil
}

func DecifrarAES_CFB(c0 []byte, key []byte, pixels []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(c0) != aes.BlockSize {
		return nil, errors.New("el vector de inicialización (IV) debe tener 16 bytes")
	}
	if len(pixels)%aes.BlockSize != 0 {
		return nil, errors.New("los datos cifrados no son múltiplos del tamaño de bloque")
	}

	plaintext := make([]byte, len(pixels))
	feedback := make([]byte, aes.BlockSize)
	copy(feedback, c0)

	for i := 0; i < len(pixels); i += aes.BlockSize {
		blockOut := make([]byte, aes.BlockSize)
		block.Encrypt(blockOut, feedback)

		chunkSize := aes.BlockSize
		if i+chunkSize > len(pixels) {
			chunkSize = len(pixels) - i
		}

		for j := 0; j < chunkSize; j++ {
			plaintext[i+j] = blockOut[j] ^ pixels[i+j]
		}

		copy(feedback, pixels[i:i+chunkSize])
	}

	return UnpadPKCS7(plaintext)
}
