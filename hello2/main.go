// go의 모든 코드는 package로 시작해야 함. 코드를 묶는 단위
// 패키지명: main은 특별한 키워드. program starting point를 포함하는 패키지
// 프로그램이 실행된다는 건 memory에 load해서 CPU가 읽음
// 어디서부터? 프로그램 시작점부터 (main). 시작점은 한 프로그램당 1개
// 즉 패키지는 main과 여러 다른 패키지들로 구성
package main

// fmt: 패키지명
// fmt 기능을 가져와서 사용하겠다.
import "fmt"

// 함수선언 함수명(main은 키워드: 프로그램 시작점)
func main() {
	// fmt에 있는 함수를 사용
	fmt.Println("Hello Go World")
}
