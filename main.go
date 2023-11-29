package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

type PrivateKey struct {
	D, N *big.Int
}

type PublicKey struct {
	E, N *big.Int
}

func KeyGen() (PrivateKey, PublicKey) {
	p, _ := rand.Prime(rand.Reader, 512)
	q, _ := rand.Prime(rand.Reader, 512)

	n := new(big.Int).Mul(p, q)
	m := new(big.Int).Mul(new(big.Int).Sub(p, big.NewInt(1)), new(big.Int).Sub(q, big.NewInt(1)))

	e := big.NewInt(65537) // Common value for e in practice

	d := new(big.Int)
	d.ModInverse(e, m)

	privateKey := PrivateKey{d, n}
	publicKey := PublicKey{e, n}

	return privateKey, publicKey
}

func Encrypt(message *big.Int, pubKey PublicKey) *big.Int {
	return new(big.Int).Exp(message, pubKey.E, pubKey.N)
}

func Decrypt(ciphertext *big.Int, privKey PrivateKey) *big.Int {
	return new(big.Int).Exp(ciphertext, privKey.D, privKey.N)
}

func TextToBigInt(text string) *big.Int {
	result := big.NewInt(0)

	for _, char := range text {
		charValue := big.NewInt(int64(char))
		result.Mul(result, big.NewInt(512))
		result.Add(result, charValue)
	}

	return result
}

func Verify(message *big.Int, decryptedMessage *big.Int) {
	if decryptedMessage.Cmp(message) == 0 {
		fmt.Println("Decryption successful.")
	} else {
		fmt.Println("Decryption unsuccessful.")
	}
}

func main() {
	privateKey, publicKey := KeyGen()

	text := "Hello!"
	message := TextToBigInt(text)

	ciphertext := Encrypt(message, publicKey)

	decryptedMessage := Decrypt(ciphertext, privateKey)

	fmt.Println("Original Message:", text)
	fmt.Println("Encrypted Message (Signature):", ciphertext)
	Verify(message, decryptedMessage)
}
