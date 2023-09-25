package main

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

func main() {
	var c string
	//var c7 string  // 변수명은 대소문자로 시작
	//var 7c string  //숫자먼저 X
	var d int
	var e bool
	var f float64
	var G = 99
	// korean := "정찬성" //문접적으로 문제없지만 came1케이스를 사용하는 관례를 따르자
	korean := "정찬성"
	fmt.Println(korean)

	fmt.Println(c, d, e, f, G)
	//var a int //declaration
	//a = 7     //assing value

	// var a int = 7 //declaration&assing value
	//var a = 7 //declaration&assing value
	a := 7 //declaration&assing value
	fmt.Println(a, reflect.TypeOf(a))

	//b := 8.34 //float64
	var b float32 = 8.34
	fmt.Println(a * int(b))
	fmt.Println(float32(a) > b)
	fmt.Println(b, reflect.TypeOf(b))

	fmt.Println('Z', '2', '\n', '김', '인', '하') // rune litarlers
	fmt.Println(reflect.TypeOf('Z'), reflect.TypeOf(2), reflect.TypeOf("Hi"), reflect.TypeOf(4.99), reflect.TypeOf(false))

	fmt.Println(math.Floor(2.17), math.Ceil(2.17), math.Sqrt(16))
	fmt.Println(strings.Title("open source\tprogamming\n\"go\"!"))
	//fmt.Println("OpenSource Progamming~", "Go")
}
