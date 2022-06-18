package main

/*
NewYear 함수를 호출하면, main 함수에서 member의 Age가 1 추가 되어야 합니다.

NewSalary 함수를 호출하면, 일시적으로 바뀐 Salary이기 때문에 main 함수에서 member의 salary는 바뀌지 않아야 합니다.

힌트: 포인터
*/

import "fmt"

type Member struct {
	Name   string
	Age    int
	Salary int
}

func NewYear(member *Member) {
	member.Age++
}

func NewSalary(member Member) {
	member.Salary += 10000
	fmt.Println(fmt.Sprintf("%s의 이번 달 salary : %d", member.Name, member.Salary))
}

func main() {
	member := Member{"Ann", 29, 10000}

	fmt.Println(fmt.Sprintf("%s의 나이 : %d", member.Name, member.Age))
	NewYear(&member)
	fmt.Println(fmt.Sprintf("새해가 된 후, %s의 나이 : %d", member.Name, member.Age))

	NewSalary(member)
	fmt.Println(fmt.Sprintf("%s의 기존 salary : %d", member.Name, member.Salary))
}
