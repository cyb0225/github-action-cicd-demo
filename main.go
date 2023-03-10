package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	var wg sync.WaitGroup
	wg.Add(len(arr))

	for i, v := range arr {
		index := i
		val := v
		go func() {
			defer wg.Done()
			// There is a problem with closed packages here
			fmt.Println(index, val)
		}()
	}

	_ = FuncA() // do not deal with the error
	fmt.Println("123")
	wg.Wait()
}

func FuncA() error {
	return nil
}
