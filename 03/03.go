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

	targetLength := len(dataArr[0])

	gamma := ""
	epsilon := ""

	for i := 0; i < targetLength; i++ {
		zeros := 0
		ones := 0

		for j := 0; j < len(dataArr); j++ {
			if len(dataArr[j]) != targetLength {
				continue
			}

			val, err := strconv.Atoi(dataArr[j][i : i+1])
			check(err)

			if val == 0 {
				zeros++
			} else if val == 1 {
				ones++
			}
		}

		if zeros > ones {
			gamma += "0"
			epsilon += "1"
		} else {
			gamma += "1"
			epsilon += "0"
		}
	}

	fmt.Print(gamma)
	fmt.Print(": ")
	gammaInt, err := strconv.ParseInt(gamma, 2, 64)
	check(err)
	fmt.Println(gammaInt)

	fmt.Print(epsilon)
	fmt.Print(": ")
	epsilonInt, err := strconv.ParseInt(epsilon, 2, 64)
	check(err)
	fmt.Println(epsilonInt)

	fmt.Println(gammaInt * epsilonInt)
}
