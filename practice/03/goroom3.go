package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
switch문을 이용하여 score를 grade로 만들어주세요.

90점 이상 A

80점 이상 B

70점 이상 C

60점 이상 D

그 이하 F
*/

func getGrade2(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
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
		fmt.Println(getGrade2(num))
	}
}
