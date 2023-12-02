package main

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("puzzle_1.input")
	// writer := bufio.NewWriter(os.Stdout)
	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d+`)
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		result := re.FindAllString(line, -1)
		digits := strings.Join(result, "")
		a, _ := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
		total += a
	}
	println(total)
}
