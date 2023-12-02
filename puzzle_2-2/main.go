package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/dlclark/regexp2" // needed for lookaheads `?=`
)

func main() {
	file, _ := os.Open("full.input")
	scanner := bufio.NewScanner(file)
	re := regexp2.MustCompile(`Game\s(?<game>\d*):|(?<blue>\d*)\s(blue)|(?<red>\d*)\s(red)|(?<green>\d*)\s(green)`, regexp2.None)
	var totalPower, blueCnt, blueMax, redCnt, redMax, greenCnt, greenMax int
	for scanner.Scan() {
		blueCnt, blueMax, redCnt, redMax, greenCnt, greenMax = 0, 0, 0, 0, 0, 0
		line := scanner.Text()
		res, _ := re.FindStringMatch(line)
		if res != nil {
			for res != nil {
				// Blue
				blueCnt, _ = strconv.Atoi(res.GroupByName("blue").Capture.String())
				if blueCnt > blueMax {
					blueMax = blueCnt
				}
				// Red
				redCnt, _ = strconv.Atoi(res.GroupByName("red").Capture.String())
				if redCnt > redMax {
					redMax = redCnt
				}

				// Green
				greenCnt, _ = strconv.Atoi(res.GroupByName("green").Capture.String())
				if greenCnt > greenMax {
					greenMax = greenCnt
				}

				res, _ = re.FindNextMatch(res)
			}
			totalPower += blueMax * redMax * greenMax
		}
	}
	fmt.Printf("totalPower: %d \n", totalPower)
}
