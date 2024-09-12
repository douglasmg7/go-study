package filemanager

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
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
