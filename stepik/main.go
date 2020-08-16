package main

import (
	"fmt"
	"os"
	"strconv"
)

func calcDigitalRoot(val string) int {
	if len(val) == 2 {
		return int((val[0] - '0') + (val[1] - '0'))
	}

	var nextVal string

	for _, ch := range val {
		tmpVal, _ := strconv.Atoi(nextVal)
		nextVal = strconv.Itoa(int(ch-'0') + tmpVal)
	}

	return calcDigitalRoot(nextVal)
}

func main() {
	var input string

	fmt.Fscan(os.Stdin, &input)

	fmt.Println(calcDigitalRoot(input))
}
