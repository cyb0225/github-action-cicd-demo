/**
@author: yeebing
@date: 2023/2/18
**/

package main

import "fmt"

var (
	// unused-var
	A = 10
)

func Test() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range arr {
		go func() {
			// There is a problem with closed packages here
			fmt.Println(i, v)
		}()
	}
}
