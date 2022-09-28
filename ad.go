package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimqianchuanSDK/conf"
	"github.com/CriarBrand/SimqianchuanSDK/utils"
	"github.com/guonaihong/gout"
)

//--------------------------------------------更新计划状态--------------------------------------------------------------------

// AdStatusUpdateReq 更新计划状态的请求结构体
type AdStatusUpdateReq struct {
	AccessToken        string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdStatusUpdateBody        // POST请求的data
}

type AdStatusUpdateBody struct {
	AdIds              []int64 `json:"ad_ids"`                         //需要更新的广告计划id，最多支持10个
	AdvertiserId       int64   `json:"advertiser_id"`                  //广告主id
	OptStatus          string  `json:"opt_status"`                     //批量更新的广告计划状态，允许值： DISABLE 暂停计划、 DELETE 删除计划、 ENABLE 启用计划、 REVIVE 复活续投计划
	ScheduleFixedRange int64   `json:"schedule_fixed_range,omitempty"` // 固定投放时长，当opt_status为 REVIVE 时必填 单位为秒，最小值为1800（0.5小时），最大值为48*1800（24小时），值必须为1800倍数，不然会报错
	ReviveBudget       float64 `json:"revive_budget,omitempty"`        // 复活预算 注： 当opt_status为REVIVE时，revive_budget与budget二者至少传一个，如二者都填，系统默认以revive_budget为准
	Budget             float64 `json:"budget,omitempty"`               // 预算，当opt_status为 REVIVE 时必填，单位为元，最多支持两位小数 当预算模式为日预算时，预算范围是300 - 9999999.99； 当预算模式为总预算时，预算范围是max(300,投放天数x100) - 9999999.99
}

// AdStatusUpdateResData 更新计划状态 的 响应结构体
type AdStatusUpdateResData struct {
	AdId   []int64                      `json:"ad_id"`  //更新成功的计划id
	Errors []AdStatusUpdateResDataError `json:"errors"` //更新失败的计划id和失败原因
}

type AdStatusUpdateResDataError struct {
	AdId         int64  `json:"ad_id"`         //更新失败的计划id
	ErrorMessage string `json:"error_message"` //更新预算失败的原因
}

