package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimqianchuanSDK/conf"
	"github.com/CriarBrand/SimqianchuanSDK/utils"
	"github.com/guonaihong/gout"
)

// -----------------------------------------------------获取广告账户报表数据----------------------------------------------

type AdvertiserReportReq struct {
	AccessToken string `json:"access_token"`
	AdvertiserReportReqBase
}

// AdvertiserReportReqBase 获取广告账户数据-请求
type AdvertiserReportReqBase struct {
	AdvertiserId int64                        `json:"advertiser_id"`         // 千川广告主账户id
	StartDate    string                       `json:"start_date"`            // 开始时间，格式 2021-04-05
	EndDate      string                       `json:"end_date"`              // 结束时间，格式 2021-04-05 ，时间跨度不能超过 180 天
	Fields       []string                     `json:"fields"`                // 需要查询的 消耗指标
	Filtering    AdvertiserReportReqFiltering `json:"filtering"`             // 过滤条件
	OrderField   string                       `json:"order_field,omitempty"` // 排序字段，允许值参考 数据指标 ，默认不传为 stat_cost
	OrderType    string                       `json:"order_type,omitempty"`  // 排序方式，允许值： ASC 升序（默认）、 DESC 降序
	Page         int64                        `json:"page,omitempty"`        // 页码，默认为 1
	PageSize     int64                        `json:"page_size,omitempty"`   // 页面大小，默认为 10 ，取值范围： 1-500
}

type AdvertiserReportReqFiltering struct {
	MarketingGoal  string  `json:"marketing_goal"`            // 营销目标，允许值： ALL ：全部 VIDEO_PROM_GOODS ：短视频带货 LIVE_PROM_GOODS ：直播间带货
	OrderPlatform  string  `json:"order_platform,omitempty"`  // 下单平台，允许值： ALL ：全部 QIANCHUAN ： 千川pc（默认） ECP_AWEME ：小店随心推
	MarketingScene string  `json:"marketing_scene,omitempty"` // 营销场景，允许值： ALL ：全部 FEED ： 通投广告 SEARCH ：搜索广告 注意：当下单平台为“小店随心推”时，不支持
	PromotionWay   string  `json:"promotion_way,omitempty"`   // 推广方式，允许值： STANDARD ：专业推广 SIMPLE ： 极速推广 注意：当下单平台为“小店随心推”时，不支持
	SmartBidType   string  `json:"smart_bid_type,omitempty"`  // 投放场景（投放方式），允许值： SMART_BID_CUSTOM ：控成本投放 SMART_BID_CONSERVATIVE ： 放量投放 注意：当下单平台为“小店随心推”或营销场景为“搜索广告”时，不支持
	Status         string  `json:"status,omitempty"`          // 按计划状态过滤，不传入即默认返回“全部（包含已删除）”，其他规则详见 【附录-广告计划查询状态】 （暂不支持“系统暂停”和“在投计划配额超限”） 注意：当下单平台为“小店随心推”时，不支持
	AwemeIds       []int64 `json:"aweme_ids,omitempty"`       // 按抖音id过滤，即关联的抖音号
}

