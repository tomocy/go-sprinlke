package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

const (
	transformPatternsFileName = "transforms.txt"
	otherWord                 = "*"
)

var transformPatterns = []string{
	otherWord,
	otherWord + otherWord,
	"go " + otherWord,
	"let's" + otherWord,
}

func init() {
	transformPatterns = append(
		transformPatterns,
		gatherTransformPatternsFromFile(transformPatternsFileName)...,
	)
}

func gatherTransformPatternsFromFile(fileName string) []string {
	transformPatterns := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		return transformPatterns
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		transformPatterns = append(transformPatterns, scanner.Text())
	}

	return transformPatterns
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		replacedText := replaceWith(scanner.Text())
		fmt.Println(replacedText)
	}
}

func replaceWith(text string) string {
	p := transformPatterns[rand.Intn(len(transformPatterns))]
	replacedText := strings.Replace(p, otherWord, text, -1)

	return replacedText
}
