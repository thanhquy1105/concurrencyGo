package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

// func updateMessage(s string) {
// 	defer wg.Done()

// 	msg = s
// }

// func Run() {
// 	msg = "Hello, world!"

// 	wg.Add(2)
// 	go updateMessage("Hello, universe!")
// 	go updateMessage("Hello, cosmo!")

// 	wg.Wait()

// 	fmt.Println(msg)
// }

func updateMessage(s string, m *sync.Mutex) {
	defer wg.Done()

	m.Lock()
	msg = s
	m.Unlock()
}

func Run() {
	msg = "Hello, world!"

	var mutex sync.Mutex

	wg.Add(2)
	go updateMessage("Hello, universe!", &mutex)
	go updateMessage("Hello, cosmo!", &mutex)

	wg.Wait()

	fmt.Println(msg)
}
