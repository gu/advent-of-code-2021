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

	aim := 0
	forward := 0
	depth := 0

	for i := 0; i < len(dataArr); i++ {

		lineArr := strings.Split(dataArr[i], " ")

		if len(lineArr) != 2 {
			continue
		}

		value, err := strconv.Atoi(lineArr[1])
		check(err)

		if lineArr[0] == "forward" {
			forward += value
			depth += aim * value
		} else if lineArr[0] == "down" {
			aim += value
		} else if lineArr[0] == "up" {
			aim -= value
		}

	}

	fmt.Println(forward)
	fmt.Println(depth)
	fmt.Println(forward * depth)
}
