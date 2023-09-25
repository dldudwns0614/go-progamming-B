// 'ls -a' 는 숨겨진 파일 보이기, 'ls'는 그냥 파일만
// 'go mod init week05' go mod 파일 만들기
// 'go run .\main.go' 파일 실행
package main

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Guess number game!")
	answer := rand.Intn(100) + 1 //1~100
	fmt.Println(answer)
	for i := 0; 1 < 10; i++ {
	fmt.Print("Input number : ")
	reader := bufio.NewReader(os.Stdin)
	inputNumberString, err := reader.ReaderString('\n')
	if err != nil {
		log.Fatal(err)
	}
	inputNumberString = strings.TrimSpace(inputNumberString)
	inputGuess, err := strconv.Atoi(inputNumberString)
	if err != nil {
		log.Fatal(err)
	}
	if inputGuess == answer{
		fmt.Println("great")
		break
	}
	else if inputGuess < answer {
		fmt.Println("lower")
	}
	else if inputGuess > answer{
		fmt.Println("higher")
	}
}
}