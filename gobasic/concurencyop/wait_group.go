package concurencyop

import (
	"fmt"
	"sync"
	"time"
)

// Wait group Implementation Code
func WaitGroupImplementation(evilNinjas []string) {

	start := time.Now()

	defer func() {

		fmt.Println("Time consume at wait group: ", time.Since(start))

	}()
	fmt.Println("Wait Group Calling Started -- -- --")
	var beeper1 sync.WaitGroup
	beeper1.Add(len(evilNinjas))

	for _, evilNinja := range evilNinjas {

		go NinjaAttackWaitGroup(evilNinja, &beeper1)
	}
	beeper1.Wait()
	fmt.Println("Wait Group Completed")

}

func NinjaAttackWaitGroup(evilNinja string, beeper *sync.WaitGroup) {

	time.Sleep(time.Second)
	fmt.Println("We killed throung WaitGroup: ", evilNinja)
	beeper.Done()
}
