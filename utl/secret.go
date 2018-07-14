package utl

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"encoding/hex"
)

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func AesEncrypt(origData []byte, secretKey string) (string, error) {
	key := []byte(secretKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS7Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return fmt.Sprintf("%x", crypted), nil
}

func AesMustEncrypt(origData []byte, secretKey string) string {
	ret, _ := AesEncrypt(origData, secretKey)
	return ret
}

func AesDecrypt(cryptedStr string, secretKey string) ([]byte, error) {
	crypted, _ := hex.DecodeString(cryptedStr)
	key := []byte(secretKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}

func AesMustDecrypt(cryptedStr string, secretKey string) []byte {
	ret, _ := AesDecrypt(cryptedStr, secretKey)
	return ret
}