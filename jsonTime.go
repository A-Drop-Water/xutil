package xutil

import (
	"fmt"
	"time"
)

// JsonTime json序列化时会变成年-月-日的形式
type JsonTime time.Time

// MarshalJSON 自己实现marshal方法
func (t JsonTime) MarshalJSON() ([]byte, error) {
	// 显示年月日
	s := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02"))
	return []byte(s), nil
}
