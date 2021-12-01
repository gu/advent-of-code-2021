package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	filename := os.Args[1]

	data, err := os.ReadFile(filename)
	check(err)

	dataStr := string(data)
	dataArr := strings.Split(dataStr, "\n")

	increasedCount := 0

	for i := 3; i < len(dataArr); i++ {

		aCurrStr := dataArr[i-1]
		aPrevStr := dataArr[i-2]
		aLastStr := dataArr[i-3]
		bCurrStr := dataArr[i]
		bPrevStr := dataArr[i-1]
		bLastStr := dataArr[i-2]

		if len(aCurrStr) > 0 && len(bCurrStr) > 0 {
			aCurr, err := strconv.Atoi(aCurrStr)
			check(err)
			aPrev, err := strconv.Atoi(aPrevStr)
			check(err)
			aLast, err := strconv.Atoi(aLastStr)
			check(err)
			bCurr, err := strconv.Atoi(bCurrStr)
			check(err)
			bPrev, err := strconv.Atoi(bPrevStr)
			check(err)
			bLast, err := strconv.Atoi(bLastStr)
			check(err)

			if bCurr+bPrev+bLast > aCurr+aPrev+aLast {
				increasedCount++
			}
		}

	}

	fmt.Println(increasedCount)
}
