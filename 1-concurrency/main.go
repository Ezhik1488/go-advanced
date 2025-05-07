package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	countNum := 10
	numsCh := make(chan int, countNum)
	resultCh := make(chan int, countNum)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	go GenerateNum(numsCh, countNum, wg)
	go PowToTwo(numsCh, resultCh, wg)

	wg.Wait()

	for sqNum := range resultCh {
		fmt.Printf("%d ", sqNum)
	}
}

func GenerateNum(numChan chan int, countNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < countNum; i++ {
		numChan <- rand.Intn(100)
	}
	close(numChan)
}

func PowToTwo(numChan chan int, resChan chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for value := range numChan {
		resChan <- value * value
	}
	close(resChan)
}
