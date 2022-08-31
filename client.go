package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimqianchuanSDK/conf"
	"github.com/CriarBrand/SimqianchuanSDK/utils"
	"github.com/guonaihong/gout/dataflow"
	"github.com/guonaihong/gout/debug"
	api "github.com/guonaihong/gout/interface"
	"math"
	"math/rand"
	"time"
)

type Client struct {
	appId          int64
	secret         string
	debug          bool
	retryTime      int
	maxWaitTime    time.Duration
	waitTime       time.Duration
	beforeRequest  []api.RequestMiddler
	beforeResponse []api.ResponseMiddler
}

func NewClient(appId int64, secret string, retry int, debug bool, waitTime, maxWaitTime time.Duration) *Client {
	return &Client{
		appId:       appId,
		secret:      secret,
		debug:       debug,
		retryTime:   retry,
		maxWaitTime: maxWaitTime,
		waitTime:    waitTime,
	}
}

func (client *Client) UseRequest(handler ...api.RequestMiddler) {
	client.beforeRequest = append(client.beforeRequest, handler...)
}

func (client *Client) UseResponse(handler ...api.ResponseMiddler) {
	client.beforeResponse = append(client.beforeResponse, handler...)
}

func (client *Client) DoRequest(df *dataflow.DataFlow, dest interface{}) error {
	if len(client.beforeRequest) > 0 {
		df = df.RequestUse(client.beforeRequest...)
	}
	if len(client.beforeResponse) > 0 {
		df = df.ResponseUse(client.beforeResponse...)
	}
	if client.debug {
		df = df.Debug(debug.Func(func(o *debug.Options) {
			o.Debug = true
			o.Color = true
		}))
	}
	retryFunc := func() error {
		return df.BindJSON(dest).Do()
	}
	retryTime := client.retryTime
	// 如果把BindJSON移到上面df处，会出现无法绑定的情况
	for err := retryFunc(); err != nil; err = retryFunc() {
		if retryTime <= 0 {
			return err
		}
		select {
		case <-time.After(client.getSleep(retryTime)):
		}
		retryTime--
	}
	return nil
}

func (client *Client) getSleep(curryTryTime int) time.Duration {
	temp := uint64(client.waitTime * time.Duration(math.Exp2(float64(curryTryTime))))
	if temp <= 0 {
		temp = uint64(client.waitTime)
	}
	temp = utils.Min(uint64(client.maxWaitTime), uint64(temp))
	//对int64边界处理, 后面使用rand.Int63n所以,最大值只能是int64的最大值防止溢出
	if temp > math.MaxInt64 {
		temp = math.MaxInt64
	}

	temp /= 2
	return time.Duration(temp) + time.Duration(rand.Int63n(int64(temp)))
}

func (client *Client) url(api string) string {
	return conf.API_HTTP_SCHEME + conf.API_HOST + api
}
