// 'ls -a' 는 숨겨진 파일 보이기, 'ls'는 그냥 파일만
// 'go mod init week05' go mod 파일 만들기
// 'go run .\main.go' 파일 실행
package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time //노란색 밑줄 상관 x
	now = time.Now()
	//var year int = now.Year()
	var year = now.Year()
	month := now.Month()
	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())
}
