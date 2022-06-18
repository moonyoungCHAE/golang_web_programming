package main

/*
map을 초기화해보세요.
key: int
value: stirng
map의 이름: myMap
myMap을 순회하여, myMap의 key, value를 출력해보세요.
*/

import "fmt"

func main() {
	myMap := make(map[int]string)
	myMap[1] = "Seoul"
	myMap[2] = "London"

	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
