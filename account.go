package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimqianchuanSDK/conf"
	"github.com/CriarBrand/SimqianchuanSDK/utils"
	"github.com/guonaihong/gout"
)

// -----------------------------------------------------获取授权时登录用户信息----------------------------------------------

type UserInfoRes struct {
	ID          int64  `json:"id"`           //用户id
	Email       string `json:"email"`        //邮箱（已经脱敏处理）
	DisplayName string `json:"display_name"` // 用户名
}

// GetUserInfo 获取授权时登录用户信息
func (client *Client) GetUserInfo(accessToken string, response *UserInfoRes) error {
	df := gout.GET(client.url(conf.API_USER_INFO)).
		SetHeader(gout.H{
			"Access-Token": accessToken,
		})
	return client.DoRequest(df, response)
}

// -----------------------------------------------------获取已授权的账户（店铺/代理商）----------------------------------------------

// AdvertiserListResData 获取已授权的账户（店铺/代理商）-返回
type AdvertiserListResData struct {
	AdvertiserId   int64  `json:"advertiser_id"`   // 账户id
	AdvertiserName string `json:"advertiser_name"` // 账户名称
	IsValid        bool   `json:"is_valid"`        // 授权有效性，返回值：true/false,用于判断当前授权关系是否仍然有效
	AccountRole    string `json:"account_role"`    // 授权账号角色，返回值：PLATFORM_ROLE_QIANCHUAN_AGENT代理商账户、PLATFORM_ROLE_SHOP_ACCOUNT 店铺账户
}

type AdvertiserListResDataCom struct {
	List []AdvertiserListResData `json:"list"`
}

// GetAdvertiserList 获取已授权的账户（店铺/代理商）
func (client *Client) GetAdvertiserList(accessToken string, response *AdvertiserListResDataCom) error {
	df := gout.GET(client.url(conf.API_ADVERTISER_LIST)).
		SetQuery(gout.H{
			"access_token": accessToken,
			"app_id":       client.appId,
			"secret":       client.secret,
		})
	return client.DoRequest(df, response)
}

// -----------------------------------------------------获取店铺/代理商账户关联的广告账户列表----------------------------------------------

// AccountAdvertiserListReq 获取店铺账户关联的广告账户列表-请求
type AccountAdvertiserListReq struct {
	ObjectId    int64  // 店铺/代理商id
	Page        uint64 // 页码.默认值: 1
	PageSize    uint64 // 页面数据量.默认值: 10， 最大值：100
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type AccountAdvertiserListResCom struct {
	List     []int64  `json:"list"`
	PageInfo PageInfo `json:"page_info"`
}

// GetShopAdvertiserList 获取店铺账户关联的广告账户列表
func (client *Client) GetShopAdvertiserList(request *AccountAdvertiserListReq, response *AccountAdvertiserListResCom) error {
	df := gout.GET(client.url(conf.API_SHOP_ADVERTISER_LIST)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(gout.H{
			"shop_id":   request.ObjectId,
			"page":      request.Page,
			"page_size": request.PageSize,
		})
	return client.DoRequest(df, response)
}

// GetAgentAdvertiserList 获取代理商账户关联的广告账户列表-请求
func (client *Client) GetAgentAdvertiserList(request *AccountAdvertiserListReq, response *AccountAdvertiserListResCom) error {
	df := gout.GET(client.url(conf.API_AGENT_ADVERTISER_LIST)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(gout.H{
			"advertiser_id": request.ObjectId,
			"page":          request.Page,
			"page_size":     request.PageSize,
		})
	return client.DoRequest(df, response)
}

// -----------------------------------------------------获取千川广告账户基础信息----------------------------------------------

// AdvertiserPublicInfoReq 获取千川广告账户基础信息-请求
type AdvertiserPublicInfoReq struct {
	AdvertiserPublicInfoReqBase
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type AdvertiserPublicInfoResData struct {
	ID                 int64  `json:"id"`
	Name               string `json:"name"`
	Company            string `json:"company"`
	FirstIndustryName  string `json:"first_industry_name"`
	SecondIndustryName string `json:"second_industry_name"`
}

type AdvertiserPublicInfoReqBase struct {
	AdvertiserIds []int64 `json:"advertiser_ids,omitempty"`
}

// GetAdvertiserPublicInfo 获取千川广告账户基础信息
func (client *Client) GetAdvertiserPublicInfo(request *AdvertiserPublicInfoReq, response *[]AdvertiserPublicInfoResData) error {
	df := gout.GET(client.url(conf.API_ADVERTISER_PUBLIC_INFO)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.AdvertiserPublicInfoReqBase))
	return client.DoRequest(df, response)
}

// -----------------------------------------------------获取在投计划配额信息----------------------------------------------

// AdvertiserAdQuotaReq 获取在投计划配额信息-请求
type AdvertiserAdQuotaReq struct {
	AdvertiserId int64
	AccessToken  string // 调用/oauth/access_token/生成的token，此token需要用户授权。
}

type AdvertiserAdQuotaResData struct {
	QuotaFeed struct {
		DeliveryInfo struct {
			AdlabNum   int64 `json:"adlab_num"`
			NoAdlabNum int64 `json:"no_adlab_num"`
			TotalNum   int64 `json:"total_num"`
		} `json:"delivery_info"`
		MonthCost int64 `json:"month_cost"`
		QuotaGear int64 `json:"quota_gear"`
		QuotaInfo struct {
			TotalNum int64 `json:"total_num"`
		} `json:"quota_info"`
		StageInfo struct {
			IsPrimary int64 `json:"is_primary"`
		} `json:"stage_info"`
	} `json:"quota_feed"`
	QuotaSearch struct {
		DeliveryInfo struct {
			AdlabNum   int64 `json:"adlab_num"`
			NoAdlabNum int64 `json:"no_adlab_num"`
			TotalNum   int64 `json:"total_num"`
		} `json:"delivery_info"`
		MonthCost int64 `json:"month_cost"`
		QuotaGear int64 `json:"quota_gear"`
		QuotaInfo struct {
			TotalNum int64 `json:"total_num"`
		} `json:"quota_info"`
		StageInfo struct {
			IsPrimary int64 `json:"is_primary"`
		} `json:"stage_info"`
	} `json:"quota_search"`
}

// GetAdvertiserAdQuota 获取在投计划配额信息
func (client *Client) GetAdvertiserAdQuota(request *AdvertiserAdQuotaReq, response *AdvertiserAdQuotaResData) error {
	df := gout.GET(client.url(conf.API_ADVERTISER_AD_QUOTA)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(gout.H{
			"advertiser_id": request.AdvertiserId,
		})
	return client.DoRequest(df, response)
}
