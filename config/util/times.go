package util

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// Htime 自定义时间类型，继承自 time.Time
type Htime struct {
	time.Time
}

var (
	// formatTime 定义时间格式化字符串
	formatTime = "2006-01-02 15:04:05"
)

// MarshalJSON 方法用于将 Htime 类型序列化为 JSON 格式
// 输出格式为带引号的指定格式的时间字符串
func (t Htime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(formatTime))
	return []byte(formatted), nil
}

// UnmarshalJSON 方法用于从 JSON 数据中反序列化 Htime 类型
// 根据指定格式解析字符串并设置到 Htime 中
func (t *Htime) UnmarshalJSON(b []byte) (err error) {
	now, err := time.ParseInLocation(`"`+formatTime+`"`, string(b), time.Local)
	if err != nil {
		return err
	}
	*t = Htime{now}
	return
}

// Value 方法用于将 Htime 类型转换为数据库驱动值
func (t Htime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time.Format(formatTime), nil
}

// Scan 方法用于从数据库数据中扫描并设置到 Htime 类型
func (t *Htime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if !ok {
		return fmt.Errorf("can not convert %v to timestamp", v)
	}
	// 确保从数据库读取的时间只包含秒级精度
	value = value.Truncate(time.Second)
	*t = Htime{value}
	return nil
}
