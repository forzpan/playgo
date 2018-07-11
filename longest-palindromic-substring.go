package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(string(maxPalindrome([]byte(""))))
	fmt.Println(string(maxPalindrome([]byte("1"))))
	fmt.Println(string(maxPalindrome([]byte("121"))))
	fmt.Println(string(maxPalindrome([]byte("123321"))))
	fmt.Println(string(maxPalindrome([]byte("1234321"))))
	fmt.Println(string(maxPalindrome([]byte("1231232111"))))
	fmt.Println(string(maxPalindrome([]byte("12321123432"))))
}

// Manacher 算法查找
func maxPalindrome(str []byte) []byte {
	//加符号，消除奇偶
	newstr := make([]byte, 2*len(str)+1)
	newstr[0] = '#'
	for i, v := range str {
		newstr[2*i+1] = v
		newstr[2*i+2] = '#'
	}
	//i位置的最大回文范围r[i]
	r := make([]int, len(newstr))
	//回文最右到达，该到达的对称轴位置，当前最大范围
	right, mid, maxRange, maxMid := 0, 0, 0, 0
	for i, N := 0, len(newstr); i < N; i++ {
		if i < right {
			r[i] = r[2*mid-i]
			if right-i < r[i] {
				r[i] = right - i
			}
		} else {
			r[i] = 0
		}
		for i-r[i]-1 >= 0 && i+r[i]+1 < N && newstr[i-r[i]-1] == newstr[i+r[i]+1] {
			r[i]++
		}
		if i+r[i] > right {
			mid, right = i, i+r[i]
		}
		if maxRange < r[i] {
			maxRange = r[i]
			maxMid = i
		}
	}
	return bytes.Replace(newstr[maxMid-maxRange:maxMid+maxRange+1], []byte{'#'}, nil, -1)
}

//动态规划 算法
func longestPalindrome(str []byte) []byte {
	n := len(str)
	pal := make([][]bool, n)
	for i := 0; i < n; i++ {
		pal[i] = make([]bool, n)
	}

	//pal[i][j] 表示s[i...j]是否是回文串
	left, right, maxLen := 0, 0, 0
	// i作为终点
	for i := 0; i < n; i++ {
		//j作为起点
		for j := i; j >= 0; j-- {
			if str[j] == str[i] && (i-j < 2 || pal[j+1][i-1]) {
				pal[j][i] = true
				if maxLen < i-j+1 {
					maxLen = i - j + 1
					left, right = j, i+1
				}
			}
		}
	}
	return str[left:right]
}
