// 'ls -a' 는 숨겨진 파일 보이기, 'ls'는 그냥 파일만
// 'go mod init week05' go mod 파일 만들기
// 'go run .\main.go' 파일 실행
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Input score : ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputScoreString = strings.TrimSpace(inputScoreString)      //remove space
	inputScore, err := strconv.ParseFloat(inputScoreString, 32) //string to 32bit float
	fmt.Println(inputScore)
	var grade string
	if inputScore >= 90 {
		grade := "A grade!"
		// fmt.Println("you got", grade)
	} else {
		grade = "under A grade..."
		//fmt.Println("you got", grade)
	}
	fmt.Println("you got", grade)
}
