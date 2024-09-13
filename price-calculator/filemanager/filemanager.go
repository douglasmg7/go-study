package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct{
	InputFilePath string
	OutputFilePath string
}

func (fm *FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Error closing the file")
			fmt.Println(err)
		}
	}(file)

	if err != nil {
		return nil, err
	}

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func (fm *FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)
	if err != nil {
		fmt.Println(err.Error())
		return errors.New("filed to create file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)
	if err != nil {
		return errors.New("file to encode data to json")
	}
	return nil
}

func New(inputPath, outputPath string) *FileManager{
	return &FileManager{
		InputFilePath: inputPath,
		OutputFilePath: outputPath,
	}
}
