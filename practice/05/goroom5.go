package main

import "fmt"

/*
numbers1은 int형의 slice입니다.
 numbers1은 1이라는 값을 가집니다.
numbers2는 int형의 slice입니다.
numbers2는 1, 2, 3이라는 값을 가집니다.
*/

func main() {
	var numbers1 []int
	numbers1 = append(numbers1, 1)

	numbers2 := []int{1, 2, 3}

	fmt.Println(numbers1)
	fmt.Println(numbers2)
}
