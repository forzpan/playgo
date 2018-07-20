package main

import (
	"fmt"
)

func main() {
	fmt.Println(mostWater([]uint{1, 2, 3, 4, 5, 6, 7, 8, 1}))
}

func mostWater(height []uint) (s, e uint) {

	if len(height) < 2 {
		panic("参数太短")
	}

	// 从两端开始寻找
	var max, i, j uint = 0, 0, uint(len(height)) - 1

	for {
		h, b := height[i], true
		if height[i] > height[j] {
			h, b = height[j], false
		}

		area := h * (j - i)
		if max < area {
			max, s, e = area, i, j
		}

		// 尝试寻找高一些的短板
		if b {
			for {
				i++
				if i == j {
					return
				}
				if height[i] > h {
					break
				}
			}
		} else {
			for {
				j--
				if i == j {
					return
				}
				if height[j] > h {
					break
				}
			}
		}
	}
}
