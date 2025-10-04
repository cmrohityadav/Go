package main

import (
	"fmt"
	"sync"
	"time"
)

type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()
		p.mu.Unlock()

	}()

	p.mu.Lock()
	p.views += 1
	// p.mu.Unlock()
}

func main() {

	myPost := post{views: 0}
	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}

	wg.Wait()

	fmt.Println(myPost.views)

	//--------Mutex-----------

	var coruptCounter int = 0

	for range 5000 {
		go func() {
			coruptCounter++
		}()
		go func() {
			coruptCounter++
		}()
	}

	time.Sleep(2 * time.Second)

	fmt.Println("corupt Counter:", coruptCounter)

	mut := &sync.Mutex{}

	var safeCounter int = 0

	for i := 0; i < 5000; i++ {
		go func() {
			mut.Lock()
			safeCounter++
			mut.Unlock()
		}()
		go func() {
			mut.Lock()
			safeCounter++
			mut.Unlock()
		}()
	}

	time.Sleep(2 * time.Second)
	fmt.Println("safe Counter:", safeCounter)


	// -------- sync.RWMutex --------

	var rw sync.RWMutex
	var rwSafeCounter int = 0

	// -------- Writers --------
	for i := 0; i < 50; i++ { // 10 writers for demo clarity
		wg.Add(1)
		go func(id int) {
			defer wg.Done();

			fmt.Printf("Writer %d trying to write...\n",id);
			fmt.Println("Waiting until all readers finish");

			rw.Lock();

			fmt.Printf("Writer %d Started to write...\n",id);
			
			rwSafeCounter++;

			fmt.Printf("Writer %d incremented counter from %d to %d\n", id, rwSafeCounter-1,rwSafeCounter);

			rw.Unlock();

		}(i)
	}

	// -------- Readers --------
	for i := 0; i < 3; i++ { // 3 readers
		go func(id int) {
			for  { 
				rw.RLock()
				fmt.Printf("Reader %d started reading, sees: %d\n", id, rwSafeCounter)
				time.Sleep(500 * time.Millisecond)
				fmt.Printf("Reader %d finished reading\n", id)
				rw.RUnlock()
				time.Sleep(200 * time.Millisecond)
			}
		}(i)
	}

	// Wait for all writers to finish
	wg.Wait()

	// Give readers some time to finish
	time.Sleep(20 * time.Second)

	fmt.Println("Final Safe Counter (RWMutex):", rwSafeCounter);

	
}
