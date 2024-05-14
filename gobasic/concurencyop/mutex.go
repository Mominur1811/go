package concurencyop

import (
	"fmt"
	"sync"
	"time"
)

type mutexImplementation struct {
	mu sync.Mutex
	v  map[string]int
}

func (mutexVar *mutexImplementation) IncrementValue(str string, beeper *sync.WaitGroup) {

	mutexVar.mu.Lock()
	mutexVar.v[str]++
	mutexVar.mu.Unlock()
	beeper.Done()
}

func MutexImplementation() {

	start := time.Now()
	defer func() {

		fmt.Println("Time consume for mutex : ", time.Since(start))

	}()
	fmt.Println("Mutex calling -- -- --")
	mutexVar := mutexImplementation{v: make(map[string]int)}

	var beeper2 sync.WaitGroup
	beeper2.Add(100)

	for i := 0; i < 100; i++ {

		go mutexVar.IncrementValue("MUTEX", &beeper2)
	}
	beeper2.Wait()
	fmt.Println("Value of key MUTEX is: ", mutexVar.v["MUTEX"])
	fmt.Println("Mutex Ended -- -- --")

}
