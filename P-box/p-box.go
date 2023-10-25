package main

import "fmt"

// Function for a straight P-box
func forwardPBox(inputData uint8) uint8 {
	pBoxPermutation := [8]int{1, 5, 2, 0, 3, 7, 4, 6} // Перестановка для прямого P-бокса

	outputData := uint8(0)
	for i := 0; i < 8; i++ {
		outputData |= ((inputData >> uint(i) & 1) << uint(pBoxPermutation[i]))
	}
	return outputData
}

// Function for the inverse P-box
func inversePBox(inputData uint8) uint8 {
	pBoxPermutation := [8]int{3, 0, 2, 4, 6, 1, 7, 5} // Перестановка для зворотного P-бокса

	outputData := uint8(0)
	for i := 0; i < 8; i++ {
		outputData |= ((inputData >> uint(i) & 1) << uint(pBoxPermutation[i]))
	}
	return outputData
}

func main() {
	inputData := uint8(0b10111100) // Приклад 8-бітових вхідних даних
	encryptedData := forwardPBox(inputData)
	decryptedData := inversePBox(encryptedData)

	fmt.Printf("Input data: %08b\n", inputData)
	fmt.Printf("Encrypted data: %08b\n", encryptedData)
	fmt.Printf("Decrypted data: %08b\n", decryptedData)
}
