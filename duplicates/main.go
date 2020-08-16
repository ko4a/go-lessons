package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"io"
)

func nativeParseString(text string)(result []int, err error)  {
	var buf[]int
	for _,number := range strings.Split(text, " "){
		value, err := strconv.Atoi(number)

		if err != nil {
			return nil, err
		}

		buf = append(buf, value)
	}
	return buf, nil
}

func customParseString(text string) (result []int, err error) {
	var buf []int
	var number string

	for pos, char := range text {
		if char != ' ' {
			number += string(char)
		}

		if char == ' ' || pos +1 == len(text) {
			i, err := strconv.Atoi(number)

			if err != nil {
				return nil, err
			}

			buf = append(buf, i)
			number = ""
		}
	}

	return buf, nil
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

func customParseAndRemoveDuplicates(input io.Reader) (result []int , err error)  {
	in := bufio.NewScanner(input)
	var buf []int

	for in.Scan() {
		buf, err = customParseString(in.Text())

		if err !=nil {
			return nil, err
		}

		buf= removeDuplicates(buf)
	}
	return buf,nil
}

func nativeParseAndRemoveDuplicates(input io.Reader)(result []int, err error)  {
	in := bufio.NewScanner(input)
	var buf []int

	for in.Scan() {
		buf, err = nativeParseString(in.Text())

		if err !=nil {
			return nil, err
		}

		buf= removeDuplicates(buf)
	}
	return buf,nil
}

func main() {
	res,err := customParseAndRemoveDuplicates(os.Stdin)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(res)

	res,err = nativeParseAndRemoveDuplicates(os.Stdin)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(res)


}
