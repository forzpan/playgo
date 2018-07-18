package main

import (
	"fmt"
	"math"
)

func main() {

	str := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	fmt.Println(take(str, 5))
}

func take(str string, rows int) string {
	tmp := make([][]byte, rows)
	var rv []byte
	for i, length := 0, len(str); i < length; i++ {
		//决定第几行
		idx := rows - 1 - int(math.Abs(float64(rows-1-i%(2*rows-2))))
		tmp[idx] = append(tmp[idx], str[i])
	}
	for n := 0; n < rows; n++ {
		rv = append(rv, tmp[n]...)
	}
	return string(rv)
}
