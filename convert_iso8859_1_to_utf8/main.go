package main

import (
	"fmt"
	"golang.org/x/text/encoding/charmap"
	"io"
	"os"
)

func main() {
	f, err := os.Open("iso.txt")
	if err != nil {
		fmt.Print("Error:", err)
	}

	out, err := os.Create("utf8.txt")
	if err != nil {
		fmt.Print("Error:", err)
	}

	r := charmap.ISO8859_1.NewDecoder().Reader(f)

	io.Copy(out, r)

	out.Close()
	f.Close()
}
