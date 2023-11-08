package main

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/sha3"
	"time"
)

const (
	w      = 64
	rounds = 24
)

type state [5][5]uint64

var roundConstants = [rounds]uint64{
	0x0000000000000001, 0x0000000000008082, 0x800000000000808a, 0x8000000080008000,
	0x000000000000808b, 0x0000000080000001, 0x8000000080008081, 0x8000000000008009,
	0x000000000000008a, 0x0000000000000088, 0x0000000080008009, 0x000000008000000a,
	0x000000008000808b, 0x800000000000008b, 0x8000000000008089, 0x8000000000008003,
	0x8000000000008002, 0x8000000000000080, 0x000000000000800a, 0x800000008000000a,
	0x8000000080008081, 0x8000000000008080, 0x0000000080000001, 0x8000000080008008,
}

func rotateLeft64(x uint64, n uint) uint64 {
	return (x << n) | (x >> (w - n))
}

func blockPermutation(s *state) {
	for round := 0; round < rounds; round++ {
		theta(s)
		rho(s)
		pi(s)
		chi(s)
		iota(s, round)
	}
}

func theta(s *state) {
	var c [5]uint64
	var d [5]uint64

	for i := 0; i < 5; i++ {
		c[i] = s[i][0] ^ s[i][1] ^ s[i][2] ^ s[i][3] ^ s[i][4]
	}

	for i := 0; i < 5; i++ {
		d[i] = c[(i+4)%5] ^ rotateLeft64(c[(i+1)%5], 1)
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			s[i][j] ^= d[i]
		}
	}
}

func rho(s *state) {
	shifts := [5][5]int{
		{0, 36, 3, 41, 18},
		{1, 44, 10, 45, 2},
		{62, 6, 43, 15, 61},
		{28, 55, 25, 21, 56},
		{27, 20, 39, 8, 14},
	}

	for t := 0; t < 24; t++ {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				index := (3*i + 2*j) % 5
				tmp := s[index][i]
				s[index][i] = rotateLeft64(s[i][j], uint(shifts[i][j]))
				s[i][j] = tmp
			}
		}
		theta(s)
	}
}

func pi(s *state) {
	var a state
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			a[j][(2*i+3*j)%5] = s[i][j]
		}
	}

	copy(s[:], a[:])
}

func chi(s *state) {
	var t [5]uint64

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			t[j] = s[i][j] ^ ((^s[i][(j+1)%5]) & s[i][(j+2)%5])
		}
		for j := 0; j < 5; j++ {
			s[i][j] = t[j]
		}
	}
}

func iota(s *state, round int) {
	s[0][0] ^= roundConstants[round]
}

func absorb(s *state, message []byte) {
	for len(message) > 0 {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				if len(message) >= 8 {
					s[i][j] ^= binary.LittleEndian.Uint64(message[:8])
					message = message[8:]
				} else {
					for k := 0; k < len(message); k++ {
						s[i][j] ^= uint64(message[k]) << (8 * k)
					}
					s[i][j] ^= 1 << (8 * len(message))
					message = nil
				}
			}
		}
		blockPermutation(s)
	}
}

func pad(message []byte, rate int) []byte {
	// Pad with "10*1"
	message = append(message, 0x06) // binary: 00000110

	// Pad with zeros
	for len(message)%rate != rate-1 {
		message = append(message, 0)
	}

	// Set the last bit to 1
	message = append(message, 0x80) // binary: 10000000

	return message
}

func squeeze(s *state, d int) []byte {
	result := make([]byte, 0, d/8)

	for len(result)*8 < d {
		for i := 0; i < 5; i++ {
			for j := 0; j < 5; j++ {
				result = append(result, byte(s[i][j]))
			}
		}
		blockPermutation(s)
	}

	return result[:d/8]
}

func sha_3(message []byte, d int) []byte {
	var s state

	message = pad(message, 136)

	absorb(&s, message)

	return squeeze(&s, d)
}

func main() {
	message1 := []byte("Hello, SHA-3!")
	startTime := time.Now()
	hash1 := sha_3(message1, 256)
	elapsedTime := time.Since(startTime)
	fmt.Printf("My SHA-3: %x\n", hash1)
	fmt.Println("The execution time of my algorithm: ", elapsedTime)

	startTime1 := time.Now()
	hash_1 := sha3.Sum256(message1)
	elapsedTime1 := time.Since(startTime1)
	fmt.Printf("Library SHA-3: %x\n", hash_1)
	fmt.Println("Runtime of the library algorithm: ", elapsedTime1)

	message2 := []byte("Testing different input data.")
	startTime2 := time.Now()
	hash2 := sha_3(message2, 256)
	elapsedTime2 := time.Since(startTime2)
	fmt.Printf("\nMy SHA-3: %x\n", hash2)
	fmt.Println("The execution time of my algorithm: ", elapsedTime2)

	startTime3 := time.Now()
	hash_2 := sha3.Sum256(message2)
	elapsedTime3 := time.Since(startTime3)
	fmt.Printf("Library SHA-3: %x\n", hash_2)
	fmt.Println("Runtime of the library algorithm: ", elapsedTime3)

	message3 := []byte{0x47, 0xd7, 0x07, 0x99, 0xe8, 0x05, 0x99, 0xd9,
		0x4c, 0x27, 0x3e, 0x81, 0x76, 0x6a, 0x3a, 0x8b,
		0x6c, 0x07, 0x78, 0xf6, 0x05, 0x5a, 0xd2, 0xf4}
	startTime4 := time.Now()
	hash3 := sha_3(message3, 256)
	elapsedTime4 := time.Since(startTime4)
	fmt.Printf("\nMy SHA-3: %x\n", hash3)
	fmt.Println("The execution time of my algorithm: ", elapsedTime4)

	startTime5 := time.Now()
	hash_3 := sha3.Sum256(message3)
	elapsedTime5 := time.Since(startTime5)
	fmt.Printf("Library SHA-3: %x\n", hash_3)
	fmt.Println("Runtime of the library algorithm: ", elapsedTime5)

	message4 := []byte("Another message for comparison.")
	startTime6 := time.Now()
	hash4 := sha_3(message4, 256)
	elapsedTime6 := time.Since(startTime6)
	fmt.Printf("\nMy SHA-3: %x\n", hash4)
	fmt.Println("The execution time of my algorithm: ", elapsedTime6)

	startTime7 := time.Now()
	hash_4 := sha3.Sum256(message4)
	elapsedTime7 := time.Since(startTime7)
	fmt.Printf("Library SHA-3: %x\n", hash_4)
	fmt.Println("Runtime of the library algorithm: ", elapsedTime7)
}
