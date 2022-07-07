package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex
var balance int

func deposit(amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("Depositing: ", amount)
	balance += amount
	fmt.Println("Current Balance: ", balance)
	mutex.Unlock()
	wg.Done()
}

func withdraw(amount int, wg *sync.WaitGroup) {
	mutex.Lock()
	fmt.Println("Withdrawing: ", amount)
	if balance > amount {
		balance -= amount
		fmt.Println("Current Balance: ", balance)
	} else {
		fmt.Println("Not Sufficent Balance.")
	}
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	balance = 500
	fmt.Println("Current Balance: ", balance)
	go deposit(1000, &wg)
	go withdraw(500, &wg)

	wg.Wait()
	fmt.Println("Current Balance: ", balance)

}
