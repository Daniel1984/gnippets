package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
)

type hashInterface interface {
	io.Writer
	Sum(b []byte) []byte
}

func main() {
	md5Hash := md5.New()
	sha256Hash := sha256.New()

	md5Msg := encryptMsg("hello world", md5Hash)
	sha256Msg := encryptMsg("hello world", sha256Hash)

	hmacKey := []byte("12a34")
	hmacHash := hmac.New(sha256.New, hmacKey)
	hmacMsg := encryptMsg("hello world", hmacHash)

	fmt.Println("md5 -", md5Msg)
	fmt.Println("sha256 -", sha256Msg)
	fmt.Println("hmac -", hmacMsg)
}

func encryptMsg(in string, h hashInterface) string {
	h.Write([]byte(in))
	hashedMsg := hex.EncodeToString(h.Sum(nil))
	return hashedMsg
}
