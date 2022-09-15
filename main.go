package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	t := time.Now()
	wg = sync.WaitGroup{}
	for i := 0; i < 100_000; i++ {
		wg.Add(1)
		go test()
	}
	wg.Wait()
	fmt.Println(time.Since(t).Milliseconds())
	time.Sleep(1000 * time.Second)
}

func test() {
	key := "testtestesstestsetsetsetsettseet"
	encrypted := Encrypt(key, whatever{A: "hmm", B: "test", C: 123})
	decrypted := Decrypt[whatever](key, encrypted)
	_ = decrypted.C
	time.Sleep(10 * time.Second)
	wg.Done()
}

type whatever struct {
	A string
	B string
	C int
}
