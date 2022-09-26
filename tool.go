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
