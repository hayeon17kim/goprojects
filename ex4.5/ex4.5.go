package main

import "fmt"

func main() {
	var a int16 = 3456
	// 상위 1바이트가 없어짐
	var b int8 = int8(a)

	fmt.Println(a, b)
}
