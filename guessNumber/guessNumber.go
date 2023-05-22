package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

var stdin = bufio.NewReader(os.Stdin)

func InputIntValue() (int, error) {
	var inputValue int
	_, err := fmt.Scanln(&inputValue)

	if err != nil {
		stdin.ReadString('\n')
	}

	return inputValue, err

}

func main() {

	rand.Seed(time.Now().UnixNano())

	targetValue := rand.Intn(100)
	cnt := 1

	for {
		fmt.Printf("숫자값을 입력하세요:")
		inputValue, err := InputIntValue()

		if err != nil {
			fmt.Println("숫자만 입력하세요.")
		} else {
			if inputValue > targetValue {
				fmt.Println("입력하신 숫자가 더 큽니다.")
			} else if inputValue < targetValue {
				fmt.Println("입력하신 숫자가 더 작습니다.")
			} else {
				fmt.Println("축하합니다. 입력하신 숫자는 ", inputValue, " 입니다. 시도한 횟수:", cnt)
				break
			}

			cnt++
		}
	}
}
