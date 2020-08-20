package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Points([]string{"1:0", "2:0", "3:0", "4:0", "2:1", "3:1", "4:1", "3:2", "4:2", "4:3"}))
	fmt.Println(Points([]string{"1:1", "2:2", "3:3", "4:4", "2:2", "3:3", "4:4", "3:3", "4:4", "4:4"}))
}

func Points(games []string) int {
	points := 0
	for _, v := range games {
		split := strings.Split(v, ":")
		var home string = split[0]
		var away string = split[1]

		homeNum, _ := strconv.ParseInt(home, 10, 64)
		awayNum, _ := strconv.ParseInt(away, 10, 64)

		if homeNum > awayNum {
			points += 3
		} else if homeNum == awayNum {
			points++
		}
	}
	return points
}
