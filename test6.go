package main

import "fmt"

type xx struct {
	aa string
	bb string
}
type student struct {
	name  string
	age   int
	grade []*xx
}

func main() {
	person := []*student{
		{"tom", 18, []*xx{{"1", "2"}, {"3", "4"}}},
		{"tom2", 19, []*xx{{"1", "2"}, {"3", "4"}}},
	}

	for i := len(person) - 1; i >= 0; i-- {
		test66(person[i])
		fmt.Println(person[i].grade)
	}
}

func test66(s *student) {
	var newXx []*xx
	s.grade = newXx
}
