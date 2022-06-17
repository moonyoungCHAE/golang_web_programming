package main

/*
주어진  숫자들: nums

nums 안에 있는 값들의 합을 구하세요.

main 함수 안에 선언해주세요.

add 함수의 parameter: nums []int, return: int


*/
import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	strNums := strings.Split(input, ",")
	var nums []int
	for _, str := range strNums {
		num, _ := strconv.Atoi(str)
		nums = append(nums, num)
	}

	add := func(nums []int) int {
		count := 0
		for _, num := range nums {
			count += num
		}
		return count
	}
	fmt.Println(add(nums))
}
