package main

import (
	"os"
)

// create file
func createFile(filePath string) {
	newFilePath := filePath + "NewFile.txt"
	_, err := os.Create(newFilePath)
	if err != nil {
		println(err)
	}
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println(err)
	}

	fileData, err := os.ReadFile(filePath + "NewFile.txt")
	if err != nil {
		println(err)
	}
	println(string(fileData))
	defer file.Close()
}

// Read from file
func readFile(filePath string) string {
	data, err := os.ReadFile(filePath)
	if err != nil {
		println(err)
	}
	return string(data)
}

// Write to file
func writeToFile(filePath string) (string, string) {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println(err)
	}
	fileContents1, err1 := os.ReadFile(filePath)
	if err1 != nil {
		println(err1)
	}
	defer file.Close()

	_, err2 := file.WriteString("\n I have been added through method call")
	if err2 != nil {
		println(err)
	}
	fileContents2, err3 := os.ReadFile(filePath)
	if err3 != nil {
		println(err3)
	}
	defer file.Close()

	return string(fileContents1), string(fileContents2)
}

// Delete Existing File
func deleteFile(filePath string) {
	os.Remove(filePath)
}

func main() {
	//println(readFile("C:\\Users\\achoukek\\Desktop\\Testing.txt"))
	deleteFile("C:\\Users\\SRS\\Desktop\\NewFile.txt")
}
