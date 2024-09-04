package util

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
)

func DesEncrypt(origData []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func DesDecrypt(crypted []byte, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = unpadding(origData)
	return origData, nil
}

func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, padText...)
}

func unpadding(src []byte) []byte {
	n := len(src)
	if n == 0 {
		return nil
	}
	padNum := int(src[n-1])
	return src[:n-padNum]
}
