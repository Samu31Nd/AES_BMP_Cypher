package aesutils

import (
	"crypto/aes"
	"errors"
	"labAES28_04/utils"
)

func CifrarAES_CBC(c0 []byte, key []byte, pixels []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(c0) != aes.BlockSize {
		return nil, errors.New("el vector de inicialización (IV) debe tener 16 bytes")
	}

	pixels = PadPKCS7(pixels, aes.BlockSize)

	ciphertext := make([]byte, len(pixels))
	prev := make([]byte, aes.BlockSize)
	copy(prev, c0)

	for start := 0; start < len(pixels); start += aes.BlockSize {
		end := start + aes.BlockSize
		blockIn := utils.XorBlocks(pixels[start:end], prev)
		blockOut := make([]byte, aes.BlockSize)
		block.Encrypt(blockOut, blockIn)
		copy(ciphertext[start:end], blockOut)
		copy(prev, blockOut)
	}

	return ciphertext, nil
}

func DecifrarAES_CBC(c0 []byte, key []byte, pixels []byte) ([]byte, error) {
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
	prev := make([]byte, aes.BlockSize)
	copy(prev, c0)

	for start := 0; start < len(pixels); start += aes.BlockSize {
		end := start + aes.BlockSize
		blockOut := make([]byte, aes.BlockSize)
		block.Decrypt(blockOut, pixels[start:end])
		copy(plaintext[start:end], utils.XorBlocks(blockOut, prev))
		copy(prev, pixels[start:end])
	}

	return UnpadPKCS7(plaintext)
}
