package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

type myaes struct {
	key []byte
}

var maes *myaes

func (this *myaes) aesEncrypt(origData []byte) (string, error) {
	block, err := aes.NewCipher(this.key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = this.pkcs5Padding(origData, blockSize)
	// origData = zeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, this.key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func (this *myaes) aesDecrypt(cd string) (string, error) {
	crypted, err := base64.StdEncoding.DecodeString(cd)
	if err != nil {
		//AesDescrypt,base65 decodeString failed.
		fmt.Println("error")
	}
	block, err := aes.NewCipher(this.key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, this.key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = this.pkcs5UnPadding(origData)
	// origData = zeroUnPadding(origData)

	return string(origData), nil
}

func (this *myaes) zeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func (this *myaes) zeroUnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func (this *myaes) pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func (this *myaes) pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func (this *myaes) setKey(key []byte) bool {
	switch len(key) {
	default:
		//"haes set key failed. key is invalid. use default value instead of."
		fmt.Errorf("error")
		keystr1 := "6mtkv"
		keystr2 := "M8jY"
		keystr3 := "qJFx"
		keystr4 := "w0T"
		keystr := keystr1 + keystr2 + keystr3 + keystr4
		this.key = []byte(keystr)
		return false
	case 16, 24, 32:
		//"Change key value"
		fmt.Println("error")
		this.key = key
		return true
	}
}

func newAes() *myaes {
	my := new(myaes)
	keystr1 := "6mtkv"
	keystr2 := "M8jY"
	keystr3 := "qJFx"
	keystr4 := "w0T"
	keystr := keystr1 + keystr2 + keystr3 + keystr4
	my.key = []byte(keystr)
	return my
}

func init() {
	maes = newAes()
}

func Encrypt(dt string) (string, error) {
	return maes.aesEncrypt([]byte(dt))
}
func Decrypt(dt string) (string, error) {
	return maes.aesDecrypt(dt)
}
func SetKey(key []byte) bool {
	return maes.setKey(key)
}
