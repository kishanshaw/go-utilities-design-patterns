package Impl

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type Singleton struct {
}

var singleton *Singleton

func GetInstance() *Singleton {
	if singleton == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleton == nil {
			fmt.Println("Creating single instance now")
			singleton = &Singleton{}
		} else {
			fmt.Println("Singleton instance already exists")
		}
	} else {
		fmt.Println("Singleton instance already exists")
	}
	return singleton
}
