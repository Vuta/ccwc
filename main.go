package main

import (
	"fmt"
	"flag"
	"os"
	"bufio"
	"io"
	"unicode"
	"errors"
)

type stats struct {
	byteCount uint
	wordCount uint
	lineCount uint
	charCount uint
}

func main() {
	cFlg := flag.Bool("c", false, "number of bytes")
	lFlg := flag.Bool("l", false, "number of lines")
	wFlg := flag.Bool("w", false, "number of words")
	mFlg := flag.Bool("m", false, "number of characters")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Println("Does not support multiple files")
		os.Exit(1)
	}

	filename := flag.Args()[0]
	f, err := os.Open(filename)
	if err != nil {
		os.Exit(1)
	}
	defer f.Close()

	stats := stats{}
	fileInfo, _ := f.Stat()
	// only works with Unix regular file, but who cares 🤷
	stats.byteCount = uint(fileInfo.Size())

	if *cFlg && !(*lFlg || *wFlg || *mFlg) {
		fmt.Printf("    %v    %s\n", stats.byteCount, filename)

		return
	}

	reader := bufio.NewReader(f)
	// accounted for consecutive spaces when counting words
	prevRune := rune(0)

	for {
		c, _, err := reader.ReadRune()

		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}

			os.Exit(1)
		}

		stats.charCount += 1

		if c == '\n' {
			stats.lineCount += 1
		}

		if unicode.IsSpace(c) && !unicode.IsSpace(prevRune) {
			stats.wordCount += 1
		}
		prevRune = c
	}

	result := "   "

	if !(*cFlg || *lFlg || *wFlg || *mFlg) {
		*cFlg = true
		*lFlg = true
		*wFlg = true
	}

	if *lFlg {
		result += fmt.Sprintf("%v    ", stats.lineCount)
	}

	if *wFlg {
		result += fmt.Sprintf("%v    ", stats.wordCount)
	}

	if *cFlg {
		result += fmt.Sprintf("%v    ", stats.byteCount)
	}

	if *mFlg {
		result += fmt.Sprintf("%v    ", stats.charCount)
	}

	result += fmt.Sprintf("%s", filename)

	fmt.Println(result)
}
