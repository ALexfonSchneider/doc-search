package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	h := sha256.New()
	h.Write([]byte("this is a password"))
	// Calculate and print the hash
	fmt.Println(
		fmt.Sprintf("%x", h.Sum(nil)))

	fmt.Println(
		base64.URLEncoding.EncodeToString(h.Sum(nil)))
}
