package input

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func SingleLineFile(path string) string {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func ParseInts(s string) []int {
	tokens := strings.Split(s, ",")
	values := make([]int, len(tokens))
	for i, t := range tokens {
		v, _ := strconv.Atoi(t)
		values[i] = v
	}
	return values
}

func Lines(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	lines := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}