package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	Create_And_Write_In_New_File()
}

func Create_And_Write_In_New_File() {

	currentTime := time.Now()
	newFile := "home-work-file.txt"
	text := "Ice-cream\n"

	// Create new file
	file, _ := os.Create(newFile)
	fmt.Println(">>File created")

	file.WriteString(text)
	fmt.Println(">>Wrote Ice cream in file")

	// Append the string to the end
	appendStr := currentTime.String()
	file.WriteString(appendStr)
	fmt.Println(">>Append current date in file")

	// find the file end number
	data, _ := ioutil.ReadFile(newFile)
	fmt.Println(data)
	slice_size := len(data)
	fmt.Println(slice_size)

	file.Close()
}
