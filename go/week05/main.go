// 'ls -a' 는 숨겨진 파일 보이기, 'ls'는 그냥 파일만
// 'go mod init week05' go mod 파일 만들기
// 'go run .\main.go' 파일 실행
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScore := reader.ReadString('\n') // 1 variable but reader.ReadString returns 2 values
	fmt.Println(inputScore)
}
