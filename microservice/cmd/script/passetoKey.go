package passetoKey

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/base64"
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

func createSecretKeyJWT() {

	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		panic(err)
	}

	secretKey := base64.StdEncoding.EncodeToString(key)
	fmt.Println("Secret Key:", secretKey)

}

func main() {
	createPasetoKey()
	createSecretKeyJWT()
}
