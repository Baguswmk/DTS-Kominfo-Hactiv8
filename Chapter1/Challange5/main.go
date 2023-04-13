package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex = &sync.Mutex{}
var wg = &sync.WaitGroup{}

func processAcak1(data interface{}, no int) {
	fmt.Println(data, " ", no)
}

func processAcak2(data interface{}, no int) {
	fmt.Println(data, " ", no)
}

func cetakAcak(data1 interface{}, data2 interface{}) {

	for i := 1; i <= 4; i++ {
		go processAcak1(data1, i)
		go processAcak2(data2, i)
	}
}

func processRapi1(data interface{}, no int) {
	defer wg.Done()
	mutex.Lock()
	fmt.Println(data, " ", no)
	mutex.Unlock()
}

func processRapi2(data interface{}, no int) {
	defer wg.Done()
	mutex.Lock()
	fmt.Println(data, " ", no)
	mutex.Unlock()
}

func cetakRapi(data1 interface{}, data2 interface{}) {

	for i := 1; i <= 4; i++ {
		wg.Add(2)
		go processRapi1(data1, i)
		go processRapi2(data2, i)
		wg.Wait()
	}

}

func main() {

	data1 := []interface{}{"bisa1", "bisa2", "bisa3"}
	data2 := []interface{}{"coba1", "coba2", "coba3"}

	fmt.Println("======================= ACAK ===========================")
	cetakAcak(data1, data2)
	time.Sleep(3 * time.Second)
	fmt.Println("======================= RAPI ===========================")
	cetakRapi(data1, data2)
	time.Sleep(3 * time.Second)
}