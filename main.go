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

var transformPatterns []string

func init() {
	transformPatterns = gatherTransformPatternsFromFile(transformPatternsFileName)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transformPatterns[rand.Intn(len(transformPatterns))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
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
