package main

/*
input이 될 수 있는 key (expectdKeys) ["a","b","c","d","e"]

expectedKeys 중에 실제로 input으로 들어온 값이 있다면 existedKey에 추가해주세요.
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

	// input 값에 있는 Key로 map에 값을 넣어준다.
	myMap := make(map[string]struct{})
	keys := strings.Split(input, ",")
	for _, key := range keys {
		myMap[key] = struct{}{}
	}

	// 모든 key의 후보
	expectedKeys := []string{"a", "b", "c", "d", "e"}
	var exitedKeys []string
	for _, key := range expectedKeys {
		_, ok := myMap[key]
		if ok {
			exitedKeys = append(exitedKeys, key)
		}
	}

	fmt.Println(exitedKeys)
}
