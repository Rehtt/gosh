package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	AES128 = 16
	AES192 = 24
	AES256 = 32
)

// PKCS5Padding 填充明文
func pKCS5Padding(plaintext []byte, blockSize int) []byte {
	padding := blockSize - len(plaintext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(plaintext, padtext...)
}

// pKCS5UnPadding 去除填充数据
func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func GoshEncrypt(origData, key []byte) ([]byte, error) {
	o := make([]byte, 3, len(origData)+3)
	copy(o, key)
	o = append(o, origData...)
	return AesEncrypt(o, key)
}

// AesEncrypt 加密
func AesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func GoshDecrypt(crypted, key []byte) (origData []byte, err error) {
	origData, err = AesDecrypt(crypted, key)
	if err != nil {
		return
	}
	for i := range origData[:3] {
		if origData[i] != key[i] {
			return nil, fmt.Errorf("error")
		}
	}
	return origData[3:], nil
}

// AesDecrypt 解密
func AesDecrypt(crypted, key []byte) (origData []byte, err error) {
	defer func() {
		if e := recover(); e != nil {
			origData = nil
			err = fmt.Errorf("%v", e)
			return
		}
	}()
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	//AES分组长度为128位，所以blockSize=16，单位字节
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) //初始向量的长度必须等于块block的长度16字节
	origData = make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS5UnPadding(origData)
	return origData, nil
}

func AesGenerateKey(b int) (out []byte) {
	m := "MiBrbpCHxjldncS4RJTuW1IPEQtgqXUhF7YOo06mKZL25NAVDfkey98G3wsavz"
	for i := 0; i < b; i++ {
		n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(m))))
		out = append(out, m[n.Int64()])
	}
	return
}
