package main

import "fmt"

// 중괄호에 속하지 않음
// 패키지 전역 변수
var g int = 10

func main() {
	// 이 함수 안에서만 유효
	var m int = 20

	{
		// 이 중괄호 안에서만 유효
		var s int = 50
		fmt.Println(m, s, g)
	}

	//m = s + 20
}
