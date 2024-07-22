package run

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
)

func Decrypt(data, key, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("blocksize: ", block.BlockSize())
	fmt.Println("ivsize: ", len(iv))
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(data, data)
	return data
}
