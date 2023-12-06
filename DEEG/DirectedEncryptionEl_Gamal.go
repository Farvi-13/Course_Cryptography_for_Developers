package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func generateRandomPrime(bits int) *big.Int {
	prime, _ := rand.Prime(rand.Reader, bits)
	return prime
}

func factorize(n *big.Int) []*big.Int {
	result := []*big.Int{}
	divisor := big.NewInt(2)

	for new(big.Int).Mod(n, divisor).Cmp(big.NewInt(0)) == 0 {
		result = append(result, new(big.Int).Set(divisor))
		n.Div(n, divisor)
	}

	for divisor.Cmp(n) < 0 {
		divisor = divisor.Add(divisor, big.NewInt(1))
		for new(big.Int).Mod(n, divisor).Cmp(big.NewInt(0)) == 0 {
			result = append(result, new(big.Int).Set(divisor))
			n.Div(n, divisor)
		}
	}

	if n.Cmp(big.NewInt(1)) > 0 {
		result = append(result, new(big.Int).Set(n))
	}

	return result
}
func findPrimitiveRoot(p *big.Int) *big.Int {
	phi := new(big.Int).Sub(p, big.NewInt(1))

	factors := factorize(phi)

	for g := big.NewInt(2); g.Cmp(p) < 0; g = g.Add(g, big.NewInt(1)) {
		isPrimitiveRoot := true

		// Check if g is coprime with φ(p)
		if new(big.Int).GCD(nil, nil, g, phi).Cmp(big.NewInt(1)) != 0 {
			continue
		}

		for _, q := range factors {
			exponent := new(big.Int).Div(phi, q)
			potentialRoot := new(big.Int).Exp(g, exponent, p)

			if potentialRoot.Cmp(big.NewInt(1)) == 0 {
				isPrimitiveRoot = false
				break
			}
		}

		if isPrimitiveRoot {
			return g
		}
	}

	return nil
}

func generatePrivateKey(p *big.Int) *big.Int {
	one := big.NewInt(1)
	pMinusOne := new(big.Int).Sub(p, one)

	privateKey, _ := rand.Int(rand.Reader, pMinusOne)
	return privateKey.Add(privateKey, one)
}

func generatePublicKey(g, a, p *big.Int) *big.Int {
	return new(big.Int).Exp(g, a, p)
}

func encrypt(message, p, g, b *big.Int) (*big.Int, *big.Int) {
	k, _ := rand.Int(rand.Reader, p)
	x := new(big.Int).Exp(g, k, p)
	temp := new(big.Int).Exp(b, k, p)
	temp.Mul(temp, message)
	y := temp.Mod(temp, p)
	return x, y
}

func decrypt(x, y, a, p *big.Int) *big.Int {
	s := new(big.Int).Exp(x, a, p)
	sInverse := new(big.Int).ModInverse(s, p)
	temp := new(big.Int).Mul(y, sInverse)
	return temp.Mod(temp, p)
}

func main() {
	bits := 2048
	p := generateRandomPrime(bits)

	g := findPrimitiveRoot(p)

	a := generatePrivateKey(p)
	b := generatePublicKey(g, a, p)

	message := big.NewInt(123456789)

	x, y := encrypt(message, p, g, b)

	decrypted := decrypt(x, y, a, p)

	fmt.Println("Original Message:", message)
	fmt.Println("Decrypted Message:", decrypted)
}
