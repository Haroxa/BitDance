package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	MaxNum := 100
	RandomNum := PoseRandomNum(MaxNum)
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Welcome to play the guessing-name!\n")
	for i := 6; i >= 1; i-- {
		GuessNum := ScanNum(reader)
		if GuessNum == -1 {
			continue
		}
		if CheckNum(RandomNum, GuessNum) {
			break
		}
		if i != 1 {
			fmt.Printf("You have %d chances to try!\n", i-1)
		} else {
			fmt.Printf("Failure , you have used up all chances!\n")
		}
	}
	fmt.Printf("The correct answer is [%d]\n", RandomNum)
}

func PoseRandomNum(MaxNum int) int {
	RandomNum := rand.Intn(MaxNum)
	//fmt.Printf("The RandomNum is [%d] \n", RandomNum)
	return RandomNum
}

func ScanNum_01() int {
	fmt.Printf("Please input your guess number:")
	GuessNum := 0
	fmt.Scanf("%d\n", &GuessNum) // 去掉末尾换行，不然会直接跳过
	//fmt.Printf("Your input is [%d] \n", GuessNum)
	return GuessNum
}

func ScanNum(reader *bufio.Reader) int {
	fmt.Printf("Please input your guess number:")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occured while reading input!\n[%s]", err)
		return -1
	}
	//fmt.Printf("Your input is [%s] \n", input)

	input = strings.TrimSuffix(input, "\r\n")

	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Printf("Invalid input. Please enter an integer value!\n[%s]", err)
		return -1
	}

	return guess
}

func CheckNum(RandomNum int, GuessNum int) bool {
	if GuessNum == RandomNum {
		fmt.Printf("Congratulation on the correct answer!\n")
		return true
	} else if GuessNum < RandomNum {
		fmt.Printf("Your guess is less than the answer!\n")
	} else {
		fmt.Printf("Your guess is more than the answer!\n")
	}
	return false
}
