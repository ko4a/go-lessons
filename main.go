package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"io"
)

func parseString2(text string)(result []int)  {
	var buf[]int
	for _,number := range strings.Split(text, " "){
		value,_ := strconv.Atoi(number)
		buf = append(buf, value)
	}
	return buf
}

func parseString(text string) (result []int) {
	var buf []int
	var number string

	for pos, char := range text {
		if char != ' ' {
			number += string(char)
		}

		if char == ' ' || pos +1 == len(text) {
			i, _ := strconv.Atoi(number)
			buf = append(buf, i)
			number = ""
		}
	}

	return buf
}

func removeDuplicates(values []int) (result []int) {
	for pos, value := range values {
		var isUniq bool = true

		for i := pos + 1; i < len(values); i++ {
			if value == values[i] {
				isUniq = false
			}
		}

		if isUniq {
			result = append(result, value)
		}

	}
	return result
}



func main() {
	in := bufio.NewScanner(os.Stdin)
	var buf []int

	for in.Scan() {
		buf = parseString(in.Text())
		buf = parseString2(in.Text())
		buf= removeDuplicates(buf)
	}

	fmt.Println(buf)
}
