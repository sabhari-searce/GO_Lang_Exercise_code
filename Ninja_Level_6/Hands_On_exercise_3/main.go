package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("Sample.txt")
	if err != nil {
		panic("Error!!!")
	}
	defer file.Close()

	_, error := file.WriteString("Content to be written before file is closed")
	if error != nil {
		panic("Error in writing to file!!!!")
	}

	data, errr := os.ReadFile("sample.txt")

	if errr != nil {
		panic("Error inreading!!")
	}
	fmt.Println(string(data))

}
