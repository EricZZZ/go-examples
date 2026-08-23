package main

import (
	"fmt"
	"runtime"
	"time"
)

func firstWrong() {
	ch := make(chan int) // unbuffered: every send needs a corresponding receive
	for i := 1; i <= 3; i++ {
		go func() {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf(" worker %d: sending...\n", i)
			ch <- i
			fmt.Printf(" worker %d: done\n", i)
		}()
	}
	fmt.Printf(" main got worker %d, returning\n", <-ch)
}

func firstRight() {
	ch := make(chan int, 3) // buffered: send can proceed without a corresponding receive
	for i := 1; i <= 3; i++ {
		go func() {
			time.Sleep(50 * time.Millisecond)
			fmt.Printf(" worker %d: sending...\n", i)
			ch <- i
			fmt.Printf(" worker %d: done\n", i)
		}()
	}
	<-ch
	fmt.Println(" main got worker, returning")
}

func main() {
	fmt.Println("WRONG (unbuffered channel):")
	before := runtime.NumGoroutine()
	firstWrong()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf(" goroutines: %d -> %d (2 stuck forever)\n\n", before, runtime.NumGoroutine())

	fmt.Println("RIGHT (buffered channel):")
	before = runtime.NumGoroutine()
	firstRight()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf(" goroutines: %d -> %d (no stuck forever)", before, runtime.NumGoroutine())
}
