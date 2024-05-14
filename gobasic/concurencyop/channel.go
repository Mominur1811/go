package concurencyop

import (
	"fmt"
	"time"
)

// Channel Implementation Code
func ChannelImplementation(ninjas []string) {

	start := time.Now()

	defer func() {

		fmt.Println("Time consume at channel: ", time.Since(start))

	}()
	fmt.Println("Channel Calling Started -- -- --")
	goChannel := make(chan string)

	for i := 0; i < len(ninjas); i++ {

		go NinjaAttackChannel(ninjas[i], goChannel)

	}

	for i := 0; i < len(ninjas); i++ {

		fmt.Println(<-goChannel)
	}

	fmt.Println("Channel Calling Completed")

}

func NinjaAttackChannel(ninja string, goChannel chan string) {

	time.Sleep(time.Second)
	goChannel <- fmt.Sprint("We killed through Channel: ", ninja)

}
