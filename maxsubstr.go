// main
package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := []byte("abcbbbbcdbef")

	fmt.Println(string(maxsubstr(str)))

	_, maxstr := hmsubstr(str)
	fmt.Println(string(maxstr))
}

// 查找
func maxsubstr(str []byte) (rv []byte) {
	dict := make(map[byte]int)
	start := 0
	for i, v := range str {
		if idx, ok := dict[v]; ok {
			if idx+1 > start {
				if len(rv) < i-start {
					rv = str[start:i]
				}
				start = idx + 1
			}

		}
		dict[v] = i
	}
	if len(rv) < len(str)-start {
		rv = str[start:]
	}
	return
}

// 返回从头开始的不重复的最长子串和整体的最长子串
func hmsubstr(str []byte) ([]byte, []byte) {
	n := len(str)
	if n == 1 {
		return str, str
	}
	if n == 2 {
		if str[0] == str[1] {
			return str[0:1], str[0:1]
		}
		return str, str
	}

	head, max := hmsubstr(str[1:])

	idx := bytes.IndexByte(head, str[0])
	// 重置head
	if -1 == idx {
		idx = len(head)
	}
	head = str[:idx+1]

	if len(head) >= len(max) {
		max = head
	}

	return head, max
}
