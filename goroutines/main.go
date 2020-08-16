package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func f(n int, id int, chanel chan int) {
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(id)
	chanel <- id
}

func main() {

	if len(os.Args) < 2 {
		panic("Usage: <time in seconds>")
	}

	seconds, err := strconv.Atoi(os.Args[1])

	if err != nil {
		panic("Usage: <time in seconds>")
	}

	chanel := make(chan int)

	for i := 0; i < 10; i++ {
		go f(seconds, i, chanel)
	}


	timeOut := time.After(2 * time.Second)

	for i:=0 ; i< 10; i++ {
		select {
		 case id:= <-chanel:
			 fmt.Println("Goroutine: ", id, "ended")
		 case <-timeOut:
		 	fmt.Println("im tired, could u be faster, please???")
		}
	}


	/// deadlock example

	/*anotherChannel := make(chan int)
	<-anotherChannel*/



}
