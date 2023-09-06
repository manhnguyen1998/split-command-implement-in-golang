package main

import (
	"fmt"
	"os"
	"testing"
)

// func TestIsValidArgumentTableDriven(t *testing.T) {
// 	var tests = []struct {
// 		arg []string
// 		want bool
// 	} {
// 		{[]string{}, false},
// 		{[]string{"main.go", "test.txt"}, false},
// 		{[]string{"main.go","-l", "test.txt"}, false},
// 		{[]string{"main.go","-l", "2"}, false},
// 		{[]string{"main.go","-l","2", "test.txt"}, true},
// 		{[]string{"main.go","-n","2", "test.txt"}, true},
// 		{[]string{"main.go","-b","2", "test.txt"}, true},
// 		{[]string{"main.go","-b","2b", "test.txt"}, true},
// 		{[]string{"main.go","-b","2B", "test.txt"}, true},
// 		{[]string{"main.go","-b","2k", "test.txt"}, true},
// 		{[]string{"main.go","-b","2K", "test.txt"}, true},
// 		{[]string{"main.go","-b","2m", "test.txt"}, true},
// 		{[]string{"main.go","-b","2M", "test.txt"}, true},
// 		{[]string{"main.go","-b","2g", "test.txt"}, true},
// 		{[]string{"main.go","-b","2G", "test.txt"}, true},
// 		{[]string{"main.go","-b","2l", "test.txt"}, false},
// 	}

// 	for _, tt := range tests {
// 		testname := fmt.Sprintf("test with %s", tt.arg)
// 		t.Run(testname, func(t *testing.T) {
// 			os.Args = tt.arg
// 			actualResult := IsValidArgument(os.Args)
// 			if actualResult != tt.want {
// 				t.Errorf("got %t, want %t", actualResult, tt.want)
// 			}
// 		})
// 	}
// }

// func TestSplitFileTableDriven(t *testing.T) {
// 	var tests = []struct {
// 		arg []string
// 		want string
// 	} {
// 		{[]string{"main.go","-l", "2", "test.txt"}, "split file by lines successfully"},
// 		{[]string{"main.go","-n", "2", "test.txt"}, "split file by chunks successfully"},
// 		{[]string{"main.go","-b", "2k", "test.txt"}, "split file by bytes successfully"},
// 	}

// 	for _, tt := range tests {
// 		testname := fmt.Sprintf("test with %s", tt.arg)
// 		t.Run(testname, func(t *testing.T) {
// 			os.Args = tt.arg
// 			actualResult := SplitFile(os.Args)
// 			if actualResult != tt.want {
// 				t.Errorf("got %s, want %s", actualResult, tt.want)
// 			}
// 		})
// 	}
// }

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
