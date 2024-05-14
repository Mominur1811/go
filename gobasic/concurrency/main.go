package main

import (
	"fmt"
	"gobasic/concurencyop"
	"time"
)

func main() {

	start := time.Now()

	defer func() {

		fmt.Println(time.Since(start))

	}()
	evilNinjas := []string{"Momin", "Anik", "Meem"}

	concurencyop.ChannelImplementation(evilNinjas) //check concurrencyop folder for these function
	fmt.Printf("\n")
	concurencyop.WaitGroupImplementation(evilNinjas) //check concurrencyop folder for these function
	fmt.Printf("\n")
	concurencyop.MutexImplementation() //check concurrencyop folder for these function
	fmt.Printf("\n")

}
