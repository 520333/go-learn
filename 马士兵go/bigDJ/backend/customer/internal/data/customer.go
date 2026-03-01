package data

import (
	"context"
	"customer/internal/biz"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

type CustomerData struct {
	data *Data
}

func NewCustomerData(data *Data) *CustomerData {
	return &CustomerData{data: data}
}

// SetVerifyCode 设置验证码方法
func (cd CustomerData) SetVerifyCode(telephone, code string, ex int64) error {
	status := cd.data.Rdb.Set(context.Background(), "CVC:"+telephone, code, time.Duration(ex)*time.Second)
	if _, err := status.Result(); err != nil {
		return err
	}
	return nil
}

// GetVerifyCode 获取号码对应的验证码
func (cd CustomerData) GetVerifyCode(telephone string) string {
	status, _ := cd.data.Rdb.Get(context.Background(), "CVC:"+telephone).Result()
	return status
}

// GetCustomerByTelephone 根据电话获取顾客信息
func (cd CustomerData) GetCustomerByTelephone(telephone string) (*biz.Customer, error) {
	// 查询基于电话
	customer := &biz.Customer{}
	result := cd.data.Mdb.Where("telephone = ?", telephone).First(customer)
	// query 执行成功，同时查询到记录
	if result.Error == nil && customer.ID > 0 {
		return customer, nil
	}
	// 有记录不存在的错误
	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 创建customer并返回
		customer.Telephone = telephone
		customer.Name = sql.NullString{Valid: false}
		customer.Email = sql.NullString{Valid: false}
		customer.Wechat = sql.NullString{Valid: false}
		resultCreate := cd.data.Mdb.Create(customer)
		if resultCreate.Error != nil {
			return nil, resultCreate.Error
		}
	}
	// 有错误 但不是记录不存在的错误，不做业务逻辑处理
	return customer, nil
}

// GenerateTokenAndSave 生成Token和存储
func (cd CustomerData) GenerateTokenAndSave(c *biz.Customer, duration time.Duration, secret string) (string, error) {
	// 一 生成token
	// 处理token中的数据
	claims := jwt.RegisteredClaims{
		Issuer:    "BigDJ", //签发机构
		Subject:   "customer-authentication",
		Audience:  []string{"customer", "other"},                // 签发给谁
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)), // 有效期
		NotBefore: nil,                                          // 何时启用
		IssuedAt:  jwt.NewNumericDate(time.Now()),               // 签发时间
		ID:        fmt.Sprintf("%d", c.ID),                      //用户id
	}
	// 生成token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 签发token
	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	// 二 存储
	c.Token = signedToken
	c.TokenCreatedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	if result := cd.data.Mdb.Save(c); result.Error != nil {
		return "", result.Error
	}
	// 操作成功
	return signedToken, nil
}
