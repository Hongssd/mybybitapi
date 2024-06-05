package mybybitapi

import (
	"fmt"
)

type BybitErrorRes struct {
	RetCode int    `json:"retCode"`
	RetMsg  string `json:"retMsg"`
}

type BybitTimeRes struct {
	Time int64 `json:"time"` //REST网关接收请求时的时间戳，Unix时间戳的微秒数格式，如 1597026383085123返回的时间是请求验证后的时间
}
type BybitRestRes[T any] struct {
	BybitErrorRes             //错误信息
	BybitTimeRes              //时间戳
	RetExtInfo    interface{} `json:"retExtInfo"`
	Result        T           `json:"result"` //请求结果
}

func handlerCommonRest[T any](data []byte, code int) (*BybitRestRes[T], error) {
	res := &BybitRestRes[T]{}

	if code == 401 {
		return nil, fmt.Errorf("response 401 code, check your api key and api secret")
	}

	if code != 200 {
		return nil, fmt.Errorf("http response code:%v", code)
	}
	err := json.Unmarshal(data, &res)
	if err != nil {
		log.Errorf("rest err data: %s", string(data))
		return nil, err
	}
	return res, err
}
func (err *BybitErrorRes) handlerError() error {
	if err.RetCode != 0 && err.RetMsg != "OK" {
		return fmt.Errorf("request error:[code:%v][message:%v]", err.RetCode, err.RetMsg)
	} else {
		return nil
	}

}
func resConvertTo[T, R any](origin *BybitRestRes[T], target R) *BybitRestRes[R] {
	newRes := &BybitRestRes[R]{
		BybitErrorRes: origin.BybitErrorRes,
		BybitTimeRes:  origin.BybitTimeRes,
		RetExtInfo:    origin.RetExtInfo,
		Result:        target,
	}
	return newRes
}
