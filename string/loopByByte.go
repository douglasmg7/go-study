package main

import "fmt"

func main() {
	// Indexing a string yields its bytes, not its characters: a string is just a bunch of bytes.
	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Println("Println:")
	fmt.Println(sample)

	fmt.Println("Byte loop:")
	for i := 0; i < len(sample); i++ {
		fmt.Printf("%x ", sample[i])
	}
	fmt.Printf("\n")

	fmt.Println("Printf with %x:")
	fmt.Printf("%x\n", sample)

	fmt.Println("Printf with % x:")
	fmt.Printf("% x\n", sample)

	fmt.Println("Printf with %q:")
	fmt.Printf("%q\n", sample)

	fmt.Println("Printf with %+q:")
	fmt.Printf("%+q\n", sample)

	// A for range loop, by contrast, decodes one UTF-8-encoded rune on each iteration.
	for index, runeValue := range sample {
		fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
	}
}
