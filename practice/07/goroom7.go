package main

import (
	"fmt"
	"strings"
)

/*
inputs에는 input으로 넣는 값들이 배열로 들어가있습니다.

for문을 돌면서, inputs에 있는 값들을 출력해주세요.
*/

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	inputs := strings.Split(input, ",")
	for _, value := range inputs {
		fmt.Println(value)
	}
}
