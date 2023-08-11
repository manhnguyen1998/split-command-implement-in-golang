// -- main.go --
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func check(e error) {
	if e != nil {
			panic(e)
	}
}

func IsValidArgument(args []string)(valid bool) {
	fmt.Println(args)
	if len(args) < 4 {
		return false
	}

	f, err := os.Open(args[3])
	if err != nil {
		fmt.Println(f, args[3], "is not found")
		return false
	}

	if args[1] == "-b" {
		var arg = args[2]
		re:=regexp.MustCompile("\\d+|\\D+")
		var arg_list = re.FindAllString(arg, -1)
		if len(arg_list) == 1{
			val, err := strconv.Atoi(arg_list[0]) 
			if err!= nil {
				fmt.Println("val", val)
				return false
			}
			return true
		}
		if len(arg_list) == 2 {
			val, err := strconv.Atoi(arg_list[0]) 
			if err!= nil {
				fmt.Println("val", val)
				return false
			}

			switch arg_list[1] {
			case "K" :
				return true
			case "k" :
				return true
			case "M" :
				return true
			case "m" :
				return true
			case "G" :
				return true
			case "g" :
				return true
			case "B" :
				return true
			case "b" :
				return true
			default:
				return false
			}
		}
	}

	if args[1] == "-l" {
		val, err := strconv.Atoi(args[2]) 
		if err!= nil {
			fmt.Println("val", val)
			return false
		}
		return true
	}

	if args[1] == "-n" {
		val, err := strconv.Atoi(args[2]) 
		if err!= nil {
			fmt.Println("val", val)
			return false
		}
		return true
	}

	return false
}

func SplitFileByLine(inputPath, outputPathPrefix string, linesPerFile int) error {
	newpath := filepath.Join(".", "split_result")
	os.MkdirAll(newpath, os.ModePerm)

	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	chunkNumber := 0
	lineCount := 0

	chunkFilePath := func() string {
		return fmt.Sprintf("%s-%03d", outputPathPrefix, chunkNumber)
	}

	chunkFile, err := os.Create(chunkFilePath())
	if err != nil {
		return err
	}
	defer chunkFile.Close()

	for scanner.Scan() {
		_, err := chunkFile.WriteString(scanner.Text() + "\n")
		if err != nil {
			return err
		}

		lineCount++
		if lineCount >= linesPerFile {
			chunkFile.Close()
			chunkNumber++
			chunkFile, err = os.Create(chunkFilePath())
			if err != nil {
				return err
			}
			lineCount = 0
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func SplitFileByChunk(inputPath, outputPathPrefix string, chunkSize int) error {
	newpath := filepath.Join(".", "split_result")
	os.MkdirAll(newpath, os.ModePerm)

	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	buffer := make([]byte, chunkSize)
	chunkNumber := 0

	for {
		n, err := inputFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		chunkFilePath := fmt.Sprintf("%s-%03d", outputPathPrefix, chunkNumber)
		chunkFile, err := os.Create(chunkFilePath)
		if err != nil {
			return err
		}

		_, err = chunkFile.Write(buffer[:n])
		if err != nil {
			chunkFile.Close()
			return err
		}

		chunkFile.Close()
		chunkNumber++
	}

	return nil
}

func SplitFileByByte(inputPath, outputPathPrefix string, optionString string) error {
	newpath := filepath.Join(".", "split_result")
	os.MkdirAll(newpath, os.ModePerm)

	re:=regexp.MustCompile("\\d+|\\D+")
	var arg_list = re.FindAllString(optionString, -1)
	size, err := strconv.Atoi(arg_list[0]) 
	if err == nil {
		fmt.Println("size", size)
	}
	var bytesPerChunk = size
	if len(arg_list) == 2 {
		fmt.Println("split by", arg_list[1])
		if arg_list[1]== "K" || arg_list[1] == "k" {
			bytesPerChunk = size * 1000
		}
		if arg_list[1]== "M" || arg_list[1] == "m" {
			bytesPerChunk = size * 1000000
		}
		if arg_list[1]== "G" || arg_list[1] == "g" {
			bytesPerChunk = size * 1000000000
		}
	}

	inputFile, err := os.Open(inputPath)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	chunkNumber := 0

	for {
		buffer := make([]byte, int(bytesPerChunk))
		n, err := inputFile.Read(buffer)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		chunkFilePath := fmt.Sprintf("%s-%03d", outputPathPrefix, chunkNumber)
		chunkFile, err := os.Create(chunkFilePath)
		if err != nil {
			return err
		}

		_, err = chunkFile.Write(buffer[:n])
		if err != nil {
			chunkFile.Close()
			return err
		}

		chunkFile.Close()
		chunkNumber++
	}

	return nil
}

func SplitFile(args []string)(str string) {
	var file_path = args[3]
	if args[1] == "-l" {
		val, err := strconv.Atoi(args[2])
		check(err)
		SplitFileByLine(file_path, "./split_result/", val)
		fmt.Println("split file by lines successfully")
		return "split file by lines successfully"
	}
	if args[1] == "-n" {
		val, err := strconv.Atoi(args[2])
		check(err)
		SplitFileByChunk(file_path, "./split_result/", val)
		fmt.Println("split file by chunks successfully")
		return "split file by chunks successfully"
	}
	if args[1] == "-b" {
		SplitFileByByte(file_path, "./split_result/", args[2])
		fmt.Println("split file by bytes successfully")
		return "split file by bytes successfully"
	}
	return "there was an error"
}

func main() {
	if !IsValidArgument(os.Args) {
		fmt.Println("Invalid argument")
		fmt.Println("usage: split [-l line_count] [-a suffix_length] [file [prefix]]")
		fmt.Println("\tsplit -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]")
		fmt.Println("\tsplit -n chunk_count [-a suffix_length] [file [prefix]]")
		return
	}

	SplitFile(os.Args)
}
