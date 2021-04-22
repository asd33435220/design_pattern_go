package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var HungrySingleInstance *single

func init(){
	once.Do(func() {
		HungrySingleInstance = &single{}
		fmt.Println("Creating HungrySingleInstance instance init")
	})
}
func getHungryInstance() *single {
	if HungrySingleInstance == nil {
		HungrySingleInstance = &single{}
		fmt.Println("Creating HungrySingleInstance instance now")
	} else {
		fmt.Println("HungrySingleInstance already created.")
	}
	return HungrySingleInstance
}
