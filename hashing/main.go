package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	md5HashMsg := hashMd5("hello world")
	fmt.Println(md5HashMsg)
}

func hashMd5(in string) string {
	h := md5.New()
	h.Write([]byte(in))
	hashedMsg := hex.EncodeToString(h.Sum(nil))
	return hashedMsg
}
