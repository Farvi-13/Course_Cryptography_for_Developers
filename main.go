package main

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec/v2"
)

type ECPoint struct {
	X *big.Int
	Y *big.Int
}

// G-generator receiving
func BasePointGGet() ECPoint {
	curve := btcec.S256()
	baseX, baseY := curve.Params().Gx, curve.Params().Gy
	return ECPoint{X: baseX, Y: baseY}
}

// ECPoint creation
func ECPointGen(x, y *big.Int) ECPoint {
	return ECPoint{X: x, Y: y}
}

// DOES P ∈ CURVE?
func IsOnCurveCheck(a ECPoint) bool {
	curve := btcec.S256()
	return curve.IsOnCurve(a.X, a.Y)
}

// P + Q
func AddECPoints(a, b ECPoint) ECPoint {
	curve := btcec.S256()
	ax, ay := curve.Add(a.X, a.Y, b.X, b.Y)
	return ECPoint{X: ax, Y: ay}
}

// 2P
func DoubleECPoints(a ECPoint) ECPoint {
	curve := btcec.S256()
	ax, ay := curve.Double(a.X, a.Y)
	return ECPoint{X: ax, Y: ay}
}

// k * P
func ScalarMult(k *big.Int, a ECPoint) ECPoint {
	curve := btcec.S256()
	ax, ay := curve.ScalarMult(a.X, a.Y, k.Bytes())
	return ECPoint{X: ax, Y: ay}
}

// Serialize point
func ECPointToString(point ECPoint) string {
	return fmt.Sprintf("(%s,%s)", point.X.String(), point.Y.String())
}

// Deserialize point
func StringToECPoint(s string) ECPoint {
	var x, y big.Int
	fmt.Sscanf(s, "(%s,%s)", &x, &y)
	return ECPoint{X: &x, Y: &y}
}

// Print point
func PrintECPoint(point ECPoint) {
	fmt.Printf("(%s,%s)\n", point.X.String(), point.Y.String())
}

func SetRandom(bits int) *big.Int {
	max := new(big.Int).Lsh(big.NewInt(1), uint(bits))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		panic(err)
	}
	return n
}

func main() {
	// Your test
	G := BasePointGGet()
	k := SetRandom(256)
	d := SetRandom(256)

	H1 := ScalarMult(d, G)
	H2 := ScalarMult(k, H1)

	H3 := ScalarMult(k, G)
	H4 := ScalarMult(d, H3)

	fmt.Println(ECPointToString(H2) == ECPointToString(H4))

	// My test
	// Test BasePointGGet
	basePoint := BasePointGGet()
	fmt.Println("Base Point G:")
	PrintECPoint(basePoint)
	fmt.Println("Is base point on curve:", IsOnCurveCheck(basePoint))

	// Test ECPointGen
	x := SetRandom(256)
	y := SetRandom(256)
	testPoint := ECPointGen(x, y)
	fmt.Println("\nTest Point:")
	PrintECPoint(testPoint)

	// Test AddECPoints
	sum := AddECPoints(basePoint, testPoint)
	fmt.Println("\nSum of Base Point G and Test Point:")
	PrintECPoint(sum)

	// Test DoubleECPoints
	doubled := DoubleECPoints(testPoint)
	fmt.Println("\nDouble of Test Point:")
	PrintECPoint(doubled)

	// Test ScalarMult
	scalar := big.NewInt(3)
	scaled := ScalarMult(scalar, testPoint)
	fmt.Println("\nScalar Multiplication (3 * Test Point):")
	PrintECPoint(scaled)

	// Test ECPointToString and StringToECPoint
	serialized := ECPointToString(testPoint)
	fmt.Println("\nSerialized Test Point:", serialized)
	deserialized := StringToECPoint(serialized)
	fmt.Println("Deserialized Test Point:")
	PrintECPoint(deserialized)
}
