package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

// for waitgroup refer to this website - https://www.geeksforgeeks.org/go-language/using-waitgroup-in-golang/

var signals = []string{"test"}
var wg sync.WaitGroup // pointer
var mut sync.Mutex // pointer


func main() {
	// go greeter("Satvik") // we are not waiting for the thread to come back and tell the result for this one way is to use time method and wait for specify time
	// greeter("srivastava")

	websiteList := []string{
		"https://google.com",
		"https://fb.com",
		"https://github.com",
	}

	for _, web := range websiteList {
		wg.Add(1)
		getStatusCode(web) // now if you use go keyword same things will happen it will not print anything buz we are not waiting for the response to come
	}
	wg.Wait() // this guy will not allow the main method to just finish bec his rest of the friends("wg.Add()") are not complete
	fmt.Println(signals)

}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(3 * time.Millisecond) // not the optimum way to handle go routine
// 		fmt.Println("Hello ", s)
// 	}
// }

func getStatusCode(endpoint string) {
	res, err := http.Get(endpoint)
	if err != nil {
		log.Fatal(err)
	} else {
		// we are locking the resources for writing
		mut.Lock()
		signals = append(signals, endpoint)
		// we are unlocking the resources so that other process can perform there task
		mut.Unlock() // this is done to prevent deadlock suitation or incorrect read write opertion on wrong data
		fmt.Printf("%d is the status code for %s\n", res.StatusCode, endpoint)
		defer wg.Done() // signal that says i am done whatever i was doing 
	}
	
}


// NOTE : mutex - mutual exclusive lock