package main

import (
	"code.test.com/go/rate"
	"fmt"
	"sync"
	"time"
)

func main() {
	//testLimitRate()
	//testLeakyRate()
	testTokenRate()
}

func testLimitRate() {
	var wg sync.WaitGroup
	var lr rate.LimitRate
	lr.Set(3, time.Second) // 1s内最多请求3次

	for i := 0; i < 100; i++ {
		wg.Add(1)

		fmt.Println("Request", i, time.Now())
		go func(i int) {
			if lr.Allow() {
				fmt.Println("Response", i, time.Now())
			}
			wg.Done()
		}(i)

		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}

// 漏桶
func testLeakyRate() {
	var wg sync.WaitGroup
	var lr rate.LeakyBucket
	lr.Set(3, 3) //每秒访问速率限制为3个请求，桶容量为3

	for i := 0; i < 100; i++ {
		wg.Add(1)

		fmt.Println("Request", i, time.Now())
		go func(i int) {
			if lr.Allow() {
				fmt.Println("Response req", i, time.Now())
			}
			wg.Done()
		}(i)

		time.Sleep(100 * time.Millisecond)
	}
	wg.Wait()
}

// 令牌桶
func testTokenRate() {
	var wg sync.WaitGroup
	var lr rate.TokenBucket
	lr.Set(1, 3, 3) //每秒访问速率限制为3个请求，桶容量为3

	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		fmt.Println("Request", i, time.Now())
		go func(i int) {
			if lr.Allow() {
				fmt.Println("Response", i, time.Now())
			}
			wg.Done()
		}(i)

		time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()
}
