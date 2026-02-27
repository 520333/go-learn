package data

import (
	"context"
	"time"
)

type CustomerData struct {
	data *Data
}

func NewCustomerData(data *Data) *CustomerData {
	return &CustomerData{data: data}
}

// SetVerifyCode 设置验证码方法
func (cd CustomerData) SetVerifyCode(telephone, code string, ex int64) error {
	const life = 60
	status := cd.data.Rdb.Set(context.Background(), "CVC:"+telephone, code, time.Duration(ex)*time.Second)
	if _, err := status.Result(); err != nil {
		return err
	}
	return nil
}
