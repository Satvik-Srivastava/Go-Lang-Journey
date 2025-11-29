package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Welcome to run condition topic")

	var score = []int{0}

	mut := &sync.Mutex{}

	// we are going to use iife
	wg := &sync.WaitGroup{}

	// adding the waitgroup

	// wg.Add(1)
	wg.Add(3)
	// the first func doesn't know that there are other func like two 3 and three 3 in the market for this we will use 
	// channel so that these function can talk to each other
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("One r")
		mut.Lock() // locking the resource so that no other go routine process can access the file to read or write 
		score = append(score, 1)
		mut.Unlock() // after successfull writing in the resource unlocking it so that other process can use them	
		wg.Done()
	}(wg, mut)
	// wg.Add(1)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Two r")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	go func(wg *sync.WaitGroup, m *sync.Mutex) {
		fmt.Println("Three r")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	wg.Wait()
	fmt.Println(score)
}
