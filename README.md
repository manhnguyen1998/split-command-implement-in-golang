# An implement of split command on Golang

## split – split a file into pieces

# Usage: 
## Split file by lines
```
go run main.go [-l line_count] [file [prefix]]
```
- Create split files line_count lines in length.
## Split file by bytes
```
split -b byte_count[K|k|M|m|G|g] [file [prefix]]
```
- Create split files byte_count bytes in length.  
  - If k or K is appended to the number, the file is split into byte_count kilobyte pieces. 
  - If m or M is appended to the number, the file is split into byte_count megabyte pieces. 
  - If g or G is appended to the number, the file is split into byte_count gigabyte pieces.
  - If b or B or nothing is appended to the number, it's default that split into byte_count byte pieces.
## Split file by chunks
```
split -n chunk_count [-a suffix_length] [file [prefix]]
```
- Split file into chunk_count smaller files.  The first n - 1 files will be of size (size of file / chunk_count ) and the last file will contain the remaining bytes.

# Test
```
go test *.go
```

# Result
- The result files will be written inside `split_result` folder
- If the folder does not exist, the code will create it immediately
