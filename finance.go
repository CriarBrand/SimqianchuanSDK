package SimqianchuanSDK

import (
	"github.com/CriarBrand/SimqianchuanSDK/conf"
	"github.com/guonaihong/gout"
)

type FinanceWalletReq struct {
	AccessToken  string `json:"access_token"`
	AdvertiserId int64  `json:"advertiser_id"`
}

type FinanceWalletResData struct {
	TotalBalanceAbs                 int64 `json:"total_balance_abs"`                   // 账户总余额
	GrantBalance                    int64 `json:"grant_balance"`                       // 赠款余额
	UnionValidGrantBalance          int64 `json:"union_valid_grant_balance"`           // 赠款余额-穿山甲-可用
	SearchValidGrantBalance         int64 `json:"search_valid_grant_balance"`          // 赠款余额-巨量搜索广告-可用
	CommonValidGrantBalance         int64 `json:"common_valid_grant_balance"`          // 赠款余额-巨量信息流广告-可用
	DefaultValidGrantBalance        int64 `json:"default_valid_grant_balance"`         // 赠款余额-通用-可用
	GeneralTotalBalance             int64 `json:"general_total_balance"`               // 通用余额
	GeneralBalanceValid             int64 `json:"general_balance_valid"`               // 通用余额-可用余额
	GeneralBalanceValidNonGrant     int64 `json:"general_balance_valid_non_grant"`     // 通用余额-可用余额-非赠款
	GeneralBalanceValidGrantUnion   int64 `json:"general_balance_valid_grant_union"`   // 通用余额-可用余额-赠款-穿山甲
	GeneralBalanceValidGrantSearch  int64 `json:"general_balance_valid_grant_search"`  // 通用余额-可用余额-赠款-巨量搜索广告
	GeneralBalanceValidGrantCommon  int64 `json:"general_balance_valid_grant_common"`  // 通用余额-可用余额-赠款-巨量信息流广告
	GeneralBalanceValidGrantDefault int64 `json:"general_balance_valid_grant_default"` // 通用余额-可用余额-赠款-通用
	GeneralBalanceInvalid           int64 `json:"general_balance_invalid"`             // 通用余额-不可用余额
	GeneralBalanceInvalidOrder      int64 `json:"general_balance_invalid_order"`       // 通用余额-不可用余额-随心推已下单
	GeneralBalanceInvalidFrozen     int64 `json:"general_balance_invalid_frozen"`      // 通用余额-不可用余额-冻结
	BrandBalance                    int64 `json:"brand_balance"`                       // 品牌余额
	BrandBalanceValid               int64 `json:"brand_balance_valid"`                 // 品牌余额-可用余额
	BrandBalanceValidNonGrant       int64 `json:"brand_balance_valid_non_grant"`       // 品牌余额-可用余额-非赠款
	BrandBalanceValidGrant          int64 `json:"brand_balance_valid_grant"`           // 品牌余额-可用余额-赠款
	BrandBalanceInvalid             int64 `json:"brand_balance_invalid"`               // 品牌余额-不可用余额
	BrandBalanceInvalidFrozen       int64 `json:"brand_balance_invalid_frozen"`        // 品牌余额-不可用余额-冻结
	DeductionCouponBalance          int64 `json:"deduction_coupon_balance"`            // 消返红包余额
	DeductionCouponBalanceAll       int64 `json:"deduction_coupon_balance_all"`        // 消返红包余额（通用）
	DeductionCouponBalanceOther     int64 `json:"deduction_coupon_balance_other"`      // 消返红包余额（代投）
	DeductionCouponBalanceSelf      int64 `json:"deduction_coupon_balance_self"`       // 消返红包余额（自投）
	GrantExpiring                   int64 `json:"grant_expiring"`                      // 15天内赠款到期金额
	ShareBalance                    int64 `json:"share_balance"`                       // 共享赠款余额
	ShareBalanceValidGrantUnion     int64 `json:"share_balance_valid_grant_union"`     // 共享赠款余额-可用余额-赠款-穿山甲
	ShareBalanceValidGrantSearch    int64 `json:"share_balance_valid_grant_search"`    // 共享赠款余额-可用余额-赠款-巨量搜索广告
	ShareBalanceValidGrantCommon    int64 `json:"share_balance_valid_grant_common"`    // 共享赠款余额-可用余额-赠款-巨量信息流广告
	ShareBalanceValidGrantDefault   int64 `json:"share_balance_valid_grant_default"`   // 共享赠款余额-可用余额-赠款-通用
	ShareBalanceValid               int64 `json:"share_balance_valid"`                 // 共享赠款余额-可用余额
	ShareBalanceExpiring            int64 `json:"share_balance_expiring"`              // 共享赠款余额-30天内到期余额
	ShareExpiringDetailList         []struct {
		Category   string `json:"category"`    // 类别，允许值： CONFIRM 站内信息流及其他 DEFAULT 通用 SEARCH 站内搜索 UNION 网盟穿山甲
		Amount     int64  `json:"amount"`      // 金额
		ExpireTime int64  `json:"expire_time"` // 到期时间

	} `json:"share_expiring_detail_list"`
}

// GetFinanceWallet 获取账户钱包信息
func (client *Client) GetFinanceWallet(request *FinanceWalletReq, response *FinanceWalletResData) error {
	df := gout.GET(client.url(conf.API_WALLET_FINANCE_GET)).
		SetHeader(gout.H{
			"Access-Token": request.AccessToken,
		}).
		SetQuery(gout.H{
			"advertiser_id": request.AdvertiserId,
		})
	return client.DoRequest(df, response)
}
