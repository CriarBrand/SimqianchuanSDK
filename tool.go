package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimqianchuanSDK/conf"
	"github.com/CriarBrand/SimqianchuanSDK/utils"
	"github.com/guonaihong/gout"
)

// -----------------------------------------------------查询抖音类目下的推荐达人----------------------------------------------

type GetAwemeCategoryTopAuthorReqBase struct {
	AdvertiserId int64    `json:"advertiser_id"`
	CategoryId   int64    `json:"category_id,omitempty"` // 类目id，一级，二级，三级类目id均可
	Behaviors    []string `json:"behaviors,omitempty"`   // 抖音用户行为类型，详见【附录-抖音达人互动用户行为类型】 默认为空,仅影响覆盖人群数
}

type GetAwemeCategoryTopAuthorReq struct {
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	GetAwemeCategoryTopAuthorReqBase
}

type GetAwemeCategoryTopAuthorRes struct {
	Authors []AwemeCategoryTopAuthor `json:"authors"`
}

type AwemeCategoryTopAuthor struct {
	AuthorName      string `json:"author_name"`        // 抖音作者名
	TotalFansNumStr string `json:"total_fans_num_str"` // 粉丝数
	CoverNumStr     string `json:"cover_num_str"`      // 覆盖人群数
	LabelId         int64  `json:"label_id"`           // 抖音号id
	AwemeId         string `json:"aweme_id"`           // 抖音id
	Avatar          string `json:"avatar"`             // 抖音头像
	CategoryName    string `json:"category_name"`      // 抖音分类
}

// GetAwemeCategoryTopAuthor 查询抖音类目下的推荐达人
func (client *Client) GetAwemeCategoryTopAuthor(request *GetAwemeCategoryTopAuthorReq, response *GetAwemeCategoryTopAuthorRes) error {
	df := gout.GET(client.url(conf.API_TOOLS_AWEME_CATEGORY_TOP_AUTHOR_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.GetAwemeCategoryTopAuthorReqBase))
	return client.DoRequest(df, response)
}

// -----------------------------------------------------查询抖音类目列表----------------------------------------------

type GetAwemeMultiLevelCategoryReqBase struct {
	AdvertiserId int64    `json:"advertiser_id"`
	Behaviors    []string `json:"behaviors,omitempty"` // 抖音用户行为类型，详见【附录-抖音达人互动用户行为类型】 默认为空,仅影响覆盖人群数
}

type GetAwemeMultiLevelCategoryReq struct {
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	GetAwemeMultiLevelCategoryReqBase
}

type GetAwemeMultiLevelCategoryRes struct {
	Categories []AwemeMultiLevelCategory `json:"categories"`
}

type AwemeMultiLevelCategory struct {
	Id          int64  `json:"id"`
	CoverNumStr string `json:"cover_num_str"`
	FansNumStr  string `json:"fans_num_str"`
	Value       string `json:"value"`
	Children    []struct {
		Id          int64  `json:"id"`
		CoverNumStr string `json:"cover_num_str"`
		FansNumStr  string `json:"fans_num_str"`
		Value       string `json:"value"`
		Children    []struct {
			Id          int64  `json:"id"`
			CoverNumStr string `json:"cover_num_str"`
			FansNumStr  string `json:"fans_num_str"`
			Value       string `json:"value"`
		} `json:"children"`
	} `json:"children"`
}

// GetAwemeMultiLevelCategory 查询抖音类目列表
func (client *Client) GetAwemeMultiLevelCategory(request *GetAwemeMultiLevelCategoryReq, response *GetAwemeMultiLevelCategoryRes) error {
	df := gout.GET(client.url(conf.API_TOOLS_AWEME_MULTI_LEVEL_CATEGORY_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.GetAwemeMultiLevelCategoryReqBase))
	return client.DoRequest(df, response)
}

// -----------------------------------------------------行为类目查询----------------------------------------------

