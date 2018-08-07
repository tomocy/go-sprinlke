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
	transformsFileName = "transforms.txt"
	otherWord          = "*"
)

var transforms []string

func init() {
	transforms = gatherTransformsFromFile(transformsFileName)
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))
	}
}

func gatherTransformsFromFile(fileName string) []string {
	transforms := make([]string, 0)
	file, err := os.Open(fileName)
	if err != nil {
		return transforms
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		transforms = append(transforms, scanner.Text())
	}

	return transforms
}
