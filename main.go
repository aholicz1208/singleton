package main

import (
	"fmt"
	"sync"
)

/*
	singleton pattern : create instance once time per executed
*/

type singleton struct {}

var (
	once     sync.Once
	instance *singleton
)

func newInstance() *singleton {
	// do one time when program executed.
	once.Do(func() {
		instance = &singleton{}
	})
	// return created instance.
	return instance
}

func main() {
	instance1  := newInstance()
	instance2  := newInstance()

	fmt.Print(instance1==instance2)
}