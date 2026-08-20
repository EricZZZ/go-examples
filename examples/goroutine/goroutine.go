package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// test_1()
	// test_2()
	// test_3()
	test_4()
}

/**
 * 测试协程
 */
func test_1() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello")
	}
	wg.Add(1)
	go sayHello()
	// sleep 0.1s 这是一种错觉，并没等待协程执行完
	// time.Sleep(100 * time.Millisecond)

	// 等待协程执行完
	wg.Wait()

	println("goroutine")
}

func test_2() {
	var wg sync.WaitGroup
	salutation := "hello"

	wg.Go(func() {
		salutation = "welcome"
	})
	wg.Wait()
	fmt.Println(salutation)
}

/*
 * 在1.22版本前打印，闭包捕获的不是当时 salutation 的“值”，而是对变量 salutation 的引用。
 * good day
 * good day
 * good day
 *
 * 1.22版本后bug修复
 */
func test_3() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Go(func() {
			fmt.Println(salutation)
		})
	}
	wg.Wait()
}

/*
 * 测量协程大小
 * 执行显示 2.566kb
 */
func test_4() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() {
		wg.Done()
		<-c
	}
	numGoroutines := 10000
	before := memConsumed()
	wg.Add(numGoroutines)
	for range numGoroutines {
		go noop()
	}
	wg.Wait()
	after := memConsumed()
	fmt.Printf("%.3fkb", float64(after-before)/float64(numGoroutines)/1024)
}
