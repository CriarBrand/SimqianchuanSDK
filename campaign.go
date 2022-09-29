package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimqianchuanSDK/conf"
	"github.com/CriarBrand/SimqianchuanSDK/utils"
	"github.com/guonaihong/gout"
)

//--------------------------------------------广告组列表获取--------------------------------------------------------------------

// CampaignListGetReq 获取广告组数据-请求
type CampaignListGetReq struct {
	AccessToken string `json:"access_token"` // 调用/oauth/access_token/生成的token，此token需要用户授权。
	CampaignListGetReqBase
}

type CampaignListGetReqBase struct {
	AdvertiserId int64                 `json:"advertiser_id"`       // 千川广告账户ID
	Filter       CampaignListGetFilter `json:"filter"`              // 过滤器，无过滤条件情况下返回“所有不包含已删除”的广告组列表
	Page         int64                 `json:"page,omitempty"`      // 页码，默认为1
	PageSize     int64                 `json:"page_size,omitempty"` // 页面大小，默认值: 10， 允许值：10、20、50、100、500、1000

}

type CampaignListGetFilter struct {
	Ids            []int64 `json:"ids,omitempty"`             // 广告组ID列表，目前只支持一个。
	Name           string  `json:"name,omitempty"`            // 广告组名称关键字，长度为1-30个字符，其中1个中文字符算2位
	MarketingGoal  string  `json:"marketing_goal"`            // 广告组营销目标，允许值：VIDEO_PROM_GOODS：短视频带货、LIVE_PROM_GOODS：直播带货
	MarketingScene string  `json:"marketing_scene,omitempty"` // 营销场景，允许值：FEED 通投广告，SEARCH 搜索广告，默认为 FEED
	Status         string  `json:"status,omitempty"`          // 广告组状态，允许值：ALL：所有包含已删除、ENABLE：启用、DISABLE：暂停、DELETE：已删除。不传入即默认返回“所有不包含已删除”
}

type CampaignListGetResData struct {
	List     []CampaignListGetResDataDetail `json:"list"`
	PageInfo PageInfo                       `json:"page_info"`
}
type CampaignListGetResDataDetail struct {
	ID             int64   `json:"id"`              // 广告组ID
	Name           string  `json:"name"`            // 广告组名称
	Budget         float64 `json:"budget"`          // 广告组预算，单位：元，精确到两位小数。
	BudgetMode     string  `json:"budget_mode"`     // 广告组预算类型
	MarketingGoal  string  `json:"marketing_goal"`  // 广告组营销目标，VIDEO_PROM_GOODS：短视频带货、LIVE_PROM_GOODS：直播带货。
	MarketingScene string  `json:"marketing_scene"` // 营销场景，允许值：FEED 通投广告，SEARCH 搜索广告
	Status         string  `json:"status"`          // 广告组状态，ALL：所有包含已删除、ENABLE：启用、DISABLE：暂停、DELETE：已删除。
	CreateDate     string  `json:"create_date"`     // 广告组创建日期, 格式：yyyy-mm-dd
}

// GetCampaignList 广告组列表获取
func (client *Client) GetCampaignList(request *CampaignListGetReq, response *CampaignListGetResData) error {
	df := gout.GET(client.url(conf.API_CAMPAIGN_LIST_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.CampaignListGetReqBase))
	return client.DoRequest(df, response)
}

//--------------------------------------------广告组状态更新--------------------------------------------------------------------

// UpdateBatchCampaignStatusReq 广告组状态更新
type UpdateBatchCampaignStatusReq struct {
	AccessToken                      string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	UpdateBatchCampaignStatusReqBody        // POST请求的data
}

type UpdateBatchCampaignStatusReqBody struct {
	AdvertiserId int64   `json:"advertiser_id"`
	CampaignIds  []int64 `json:"campaign_ids"` // 广告组ID，不超过10个，操作更新的广告组ID需要属于千川账户ID否则会报错；
	OptStatus    string  `json:"opt_status"`   // 操作类型，允许值: "ENABLE"：启用, "DELETE"：删除, "DISABLE"：暂停；对于删除的广告组不可进行任何操作。
}

type UpdateBatchCampaignStatusRes struct {
	Success []int64                              `json:"success"` // 更新成功的广告组ID列表
	Errors  []UpdateBatchCampaignStatusResErrors `json:"errors"`  // 更新失败的广告组列表
}

type UpdateBatchCampaignStatusResErrors struct {
	CampaignId   int64  `json:"campaign_id"`   // 更新失败广告组ID
	ErrorMessage string `json:"error_message"` // 更新失败的原因
}

// UpdateBatchCampaignStatus 广告组状态更新
func (client *Client) UpdateBatchCampaignStatus(request *UpdateBatchCampaignStatusReq, response *UpdateBatchCampaignStatusRes) error {
	df := gout.POST(client.url(conf.API_BATCH_CAMPAIGN_STATUS_UPDATE)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetJSON(request.UpdateBatchCampaignStatusReqBody)
	return client.DoRequest(df, response)
}

//--------------------------------------------广告组更新--------------------------------------------------------------------

// UpdateCampaignReq 广告组更新
type UpdateCampaignReq struct {
	AccessToken           string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	UpdateCampaignReqBody        // POST请求的data
}

type UpdateCampaignReqBody struct {
	AdvertiserId int64   `json:"advertiser_id"`
	CampaignId   int64   `json:"campaign_id"`             // 广告组ID
	BudgetMode   string  `json:"budget_mode,omitempty"`   // 预算类型，详见【附录-预算类型】，允许值： BUDGET_MODE_DAY 日预算，BUDGET_MODE_INFINITE 预算不限
	Budget       float64 `json:"budget,omitempty"`        // 广告组预算，最多支持两位小数，当budget_mode为BUDGET_MODE_DAY时必填，预算单次修改幅度不能低于100元，且日预算不少于300元
	CampaignName string  `json:"campaign_name,omitempty"` // 广告组名称，长度为1-100个字符，其中1个中文字符算2位 需要注意：广告组名称不修改的话，可不填。填入的话，需与原广告组名称不同，否则报错
}

type UpdateCampaignRes struct {
	CampaignId int64 `json:"campaign_id"`
}

// UpdateCampaign 广告组更新
func (client *Client) UpdateCampaign(request *UpdateCampaignReq, response *UpdateCampaignRes) error {
	df := gout.POST(client.url(conf.API_CAMPAIGN_UPDATE)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetJSON(request.UpdateCampaignReqBody)
	return client.DoRequest(df, response)
}
