package main

import "fmt"

// 매개변수 타입이 같으면 간단히 (a, b int)로 표현 가능
// 반환값이 여러개일 수 있음
func Divide(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	// 콤마로 구분
	return a / b, true
}

func main() {
	c, success := Divide(9, 3)
	fmt.Println(c, success)

	d, success := Divide(9, 0)
	fmt.Println(d, success)
}