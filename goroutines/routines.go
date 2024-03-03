package main

import (
	"fmt"
	"net/http"
	"sync"
)

var signals =[]string{"test"}
var wg sync.WaitGroup //these are pointers
var mut sync.Mutex  //pointers

func main() {
	// go greet("viraj")
	// greet("message")
	website := []string{
		"https://google.com",
		"https://lco.dev",
		"https://fb.com",
		"https://go.dev",
	}

	for _, web := range website{
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
}

// func greet(s string){
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(5 * time.Microsecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Printf("Unable to render")
	}else{
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}
}
