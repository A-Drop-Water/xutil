package xutil

import (
	"fmt"
	"math/rand"
)

// GenerateRandomDigitCode 生成随机width位的数字
func GenerateRandomDigitCode(width int) string {
	if width <= 0 {
		return ""
	}
	maxN := 1
	for i := 0; i < width; i++ {
		maxN *= 10
	}
	// 生成随机数
	randomNum := rand.Intn(maxN)
	// 格式化为指定位数，不足位数前面补0
	format := fmt.Sprintf("%%0%dd", width)
	return fmt.Sprintf(format, randomNum)
}
