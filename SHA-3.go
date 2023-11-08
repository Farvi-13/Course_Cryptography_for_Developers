package main

import (
	"encoding/binary"
	"fmt"
	"golang.org/x/crypto/sha3"
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

	// Pad the message
	message = pad(message, 136)

	// Absorb
	absorb(&s, message)

	// Squeeze
	return squeeze(&s, d)
}

func main() {
	message := []byte("Hello!")
	hash1 := sha_3(message, 256)
	fmt.Printf("My SHA-3: %x\n", hash1)

	hash_1 := sha3.Sum256(message)
	fmt.Printf("Library SHA-3: %x\n", hash_1)

	message1 := []byte{0x47, 0xd7, 7, 0x99, 0xe8, 5, 0x99, 0xd9,
		0x4c, 0x27, 0x3e, 0x81, 0x76, 0x6a, 0x3a, 0x8b,
		0x6c, 7, 0x78, 0xf6, 0x5, 0x5a, 0xd2, 0xf4}
	hash2 := sha_3(message1, 256)
	fmt.Printf("My SHA-3: %x\n", hash2)

	hash_2 := sha3.Sum256(message1)
	fmt.Printf("Library SHA-3: %x\n", hash_2)
}
