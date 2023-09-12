package main

import (
	"fmt"
	"os"
	"testing"
)

func TestSplitFileByChunk(t *testing.T) {
	inputPath := "test.txt"
	outputPathPrefix := "split_result/"
	chunkSize := 2

	err := splitFileByChunks(inputPath, outputPathPrefix, chunkSize)
	if err != nil {
		t.Errorf("Error splitting file: %v", err)
	}

	// Verify the number of generated output files
	for i := 0; i < 3; i++ {
		chunkFilePath := fmt.Sprintf("%s-%03d", outputPathPrefix, i)
		_, err := os.Stat(chunkFilePath)
		if err != nil {
			t.Errorf("Output file %s not found: %v", chunkFilePath, err)
		}
	}
}

func TestSplitFileByLine(t *testing.T) {
	inputPath := "test.txt"
	outputPathPrefix := "split_result/"
	linesPerChunk := 2

	err := splitFileByLines(inputPath, outputPathPrefix, linesPerChunk)
	if err != nil {
		t.Errorf("Error splitting file: %v", err)
	}

	// Verify the number of generated output files
	for i := 0; i < 2; i++ {
		chunkFilePath := fmt.Sprintf("%s-%03d", outputPathPrefix, i)
		_, err := os.Stat(chunkFilePath)
		if err != nil {
			t.Errorf("Output file %s not found: %v", chunkFilePath, err)
		}
	}
}

func TestSplitFileByByte(t *testing.T) {
	inputPath := "test.txt"
	outputPathPrefix := "split_result/"
	bytesPerFile := 2

	err := splitFileByBytes(inputPath, outputPathPrefix, bytesPerFile)
	if err != nil {
		t.Errorf("Error splitting file: %v", err)
	}

	// Verify the number of generated output files
	for i := 0; i < 3; i++ {
		chunkFilePath := fmt.Sprintf("%s-%03d", outputPathPrefix, i)
		_, err := os.Stat(chunkFilePath)
		if err != nil {
			t.Errorf("Output file %s not found: %v", chunkFilePath, err)
		}
	}
}
