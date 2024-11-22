package main

import (
	"fmt"
	"sync"
	"time"
)

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Olá, mundo!")
		waitGroup.Done()
	}()

	go func() {
		escrever("Hello, World!")
		waitGroup.Done()
	}()

	go func() {
		escrever("こんにちは世界")
		waitGroup.Done()
	}()

	go func() {
		escrever("Γεια σου Κόσμο")
	}()

	waitGroup.Wait()

}
