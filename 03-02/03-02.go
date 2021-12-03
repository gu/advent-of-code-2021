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
	element := os.Args[2]

	data, err := os.ReadFile(filename)
	check(err)

	dataStr := string(data)
	dataArr := strings.Split(dataStr, "\n")

	targetLength := len(dataArr[0])

	startingArr := dataArr

	for i := 0; i < targetLength; i++ {

		zeros := 0
		ones := 0

		for j := 0; j < len(startingArr); j++ {
			if len(startingArr[j]) != targetLength {
				continue
			}

			val, err := strconv.Atoi(startingArr[j][i : i+1])
			check(err)

			if val == 0 {
				zeros++
			} else if val == 1 {
				ones++
			}
		}

		filterVal := 1

		if element == "o2" {

			if zeros > ones {
				filterVal = 0
			}

		} else if element == "co2" {

			filterVal = 0

			if ones < zeros {
				filterVal = 1
			}

		}

		nextArr := make([]string, 0)

		for j := 0; j < len(startingArr); j++ {
			if len(startingArr[j]) != targetLength {
				continue
			}

			val, err := strconv.Atoi(startingArr[j][i : i+1])
			check(err)

			if val == filterVal {
				nextArr = append(nextArr, startingArr[j])
			}
		}

		startingArr = nextArr

		if len(startingArr) == 1 {
			break
		}
	}

	fmt.Println(startingArr)
	val, err := strconv.ParseInt(startingArr[0], 2, 64)
	check(err)
	fmt.Println(val)

}
