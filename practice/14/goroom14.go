package main

/*
generate 함수에서 inputs에 있는 채널에 넣어줍니다.

inputCh에 들어온 값을 출력해보세요.
*/
import (
	"fmt"
	"strings"
)

func main() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}
	inputs := strings.Split(input, ",")
	inputCh := generate(inputs)
	for value := range inputCh {
		fmt.Println(value)
	}
}

func generate(inputs []string) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		for _, input := range inputs {
			ch <- input
		}
	}()
	return ch
}
