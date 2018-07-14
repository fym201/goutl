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

type AesCryptPKCS5 struct {
	key []byte
	iv  []byte
}

func NewAesCryptPKCS5(key, iv []byte) *AesCryptPKCS5 {
	return &AesCryptPKCS5{key:key, iv:iv}
}

func (a *AesCryptPKCS5) Encrypt(data []byte) ([]byte, error) {
	aesBlockEncrypt, err := aes.NewCipher(a.key)
	content := PKCS5Padding(data, aesBlockEncrypt.BlockSize())
	encrypted := make([]byte, len(content))
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesEncrypt := cipher.NewCBCEncrypter(aesBlockEncrypt, a.iv)
	aesEncrypt.CryptBlocks(encrypted, content)
	return encrypted, nil
}

func (a *AesCryptPKCS5) Decrypt(src []byte) (data []byte, err error) {
	decrypted := make([]byte, len(src))
	var aesBlockDecrypt cipher.Block
	aesBlockDecrypt, err = aes.NewCipher(a.key)
	if err != nil {
		println(err.Error())
		return nil, err
	}
	aesDecrypt := cipher.NewCBCDecrypter(aesBlockDecrypt, a.iv)
	aesDecrypt.CryptBlocks(decrypted, src)
	return PKCS5Trimming(decrypted), nil
}


func PKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}