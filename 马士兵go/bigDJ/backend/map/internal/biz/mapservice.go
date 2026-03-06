package biz

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/go-kratos/kratos/v2/log"
)

type MapServiceBiz struct {
	log *log.Helper
}

func NewMapServiceBiz(logger log.Logger) *MapServiceBiz {
	return &MapServiceBiz{log: log.NewHelper(logger)}
}

// GetDriverInfo 获取驾驶信息
func (msbiz *MapServiceBiz) GetDriverInfo(origin, destination string) (string, string, error) {
	api := "https://restapi.amap.com/v3/direction/driving"
	//key=cc219593c0f7bf2f07d2962ff8b609d1&origin=116.481028,39.989643&destination=116.434446,39.90816&extensions=base
	key := "cc219593c0f7bf2f07d2962ff8b609d1"
	parameters := fmt.Sprintf("origin=%s&destination=%s&extensions=base&key=%s", origin, destination, key)
	url := api + "?" + parameters
	resp, err := http.Get(url)
	if err != nil {
		return "", "", err
	}
	defer func() {
		_ = resp.Body.Close()
	}()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", "", err
	}

	//fmt.Println(string(body))
	ddResp := &DirectionDrivingResp{}
	if err = json.Unmarshal(body, ddResp); err != nil {
		return "", "", err
	}
	if ddResp.Status == "0" {
		return "", "", errors.New(ddResp.Info)
	}

	return ddResp.Route.Paths[0].Distance, ddResp.Route.Paths[0].Duration, nil
}

type DirectionDrivingResp struct {
	Status   string `json:"status,omitempty"`
	Info     string `json:"info,omitempty"`
	Infocode string `json:"infocode,omitempty"`
	Count    string `json:"count,omitempty"`
	Route    struct {
		Origin      string `json:"origin,omitempty"`
		Destination string `json:"destination,omitempty"`
		Paths       []Path `json:"paths,omitempty"`
	} `json:"route"`
}

type Path struct {
	Distance string `json:"distance,omitempty"`
	Duration string `json:"duration,omitempty"`
	Strategy string `json:"strategy,omitempty"`
}
