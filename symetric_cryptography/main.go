package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
)

func main() {
	key := "ggtt1166ll995579ggtt1166ll995579"
	message := "hello world"
	encryptedMessage, err := encryptSymetricCrypto(key, message)
	if err != nil {
		panic(err)
	}
	fmt.Println(encryptedMessage)

	decryptedMessage, err := decryptSymetricCrypto(key, encryptedMessage)
	if err != nil {
		panic(err)
	}

	fmt.Println(decryptedMessage)
}

func encryptSymetricCrypto(key, msg string) (string, error) {
	keyLen := len(key)

	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		return "", errors.New("Key must be of a lengh 16, 24 or 32")
	}

	bc, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	text := []byte(msg)
	cipherText := make([]byte, aes.BlockSize+len(text))
	iv := cipherText[:aes.BlockSize]
	_, err = io.ReadFull(rand.Reader, iv)
	if err != nil {
		return "", err
	}

	cfb := cipher.NewCFBEncrypter(bc, iv)
	cfb.XORKeyStream(cipherText[aes.BlockSize:], text)

	return base64.StdEncoding.EncodeToString(cipherText), nil
}

func decryptSymetricCrypto(key, msg string) (string, error) {
	keyLen := len(key)

	if keyLen != 16 && keyLen != 24 && keyLen != 32 {
		return "", errors.New("Key must be of a lengh 16, 24 or 32")
	}

	encrypted, _ := base64.StdEncoding.DecodeString(msg)

	bc, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	if len(encrypted) < aes.BlockSize {
		return "", errors.New("cipher text too short")
	}

	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(bc, iv)
	cfb.XORKeyStream(encrypted, encrypted)
	return string(encrypted), nil
}
