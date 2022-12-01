package utils

import (
	"anagram/entity"
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
)

const (
	minLetterCount = 2
)

// HandleFlag gets the file path
func HandleFlag() (string, error) {
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return "", fmt.Errorf("file path not given")
	}

	return args[0], nil
}

// ReadFileContent reads
func ReadFileContent(filePath string) ([]entity.Content, error) {
	contents := make([]entity.Content, 0)
	readFile, err := os.Open(filePath)
	if err != nil {
		return contents, err
	}
	sc := bufio.NewScanner(readFile)
	letterCounter := minLetterCount
	c := entity.Content{
		Words: []string{},
	}
	for sc.Scan() {
		word := sc.Text()
		if len(word) == letterCounter {
			c.Words = append(c.Words, word)
		} else {
			letterCounter = len(word)
			if len(c.Words) > 1 {
				contents = append(contents, c)
			}
			c = entity.Content{
				Words: []string{word},
			}
		}
	}
	if len(c.Words) > 1 {
		contents = append(contents, c)
	}

	return contents, nil
}

// SortStringByCharacter converts string to sorted rune
func SortStringByCharacter(s string) string {
	r := stringToRuneSlice(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})

	return string(r)
}

func stringToRuneSlice(s string) []rune {
	var r []rune
	for _, runeValue := range s {
		r = append(r, runeValue)
	}

	return r
}
