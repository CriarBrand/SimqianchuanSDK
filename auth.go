package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimqianchuanSDK/conf"
	"github.com/guonaihong/gout"
)

// ------------------------------------------------------获取token---------------------------------------------------------------

// AccessTokenResData access_token返回
type AccessTokenResData struct {
	AccessToken           string `json:"access_token"`             // 用于验证权限的token
	ExpiresIn             uint64 `json:"expires_in"`               // access_token剩余有效时间,单位(秒)
	RefreshToken          string `json:"refresh_token"`            // 刷新access_token,用于获取新的access_token和refresh_token，并且刷新过期时间
	RefreshTokenExpiresIn uint64 `json:"refresh_token_expires_in"` // refresh_token剩余有效时间,单位(秒)
}

type accessTokenBody struct {
	AppId     int64  `json:"app_id"`
	Secret    string `json:"secret"`
	GrantType string `json:"grant_type"`
	AuthCode  string `json:"auth_code"`
}

func (client *Client) GetAccessToken(authCode string, response *AccessTokenResData) error {
	df := gout.POST(client.url(conf.API_OAUTH_ACCESS_TOKEN)).
		SetJSON(accessTokenBody{
			AppId:     client.appId,
			Secret:    client.secret,
			GrantType: "auth_code",
			AuthCode:  authCode,
		})
	return client.DoRequest(df, response)
}

// ------------------------------------------------------刷新token---------------------------------------------------------------

// RefreshTokenResData 刷新access_token返回
type RefreshTokenResData AccessTokenResData

type refreshTokenBody struct {
	AppId        int64  `json:"app_id"`
	Secret       string `json:"secret"`
	GrantType    string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

func (client *Client) RefreshToken(refreshToken string, response *RefreshTokenResData) error {
	df := gout.POST(client.url(conf.API_OAUTH_REFRESH_TOKEN)).
		SetJSON(refreshTokenBody{
			AppId:        client.appId,
			Secret:       client.secret,
			GrantType:    "refresh_token",
			RefreshToken: refreshToken,
		})
	return client.DoRequest(df, response)
}
