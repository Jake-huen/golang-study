package main

import (
	"fmt"
	"time"
)

func main() {
	// 변수 선언
	var a int = 20
	// 타입도 알아서 유추해줌
	b := 100
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println("hello")

	// if문
	if b > 20 { // 소괄호 없다
		fmt.Println("커용")
	}

	// 배열 선언1
	어레이 := [3]int{1, 2, 3}
	fmt.Println(어레이)

	// 배열 선언2
	어레이2 := []int{1, 2, 3, 4}
	fmt.Println(어레이2)

	// Map
	맵 := make(map[string]int)
	맵["age"] = 123
	fmt.Println(맵["age"])

	// 포인터
	c := 10
	d := &c
	fmt.Println(d)
	fmt.Println(*d)

	// 함수 실행
	함수()

	// 함수 동시실행
	go 함수1()
	함수2()
}

func 함수() {
	fmt.Println("함수 만들기")
}

func 함수1() {
	for i := 1; i < 100; i++ {
		fmt.Println("함수 동시실행1")
		time.Sleep(time.Second * 1)
	}
}

func 함수2() {
	for i := 1; i < 100; i++ {
		fmt.Println("함수 동시실행2")
		time.Sleep(time.Second * 1)
	}

}