// UpdateAdStatus 更新计划状态
func (client *Client) UpdateAdStatus(request *AdStatusUpdateReq, response *AdStatusUpdateResData) error {
	df := gout.POST(client.url(conf.API_AD_STATUS_UPDATE)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetJSON(request.AdStatusUpdateBody)
	return client.DoRequest(df, response)
}

//--------------------------------------------更新计划预算--------------------------------------------------------------------

// AdBudgetUpdateReq 更新计划预算 的 请求结构体
type AdBudgetUpdateReq struct {
	AccessToken        string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdBudgetUpdateBody        // POST请求的data
}

type AdBudgetUpdateBody struct {
	AdvertiserId int64                    `json:"advertiser_id"` //广告主id
	Data         []AdBudgetUpdateBodyData `json:"data"`          //更新预算的计划id和预算价格列表，最多支持10个
}

type AdBudgetUpdateBodyData struct {
	AdId   int64   `json:"ad_id"`  //广告计划id
	Budget float64 `json:"budget"` //更新后的预算，最多只有两位小数
}

type AdBudgetUpdateResData AdStatusUpdateResData

// UpdateAdBudget 更新计划预算
func (client *Client) UpdateAdBudget(request *AdBudgetUpdateReq, response *AdBudgetUpdateResData) error {
	df := gout.POST(client.url(conf.API_AD_BUDGET_UPDATE)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetJSON(request.AdBudgetUpdateBody)
	return client.DoRequest(df, response)
}

//--------------------------------------------更新计划出价--------------------------------------------------------------------

// AdBidUpdateReq 更新计划出价 的 请求结构体
type AdBidUpdateReq struct {
	AccessToken     string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdBidUpdateBody        // POST请求的data
}

type AdBidUpdateBody struct {
	AdvertiserId int64                 `json:"advertiser_id"` //广告主id
	Data         []AdBidUpdateBodyData `json:"data"`          //更新预算的计划id和预算价格列表，最多支持10个
}

type AdBidUpdateBodyData struct {
	AdId int64   `json:"ad_id"` //广告计划id
	Bid  float64 `json:"bid"`   //计划更新之后的出价，最多只有两位小数
}

type AdBidUpdateResData AdStatusUpdateResData

// UpdateAdBid 更新计划预算
func (client *Client) UpdateAdBid(request *AdBidUpdateReq, response *AdBidUpdateResData) error {
	df := gout.POST(client.url(conf.API_AD_BID_UPDATE)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetJSON(request.AdBidUpdateBody)
	return client.DoRequest(df, response)
}

//--------------------------------------------获取账户下计划列表（不含创意）--------------------------------------------------------------------

// AdListGetReq 获取账户下计划列表（不含创意）
type AdListGetReq struct {
	AccessToken string `json:"access_token"` // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdListGetReqBase
}

type AdListGetReqBase struct {
	AdvertiserId     int64              `json:"advertiser_id"`                // 千川广告账户ID
	RequestAwemeInfo int64              `json:"request_aweme_info,omitempty"` // 是否包含抖音号信息，允许值：0：不包含；1：包含；默认不返回
	Filtering        AdListGetFiltering `json:"filtering"`                    // 过滤器，无过滤条件情况下返回“所有不包含已删除”的广告组列表
	Page             int64              `json:"page,omitempty"`               // 页码，默认为1
	PageSize         int64              `json:"page_size,omitempty"`          // 页面大小，默认值: 10， 允许值：10、20、50、100、500、1000

}

type AdListGetFiltering struct {
	Ids               []int64 `json:"ids,omitempty"`                  // 按计划ID过滤，list长度限制 1-100
	AdName            string  `json:"ad_name,omitempty"`              // 按计划名称过滤，长度为1-30个字符
	Status            string  `json:"status,omitempty"`               // 按计划状态过滤，不传入即默认返回“所有不包含已删除”，其他规则详见【附录-广告计划查询状态】
	PromotionWay      string  `json:"promotion_way,omitempty"`        //按推广方式过滤，允许值：STANDARD专业推广、SIMPLE极速推广
	MarketingGoal     string  `json:"marketing_goal"`                 // 按营销目标过滤，允许值：VIDEO_PROM_GOODS：短视频带货；LIVE_PROM_GOODS：直播带货
	MarketingScene    string  `json:"marketing_scene,omitempty"`      // 按营销场景过滤，允许值：ALL 全部，FEED 通投广告，SEARCH 搜索广告，默认为FEED
	CampaignId        int64   `json:"campaign_id,omitempty"`          // 按广告组ID过滤
	AdCreateStartDate string  `json:"ad_create_start_date,omitempty"` // 计划创建开始时间，格式："yyyy-mm-dd"
	AdCreateEndDate   string  `json:"ad_create_end_date,omitempty"`   // 计划创建结束时间，与ad_create_start_date搭配使用，格式："yyyy-mm-dd"，时间跨度不能超过180天
	AdModifyTime      string  `json:"ad_modify_time,omitempty"`       // 计划修改时间，精确到小时，格式："yyyy-mm-dd HH"
	AwemeId           int64   `json:"aweme_id,omitempty"`             //根据抖音号过滤
	AutoManageFilter  string  `json:"auto_manage_filter,omitempty"`   //按是否为托管计划过滤，允许值：ALL ：不限，AUTO_MANAGE ：托管计划，NORMAL ：非托管计划，默认为ALL
}

type AdListGetResData struct {
	List     []AdListGetResDataDetail `json:"list"`
	FailList []int64                  `json:"fail_list"` // 获取失败的计划ID列表
	PageInfo PageInfo                 `json:"page_info"`
}

type AdListGetResDataDetail struct {
	AdId            int64                           `json:"ad_id"`
	CampaignId      int64                           `json:"campaign_id"`
	MarketingGoal   string                          `json:"marketing_goal"`
	PromotionWay    string                          `json:"promotion_way"`
	MarketingScene  string                          `json:"marketing_scene"`
	Name            string                          `json:"name"`
	Status          string                          `json:"status"`
	OptStatus       string                          `json:"opt_status"`
	AdCreateTime    string                          `json:"ad_create_time"`
	AdModifyTime    string                          `json:"ad_modify_time"`
	LabAdType       string                          `json:"lab_ad_type"`
	ProductInfo     []AdListGetResDataProductInfo   `json:"product_info"`
	AwemeInfo       []AdListGetResDataAwemeInfo     `json:"aweme_info"`
	DeliverySetting AdListGetResDataDeliverySetting `json:"delivery_setting"`
}

type AdListGetResDataProductInfo struct {
	Id                  int64   `json:"id"`
	Name                string  `json:"name"`
	DiscountPrice       float64 `json:"discount_price"`        // 售价，已废弃
	Img                 string  `json:"img"`                   // 商品主图
	MarketPrice         float64 `json:"market_price"`          //  原价，单位为元
	DiscountLowerPrice  float64 `json:"discount_lower_price"`  // 折扣价区间最小值，单位为元
	DiscountHigherPrice float64 `json:"discount_higher_price"` // 折扣价区间最大值，单位为元
}
type AdListGetResDataAwemeInfo struct {
	AwemeId     int64  `json:"aweme_id"`
	AwemeName   string `json:"aweme_name"`
	AwemeShowId string `json:"aweme_show_id"`
	AwemeAvatar string `json:"aweme_avatar"`
}
type AdListGetResDataDeliverySetting struct {
	DeepExternalAction string  `json:"deep_external_action"` // 深度转化目标，详见 【附录-枚举值】
	DeepBidType        string  `json:"deep_bid_type"`        // 深度出价方式 仅当深度转化目标"deep_external_action"为 AD_CONVERT_TYPE_LIVE_PAY_ROI 时有效 枚举值： MIN 等同于PC端，转化目标设置为“支付ROI”
	RoiGoal            float64 `json:"roi_goal"`             // 支付ROI目标，最多支持两位小数，0.01～100 注意： 按展示付费(oCPM)，根据 【保障规则】 提供保障福利，请谨慎修改支付ROI目标和定向，以免失去保障资格。
	SmartBidType       string  `json:"smart_bid_type"`       // 投放场景
	ExternalAction     string  `json:"external_action"`      // 转化目标
	Budget             float64 `json:"budget"`               // 预算
	ReviveBudget       float64 `json:"revive_budget"`        // 复活预算
	BudgetMode         string  `json:"budget_mode"`          // 预算类型
	CpaBid             float64 `json:"cpa_bid"`              // 转化出价
	StartTime          string  `json:"start_time"`           // 投放开始时间
	EndTime            string  `json:"end_time"`             // 投放结束时间
}

// GetAdList 获取账户下计划列表（不含创意）
func (client *Client) GetAdList(request *AdListGetReq, response *AdListGetResData) error {
	df := gout.GET(client.url(conf.API_AD_LIST_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.AdListGetReqBase))
	return client.DoRequest(df, response)
}

//--------------------------------------------获取计划详情（含创意信息）--------------------------------------------------------------------

// GetAdDetailReq 获取计划详情（含创意信息）
type GetAdDetailReq struct {
	AccessToken string `json:"access_token"` // 调用/oauth/access_token/生成的token，此token需要用户授权。
	GetAdDetailReqBase
}

type GetAdDetailReqBase struct {
	AdvertiserId int64 `json:"advertiser_id"` // 千川广告账户ID
	AdId         int64 `json:"ad_id"`         // 计划ID
}

type GetAdDetailRes struct {
	Name                          string                                  `json:"name"`                             // 计划名称
	AdCreateTime                  string                                  `json:"ad_create_time,omitempty"`         // 计划创建时间
	AdId                          int64                                   `json:"ad_id"`                            // 计划ID
	FirstIndustryId               int64                                   `json:"first_industry_id,omitempty"`      //创意一级行业ID
	MarketingGoal                 string                                  `json:"marketing_goal,omitempty"`         // 营销目标
	CreativeMaterialMode          string                                  `json:"creative_material_mode,omitempty"` // 创意呈现方式
	Status                        string                                  `json:"status,omitempty"`                 // 计划投放状态,详见【附录-枚举值】
	LabAdType                     string                                  `json:"lab_ad_type,omitempty"`            // 托管计划类型，NOT_LAB_AD：非托管计划，LAB_AD：托管计划
	DynamicCreative               *int64                                  `json:"dynamic_creative,omitempty"`       // 是否启用动态创意，0 关闭、1 开启
	PromotionWay                  string                                  `json:"promotion_way,omitempty"`          // 推广方式
	SecondIndustryId              int64                                   `json:"second_industry_id,omitempty"`     // 创意二级行业ID
	AdKeywords                    []string                                `json:"ad_keywords,omitempty"`            // 创意标签
	CampaignId                    int64                                   `json:"campaign_id,omitempty"`            // 广告组ID（若为托管计划，则返回null）
	IsIntelligent                 *int64                                  `json:"is_intelligent,omitempty"`         // 是否启用智选流量，0 关闭、1 开启
	AdModifyTime                  string                                  `json:"ad_modify_time,omitempty"`         // 计划修改时间
	TrackUrl                      interface{}                             `json:"track_url,omitempty"`
	ThirdIndustryId               int64                                   `json:"third_industry_id,omitempty"`                // 创意三级行业ID
	MarketingScene                string                                  `json:"marketing_scene,omitempty"`                  // 营销场景，FEED 通投广告，SEARCH 搜索广告
	IsHomepageHide                *int64                                  `json:"is_homepage_hide,omitempty"`                 // 抖音主页是否隐藏视频
	CreativeAutoGenerate          *int64                                  `json:"creative_auto_generate,omitempty"`           // 是否开启「生成更多创意」
	OptStatus                     string                                  `json:"opt_status,omitempty"`                       // 计划操作状态,详见【附录-枚举值】
	ProgrammaticCreativeTitleList []AdCreateProgrammaticCreativeTitleList `json:"programmatic_creative_title_list,omitempty"` // 程序化创意标题信息
	ProgrammaticCreativeMediaList []AdCreateProgrammaticCreativeMediaList `json:"programmatic_creative_media_list,omitempty"` // 程序化创意素材信息
	ProgrammaticCreativeCard      *AdCreateProgrammaticCreativeCard       `json:"programmatic_creative_card,omitempty"`       // 程序化创意推广卡片信息
	Audience                      AdDetailAudience                        `json:"audience"`                                   // 定向人群设置
	DeliverySetting               AdDetailDeliverySetting                 `json:"delivery_setting"`                           // 投放设置
	ProductInfo                   []AdListGetResDataProductInfo           `json:"product_info,omitempty"`                     // 商品列表
	Keywords                      []struct {
		Id        int64  `json:"id,omitempty"`         // 关键词id
		WordId    int64  `json:"word_id,omitempty"`    // 词id，不同计划下如果关键词字面相同，词id会相同
		Word      string `json:"word,omitempty"`       // 关键词字面，长度不超过30，一个汉字长度计为1，一个英文字符长度计为0.5，不能包含emoji 当keywords入参时必填
		MatchType string `json:"match_type,omitempty"` // 匹配类型，允许值: PHRASE 短语匹配，EXTENSIVE 广泛匹配，PRECISION 精准匹配 当keywords入参时必填
		Status    string `json:"status,omitempty"`     // 关键词状态 CONFIRM 审核通过且可代入 REJECT 审核拒绝 AUDIT 新建审核中 DELETE 已删除 PAUSED 词暂停

	} `json:"keywords,omitempty"`
	PivativeWords *struct {
		PhraseWords  []string `json:"phrase_words,omitempty"`  // 短语否定词列表
		PreciseWords []string `json:"precise_words,omitempty"` // 精确否定词列表
	} `json:"pivative_words,omitempty"` // 搜索否定词

	AwemeInfo []struct {
		AwemeShowId string `json:"aweme_show_id,omitempty"` // 抖音号，即客户在手机端感知到的抖音号，向客户批量抖音号时请使用该字段
		AwemeAvatar string `json:"aweme_avatar,omitempty"`  // 抖音号头像
		AwemeId     int64  `json:"aweme_id,omitempty"`      // 抖音ID
		AwemeName   string `json:"aweme_name,omitempty"`    // 抖音号昵称
	} `json:"aweme_info,omitempty"` // 计划中关联的抖音号信息
	CreativeList []struct {
		CreativeId         int64  `json:"creative_id,omitempty"`          // 创意ID，程序化创意审核通过后才会生成创意ID
		ImageMode          string `json:"image_mode,omitempty"`           // 创意素材类型
		CreativeCreateTime string `json:"creative_create_time,omitempty"` // 创意创建时间
		CreativeModifyTime string `json:"creative_modify_time,omitempty"` // 创意修改时间
		VideoMaterial      struct {
			Id int64 `json:"id,omitempty"` // 底层数据id，无实际用途（注：非素材ID）
			*AdCreateCustomVideoMaterial
			IsAutoGenerate *int64 `json:"is_auto_generate,omitempty"` // 是否为派生创意标识，1：是，0：不是
		} `json:"video_material,omitempty"` // 视频类型素材
		ImageMaterial struct {
			Id             int64    `json:"id,omitempty"`               // 底层数据id，无实际用途（注：非素材ID）
			ImageIds       []string `json:"image_ids,omitempty"`        // 图片ID列表
			IsAutoGenerate *int64   `json:"is_auto_generate,omitempty"` // 是否为派生创意标识，1：是，0：不是
		} `json:"image_material,omitempty"` // 图片类型素材
		TitleMaterial struct {
			Id int64 `json:"id,omitempty"` // 素材唯一标识
			*AdCreateTitleMaterial
		} `json:"title_material,omitempty"` // 标题类型素材，若选择了抖音号上的视频，不支持修改标题
		PromotionCardMaterial struct {
			Id          int64 `json:"id,omitempty"`           // 素材唯一标识
			ComponentId int64 `json:"component_id,omitempty"` // 组件唯一标识
			*AdCreatePromotionCardMaterial
		} `json:"promotion_card_material,omitempty"` // 推广卡片素材
	} `json:"creative_list,omitempty"` // 创意信息（若为托管计划，则返回空数组）
	RoomInfo []struct {
		AnchorId     int64  `json:"anchor_id,omitempty"`     // 主播ID
		RoomStatus   string `json:"room_status,omitempty"`   // 直播间状态（若未开播，则返回NULL）
		AnchorName   string `json:"anchor_name,omitempty"`   // 主播名称
		RoomTitle    string `json:"room_title,omitempty"`    // 直播间名称（若未开播，则返回NULL）
		AnchorAvatar string `json:"anchor_avatar,omitempty"` // 主播头像
	} `json:"room_info,omitempty"` // 直播间列表
}

type AdDetailAudience struct {
	AdCreateAudience
	InactiveRetargetingTags []struct {
		RetargetingTag int64  `json:"retargeting_tag,omitempty"` // 人群包id
		Name           string `json:"name,omitempty"`            // 人群包名称
		InactiveType   string `json:"inactive_type,omitempty"`   // 失效类型，EXPIRE 人群包过期，TAG_OFFLINE 人群包tag下线，MANUAL_OFFLINE 精选人群包手动下线
	} `json:"inactive_retargeting_tags,omitempty"` // 失效的人群包列表
}

type AdDetailDeliverySetting struct {
	AdCreateDeliverySetting
	ReviveBudget float64 `json:"revive_budget,omitempty"` // 复活预算
	DeepCpaBid   float64 `json:"deep_cpa_bid,omitempty"`
}

// GetAdDetail 获取计划详情（含创意信息）
func (client *Client) GetAdDetail(request *GetAdDetailReq, response *GetAdDetailRes) error {
	df := gout.GET(client.url(conf.API_AD_DETAIL_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.GetAdDetailReqBase))
	return client.DoRequest(df, response)
}

//--------------------------------------------创建计划（含创意生成规则）--------------------------------------------------------------------

// AdCreateReq 创建计划-请求
type AdCreateReq struct {
	AccessToken  string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdCreateBody        // POST请求的data
}

type AdCreateBody struct {
	AdvertiserId    int64                   `json:"advertiser_id"`            // 千川广告主账户id
	MarketingGoal   string                  `json:"marketing_goal"`           // 营销目标，允许值：VIDEO_PROM_GOODS 短视频带货、LIVE_PROM_GOODS 直播带货
	PromotionWay    string                  `json:"promotion_way"`            // 推广方式 ，允许值：STANDARD专业版、SIMPLE极速版
	MarketingScene  string                  `json:"marketing_scene"`          // 营销场景，允许值：FEED 通投广告，SEARCH 搜索广告
	Name            string                  `json:"name"`                     // 计划名称，长度为1-100个字符，其中1个汉字算2位字符。名称不可重复，否则会报错
	CampaignId      int64                   `json:"campaign_id,omitempty"`    // 千川广告组id 注意：当开启计划托管时，不支持
	IsIntelligent   *int64                  `json:"is_intelligent,omitempty"` // 是否启用智选流量，当“营销场景”为“搜索广告”时必填，允许值： 0 关闭、1 开启
	AwemeId         int64                   `json:"aweme_id"`
	ProductIds      []int64                 `json:"product_ids,omitempty"` // 商品id列表，即准备推广的商品列表，可通过【查询店铺商品列表】接口获取名下可推广商品(目前仅支持推一个商品，但需以数组入参)
	LabAdType       string                  `json:"lab_ad_type,omitempty"` // 是否开启计划托管，允许值： NOT_LAB_AD 非托管计划  LAB_AD 托管计划  注意：1. 当营销目标为VIDEO_PROM_GOODS（短视频带货）且推广方式为STANDARD（专业版） 时，必填  2.bind_type（抖音号）为：OFFICIAL或SELF，抖音号关系类型参考【附录-抖音号授权类型】
	DeliverySetting AdCreateDeliverySetting `json:"delivery_setting"`
	Audience        AdCreateAudience        `json:"audience"`
	AdCreateCreative
	Keywords []Keywords `json:"keywords,omitempty"` // 仅搜索广告支持，关键词列表，最多可添加1000个关键词
}

type AdCreateDeliverySetting struct {
	SmartBidType          string  `json:"smart_bid_type"`                     // 投放场景（出价方式），详见【附录-自动出价类型】，允许值：SMART_BID_CUSTOM控成本投放、SMART_BID_CONSERVATIVE 放量投放控成本投放：控制成本，尽量消耗完预算放量投放：接受成本上浮，尽量消耗更多预算
	FlowControlMode       string  `json:"flow_control_mode,omitempty"`        // 投放速度，详见【附录-计划投放速度类型】仅当 smart_bid_type 为SMART_BID_CUSTOM 时需传值，允许值：FLOW_CONTROL_MODE_FAST 尽快投放（默认值）、FLOW_CONTROL_MODE_BALANCE 均匀投放、FLOW_CONTROL_MODE_SMOOTH 优先低成本，对应千川后台「严格控制成本上限」勾选项
	ExternalAction        string  `json:"external_action"`                    // 转化目标短视频带货目的允许值：AD_CONVERT_TYPE_SHOPPING 商品购买、AD_CONVERT_TYPE_QC_FOLLOW_ACTION 粉丝提升、AD_CONVERT_TYPE_QC_MUST_BUY 点赞评论直播带货目的允许值：AD_CONVERT_TYPE_LIVE_ENTER_ACTION 进入直播间、AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION 直播间商品点击、AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION 直播间下单、AD_CONVERT_TYPE_NEW_FOLLOW_ACTION 直播间粉丝提升、AD_CONVERT_TYPE_LIVE_COMMENT_ACTION 直播间评论、AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY直播间成交
	DeepExternalAction    string  `json:"deep_external_action,omitempty"`     // 深度转化目标，对应千川后台「期待同时优化」注意：1. 仅直播带货场景支持2. 当 smart_bid_type 为SMART_BID_CUSTOM 且 flow_control_mode 为 FLOW_CONTROL_MODE_SMOOTH 亦不支持深度转化目标允许值：AD_CONVERT_TYPE_LIVE_SUCCESSORDER_ACTION 直播间下单若不传，则不生效；若传入，则仅当转化目标为AD_CONVERT_TYPE_LIVE_ENTER_ACTION、AD_CONVERT_TYPE_LIVE_CLICK_PRODUCT_ACTION 时生效
	DeepBidType           string  `json:"deep_bid_type,omitempty"`            // 深度出价方式 仅当深度转化目标为 AD_CONVERT_TYPE_LIVE_PAY_ROI 时，必填；否则，填入也会报错 允许值： MIN等同于PC端，转化目标设置为“支付ROI”
	RoiGoal               float64 `json:"roi_goal,omitempty"`                 // 支付ROI目标，最多支持两位小数，0.01～100 仅当转化目标为 AD_CONVERT_TYPE_LIVE_SUCCESSORDER_PAY 且深度转化目标为 AD_CONVERT_TYPE_LIVE_PAY_ROI 且deep_bid_type为MIN时，必填 注意： 按展示付费(oCPM)，根据【保障规则】提供保障福利，请谨慎修改支付ROI目标和定向，以免失去保障资格。
	Budget                float64 `json:"budget"`                             // 预算，最多支持两位小数当预算模式为日预算时，预算范围是300 - 9999999.99；当预算模式为总预算时，预算范围是max(300,投放天数x100) - 9999999.99
	BudgetMode            string  `json:"budget_mode"`                        // 预算类型（创建后不可修改），详见【附录-预算类型】，允许值：BUDGET_MODE_DAY 日预算，BUDGET_MODE_TOTAL 总预算
	CpaBid                float64 `json:"cpa_bid,omitempty"`                  // 转化出价，出价不能大于预算仅当 smart_bid_type 为SMART_BID_CUSTOM 时需传值
	VideoScheduleType     string  `json:"video_schedule_type,omitempty"`      // 短视频投放日期选择方式，仅短视频带货场景需入参，允许值：SCHEDULE_FROM_NOW 从今天起长期投放（总预算模式下不支持）、SCHEDULE_START_END 设置开始和结束日期
	LiveScheduleType      string  `json:"live_schedule_type,omitempty"`       // 直播间投放时段选择方式，仅直播带货场景需入参，允许值：SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段、SCHEDULE_TIME_FIXEDRANGE 固定时长
	StartTime             string  `json:"start_time,omitempty"`               // 投放起始时间，形式如：2017-01-01广告投放起始时间不允许修改。当video_schedule_type为SCHEDULE_START_END 设置开始和结束日期时需传入。当live_schedule_type 为SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段时必填；当 live_schedule_type 为SCHEDULE_TIME_FIXEDRANGE固定时长时不能传入
	EndTime               string  `json:"end_time,omitempty"`                 // 投放结束时间，形式如：2017-01-01结束时间不能比起始时间早。当video_schedule_type为SCHEDULE_START_END 设置开始和结束日期时需传入。当live_schedule_type 为SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段时必填；当 live_schedule_type 为SCHEDULE_TIME_FIXEDRANGE固定时长时不能传入
	ScheduleTime          string  `json:"schedule_time,omitempty"`            // 投放时段，当 live_schedule_type 为SCHEDULE_TIME_WEEKLY_SETTING 时生效默认全时段投放，格式是48*7位字符串，且都是0或1。也就是以半个小时为最小粒度，周一至周日每天分为48个区段，0为不投放，1为投放，不传、全传0、全传1均代表全时段投放。例如：填写"000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000000000000000000000000001111000000000000000000000"，则投放时段为周一到周日的11:30~13:30
	ScheduleFixedRange    int64   `json:"schedule_fixed_range,omitempty"`     // 固定投放时长当 live_schedule_type 为 SCHEDULE_TIME_FIXEDRANGE 时必填；当live_schedule_type 为SCHEDULE_TIME_ALLDAY 全天、SCHEDULE_TIME_WEEKLY_SETTING 指定时间段时不能传入。单位为秒，最小值为1800（0.5小时），最大值为48*1800（24小时），值必须为1800倍数，不然会报错
	EnableAutoPause       *int64  `json:"enable_auto_pause,omitempty"`        // 是否启用超成本自动暂停，允许值： 0 关闭 1 开启 注意：仅托管计划支持
	AutoManageStrategyCmd *int64  `json:"auto_manage_strategy_cmd,omitempty"` // 托管策略，允许值： 0 优先跑量 1 优先成本 注意：仅托管计划支持
	EnableFollowMaterial  *int64  `json:"enable_follow_material,omitempty"`   // 是否优质素材自动同步投放，允许值： 0 关闭 1 开启 注意：仅托管计划支持
}

type AdCreateAudience struct {
	AudienceMode           string   `json:"audience_mode,omitempty"`            //人群定向模式 当promotion_way为STANDARD专业推广，千川策略赋默认值自定义，无需传值 当promotion_way为SIMPLE极速推广，需入参，允许值：AUTO智能推荐、CUSTOM自定义
	OrientationId          int64    `json:"orientation_id,omitempty"`           // 定向包id 注意： 1、仅专业推广支持，极速推广不支持 2、若传入，则表示使用定向包 3、一个定向包最多支持同时应用至1500个计划（不包括已删除计划） 4、若该定向包包含失效人群包（过期、标签下线、精选人群下线）则创建计划失败
	ExcludeLimitedRegion   *int64   `json:"exclude_limited_region,omitempty"`   // 排除限运地区，允许值： 0：否，默认值 1：是 注： 1、仅同时满足以下条件时，设置为“1”才有效： - 营销目标为短视频带货 - 地域定向类型为“不限”/地域定向的用户状态类型为“正在该地区的用户” 2、当“可放开定向列表”为REGION且排除限运地区时，依旧会探索限运地区的目标人群
	District               string   `json:"district,omitempty"`                 // 地域定向类型，配合 city 字段使用，允许值：CITY 省市， COUNTY 区县， NONE 不限默认值为NONE
	City                   []int64  `json:"city,omitempty"`                     // 具体定向的城市列表，当 district 为COUNTY，CITY为必填，枚举值详见【附件-city.json】省市的传法："city" : [12], "district" : "CITY"区县的传法："city" : [130102], "district" : "COUNTY"
	LocationType           string   `json:"location_type,omitempty"`            // 地域定向的用户状态类型，当 district 为COUNTY，CITY为必填，允许值：CURRENT 正在该地区的用户、HOME 居住在该地区的用户、TRAVEL 到该地区旅行的用户、ALL 该地区内的所有用户
	Gender                 string   `json:"gender,omitempty"`                   // 性别，允许值：GENDER_FEMALE 女性， GENDER_MALE 男性，NONE 不限
	Age                    []string `json:"age,omitempty"`                      // 年龄，详见【附录-受众年龄区间】，允许值：AGE_BETWEEN_18_23, AGE_BETWEEN_24_30、AGE_BETWEEN_31_40、AGE_BETWEEN_41_49、AGE_ABOVE_50
	AwemeFanBehaviors      []string `json:"aweme_fan_behaviors,omitempty"`      // 抖音用户行为类型，详见【附录-抖音达人互动用户行为类型】
	AwemeFanBehaviorsDays  string   `json:"aweme_fan_behaviors_days,omitempty"` // 抖音达人互动用户行为天数
	AwemeFanCategories     []int64  `json:"aweme_fan_categories,omitempty"`     // 抖音达人分类ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向），可通过【工具-抖音达人-查询抖音类目列表】接口获取
	AwemeFanAccounts       []int64  `json:"aweme_fan_accounts,omitempty"`       // 抖音达人ID列表，与aweme_fan_behaviors同时设置才会生效（抖音达人定向），可通过【工具-抖音达人-查询抖音类目下的推荐达人】接口获取
	AutoExtendEnabled      *int64   `json:"auto_extend_enabled,omitempty"`      // 是否启用智能放量，允许值：0 关闭、1 开启
	AutoExtendTargets      []string `json:"auto_extend_targets,omitempty"`      // 可放开定向列表。当auto_extend_enabled=1 时必填。允许值：AGE 年龄、REGION 地域、GENDER 性别、INTEREST_ACTION 行为兴趣 、CUSTOM_AUDIENCE 更多人群-自定义人群
	Platform               []string `json:"platform,omitempty"`                 // 投放平台列表，允许值：ANDROID、 IOS、不传值为全选
	SmartInterestAction    string   `json:"smart_interest_action,omitempty"`    // 行为兴趣意向定向模式，允许值：RECOMMEND系统推荐，CUSTOM 自定义；不传值则为不限制需要注意：如果设置RECOMMEND，则传入action_scene、action_days、action_categories、action_words、 interest_categories、interest_words字段都无效
	ActionScene            []string `json:"action_scene,omitempty"`             // 行为场景，详见【附录-行为场景】，smart_interest_actionCUSTOM时有效，允许值：E-COMMERCE 电商互动行为、NEWS 资讯互动行为、APP APP推广互动行为
	ActionDays             int64    `json:"action_days,omitempty"`              // 用户发生行为天数，当 smart_interest_action 传 CUSTOM 时有效允许值：7, 15, 30, 60, 90, 180, 365
	ActionCategories       []int64  `json:"action_categories,omitempty"`        // 行为类目词，当 smart_interest_action 传 CUSTOM 时有效行为类目可以通过【工具-行为兴趣词管理-行为类目查询】获取
	ActionWords            []int64  `json:"action_words,omitempty"`             // 行为关键词，当 smart_interest_action 传 CUSTOM 时有效行为关键词可以通过【工具-行为兴趣词管理-行为关键词查询】获取
	InterestCategories     []int64  `json:"interest_categories,omitempty"`      // 兴趣类目词，当 smart_interest_action 传 CUSTOM 时有效兴趣类目可以通过【工具-行为兴趣词管理-兴趣类目查询】获取
	InterestWords          []int64  `json:"interest_words,omitempty"`           // 兴趣关键词，当 smart_interest_action 传 CUSTOM 时有效行为关键词可以通过【工具-行为兴趣词管理-行为关键词查询】获取
	Ac                     []string `json:"ac,omitempty"`                       // 网络类型, 详见【附录-受众网络类型】，允许值:WIFI、2G、3G、4G。 不传值或全传为全选
	RetargetingTagsInclude []int64  `json:"retargeting_tags_include,omitempty"` // 定向人群包id列表，长度限制 0-200。定向人群包可以通过【工具-DMP人群管理-获取人群包列表】获取
	RetargetingTagsExclude []int64  `json:"retargeting_tags_exclude,omitempty"` // 排除人群包id列表，长度限制 0-200。排除人群包可以通过【工具-DMP人群管理-获取人群包列表】获取
	LivePlatformTags       []string `json:"live_platform_tags,omitempty"`       // 直播带货平台精选人群包，当marketing_goal=LIVE_PROM_GOODS时有效，默认为全不选。允许值：LARGE_FANSCOUNT 高关注人群、ABNORMAL_ACTIVE高活跃人群、AWEME_FANS抖音号粉丝
}

type AdCreateCreative struct {
	CreativeMaterialMode          string                                  `json:"creative_material_mode"`                     // 创意呈现方式，允许值：CUSTOM_CREATIVE 自定义创意、PROGRAMMATIC_CREATIVE 程序化创意
	FirstIndustryId               int64                                   `json:"first_industry_id,omitempty"`                // 创意一级行业ID。可从【获取行业列表】接口获取
	SecondIndustryId              int64                                   `json:"second_industry_id,omitempty"`               // 创意二级行业ID。可从【获取行业列表】接口获取
	ThirdIndustryId               int64                                   `json:"third_industry_id,omitempty"`                // 创意三级行业ID。可从【获取行业列表】接口获取
	AdKeywords                    []string                                `json:"ad_keywords,omitempty"`                      // 创意标签。最多20个标签，且每个标签长度要求为1~20个字符，汉字算2个字符
	CreativeList                  []AdCreateCreativeList                  `json:"creative_list,omitempty"`                    // 自定义素材信息
	CreativeAutoGenerate          *int64                                  `json:"creative_auto_generate,omitempty"`           // 是否开启「生成更多创意」
	ProgrammaticCreativeMediaList []AdCreateProgrammaticCreativeMediaList `json:"programmatic_creative_media_list,omitempty"` // 程序化创意素材信息
	ProgrammaticCreativeTitleList []AdCreateProgrammaticCreativeTitleList `json:"programmatic_creative_title_list,omitempty"` // 程序化创意标题信息
	ProgrammaticCreativeCard      *AdCreateProgrammaticCreativeCard       `json:"programmatic_creative_card,omitempty"`       // 程序化创意推广卡片信息
	IsHomepageHide                *int64                                  `json:"is_homepage_hide,omitempty"`                 // 抖音主页是否隐藏视频
	DynamicCreative               *int64                                  `json:"dynamic_creative,omitempty"`                 // 是否启用动态创意，允许值：0 关闭、1 开启 当“营销场景”为“搜索广告”时必填 当“营销场景”为“通投广告”时，不支持传该字段，否则会报错
}

// AdCreateCreativeList 广告创意 - creative_list
type AdCreateCreativeList struct {
	ImageMode             string                         `json:"image_mode,omitempty"`              // 创意素材类型
	VideoMaterial         *AdCreateCustomVideoMaterial   `json:"video_material,omitempty"`          // 视频类型素材
	ImageMaterial         *AdCreateImageMaterial         `json:"image_material,omitempty"`          // 图片类型素材
	TitleMaterial         *AdCreateTitleMaterial         `json:"title_material,omitempty"`          // 标题类型素材，若选择了抖音号上的视频，不支持修改标题
	PromotionCardMaterial *AdCreatePromotionCardMaterial `json:"promotion_card_material,omitempty"` // 推广卡片素材
}

// AdCreateCustomVideoMaterial 广告创意 - 视频类型素材
type AdCreateCustomVideoMaterial struct {
	VideoId      string `json:"video_id,omitempty"`       // 视频ID
	VideoCoverId string `json:"video_cover_id,omitempty"` // 视频封面ID
	AwemeItemId  int64  `json:"aweme_item_id,omitempty"`  // 抖音视频ID
}

// AdCreateImageMaterial 广告创意 - 图片类型素材
type AdCreateImageMaterial struct {
	ImageIds []string `json:"image_ids,omitempty"` // 图片ID列表
}

// AdCreateTitleMaterial 广告创意 - 标题类型素材，若选择了抖音号上的视频，不支持修改标题
type AdCreateTitleMaterial struct {
	Title        string                 `json:"title,omitempty"`         // 创意标题
	DynamicWords []AdCreateDynamicWords `json:"dynamic_words,omitempty"` // 动态词包对象列表
}

type AdCreateDynamicWords struct {
	WordId      int64  `json:"word_id,omitempty"`      // 动态词包ID
	DictName    string `json:"dict_name,omitempty"`    // 创意词包名称
	DefaultWord string `json:"default_word,omitempty"` // 创意词包默认词
}

// AdCreatePromotionCardMaterial 广告创意 - 推广卡片素材
type AdCreatePromotionCardMaterial struct {
	Title                   string   `json:"title,omitempty"`                     // 推广卡片标题
	SellingPoints           []string `json:"selling_points,omitempty"`            // 推广卡片卖点列表
	ImageId                 string   `json:"image_id,omitempty"`                  // 推广卡片配图
	ActionButton            string   `json:"action_button,omitempty"`             // 推广卡片行动号召按钮文案
	ButtonSmartOptimization int64    `json:"button_smart_optimization,omitempty"` // 是否对行动号召按钮文案启用智能优选
}

// AdCreateProgrammaticCreativeMediaList 广告创意 - 程序化创意素材信息
type AdCreateProgrammaticCreativeMediaList struct {
	ImageMode      string   `json:"image_mode,omitempty"`          // 创意素材类型，支持视频和图片
	VideoId        string   `json:"video_id,omitempty"`            // 视频ID
	VideoCoverId   string   `json:"video_cover_id,omitempty"`      // 视频封面ID
	ImageIds       []string `json:"image_ids,omitempty,omitempty"` // 图片ID列表
	IsAutoGenerate int64    `json:"is_auto_generate,omitempty"`    // 是否为派生创意标识，1：是，0：不是
}

// AdCreateProgrammaticCreativeTitleList 广告创意 - 程序化创意标题信息
type AdCreateProgrammaticCreativeTitleList struct {
	Title        string                 `json:"title,omitempty"`         // 创意标题
	DynamicWords []AdCreateDynamicWords `json:"dynamic_words,omitempty"` // 动态词包对象列表
}

// AdCreateProgrammaticCreativeCard 广告创意 - 程序化创意推广卡片信息
type AdCreateProgrammaticCreativeCard struct {
	PromotionCardTitle                   string   `json:"promotion_card_title,omitempty"`                     // 推广卡片标题，最多7个字
	PromotionCardSellingPoints           []string `json:"promotion_card_selling_points,omitempty"`            // 推广卡片卖点列表，卖点文字长度要求为12~18个字符，汉字算2个字符
	PromotionCardImageId                 string   `json:"promotion_card_image_id,omitempty"`                  // 推广卡片配图，可通过【获取图片素材】接口获得图片素材id
	PromotionCardActionButton            string   `json:"promotion_card_action_button,omitempty"`             // 推广卡片行动号召按钮文案
	PromotionCardButtonSmartOptimization int64    `json:"promotion_card_button_smart_optimization,omitempty"` // 是否对行动号召按钮文案启用智能优选
}

type Keywords struct {
	Word      string `json:"word,omitempty"`       // 关键词字面，长度不超过30，一个汉字长度计为1，一个英文字符长度计为0.5，不能包含emoji 当keywords入参时必填
	MatchType string `json:"match_type,omitempty"` // 匹配类型，允许值: PHRASE 短语匹配，EXTENSIVE 广泛匹配，PRECISION 精准匹配 当keywords入参时必填
}

type AdCreateResData struct {
	AdId int64 `json:"ad_id"` // 创建的计划id
}

// CreateAd 创建计划（含创意生成规则）
func (client *Client) CreateAd(request *AdCreateReq, response *AdCreateResData) error {
	df := gout.POST(client.url(conf.API_AD_CREATE)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetJSON(request.AdCreateBody)
	return client.DoRequest(df, response)
}
