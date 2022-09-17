package main

import "fmt"

// iota는 소괄호를 벗어나면 초기화됨
// 0부터 시작해 1씩 증가
// 첫번째 값과 똑같은 규칙이 계속된다면 타입과 iota 생략 가능

const PI = 3.14
const FloatPI float64 = 3.14

func main() {
	var a int = PI * 100
	var b int = FloatPI * 100
	fmt.Println(a)
	fmt.Println(b) // 에러!
}