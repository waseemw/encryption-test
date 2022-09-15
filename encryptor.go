package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"encoding/json"
	"math/rand"
)

func Encrypt[T any](keyStr string, obj T) string {
	iv := make([]byte, 16)
	rand.Read(iv)

	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	c, err := aes.NewCipher([]byte(keyStr))
	if err != nil {
		panic(err)
	}

	left := len(bytes) % c.BlockSize()
	if left > 0 {
		more := make([]byte, c.BlockSize()-left)
		for i := 0; i < len(more); i++ {
			more[i] = byte(' ')
		}
		bytes = append(bytes, more...)
	}

	blockMode := cipher.NewCBCEncrypter(c, iv)
	res := make([]byte, len(bytes))
	blockMode.CryptBlocks(res, bytes)

	return hex.EncodeToString(append(iv, res...))
}

func Decrypt[T any](keyStr string, str string) T {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		panic(err)
	}
	c, err := aes.NewCipher([]byte(keyStr))
	if err != nil {
		panic(err)
	}
	blockMode := cipher.NewCBCDecrypter(c, bytes[:16])
	decryptedBytes := make([]byte, len(bytes[16:]))
	blockMode.CryptBlocks(decryptedBytes, bytes[16:])

	var res T
	if err := json.Unmarshal(decryptedBytes, &res); err != nil {
		panic(err)
	}
	return res
}
