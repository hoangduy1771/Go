package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// memory()
	files()
}

func memory() {
	var XX [10]int
	fmt.Println("XX has the Length:", len(XX), "Capacity:", cap(XX))
	fmt.Println(XX)
	XX[3] = 4
	fmt.Println(XX)

	fmt.Println("------------------------------------------")

	var YY = make([]int, 0, 10)
	YY = append(YY, 10)
	fmt.Println("YY has the Length:", len(YY), "Capacity:", cap(YY), "Data:", YY)

	fmt.Println("------------------------------------------")

	var ZZ = make([]int, 11, 12)
	fmt.Println("ZZ has the Length:", len(ZZ), "Capacity:", cap(ZZ), "Data:", ZZ)
}

func files() {
	var totalBytesWritten int
	newFile := "file1.txt"
	file, err := os.Create(newFile)
	if err != nil {
		log.Println(err)
		return
	}

	helloWorld := "Hello world\n"
	numberOfBytesWrittenToFile, _ := file.WriteString(helloWorld)
	totalBytesWritten += numberOfBytesWrittenToFile
	numberOfBytesWrittenToFile, _ = file.WriteString(helloWorld)

	log.Println("WROTE", totalBytesWritten, "Bypes to file >> Data:", helloWorld)

	// Go back to the start of the file
	file.Seek(0, 0)

	// Read the full file
	fileBytes, _ := io.ReadAll(file)
	fmt.Println("File byte: ", fileBytes)

	for i := 0; i < 10; i++ {
		OPEN_FILE_AN_WRITE_STRING(file, helloWorld)
	}

	file.Close()

}

func OPEN_FILE_AN_WRITE_STRING(file *os.File, data string) {
	_, err := file.WriteString(data)
	if err != nil {
		log.Println(err)
		return
	}
}
