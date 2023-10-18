package MyBigInt

import (
	"fmt"
	"math/big"
)

type LargeNumber struct {
	data []uint64
}

func NewLargeNumber() *LargeNumber {
	return &LargeNumber{data: make([]uint64, 0)}
}

func (ln *LargeNumber) SetHex(hexString string) error {
	hexInt, success := new(big.Int).SetString(hexString, 16)
	if !success {
		return fmt.Errorf("Invalid hexadecimal string: %s", hexString)
	}

	ln.data = make([]uint64, 0)
	for hexInt.BitLen() > 0 {
		ln.data = append(ln.data, hexInt.Uint64())
		hexInt.Rsh(hexInt, 64)
	}

	return nil
}

func (ln *LargeNumber) GetHex() string {
	hexInt := new(big.Int)
	for i := len(ln.data) - 1; i >= 0; i-- {
		hexInt.Lsh(hexInt, 64)
		hexInt.Add(hexInt, new(big.Int).SetUint64(ln.data[i]))
	}
	return fmt.Sprintf("%X", hexInt)
}

func TestLargeNumber(hexString string) {
	ln := NewLargeNumber()
	inputHex := ln.SetHex(hexString)
	if inputHex != nil {
		fmt.Println("Error:", inputHex)
		return
	}

	outputHex := ln.GetHex()
	fmt.Printf("Input: %s\nOutput: %s\n", hexString, outputHex)

	if hexString == outputHex {
		fmt.Println("Test Passed!")
	} else {
		fmt.Println("Test Failed!")
	}
}

func (ln *LargeNumber) INV(other *LargeNumber) *LargeNumber {
	// Bitwise inversion
	result := NewLargeNumber()
	result.data = make([]uint64, len(other.data))
	for i := range result.data {
		result.data[i] = ^other.data[i]
	}
	return result
}

func (ln *LargeNumber) XOR(a *LargeNumber, b *LargeNumber) *LargeNumber {
	// Bitwise XOR
	minLen := min(len(a.data), len(b.data))
	result := NewLargeNumber()
	result.data = make([]uint64, minLen)
	for i := 0; i < minLen; i++ {
		result.data[i] = a.data[i] ^ b.data[i]
	}

	return result
}

func (ln *LargeNumber) OR(a *LargeNumber, b *LargeNumber) *LargeNumber {
	// Bitwise OR
	minLen := min(len(a.data), len(b.data))
	result := NewLargeNumber()
	result.data = make([]uint64, minLen)
	for i := 0; i < minLen; i++ {
		result.data[i] = a.data[i] | b.data[i]
	}

	return result
}

func (ln *LargeNumber) AND(a *LargeNumber, b *LargeNumber) *LargeNumber {
	// Bitwise AND
	minLen := min(len(a.data), len(b.data))
	result := NewLargeNumber()
	result.data = make([]uint64, minLen)
	for i := 0; i < minLen; i++ {
		result.data[i] = a.data[i] & b.data[i]
	}

	return result
}

func (ln *LargeNumber) ShiftR(a *LargeNumber, n int) *LargeNumber {
	// Right shift by n bits
	wordsToShift := n / 64
	bitsToShift := uint(n % 64)
	result := NewLargeNumber()
	result.data = make([]uint64, len(a.data))

	if wordsToShift < len(a.data) {
		copy(result.data, a.data[wordsToShift:])
	}

	if bitsToShift > 0 {
		carry := uint64(0)
		for i := 0; i < len(result.data); i++ {
			temp := result.data[i]
			result.data[i] >>= bitsToShift
			result.data[i] |= carry
			carry = temp << (64 - bitsToShift)
		}
	}

	return result
}

func (ln *LargeNumber) ShiftL(a *LargeNumber, n int) *LargeNumber {
	// Left shift by n bits
	wordsToShift := n / 64
	bitsToShift := uint(n % 64)
	result := NewLargeNumber()
	result.data = make([]uint64, len(a.data))

	if bitsToShift > 0 {
		carry := uint64(0)
		for i := len(a.data) - 1; i >= 0; i-- {
			temp := a.data[i]
			result.data[i] = a.data[i] << bitsToShift
			result.data[i] |= carry
			carry = temp >> (64 - bitsToShift)
		}
	}

	if wordsToShift > 0 || bitsToShift > 0 {
		result.data = append(make([]uint64, wordsToShift), result.data...)
	}

	return result
}

func (ln *LargeNumber) ADD(other *LargeNumber) *LargeNumber {
	result := NewLargeNumber()
	carry := uint64(0)

	maxLen := max(len(ln.data), len(other.data))

	for i := 0; i < maxLen; i++ {
		var a, b uint64
		if i < len(ln.data) {
			a = ln.data[i]
		}
		if i < len(other.data) {
			b = other.data[i]
		}

		sum := a + b + carry
		result.data = append(result.data, sum)

		carry = sum >> 64
	}

	if carry > 0 {
		result.data = append(result.data, carry)
	}

	return result
}

func (ln *LargeNumber) SUB(other *LargeNumber) *LargeNumber {
	result := NewLargeNumber()
	borrow := uint64(0)

	maxUint64 := ^uint64(0)

	maxLen := max(len(ln.data), len(other.data))

	for i := 0; i < maxLen; i++ {
		var a, b uint64
		if i < len(ln.data) {
			a = ln.data[i]
		}
		if i < len(other.data) {
			b = other.data[i]
		}

		if a < b || (a == b && borrow > 0) {
			a += maxUint64
			borrow = 1
		} else {
			borrow = 0
		}
		sub := a - b
		result.data = append(result.data, sub)
	}

	return result
}

// У МЕНЕ ВІН НЕ ВИХОДИТЬ ЧОГОСЬ, АЛЕ Є ЯКАСЬ ЦЯ СОМНІТЄЛЬНА РЕАЛІЗАЦІЯ НА ПРОСТОРАХ ІНТЕРНЕТА
func (ln *LargeNumber) MOD(divisor *LargeNumber) (*LargeNumber, error) {
	if divisor.IsZero() {
		return nil, fmt.Errorf("Division by zero is not allowed")
	}

	result := NewLargeNumber()
	temp := NewLargeNumber()

	for i := len(ln.data) - 1; i >= 0; i-- {
		// Ensure temp.data has enough capacity to hold the shifted result
		if i < len(temp.data) {
			temp.data[i] = ln.data[i]
		} else {
			temp.data = append(temp.data, ln.data[i])
		}

		for j := 0; j < 64; j++ {
			if temp.Cmp(divisor) >= 0 {
				temp = temp.SUB(divisor)
			}
			temp.ShiftL(result, 1)
		}
	}

	return result, nil
}

func (ln *LargeNumber) IsZero() bool {
	for _, val := range ln.data {
		if val != 0 {
			return false
		}
	}
	return true
}

func (ln *LargeNumber) Cmp(other *LargeNumber) int {
	lenA := len(ln.data)
	lenB := len(other.data)

	if lenA > lenB {
		return 1
	} else if lenA < lenB {
		return -1
	}

	for i := lenA - 1; i >= 0; i-- {
		if ln.data[i] > other.data[i] {
			return 1
		} else if ln.data[i] < other.data[i] {
			return -1
		}
	}

	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
