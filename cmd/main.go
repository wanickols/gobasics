package main

//Concurrency Example
import (
	"fmt"
	"sync"
	"time"
)

var mtx = sync.Mutex{}
var wg = sync.WaitGroup{}
var dbData = []string{"id1", "id2", "id3", "id4", "id5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)
		go dbCall(i)
	}

	wg.Wait()
	fmt.Printf("\nTotal excecution time: %v", time.Since(t0))
	fmt.Printf("\nThe results are %v", results)
}

func dbCall(i int) {
	var delay float32 = 2000
	time.Sleep((time.Duration(delay) * time.Millisecond))
	fmt.Println("The result from thie database is:", dbData[i])
	mtx.Lock()
	results = append(results, dbData[i])
	mtx.Unlock()
	wg.Done()
}
