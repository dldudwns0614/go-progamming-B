// 'ls -a' 는 숨겨진 파일 보이기, 'ls'는 그냥 파일만
// 'go mod init week05' go mod 파일 만들기
// 'go run .\main.go' 파일 실행
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Guess number game!")
	answer := rand.Intn(100) + 1 //1~100
	fmt.Println(answer)
}
