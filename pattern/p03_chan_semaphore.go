package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var sem = make(chan struct{}, 3)
var job = make(chan int, 10)

func main() {
	go func() {
		for i := 0; i < 8; i++ {
			job <- (i + 1)
		}
		close(job)
	}()

	var wg sync.WaitGroup

	for j := range job {
		wg.Add(1)
		go func(v int) {
			sem <- struct{}{}
			log.Println("worker", v) // 参数j，不能用闭包的方式使用
			time.Sleep(2 * time.Second)
			<-sem
			wg.Done()
		}(j) // 这里必须用参数的方式使用j
	}

	wg.Wait()
	fmt.Println("finished")
}
