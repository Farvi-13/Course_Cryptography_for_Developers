package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func main() {
	p, g := generateParams()

	a, b := generateKeys(p, g)

	message := "Hello!"

	r, s := signMessage(message, p, g, a)

	isValid := verifySignature(message, p, g, b, r, s)
	fmt.Println("Is signature valid?", isValid)

	damagedMessage := "Hello, World?"
	damagedIsValid := verifySignature(damagedMessage, p, g, b, r, s)
	fmt.Println("Is signature valid with damaged data?", damagedIsValid)
}

func generateParams() (*big.Int, *big.Int) {
	minBits := 2048

	// Generate a random prime number p
	p, err := rand.Prime(rand.Reader, minBits)
	if err != nil {
		panic(err)
	}

	g := findPrimitiveRoot(p)

	return p, g
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
	// Calculate Euler's totient function φ(p) for a prime p
	phi := new(big.Int).Sub(p, big.NewInt(1))

	// Factorize φ(p) to find its prime factors
	factors := factorize(phi)

	// Iterate through numbers to find the smallest primitive root
	for g := big.NewInt(2); g.Cmp(p) < 0; g = g.Add(g, big.NewInt(1)) {
		isPrimitiveRoot := true

		// Check if g is coprime with φ(p)
		if new(big.Int).GCD(nil, nil, g, phi).Cmp(big.NewInt(1)) != 0 {
			continue
		}

		// Check if g is a primitive root
		for _, q := range factors {
			exponent := new(big.Int).Div(phi, q)
			potentialRoot := new(big.Int).Exp(g, exponent, p)

			if potentialRoot.Cmp(big.NewInt(1)) == 0 {
				isPrimitiveRoot = false
				break
			}
		}

		// If g passes all tests, it is a primitive root
		if isPrimitiveRoot {
			return g
		}
	}

	return nil
}

func generateKeys(p, g *big.Int) (*big.Int, *big.Int) {
	a, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
	if err != nil {
		panic(err)
	}

	b := new(big.Int).Exp(g, a, p)

	return a, b
}

func signMessage(message string, p, g, a *big.Int) (*big.Int, *big.Int) {
	hash := sha256.Sum256([]byte(message))
	h := new(big.Int).SetBytes(hash[:])

	k, err := rand.Int(rand.Reader, new(big.Int).Sub(p, big.NewInt(1)))
	if err != nil {
		panic(err)
	}

	r := new(big.Int).Exp(g, k, p)

	ainv := new(big.Int).ModInverse(a, new(big.Int).Sub(p, big.NewInt(1)))
	s := new(big.Int).Mul(ainv, h)
	s.Sub(s, new(big.Int).Mul(a, r))
	s.Mul(s, new(big.Int).ModInverse(k, new(big.Int).Sub(p, big.NewInt(1))))
	s.Mod(s, new(big.Int).Sub(p, big.NewInt(1)))

	return r, s
}

func verifySignature(message string, p, g, b, r, s *big.Int) bool {
	hash := sha256.Sum256([]byte(message))
	h := new(big.Int).SetBytes(hash[:])

	y := new(big.Int).ModInverse(b, p)

	sinv := new(big.Int).ModInverse(s, new(big.Int).Sub(p, big.NewInt(1)))
	u1 := new(big.Int).Mul(h, sinv)
	u1.Mod(u1, new(big.Int).Sub(p, big.NewInt(1)))

	u2 := new(big.Int).Mul(r, sinv)
	u2.Mod(u2, new(big.Int).Sub(p, big.NewInt(1)))

	v1 := new(big.Int).Exp(g, u1, p)
	v2 := new(big.Int).Exp(y, u2, p)
	v := new(big.Int).Mul(v1, v2)
	v.Mod(v, p)

	return v.Cmp(r) == 0
}
