package main

import "fmt"

func main() {

	m := []int{1}
	n := []int{2}
	fmt.Println(mergeSortedSlice(m, m))
	fmt.Println(mergeSortedSlice(m, n))
	fmt.Println(mergeSortedSlice(n, m))

	m = []int{1, 3, 5, 7}
	n = []int{10, 12, 15, 17}
	fmt.Println(mergeSortedSlice(m, m))
	fmt.Println(mergeSortedSlice(m, n))
	fmt.Println(mergeSortedSlice(n, m))

	m = []int{1, 3, 3, 7, 9, 11, 13, 15}
	n = []int{2, 3, 3, 3, 3, 12, 15, 17}
	fmt.Println(mergeSortedSlice(m, m))
	fmt.Println(mergeSortedSlice(m, n))
	fmt.Println(mergeSortedSlice(n, m))
}

func mergeSortedSlice(m []int, n []int) []int {

	mlen, nlen := len(m), len(n)
	r := make([]int, 0, mlen+nlen)
	if m[mlen-1] <= n[0] {
		r = append(r, m...)
		r = append(r, n...)
		return r
	}
	if n[nlen-1] <= m[0] {
		r = append(r, n...)
		r = append(r, m...)
		return r
	}

	for {
		down, up := binarySearchSlice(m, n[0])
		r = append(r, down...)
		if up == nil {
			r = append(r, n...)
			return r
		}
		m, n = n, up
	}
}

//使用二分法从有序数组中返回小于等于这个数字的slice
func binarySearchSlice(base []int, x int) ([]int, []int) {
	down, up := 0, len(base)
	if x < base[0] {
		return nil, base
	} else if x >= base[up-1] {
		return base, nil
	}

	for {
		bidx := (down + up) / 2
		if x >= base[bidx] {
			down = bidx
		} else {
			up = bidx
		}
		if down+1 == up {
			return base[:up], base[up:]
		}
	}
}
