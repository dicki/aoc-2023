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
	var gameId, validGameCnt, blueCnt, blueTot, redCnt, redTot, greenCnt, greenTot int
	var blueMax, redMax, greenMax int = 14, 12, 13
	var invalidGame bool
	for scanner.Scan() {
		invalidGame = false
		line := scanner.Text()
		res, _ := re.FindStringMatch(line)
		if res != nil {
			gameId, _ = strconv.Atoi(res.GroupByName("game").Capture.String())
			fmt.Printf("game: %d \n", gameId)
			for res != nil {
				// Blue
				blueCnt, _ = strconv.Atoi(res.GroupByName("blue").Capture.String())
				blueTot += blueCnt
				if blueCnt > blueMax {
					invalidGame = true
					break
				}

				// Red
				redCnt, _ = strconv.Atoi(res.GroupByName("red").Capture.String())
				redTot += redCnt
				if redCnt > redMax {
					invalidGame = true
					break
				}

				// Green
				greenCnt, _ = strconv.Atoi(res.GroupByName("green").Capture.String())
				greenTot += greenCnt
				if greenCnt > greenMax {
					invalidGame = true
					break
				}

				res, _ = re.FindNextMatch(res)
			}
			if !invalidGame {
				// fmt.Printf("game: %d, blue: %d, red: %d, green: %d \n", gameId, blueTot, redTot, greenTot)
				validGameCnt += gameId
			}
		}
	}
	fmt.Printf("validGameCnt: %d \n", validGameCnt)
}
