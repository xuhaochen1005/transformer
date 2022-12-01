// Package model // 支持时间戳
package model

import (
	"database/sql/driver"
	"strconv"
	"time"

	"transformer_mz/libs/errors"
)

type Timestamp struct {
	time.Time
}

func (t Timestamp) String() string {
	return t.Format(time.RFC3339)
}

func (t Timestamp) Value() (driver.Value, error) {
	if t.Equal(time.Time{}) {
		return nil, nil
	}
	return t.Unix(), nil
}

func (t *Timestamp) Scan(value interface{}) error {
	switch value.(type) {
	case int64:
		t.Time = time.Unix(value.(int64), 0)
	case []uint8:
		val, err := strconv.ParseInt(string(value.([]uint8)), 10, 64)
		if err != nil {
			return errors.New(err, "内部服务错误，请联系管理员进行处理")
		}
		t.Time = time.Unix(val, 0)
	default:
		return errors.New("无法将数值转化为timestamp", "内部服务错误，请联系管理员进行处理")
	}
	return nil
}
