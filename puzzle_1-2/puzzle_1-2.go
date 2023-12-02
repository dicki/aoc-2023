package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/dlclark/regexp2" // needed for lookaheads `?=`
)

func main() {
	file, _ := os.Open("puzzle_1-2.input")
	scanner := bufio.NewScanner(file)
	re := regexp2.MustCompile(`(?=(\d|one|two|three|four|five|six|seven|eight|nine))`, regexp2.None)
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		res, _ := re.FindStringMatch(line)
		var digits string
		for res != nil {
			matchedString := res.GroupByNumber(1).Capture.String()
			switch matchedString {
			case "one":
				digits += "1"
			case "two":
				digits += "2"
			case "three":
				digits += "3"
			case "four":
				digits += "4"
			case "five":
				digits += "5"
			case "six":
				digits += "6"
			case "seven":
				digits += "7"
			case "eight":
				digits += "8"
			case "nine":
				digits += "9"
			default:
				digits += matchedString
			}
			res, _ = re.FindNextMatch(res)
		}

		a, _ := strconv.Atoi(string(digits[0]) + string(digits[len(digits)-1]))
		total += a
	}
	println(total)
}
