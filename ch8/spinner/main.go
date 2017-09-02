// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 218.

// Spinner displays an animation while computing the 45th Fibonacci number.
package main

import (
	"fmt"
	"time"
)

var nx int

//!+
func main() {
	go wrap()(100 * time.Millisecond)
        fmt.Println(nx)

	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}

// Hzh测试：
// go语句的部分会立即执行
// 但会提取最终的函数调用后才开启一个新的goroutine并返回。
// 这与defer的语法逻辑一样。
func wrap() func (time.Duration) {
    nx += 99
    time.Sleep(2 * time.Second)
    return spinner
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)  // 或"\b%c"
			time.Sleep(delay)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

//!-
