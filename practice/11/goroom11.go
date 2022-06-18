package main

/*
NewYear() 함수가 호출되면, main 함수에서도 member의 Age가 증가해야 합니다.

NewSalary() 함수는 일시적으로 salary가 증가한 것이기 때문에, main 함수에서는 Salary의 값이 바뀌면 안됩니다.

힌트:

Pointer Receiver, Value Receiver
*/

import "fmt"

type Member struct {
	Name   string
	Age    int
	Salary int
}

func (member *Member) NewYear() {
	member.Age++
}

func (member Member) NewSalary() {
	member.Salary += 10000
	fmt.Println(fmt.Sprintf("%s의 이번 달 salary : %d", member.Name, member.Salary))
}

func main() {
	member := Member{"Ann", 29, 10000}

	fmt.Println(fmt.Sprintf("%s의 나이 : %d", member.Name, member.Age))
	member.NewYear()
	fmt.Println(fmt.Sprintf("새해가 된 후, %s의 나이 : %d", member.Name, member.Age))

	member.NewSalary()
	fmt.Println(fmt.Sprintf("%s의 기존 salary : %d", member.Name, member.Salary))
}