type AdvertiserReportResDetail struct {
	AdvertiserId               int64   `json:"advertiser_id"`
	StatCost                   float64 `json:"stat_cost"`
	ShowCnt                    int64   `json:"show_cnt"`
	Ctr                        float64 `json:"ctr"`
	CpmPlatform                float64 `json:"cpm_platform"`
	ClickCnt                   int64   `json:"click_cnt"`
	PayOrderCount              int64   `json:"pay_order_count"`
	CreateOrderAmount          float64 `json:"create_order_amount"`
	CreateOrderCount           int64   `json:"create_order_count"`
	PayOrderAmount             float64 `json:"pay_order_amount"`
	CreateOrderRoi             float64 `json:"create_order_roi"`
	PrepayAndPayOrderRoi       float64 `json:"prepay_and_pay_order_roi"`
	PrepayOrderCount           int64   `json:"prepay_order_count"`
	PrepayOrderAmount          float64 `json:"prepay_order_amount"`
	DyFollow                   int64   `json:"dy_follow"`
	ConvertCnt                 int64   `json:"convert_cnt"`
	ConvertCost                float64 `json:"convert_cost"`
	ConvertRate                float64 `json:"convert_rate"`
	DyShare                    int64   `json:"dy_share"`
	DyComment                  int64   `json:"dy_comment"`
	DyLike                     int64   `json:"dy_like"`
	LivePayOrderCostPerOrder   float64 `json:"live_pay_order_cost_per_order"`
	LubanLiveEnterCnt          int64   `json:"luban_live_enter_cnt"`
	LiveWatchOneMinuteCount    int64   `json:"live_watch_one_minute_count"`
	LiveFansClubJoinCnt        int64   `json:"live_fans_club_join_cnt"`
	LubanLiveSlidecartClickCnt int64   `json:"luban_live_slidecart_click_cnt"`
	LubanLiveClickProductCnt   int64   `json:"luban_live_click_product_cnt"`
	LubanLiveCommentCnt        int64   `json:"luban_live_comment_cnt"`
	LubanLiveShareCnt          int64   `json:"luban_live_share_cnt"`
	LubanLiveGiftCnt           int64   `json:"luban_live_gift_cnt"`
	LubanLiveGiftAmount        float64 `json:"luban_live_gift_amount"`
	TotalPlay                  int64   `json:"total_play"`
	PlayDuration3S             int64   `json:"play_duration_3s"`
	Play25FeedBreak            int64   `json:"play_25_feed_break"`
	Play50FeedBreak            int64   `json:"play_50_feed_break"`
	Play75FeedBreak            int64   `json:"play_75_feed_break"`
	PlayOver                   int64   `json:"play_over"`
	PlayOverRate               float64 `json:"play_over_rate"`
}

type AdvertiserReportResData struct {
	List     []AdvertiserReportResDetail `json:"list"`
	PageInfo PageInfo                    `json:"page_info"`
}

// GetAdvertiserReport 获取广告账户报表数据
func (client *Client) GetAdvertiserReport(request *AdvertiserReportReq, response *AdvertiserReportResData) error {
	df := gout.GET(client.url(conf.API_REPORT_ADVERTISER_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.AdvertiserReportReqBase))
	return client.DoRequest(df, response)
}

// -----------------------------------------------------获取广告计划数据-------------------------------------------------

type AdReportReq struct {
	AccessToken string `json:"access_token"`
	AdReportReqBase
}

// AdReportReqBase 获取广告计划数据-请求
type AdReportReqBase struct {
	AdvertiserId int64                `json:"advertiser_id"`         // 千川广告主账户id
	StartDate    string               `json:"start_date"`            // 开始时间，格式 2021-04-05
	EndDate      string               `json:"end_date"`              // 结束时间，格式 2021-04-05 ，时间跨度不能超过 180 天
	Fields       []string             `json:"fields"`                // 需要查询的 消耗指标
	Filtering    AdReportReqFiltering `json:"filtering"`             // 过滤条件
	OrderField   string               `json:"order_field,omitempty"` // 排序字段，允许值参考 数据指标 ，默认不传为 stat_cost
	OrderType    string               `json:"order_type,omitempty"`  // 排序方式，允许值： ASC 升序（默认）、 DESC 降序
	Page         int64                `json:"page,omitempty"`        // 页码，默认为 1
	PageSize     int64                `json:"page_size,omitempty"`   // 页面大小，默认为 10 ，取值范围： 1-500
}

