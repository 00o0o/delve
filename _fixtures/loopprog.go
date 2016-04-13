package main

import (
	"fmt"
	"time"
)

func loop() {
	i := 0
	for {
		i++
		if (i % 100000000) == 0 {
			fmt.Println(i)
		}
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println(i)
}

func main() {
	fmt.Println("past main")
	loop()
}
