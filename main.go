package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	for _, v := range arr {
		go func() {
			// There is a problem with closed packages here
			fmt.Println(v)
		}()
	}

}