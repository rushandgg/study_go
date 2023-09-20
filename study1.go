package main

func main() {

	// 1st day : Install golang and Print
	println("first programing with go!!")

	// 2nd day : Var and Const

	// var 변수, 선언해놓고 나중에 할당 가능
	// 초기값 설정 안하면 Zero Value를 기본적으로 할당 (숫자형은 0, bool은 false, string은 "")
	var a int
	var f float32 = 11.

	a = 10
	f = 12.0

	println(a)
	println(f)

	// 한꺼번에 선언 가능
	var i, j, k int = 3, 4, 5
	println(i, j, k)
	println(i + j + k)

	// 상수, 할당해놓고 나면 나중에 변경 불가능
	// Go에서는 할당되는 값을 보고 타입을 추론할 수 있음
	const c int = 10
	const s string = "const string"
	println(c, s)

	// var처럼 const도 한꺼번에 가능
	const (
		V = "visa"
		M = "Master"
		A = "American Express"
	)
	println(V, M, A)

	// iota 라는 identifier를 사용하면 상수값을 0부터 순차적으로 할당함
	const (
		apple  = iota // 0, iota is identifier
		grape         // 1
		orange        // 2
	)
	println(apple, grape, orange)
}
