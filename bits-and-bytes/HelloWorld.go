package main

import (
	"fmt"
	"math"  
)

func main() {
	fmt.Println("Hello World")
	showVariables()

}

var GlobalA string = "this is a global variable"

const GlobalConstantA string = "This is a global constant"
const PI float32 = 3.14

func showVariables() {
	// BASIC VARIABLE TYPES
	PI := 3.14
	var LocalA string = "this is a local variable"
	fmt.Println(GlobalA)
	fmt.Println(LocalA)
	fmt.Println(PI)

	fmt.Println("=============================")

	// NON INT DATA TYPES
	var XString string = "I am a string"
	fmt.Printf("String: %s Bytes:%d \n", XString, len(XString))
	var XFloat32 float32 = 3.1
	fmt.Printf("XFloat32: %f Binary:%b \n", XFloat32, math.Float32bits(XFloat32))
	var XFloat64 float64 = 3.1
	fmt.Printf("XFloat64: %f Binary:%b \n", XFloat64, math.Float64bits(XFloat64))
	var XByte byte = 1
	fmt.Printf("XByte: %d Binary:%08b \n", XByte, XByte)
	var XSlice []byte
	fmt.Printf("Slide before data:%d Binary:%08b \n", XSlice, XSlice)
	XSlice = append(XSlice, 1, 2, 3)
	fmt.Printf("Slide after data:%d Binary:%08b \n", XSlice, XSlice)
	XSlice = append([]byte{4, 5, 6}, XSlice...)
	fmt.Printf("Slide after data 2nd time:%d Binary:%08b \n", XSlice, XSlice)
	XSlice = append(XSlice, 7, 8, 9)
	fmt.Printf("Slide after data 3nd time:%d Binary:%08b \n", XSlice, XSlice)

	fmt.Println("=============================")

	// LOOPS
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	// for i := 0; i < len(XSlice); i++ {
	// 	fmt.Println("Loop index:", i, "Value:", XSlice[i])
	// }

	for index := range XSlice {
		fmt.Println("Loop index: ", index, "Value:", XSlice[index], "Length:", len(XSlice))
	}

}
