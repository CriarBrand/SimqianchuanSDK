package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimqianchuanSDK/conf"
	"github.com/guonaihong/gout"
)

type AuthorizedAwemeReq struct {
	AccessToken  string `json:"access_token"`
	AdvertiserId int64  `json:"advertiser_id"`
	Page         int64  `json:"page,omitempty"`      // 页码，默认为1
	PageSize     int64  `json:"page_size,omitempty"` // 页面大小，默认值：10，最大值：100
}

type AuthorizedAwemeResData struct {
	AwemeIdList []AwemeIdListDetail `json:"aweme_id_list"`
	PageInfo    PageInfo            `json:"page_info"`
}

type AwemeIdListDetail struct {
	AwemeAvatar string   `json:"aweme_avatar"`
	AwemeId     int64    `json:"aweme_id"`
	AwemeShowId string   `json:"aweme_show_id"`
	AwemeName   string   `json:"aweme_name"`
	AwemeStatus string   `json:"aweme_status"`
	BindType    []string `json:"bind_type"`
}

// GetAuthorizedAweme 获取千川账户下已授权抖音号
func (client *Client) GetAuthorizedAweme(request *AuthorizedAwemeReq, response *AuthorizedAwemeResData) error {
	df := gout.GET(client.url(conf.API_AWEME_AUTHORIZED_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(gout.H{
			"advertiser_id": request.AdvertiserId,
		})
	return client.DoRequest(df, response)
}
