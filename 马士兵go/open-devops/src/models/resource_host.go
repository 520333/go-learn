package models

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"time"
)

// AgentCollectInfo 机器shell采集到的字段
type AgentCollectInfo struct {
	SN       string `json:"sn"`       // sn号
	CPU      string `json:"cpu"`      // cpu核心数
	Mem      string `json:"mem"`      // 内存g数
	Disk     string `json:"disk"`     // 磁盘g数
	IpAddr   string `json:"ip_addr"`  // ip
	HostName string `json:"hostname"` // hostname
}

type ResourceHost struct {
	// 公共字段
	Id         int64           `json:"id"`
	Uid        string          `json:"uid"`
	Hash       string          `json:"hash"`
	Name       string          `json:"name"`
	PrivateIps json.RawMessage `json:"private_ips"`
	Tags       json.RawMessage `json:"tags"`

	// 公有云字段
	CloudProvider    string          `json:"cloud_provider"`
	ChargingMode     string          `json:"charging_mode"`
	Region           string          `json:"region"`
	AccountId        int64           `json:"account_id"`
	VpcId            string          `json:"vpc_id"`
	SubnetId         string          `json:"subnet_id"`
	SecurityGroups   string          `json:"security_groups"`
	Status           string          `json:"status"`
	InstanceType     string          `json:"instance_type"`
	PublicIps        json.RawMessage `json:"public_ips"`
	AvailabilityZone string          `json:"availability_zone"`

	// 机器采集到的字段
	SN         string    `json:"sn" xorm:"-"`       // sn号
	CPU        string    `json:"cpu" xorm:"cpu"`    // cpu核数
	Mem        string    `json:"mem"`               // 内存g数
	Disk       string    `json:"disk"`              // 磁盘g数
	IpAddr     string    `json:"ip_addr" xorm:"-"`  // ip
	HostName   string    `json:"hostname" xorm:"-"` // hostname
	CreateTime time.Time `json:"create_time" xorm:"create_time created"`
	UpdateTime time.Time `json:"update_time" xorm:"update_time updated"`
}

// GenHash 是判断这个资源是否发生变化的函数
func (rh *ResourceHost) GenHash() string {
	h := md5.New()
	h.Write([]byte(rh.SN))
	h.Write([]byte(rh.Name))
	h.Write([]byte(rh.IpAddr))
	h.Write([]byte(rh.CPU))
	h.Write([]byte(rh.Mem))
	h.Write([]byte(rh.Disk))
	return hex.EncodeToString(h.Sum(nil))
}

func (rh *ResourceHost) GetOne() (*ResourceHost, error) {

	has, err := DB["stree"].Get(rh)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, nil
	}
	return rh, nil

}

func (rh *ResourceHost) AddOne() error {
	_, err := DB["stree"].InsertOne(rh)
	return err
}

func (rh *ResourceHost) Update() (bool, error) {
	rowAffected, err := DB["stree"].ID(rh.Id).Update(rh)
	//rowAffected, err := DB["stree"].Update(rh)
	if err != nil {
		return false, err
	}
	if rowAffected > 0 {
		return true, nil
	}
	return false, nil
}
