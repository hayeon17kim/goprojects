package main

import "fmt"

func main() {
	math := 80
	eng := 74
	history := 95
	fmt.Println("평균 점수: ", (math + eng + history) / 3)

	math = 88
	eng = 92
	history = 53
	fmt.Println("평균 점수: ", (math + eng + history) / 3)
}