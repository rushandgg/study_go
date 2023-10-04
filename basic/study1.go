package main

import (
	"fmt"
)

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

	// 3rd day : Data Type

	// bool
	// string : 한번 생성되면 수정 불가능
	// int, int8, int32, int64
	// uint uint8, uint16, uint32, uint64, uintptr
	// float32, float64, complex64, complex128
	// byte : uint8과 동일, 바이트 코드에 사용
	// rune : int32와 동일, 유니코드 코드포인트에 사용

	// 문자열
	// '' 은 Raw String Literal, 별도로 해석X, Raw String 그대로의 값을 갖음, \n 이 있을 경우 NewLine으로 해석X

	rawLiteral := `야옹\n
	야옹\n
	야옹`

	fmt.Println(rawLiteral)

	fmt.Println()

	interLiteral := "야옹\n야아아아아옹\n야옹"

	fmt.Println(interLiteral)

	// 데이터 타입 변환

	var i2 int = 100
	var u2 uint = uint(i2)
	var f2 float32 = float32(i2)

	println(f2, u2)

	// 4th day : Operator

	// 산술연산자
	f3 := 1
	g3 := 2
	d3 := (f3 + g3) / 3
	println(d3)

	// 관계연산자
	println(f3 == g3)
	println(f3 != g3)
	println(f3 >= g3)

	// Bitwise 연산자
	println((f3 & g3) << 5)

	// 할당연산자
	a = 100
	println(a)
	a *= 10
	println(a)
	a >>= 2
	println(a)
	a |= 1
	println(a)

	// 포인터연산자
	var t int = 10
	var p = &t // k의 주소를 할당
	println(p)
	println(*p) // p가 가리키는 주소에 있는 실제 내용을 출력

	// 5th day : if statement

	// if문 : 해당 조건이 맞으면 {} 안의 내용을 실행, {}의 시작을 if문과 같은 라인에 두어야 함, if문의 조건식은 반드시 Boolean 식으로 표현해야 함
	if t == 1 {
		println("야옹")
	} else {
		println("야옹2")
	}

	// swtich문

	var name string
	var category = 1

	switch category {
	case 1:
		name = "고"
	case 2:
		name = "양"
	case 3, 4:
		name = "E"
	default:
		name = "Other"
	}
	println(name)

	grade(77)

	// 6th day : iteration

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	println(sum)

	// for문 무한루프는 초기값; 조건식; 증감을 모두 생략하면 됨

	// for range 문, 컬렉션으로 부터 위치인덱스와 값을 하나씩 가져옴
	names := []string{"홍길동", "이순신", "강감찬"}
	for index, name := range names {
		println(index, name)
	}

	// break, continue, goto
	// break : for, switch, select문에 사용 가능, 루프를 빠져나옴, 이중 for문의 경우 가장 안쪽 for문만 빠져나온다
	// continue : for문에서만 사용 가능, for루프의 시작으로 감
	// goto : 특정 label로 이동한다.

}
func grade(score int) {
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")

	}
}
