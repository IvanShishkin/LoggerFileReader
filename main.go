package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	//"github.com/stoicperlman/fls"
)

const lineLegth = 25

func main() {
	// open the file
	file, err := os.Open("ss2.json.log")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	o2, err := lineCounter(file)
	fmt.Println(o2)

	//o3, err := file.see

	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		//fmt.Println(fileScanner.Text())
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}

func readLine(line int, f *os.File) (string, error) {
	lineBuffer := make([]byte, 24)
	f.Seek(int64(line*lineLegth), 0)
	_, err := f.Read(lineBuffer)
	return string(lineBuffer), err
}

func lineCounter(r io.Reader) (int, error) {

	var count int
	const lineBreak = '\n'

	buf := make([]byte, bufio.MaxScanTokenSize)

	for {
		bufferSize, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return 0, err
		}

		var buffPosition int
		for {
			i := bytes.IndexByte(buf[buffPosition:], lineBreak)
			if i == -1 || bufferSize == buffPosition {
				break
			}
			buffPosition += i + 1
			count++
		}
		if err == io.EOF {
			break
		}
	}

	return count, nil
}
