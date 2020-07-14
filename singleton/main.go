package main

import (
	"sync"
)

type single struct {
}

var singleInstance *single
var once sync.Once

func getInstance() *single {
	if singleInstance == nil {
		once.Do(func() { singleInstance = &single{} })
	}

	return singleInstance
}

func main() {
	for i := 0; i < 100; i++ {
		go getInstance()
	}
}
