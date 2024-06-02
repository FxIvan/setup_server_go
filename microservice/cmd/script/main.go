package main

import (
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
)

func createPasetoKey() {
	publicKey, privateKey, err := ed25519.GenerateKey(nil)
	if err != nil {
		panic(err)
	}

	publicKeyHex := hex.EncodeToString(publicKey)
	privateKeyHex := hex.EncodeToString(privateKey)

	fmt.Println("Public Key:", publicKeyHex)
	fmt.Println("Private Key:", privateKeyHex)
}

func main() {
	createPasetoKey()
}