type AdReportReqFiltering struct {
	AdIds          []int64 `json:"ad_ids,omitempty"`          // 广告计划id列表，最多支持100个
	MarketingGoal  string  `json:"marketing_goal"`            // 营销目标，允许值： ALL ：全部 VIDEO_PROM_GOODS ：短视频带货 LIVE_PROM_GOODS ：直播间带货
	OrderPlatform  string  `json:"order_platform,omitempty"`  // 下单平台，允许值： ALL ：全部 QIANCHUAN ： 千川pc（默认） ECP_AWEME ：小店随心推
	MarketingScene string  `json:"marketing_scene,omitempty"` // 营销场景，允许值： ALL ：全部 FEED ： 通投广告 SEARCH ：搜索广告 注意：当下单平台为“小店随心推”时，不支持
	PromotionWay   string  `json:"promotion_way,omitempty"`   // 推广方式，允许值： STANDARD ：专业推广 SIMPLE ： 极速推广 注意：当下单平台为“小店随心推”时，不支持
	SmartBidType   string  `json:"smart_bid_type,omitempty"`  // 投放场景（投放方式），允许值： SMART_BID_CUSTOM ：控成本投放 SMART_BID_CONSERVATIVE ： 放量投放 注意：当下单平台为“小店随心推”或营销场景为“搜索广告”时，不支持
	Status         string  `json:"status,omitempty"`          // 按计划状态过滤，不传入即默认返回“全部（包含已删除）”，其他规则详见 【附录-广告计划查询状态】 （暂不支持“系统暂停”和“在投计划配额超限”） 注意：当下单平台为“小店随心推”时，不支持
}

type AdReportResDetail struct {
	AdvertiserId               int64   `json:"advertiser_id"`
	AdId                       int64   `json:"ad_id"`
	StatCost                   float64 `json:"stat_cost"`
	ShowCnt                    int64   `json:"show_cnt"`
	Ctr                        float64 `json:"ctr"`
	CpmPlatform                float64 `json:"cpm_platform"`
	ClickCnt                   int64   `json:"click_cnt"`
	PayOrderCount              int64   `json:"pay_order_count"`
	CreateOrderAmount          float64 `json:"create_order_amount"`
	CreateOrderCount           int64   `json:"create_order_count"`
	PayOrderAmount             float64 `json:"pay_order_amount"`
	CreateOrderRoi             float64 `json:"create_order_roi"`
	PrepayAndPayOrderRoi       float64 `json:"prepay_and_pay_order_roi"`
	PrepayOrderCount           int64   `json:"prepay_order_count"`
	PrepayOrderAmount          float64 `json:"prepay_order_amount"`
	DyFollow                   int64   `json:"dy_follow"`
	ConvertCnt                 int64   `json:"convert_cnt"`
	ConvertCost                float64 `json:"convert_cost"`
	ConvertRate                float64 `json:"convert_rate"`
	DyShare                    int64   `json:"dy_share"`
	DyComment                  int64   `json:"dy_comment"`
	DyLike                     int64   `json:"dy_like"`
	LivePayOrderCostPerOrder   float64 `json:"live_pay_order_cost_per_order"`
	LubanLiveEnterCnt          int64   `json:"luban_live_enter_cnt"`
	LiveWatchOneMinuteCount    int64   `json:"live_watch_one_minute_count"`
	LiveFansClubJoinCnt        int64   `json:"live_fans_club_join_cnt"`
	LubanLiveSlidecartClickCnt int64   `json:"luban_live_slidecart_click_cnt"`
	LubanLiveClickProductCnt   int64   `json:"luban_live_click_product_cnt"`
	LubanLiveCommentCnt        int64   `json:"luban_live_comment_cnt"`
	LubanLiveShareCnt          int64   `json:"luban_live_share_cnt"`
	LubanLiveGiftCnt           int64   `json:"luban_live_gift_cnt"`
	LubanLiveGiftAmount        float64 `json:"luban_live_gift_amount"`
	TotalPlay                  int64   `json:"total_play"`
	PlayDuration3S             int64   `json:"play_duration_3s"`
	Play25FeedBreak            int64   `json:"play_25_feed_break"`
	Play50FeedBreak            int64   `json:"play_50_feed_break"`
	Play75FeedBreak            int64   `json:"play_75_feed_break"`
	PlayOver                   int64   `json:"play_over"`
	PlayOverRate               float64 `json:"play_over_rate"`
}

type AdReportResData struct {
	List     []AdReportResDetail `json:"list"`
	PageInfo PageInfo            `json:"page_info"`
}

// GetAdReport 获取广告计划报表数据
func (client *Client) GetAdReport(request *AdReportReq, response *AdReportResData) error {
	df := gout.GET(client.url(conf.API_REPORT_AD_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.AdReportReqBase))
	return client.DoRequest(df, response)
}
