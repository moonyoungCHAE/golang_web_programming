package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
strNumbers은 string을 , 단위로 분리한 상태의 string 타입의 slice입니다.

strNumbers를 int형으로 바꾼 후 (numbers) numbers의 index 2부터 4까지의 값을 출력해주세요.
*/
func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	var numbers []int
	strNumbers := strings.Split(input, ",")
	for _, strNum := range strNumbers {
		strNum, _ := strconv.Atoi(strNum)
		numbers = append(numbers, strNum)
	}
	fmt.Println(numbers[2:5])
}
