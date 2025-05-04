package aesutils

import (
	"bytes"
	"crypto/aes"
	"errors"
)

func PadPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

func CifrarAES_ECB(key []byte, pixels []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	pixels = PadPKCS7(pixels, aes.BlockSize)

	ciphertext := make([]byte, len(pixels))

	for start := 0; start < len(pixels); start += aes.BlockSize {
		end := start + aes.BlockSize
		block.Encrypt(ciphertext[start:end], pixels[start:end])
	}

	return ciphertext, nil
}

func UnpadPKCS7(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("datos vacíos")
	}
	padding := int(data[len(data)-1])

	if padding <= 0 || padding > aes.BlockSize {
		return nil, errors.New("padding inválido")
	}

	// Validar que todos los últimos bytes sean iguales al padding
	for i := 1; i <= padding; i++ {
		if data[len(data)-i] != byte(padding) {
			return nil, errors.New("contenido del padding inválido")
		}
	}
	return data[:len(data)-padding], nil
}

func DecifrarAES_ECB(key []byte, pixels []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(pixels)%aes.BlockSize != 0 {
		return nil, errors.New("los datos cifrados no tienen un tamaño válido")
	}

	plaintext := make([]byte, len(pixels))

	for start := 0; start < len(pixels); start += aes.BlockSize {
		end := start + aes.BlockSize
		block.Decrypt(plaintext[start:end], pixels[start:end])
	}

	plaintext, err = UnpadPKCS7(plaintext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
