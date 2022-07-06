package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func worker(id int) {
	rand.Seed(time.Now().UnixNano())
	a := rand.Intn(2)

	time.Sleep(time.Duration(a) * time.Second)

}

func main() {

	rand.Seed(time.Now().UnixNano())
	var sum int
	var wg sync.WaitGroup
	wg.Add(200)

	for i := 1; i <= 200; i++ {
		go func() {
			defer wg.Done()
			worker(i)
			sum += rand.Intn(10)
		}()
	}

	wg.Wait()
	fmt.Println("Total Rating of 200 students: ", sum)
	fmt.Println("Average Rating of 200 students: ", sum/200)
	
}
