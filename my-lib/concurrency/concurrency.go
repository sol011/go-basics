package concurrency

import (
	"fmt"
	"sync"
)

func pushToChan(s []int, c chan int, idx int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("running the goroutine for calculating sum from index", idx)
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
	c <- 100
	c <- 1
	c <- 500
	// close(c)
}

func DeadlockReadingFromForeverEmptyChan() {
	var c = make(chan int, 4)
	var x = <-c
	fmt.Println(x)
}

func SimpleWait() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int, 20)
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go pushToChan(s[:len(s)/2], c, 0, &wg)
	wg.Add(1)
	go pushToChan(s[len(s)/2:], c, len(s)/2, &wg)

	wg.Wait()
	close(c)
	for e := range c {
		fmt.Println(e)
	}
}

func UnbufferedChansBlockSendUntilReceive() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	var wg = sync.WaitGroup{}
	wg.Add(1)
	go pushToChan(s[:len(s)/2], c, 0, &wg)
	wg.Add(1)
	go pushToChan(s[len(s)/2:], c, len(s)/2, &wg)

	wg.Wait()
	close(c)
	for e := range c {
		fmt.Println(e)
	}
}
