package input

import (
	"bufio"
	"fmt"
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
	tokens := strings.Split(strings.TrimSpace(s), ",")
	values := make([]int, len(tokens))
	for i, t := range tokens {
		v, err := strconv.Atoi(t)
		if err != nil {
			log.Fatal(fmt.Sprintf("Unable to convert %s to int", t) )
		}
		values[i] = v
	}
	return values
}

func Digits(s string) []int {
	digits := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		digits[i] = int(s[i]-48) // this makes me slightly uncomfortable.
	}
	return digits
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
