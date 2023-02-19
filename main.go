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
		go func() {
			defer wg.Done()
			// There is a problem with closed packages here
			fmt.Println(i, v)
		}()
	}

	wg.Wait()
}
