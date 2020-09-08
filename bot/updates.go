package main

import "fmt"

func (update *Update ) process ()   {
	fmt.Println("processing" + update.Message.Text)
}
