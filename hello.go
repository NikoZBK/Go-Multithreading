package main

import (
	"fmt"
	"time"
)

func someTask(id int, data chan int) {
	for taskId := range data {
		time.Sleep(2 * time.Second)
		fmt.Printf("Task %d executed\n", data)
	}
}

func main() {

	for i := range 10 {
		someTask(i)
	}

}
