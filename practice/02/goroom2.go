package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
if/if-else/else 등을 이용하여 score를 grade로 만들어주세요.

90점 이상 A

80점 이상 B

70점 이상 C

60점 이상 D

그 이하 F
*/

func getGrade(score int) string {
	if score >= 90 {
		return "A"
	} else if score >= 80 {
		return "B"
	} else if score >= 70 {
		return "C"
	} else if score >= 60 {
		return "D"
	} else {
		return "F"
	}
}

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	strNums := strings.Split(input, ",")

	for _, str := range strNums {
		num, _ := strconv.Atoi(str)
		fmt.Println(getGrade(num))
	}
}
