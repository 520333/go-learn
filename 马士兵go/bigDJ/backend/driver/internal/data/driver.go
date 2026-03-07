package data

import (
	"context"
	"database/sql"
	"driver/api/verityCode"
	"driver/internal/biz"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/hashicorp/consul/api"
)

type DriverData struct {
	data *Data
}

func NewDriverInterface(data *Data) biz.DriverInterface {
	return &DriverData{data: data}
}

// GetToken 获取token实现
func (dt *DriverData) GetToken(ctx context.Context, tel string) (string, error) {
	driver := biz.Driver{}
	if err := dt.data.Mdb.Where("telephone = ?", tel).First(&driver).Error; err != nil {
		return "", err
	}

	return driver.Token.String, nil
}

// GetSaveVerifyCode 获取verifyCode服务中的村粗的验证码
func (dt *DriverData) GetSaveVerifyCode(ctx context.Context, tel string) (string, error) {
	return dt.data.Rdb.Get(ctx, "DVC:"+tel).Result()
}

func (dt *DriverData) SaveToken(ctx context.Context, tel, token string) error {
	// 先获取司机信息
	driver := biz.Driver{}
	if err := dt.data.Mdb.Where("telephone = ?", tel).First(&driver).Error; err != nil {
		return err
	}
	// 再更新
	driver.Token = sql.NullString{String: token, Valid: true}
	if err := dt.data.Mdb.Save(driver).Error; err != nil {
		return err
	}
	return nil
}

func (dt *DriverData) InitDriverInfo(ctx context.Context, tel string) (*biz.Driver, error) {
	driver := &biz.Driver{}
	driver.Telephone = tel
	driver.Status = sql.NullString{String: "stop", Valid: true}
	if err := dt.data.Mdb.Create(&driver).Error; err != nil {
		return nil, err
	}
	return &biz.Driver{}, nil
}

func (dt *DriverData) CheckVerifyCode(ctx context.Context, telephone string) (*biz.Driver, error) {
	return nil, nil
}

func (dt *DriverData) GetVerifyCode(ctx context.Context, tel string) (string, error) {
	consulConfig := api.DefaultConfig()
	consulConfig.Address = dt.data.cs.Consul.Address
	consulClient, err := api.NewClient(consulConfig)
	dis := consul.New(consulClient)
	if err != nil {
		return "", err
	}
	endponit := "discovery:///VerifyCode"
	conn, err := grpc.DialInsecure(
		ctx,
		grpc.WithEndpoint(endponit),
		grpc.WithDiscovery(dis),
	)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = conn.Close()
	}()
	client := verityCode.NewVerityCodeClient(conn)
	reply, err := client.GetVerityCode(ctx, &verityCode.GetVerityCodeRequest{
		Length: 6,
		Type:   1,
	})
	if err != nil {
		return "", err
	}
	status := dt.data.Rdb.Set(ctx, "DVC:"+tel, reply.Code, 60*time.Second)
	if _, err = status.Result(); err != nil {
		return "", err
	}
	return reply.Code, nil
}

func (dt *DriverData) FetchVerifyCode(ctx context.Context, telephone string) (string, error) {
	status := dt.data.Rdb.Get(context.Background(), "DVC:"+telephone)
	code, err := status.Result()
	if err != nil {
		return "", err
	}
	return code, nil
}

func (dt *DriverData) FetchInfoByTel(ctx context.Context, tel string) (*biz.Driver, error) {
	driver := &biz.Driver{}
	if err := dt.data.Mdb.Where("telephone = ?", tel).First(driver).Error; err != nil {
		return nil, err
	}
	return driver, nil
}

func (dt *DriverData) DelToken(tel interface{}) error {
	driver := &biz.Driver{}
	// 找到customer
	if err := dt.data.Mdb.Where("telephone = ?", tel).First(driver).Error; err != nil {
		return err
	}
	// 删除customer的token
	driver.Token = sql.NullString{String: "", Valid: false}
	return dt.data.Mdb.Save(driver).Error
}
