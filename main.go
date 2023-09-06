// -- main.go --
package main

import (
	"flag"
	"fmt"
)

func main() {
	linesPerFile := flag.Int("l", 0, "number of lines per output file")
	numberOfChunks := flag.Int("n", 0, "number of chunks to split into")  
	bytesPerFile := flag.Int("b", 0, "number of bytes per output file")
	flag.Parse()

	if *linesPerFile <= 0 && *numberOfChunks <= 0 && *bytesPerFile <= 0 {
		fmt.Println("You must specify one of -l, -n, or -b options with a positive value.")
		return
	}

	specifiedOptions := 0
	if *linesPerFile != 0 {
		specifiedOptions++
	}
	if *numberOfChunks != 0 {
		specifiedOptions++
	}
	if *bytesPerFile != 0 {
		specifiedOptions++
	}

	if specifiedOptions > 1 {
		fmt.Println("Error: You cannot specify more than one of -l, -n, and -b options simultaneously.")
		return
	}

	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Usage: split [-l lines] [-n parts] [-b bytes] <file>") 
		return
	}
	fileName := args[0]

	if err := SplitFile(fileName, *linesPerFile, *numberOfChunks, *bytesPerFile); err != nil {
		fmt.Println("Error splitting file:", err)
	}

	// if !IsValidArgument(os.Args) {
	// 	fmt.Println("Invalid argument")
	// 	fmt.Println("usage: split [-l line_count] [-a suffix_length] [file [prefix]]")
	// 	fmt.Println("\tsplit -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]")
	// 	fmt.Println("\tsplit -n chunk_count [-a suffix_length] [file [prefix]]")
	// 	return
	// }

	// SplitFile(os.Args)
}
