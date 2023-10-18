package main

import (
	"Lab_2/MyBigInt"
	"fmt"
)

func main() {
	MyBigInt.TestLargeNumber("1234ABCD58EF01")
	MyBigInt.TestLargeNumber("AA")
	MyBigInt.TestLargeNumber("1")

	fmt.Printf("\n")

	numberA := MyBigInt.NewLargeNumber()
	numberB := MyBigInt.NewLargeNumber()
	numberA.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	numberB.SetHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")

	// INV
	resultINV := numberA.INV(numberA)
	fmt.Printf("INV Result: %s\n", resultINV.GetHex())

	// XOR
	resultXOR := numberA.XOR(numberA, numberB)
	fmt.Printf("XOR Result: %s\n", resultXOR.GetHex())

	// OR
	resultOR := numberA.OR(numberA, numberB)
	fmt.Printf("OR Result: %s\n", resultOR.GetHex())

	// AND
	resultAND := numberA.AND(numberA, numberB)
	fmt.Printf("AND Result: %s\n", resultAND.GetHex())

	// shiftR
	resultShiftR := numberA.ShiftR(numberA, 5)
	fmt.Printf("ShiftR Result: %s\n", resultShiftR.GetHex())

	// shiftL
	resultShiftL := numberA.ShiftL(numberA, 3)
	fmt.Printf("ShiftL Result: %s\n", resultShiftL.GetHex())

	fmt.Printf("\n")

	numberC := MyBigInt.NewLargeNumber()
	numberD := MyBigInt.NewLargeNumber()
	numberC.SetHex("36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80")
	numberD.SetHex("70983d692f648185febe6d6fa607630ae68649f7e6fc45b94680096c06e4fadb")

	numberF := MyBigInt.NewLargeNumber()
	numberG := MyBigInt.NewLargeNumber()
	numberF.SetHex("33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc")
	numberG.SetHex("22e962951cb6cd2ce279ab0e2095825c141d48ef3ca9dabf253e38760b57fe03")

	resultADD := numberC.ADD(numberD)
	fmt.Printf("ADD Result: %s\n", resultADD.GetHex())

	resultSUB := numberF.SUB(numberG)
	fmt.Printf("SUB Result: %s\n", resultSUB.GetHex())

	resultMOD, err := numberF.MOD(numberG)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("MOD Result: %s\n", resultMOD.GetHex())
}
