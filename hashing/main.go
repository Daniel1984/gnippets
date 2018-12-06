package main

import (
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

	md5Msg := encryotMsg("hello world", md5Hash)
	sha256Msg := encryotMsg("hello world", sha256Hash)

	fmt.Println("md5 -", md5Msg)
	fmt.Println("sha256 -", sha256Msg)
}

func encryotMsg(in string, h hashInterface) string {
	h.Write([]byte(in))
	hashedMsg := hex.EncodeToString(h.Sum(nil))
	return hashedMsg
}
