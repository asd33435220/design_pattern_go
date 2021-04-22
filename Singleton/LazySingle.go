package main

import (
	"fmt"
	"sync"
)

var lock = sync.Mutex{}

type single struct {
}

var LazySingleInstance *single

func getLazyInstance() *single {
	if LazySingleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if LazySingleInstance == nil {
			fmt.Println("Creating LazySingleInstance instance now")
			LazySingleInstance = &single{}
		} else {
			fmt.Println("LazySingleInstance already created.")
		}
	} else {
		fmt.Println("LazySingleInstance already created.")
	}
	return LazySingleInstance
}
