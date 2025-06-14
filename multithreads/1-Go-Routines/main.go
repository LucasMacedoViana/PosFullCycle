package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go task("A")
	go task("B")
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d: task C is running\n", i)
			time.Sleep(1 * time.Second)
		}
	}()
	time.Sleep(15 * time.Second)
}
