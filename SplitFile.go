package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func splitFileByLines(inputFileName string, outputPathPrefix string, linesPerFile int) error {
	newpath := filepath.Join(".", "split_result")
	os.MkdirAll(newpath, os.ModePerm)

	inputFile, err := os.Open(inputFileName)
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

func splitFileByChunks(inputFileName string, outputPathPrefix string, numberOfChunks int) error {
	newpath := filepath.Join(".", "./split_result")
	os.MkdirAll(newpath, os.ModePerm)

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return err
	}
	defer inputFile.Close()

	buffer := make([]byte, numberOfChunks)
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

func splitFileByBytes(inputFileName string, outputPathPrefix string, bytesPerChunk int) error {
	newpath := filepath.Join(".", "split_result")
	os.MkdirAll(newpath, os.ModePerm)

	// NOTE: option for kilobytes, megabytes, gigabytes
	// re:=regexp.MustCompile("\\d+|\\D+")
	// var arg_list = re.FindAllString(optionString, -1)
	// size, err := strconv.Atoi(arg_list[0]) 
	// if err == nil {
	// 	fmt.Println("size", size)
	// }
	// var bytesPerChunk = size
	// if len(arg_list) == 2 {
	// 	fmt.Println("split by", arg_list[1])
	// 	if arg_list[1]== "K" || arg_list[1] == "k" {
	// 		bytesPerChunk = size * 1000
	// 	}
	// 	if arg_list[1]== "M" || arg_list[1] == "m" {
	// 		bytesPerChunk = size * 1000000
	// 	}
	// 	if arg_list[1]== "G" || arg_list[1] == "g" {
	// 		bytesPerChunk = size * 1000000000
	// 	}
	// }

	inputFile, err := os.Open(inputFileName)
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

func SplitFile(inputFileName string, linesPerFile, numberOfChunks int, bytesPerFile int) error {
	if linesPerFile < 0 || numberOfChunks < 0 || bytesPerFile < 0 {
		return fmt.Errorf("linesPerFile, numberOfFiles, and bytesPerFile must be non-negative")
	}
	if bytesPerFile == 0 && linesPerFile == 0 && numberOfChunks == 0 {
		return fmt.Errorf("At least one of linesPerFile, numberOfFiles, or bytesPerFile must be greater than zero")
	}

	if linesPerFile > 0 {
		return splitFileByLines(inputFileName, "split_result/", linesPerFile)
	} else if numberOfChunks > 0 {
		return splitFileByChunks(inputFileName, "split_result/", numberOfChunks)
	} else {
			return splitFileByBytes(inputFileName,"split_result/", bytesPerFile) 
		}
}
