package biz

import (
	"context"
	"database/sql"
	"regexp"
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

const SecretKey = "driver secret key"

// 司机表模型
type Driver struct {
	// 基础模型
	gorm.Model
	// 业务模型
	DriverWork
	// 关联部分
}

// 司机的业务模型
type DriverWork struct {
	Telephone     string         `gorm:"type:varchar(16);uniqueIndex;" json:"telephone"`
	Token         sql.NullString `gorm:"type:varchar(2047);" json:"token"`
	Name          sql.NullString `gorm:"type:varchar(255);index;" json:"name"`
	Status        sql.NullString `gorm:"type:enum('out', 'in', 'listen', 'stop');" json:"status"`
	IdNumber      sql.NullString `gorm:"type:char(18);uniqueIndex;" json:"id_number"`
	IdImageA      sql.NullString `gorm:"type:varchar(255);" json:"id_image_a"`
	LicenseImageA sql.NullString `gorm:"type:varchar(255);" json:"license_image_a"`
	LicenseImageB sql.NullString `gorm:"type:varchar(255);" json:"license_image_b"`
	DistinctCode  sql.NullString `gorm:"type:varchar(16);index;" json:"distinct_code"`
	TelephoneBak  sql.NullString `gorm:"type:varchar(16);index;" json:"telephone_bak"`
	AuditAt       sql.NullTime   `gorm:"index;" json:"audit_at"`
}

// 司机状态常量
const DriverStatusOut = "out"
const DriverStatusIn = "in"
const DriverStatusListen = "listen"
const DriverStatusStop = "stop"
const DriverTokenLife = 1 * 30 * 24 * 3600

type DriverBiz struct {
	Di DriverInterface
}

type DriverInterface interface {
	GetVerifyCode(context.Context, string) (string, error)
	FetchVerifyCode(context.Context, string) (string, error)
	FetchInfoByTel(context.Context, string) (*Driver, error)
	InitDriverInfo(context.Context, string) (*Driver, error)
	CheckVerifyCode(context.Context, string) (*Driver, error)
	GetSaveVerifyCode(context.Context, string) (string, error)
	SaveToken(context.Context, string, string) error
	GetToken(context.Context, string) (string, error)
	DelToken(interface{}) error
}

func NewDriverBiz(di DriverInterface) *DriverBiz {
	return &DriverBiz{
		Di: di,
	}
}

func (db *DriverBiz) DelToken(tel interface{}) error {
	return db.Di.DelToken(tel)
}

// CheckLogin 验证码登录信息方法
func (db *DriverBiz) CheckLogin(ctx context.Context, tel, verityCode string) (string, error) {
	code, err := db.Di.GetSaveVerifyCode(ctx, tel)
	if err != nil {
		return "", err
	}
	if verityCode != code {
		return "", errors.New(1, "verify code error", "")
	}
	// 生成token
	token, err := generateJWT(tel)
	if err != nil {
		return "", err
	}
	// 存储到driver表中
	if err = db.Di.SaveToken(ctx, tel, token); err != nil {
		return "", err
	}
	return token, nil
}
func generateJWT(tel string) (string, error) {
	claims := jwt.RegisteredClaims{
		Issuer:    "bigDJ",
		Subject:   "driver authentication",
		Audience:  []string{"driver"},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(DriverTokenLife * time.Second)),
		NotBefore: nil,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		ID:        tel,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// InitDriverInfo 将司机信息入库
func (db *DriverBiz) InitDriverInfo(ctx context.Context, tel string) (*Driver, error) {
	if tel == "" {
		return nil, errors.New(1, "telephone is empty", "")
	}
	return db.Di.InitDriverInfo(ctx, tel)
}

func (db *DriverBiz) GetVerifyCode(ctx context.Context, tel string) (string, error) {
	// 一 校验手机号
	pattern := `^(13\d|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18\d|19[0-35-9])\d{8}$`
	regexpPattern := regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(tel) {
		return "", errors.New(200, "driver telephone error", "")
	}
	return db.Di.GetVerifyCode(ctx, tel)
}

func (db *DriverBiz) CheckVerifyCode(ctx context.Context, tel, code string) bool {
	// 一 校验手机号
	pattern := `^(13\d|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18\d|19[0-35-9])\d{8}$`
	regexpPattern := regexp.MustCompile(pattern)
	if !regexpPattern.MatchString(tel) {
		return false
	}
	code = strings.TrimSpace(code)
	if len(code) == 0 {
		return false
	}
	verityCode, err := db.Di.GetVerifyCode(ctx, tel)
	if err != nil {
		return false
	}
	if verityCode != code {
		return true
	}
	return false
}

func (db *DriverBiz) LoadOrCreateByTel(ctx context.Context, tel string) (*Driver, error) {
	driver, err := db.Di.FetchInfoByTel(ctx, tel)
	if err == nil {
		return driver, err
	}
	//driver = &Driver{}
	//driver.Telephone = tel
	//driver, err := db.Di.CreateDriver(ctx, driver)
	//if err != nil {
	//	return driver, nil
	//}
	return driver, nil

}
