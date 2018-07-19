package main

import (
	"fmt"
)

func main() {

	fmt.Println(ReverseInt(8085774586302733229))
}

func ReverseInt(num int) int {

	sign := 1
	if num < 0 {
		sign = -1
		num = -1 * num
	}

	rv := 0
	for num != 0 {
		rv = rv*10 + num%10
		num /= 10
	}
	rv *= sign
	if sign > 0 && rv < 0 || sign < 0 && rv > 0 {
		panic("超过表示范围")
	}

	return rv

}
