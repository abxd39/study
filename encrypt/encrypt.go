package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
)

type Aesu struct {
}

//加密
func (a *Aesu) AesEncrypt(origData []byte, key string, iv string) ([]byte, error) {
	return a.aesEncryptPkcs5(origData, []byte(key), []byte(iv))
}

func (a *Aesu) aesEncryptPkcs5(origData, key, iv []byte) ([]byte, error) {
	return a.encrypt(origData, key, iv, a.pKCS5Padding)
}

func (a *Aesu) encrypt(origData, key, iv []byte, paddingFunc func([]byte, int) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	origData = paddingFunc(origData, block.BlockSize())
	crypted := make([]byte, block.BlockSize()+len(origData))
	//iv = origData[:block.BlockSize()]
	cipher.NewCBCEncrypter(block, iv).CryptBlocks(crypted[block.BlockSize():], origData)
	return crypted[block.BlockSize():], nil
}

func (a *Aesu) pKCS5Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

//解密
func (a *Aesu) AesDecrypt(crypted []byte, key string, iv string) ([]byte, error) {
	return a.aesDecryptPkcs5(crypted, []byte(key), []byte(iv))
}

func (a *Aesu) aesDecryptPkcs5(crypted, key, iv []byte) ([]byte, error) {
	return a.decrypt(crypted, key, iv, a.pKCS5UnPadding)
}

func (a *Aesu) decrypt(ciphertext, key, iv []byte, unPaddingFunc func([]byte) []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	cipher.NewCBCDecrypter(block, iv).CryptBlocks(ciphertext, ciphertext)
	if int(ciphertext[len(ciphertext)-1]) > len(ciphertext) {
		return nil, errors.New("aes decrypt failed")
	}
	plaintext := unPaddingFunc(ciphertext)
	log.Println(base64.StdEncoding.EncodeToString(plaintext))
	log.Println(string(plaintext))
	return bytes.Trim(plaintext, "\x00"), nil
}

func (a *Aesu) pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	if length < unPadding {
		return []byte("unPadding is error")
	}
	return origData[:(length - unPadding)]

}