type GetInterestActionActionCategoryReqBase struct {
	AdvertiserId int64    `json:"advertiser_id"`
	ActionScene  []string `json:"action_scene"` // 行为场景，详见【附录-行为场景】 允许值: "E-COMMERCE","NEWS","APP"
	ActionDays   int64    `json:"action_days"`  // 行为天数,默认值: 7、15、30、60、90、180、365
}

type GetInterestActionActionCategoryReq struct {
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	GetInterestActionActionCategoryReqBase
}

type GetInterestActionActionCategoryRes struct {
	Num      string `json:"num"`
	Children []struct {
		Num      string `json:"num"`
		Children []struct {
			Num      string `json:"num"`
			Children []struct {
				Num  string `json:"num"`
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"children,omitempty"`
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"children,omitempty"`
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"children"`
	Name string `json:"name"`
	Id   string `json:"id"`
}

// GetInterestActionActionCategory 行为类目查询
func (client *Client) GetInterestActionActionCategory(request *GetInterestActionActionCategoryReq, response *[]GetInterestActionActionCategoryRes) error {
	df := gout.GET(client.url(conf.API_INTEREST_ACTION_ACTION_CATEGORY)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.GetInterestActionActionCategoryReqBase))
	return client.DoRequest(df, response)
}

// -----------------------------------------------------行为关键词查询----------------------------------------------

type GetInterestActionActionKeywordReqBase struct {
	AdvertiserId int64    `json:"advertiser_id"`
	ActionScene  []string `json:"action_scene"` // 行为场景，详见【附录-行为场景】 允许值: "E-COMMERCE","NEWS","APP"
	ActionDays   int64    `json:"action_days"`  // 行为天数,默认值: 7、15、30、60、90、180、365
	QueryWords   string   `json:"query_words"`  // 关键词
}

type GetInterestActionActionKeywordReq struct {
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	GetInterestActionActionKeywordReqBase
}

type GetInterestActionActionKeywordRes struct {
	List []InterestActionActionKeyword `json:"list"` // 词包列表
}

type InterestActionActionKeyword struct {
	Num  string `json:"num"`  // 关键词数目
	Id   string `json:"id"`   // 关键词id
	Name string `json:"name"` // 关键词名称
}

// GetInterestActionActionKeyword 行为关键词查询
func (client *Client) GetInterestActionActionKeyword(request *GetInterestActionActionKeywordReq, response *GetInterestActionActionKeywordRes) error {
	df := gout.GET(client.url(conf.API_TOOLS_INTEREST_ACTION_ACTION_KEYWORD)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.GetInterestActionActionKeywordReqBase))
	return client.DoRequest(df, response)
}

// -----------------------------------------------------兴趣类目查询----------------------------------------------

type GetInterestActionInterestCategoryReq struct {
	AccessToken  string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	AdvertiserId int64  `json:"advertiser_id"`
}

type GetInterestActionInterestCategoryRes struct {
	Num      string `json:"num"`
	Children []struct {
		Num      string `json:"num"`
		Children []struct {
			Num      string `json:"num"`
			Children []struct {
				Num  string `json:"num"`
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"children,omitempty"`
			Name string `json:"name"`
			Id   string `json:"id"`
		} `json:"children,omitempty"`
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"children"`
	Name string `json:"name"`
	Id   string `json:"id"`
}

// GetInterestActionInterestCategory 兴趣类目查询
func (client *Client) GetInterestActionInterestCategory(request *GetInterestActionInterestCategoryReq, response *[]GetInterestActionInterestCategoryRes) error {
	df := gout.GET(client.url(conf.API_TOOLS_INTEREST_ACTION_INTEREST_CATEGORY)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(gout.H{
			"advertiser_id": request.AdvertiserId,
		})
	return client.DoRequest(df, response)
}

// -----------------------------------------------------兴趣关键词查询----------------------------------------------

type GetInterestActionInterestKeywordReqBase struct {
	AdvertiserId int64  `json:"advertiser_id"`
	QueryWords   string `json:"query_words"` // 关键词
}

type GetInterestActionInterestKeywordReq struct {
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	GetInterestActionInterestKeywordReqBase
}

type GetInterestActionInterestKeywordRes struct {
	List []InterestActionInterestKeyword `json:"list"` // 词包列表
}

type InterestActionInterestKeyword struct {
	Num  string `json:"num"`  // 关键词数目
	Id   string `json:"id"`   // 关键词id
	Name string `json:"name"` // 关键词名称
}

// GetInterestActionInterestKeyword 兴趣关键词查询
func (client *Client) GetInterestActionInterestKeyword(request *GetInterestActionInterestKeywordReq, response *GetInterestActionInterestKeywordRes) error {
	df := gout.GET(client.url(conf.API_TOOLS_INTEREST_ACTION_INTEREST_KEYWORD)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.GetInterestActionInterestKeywordReqBase))
	return client.DoRequest(df, response)
}

// -----------------------------------------------------查询创编可用人群----------------------------------------------

type GetDmpAudiencesReqBase struct {
	AdvertiserId        int64 `json:"advertiser_id"`
	RetargetingTagsType int64 `json:"retargeting_tags_type"` // 人群包类型，枚举值：0：不限营销目标的平台精选人群包，1：自定义人群包
	Offset              int64 `json:"offset,omitempty"`      // 偏移,类似于SQL中offset(起始为0,翻页时new_offset=old_offset+limit），默认值：0，取值范围:≥ 0
	Limit               int64 `json:"limit,omitempty"`       // 返回数据量，默认值：100，取值范围：1-100
}

type GetDmpAudiencesReq struct {
	AccessToken string // 调用/oauth/access_token/生成的token，此token需要用户授权。
	GetDmpAudiencesReqBase
}

type GetDmpAudiencesRes struct {
	Offset          int64                         `json:"offset"`           // 下一次查询的偏移,类似于SQL中offset(起始为0,翻页时new_offset=old_offset+limit），返回0时，代表已查询到最后一页
	TotalNum        int64                         `json:"total_num"`        // 总的人群包数量
	RetargetingTags []DmpAudiencesRetargetingTags `json:"retargeting_tags"` // 人群包列表
}

type DmpAudiencesRetargetingTags struct {
	HasOfflineTag      int64       `json:"has_offline_tag"`      // 是否包含已下线标签，0 不包含，1 包含
	RetargetingTagsId  int64       `json:"retargeting_tags_id"`  // 人群包id
	IsCommon           int64       `json:"is_common"`            // 0 该人群包不支持通投，1 该人群包支持通投，注意：不支持通投的人群包不能在千川平台创建计划，否则会报错。
	Name               string      `json:"name"`                 // 人群包名称
	RetargetingTagsTip string      `json:"retargeting_tags_tip"` // 人群包说明
	RetargetingTagsOp  string      `json:"retargeting_tags_op"`  // 人群包可选的定向规则，枚举值：INCLUDE只支持定向，EXCLUDE只支持排除，ALL支持两种规则。 当source为RETARGETING_TAGS_TYPE_PLATFORM时，只支持INCLUDE或EXCLUDE；当source为RETARGETING_TAGS_TYPE_CUSTOM时，支持ALL
	Status             int64       `json:"status"`               // 人群包状态，详见【附录-DMP相关-人群包状态】
	Source             interface{} `json:"source"`               // 人群包来源，自定义类详见【附录-DMP相关-人群包来源】，平台精选类返回空值
	CoverNum           int64       `json:"cover_num"`            // 预估人群包覆盖人群数目
}

// GetDmpAudiences 查询创编可用人群
func (client *Client) GetDmpAudiences(request *GetDmpAudiencesReq, response *GetDmpAudiencesRes) error {
	df := gout.GET(client.url(conf.API_DMP_AUDIENCES_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(utils.BuildQueryToMap(request.GetDmpAudiencesReqBase))
	return client.DoRequest(df, response)
}
