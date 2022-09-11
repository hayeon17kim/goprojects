package main

import "fmt"

func main() {
	// assign operator
	// 메모리에 정수값을 담을 수 있는 공간을 마련하고 10 넣음
	var a int = 10
	var msg string = "Hello variable"

	// a 가 가리키는 공간을 찾아 20 save
	a = 20
	msg = "Good Morning"

	fmt.Println(a, msg)
}
