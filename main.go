package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Inter your text: ")
		text, _ := reader.ReadString('\n')
		fmt.Println(textModifier(text))
	}
}

func textModifier(str string) string {
	var text string
	text = replaceSpaces(str)
	text = replaceByMinus(text)
	text = replacePlus(text)
	text = findNumbers(text)

	return text
}

func replaceSpaces(str string) string {
	newStr := strings.Split(str, "")

	var res string

	for i := 0; i < len(newStr); i++ {

		if newStr[i] == " " {
			if i+1 < len(newStr) && newStr[i+1] == " " {
				continue
			}
			res += newStr[i]
		} else {
			res += newStr[i]
		}
	}

	return res
}

func replaceByMinus(str string) string {
	newStr := strings.Split(str, "")

	var res string

	var next string
	for i := 0; i < len(newStr); i++ {

		if i+1 < len(newStr) {
			if newStr[i+1] == "-" {
				next = newStr[i+2]
				res += next
				res += newStr[i]
				continue
			}
			if newStr[i] == "-" {
				continue
			}
			if newStr[i] == next {
				next = ""
				continue
			}

			res += newStr[i]
		}
	}

	return res
}

func replacePlus(str string) string {
	newStr := strings.Split(str, "")

	var result string

	for i := 0; i < len(newStr); i++ {
		if newStr[i] == "+" {
			result += "!"
			continue
		}
		result += newStr[i]
	}
	return result
}

func findNumbers(str string) string {
	newStr := strings.Split(str, "")

	var sum int
	var result string

	for i := 0; i < len(newStr); i++ {

		if num, err := strconv.Atoi(newStr[i]); err == nil {
			if num >= 0 && num <= 9 {
				sum += num
			}
			continue
		}

		result += newStr[i]

	}

	if sum != 0 {
		sumStr := strconv.Itoa(sum)
		resStr := strings.TrimSpace(result)
		resStr = resStr + " " + sumStr
		return resStr
	}

	return result
}
