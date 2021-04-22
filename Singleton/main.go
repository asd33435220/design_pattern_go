package main

import "time"

func main() {

	for i := 0; i < 10; i++ {
		go getHungryInstance()
	}
	time.Sleep(time.Second*3)
}
