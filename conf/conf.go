package conf

const (

	// API_HOST OpenAPI HOST
	API_HOST = "ad.oceanengine.com"

	// API_HTTP_SCHEME 协议
	API_HTTP_SCHEME = "https://"

	// API_OAUTH_CONNECT 生成授权链接
	API_OAUTH_CONNECT = "/openapi/qc/audit/oauth.html"

	// API_OAUTH_ACCESS_TOKEN 获取access_token
	API_OAUTH_ACCESS_TOKEN = "/open_api/oauth2/access_token/"

	// API_OAUTH_REFRESH_TOKEN 刷新access_token
	API_OAUTH_REFRESH_TOKEN = "/open_api/oauth2/refresh_token/"

	// API_ADVERTISER_LIST 获取已授权的账户（店铺/代理商）
	API_ADVERTISER_LIST = "/open_api/oauth2/advertiser/get/"

	// API_ADVERTISER_AD_QUOTA 获取在投计划配额信息
	API_ADVERTISER_AD_QUOTA = "/open_api/v1.0/qianchuan/ad/quota/get/"

	// API_SHOP_ADVERTISER_LIST 获取店铺账户关联的广告账户列表
	API_SHOP_ADVERTISER_LIST = "/open_api/v1.0/qianchuan/shop/advertiser/list/"

	// API_AGENT_ADVERTISER_LIST 获取代理商账户关联的广告账户列表
	API_AGENT_ADVERTISER_LIST = "/open_api/2/agent/advertiser/select/"

	// API_USER_INFO 获取授权时登录用户信息
	API_USER_INFO = "/open_api/2/user/info/"

	// API_SHOP_ACCOUNT_INFO 获取店铺账户信息
	API_SHOP_ACCOUNT_INFO = "/open_api/v1.0/qianchuan/shop/get/"

	// API_AGENT_INFO 获取代理商账户信息
	API_AGENT_INFO = "/open_api/2/agent/info/"

	// API_ADVERTISER_PUBLIC_INFO 获取千川广告账户基础信息
	API_ADVERTISER_PUBLIC_INFO = "/open_api/2/advertiser/public_info/"

	// API_ADVERTISER_INFO 获取千川广告账户全量信息
	API_ADVERTISER_INFO = "/open_api/2/advertiser/info/"

	// API_REPORT_ADVERTISER_GET 获取广告账户数据
	API_REPORT_ADVERTISER_GET = "/open_api/v1.0/qianchuan/report/advertiser/get/"

	// API_LQ_AD 获取低效计划
	API_LQ_AD = "/open_api/v1.0/qianchuan/lq_ad/get/"

	// API_REPORT_AD_GET 获取广告计划数据
	API_REPORT_AD_GET = "/open_api/v1.0/qianchuan/report/ad/get/"

	// API_REPORT_CREATIVE_GET 获取广告创意数据
	API_REPORT_CREATIVE_GET = "/open_api/v1.0/qianchuan/report/creative/get/"

	// API_CAMPAIGN_CREATE 广告组创建
	API_CAMPAIGN_CREATE = "/open_api/v1.0/qianchuan/campaign/create/"

	// API_CAMPAIGN_UPDATE 广告组更新
	API_CAMPAIGN_UPDATE = "/open_api/v1.0/qianchuan/campaign/update/"

	// API_BATCH_CAMPAIGN_STATUS_UPDATE 广告组状态更新
	API_BATCH_CAMPAIGN_STATUS_UPDATE = "/open_api/v1.0/qianchuan/batch_campaign_status/update/"

	// API_CAMPAIGN_LIST_GET 广告组列表获取
	API_CAMPAIGN_LIST_GET = "/open_api/v1.0/qianchuan/campaign_list/get/"

	// API_AD_CREATE 创建计划（含创意生成规则）
	API_AD_CREATE = "/open_api/v1.0/qianchuan/ad/create/"

	// API_AD_UPDATE 更新计划（含创意生成规则）
	API_AD_UPDATE = "/open_api/v1.0/qianchuan/ad/update/"

	// API_AD_STATUS_UPDATE 更新计划状态
	API_AD_STATUS_UPDATE = "/open_api/v1.0/qianchuan/ad/status/update/"

	// API_AD_BUDGET_UPDATE 更新计划预算
	API_AD_BUDGET_UPDATE = "/open_api/v1.0/qianchuan/ad/budget/update/"

	// API_AD_BID_UPDATE 更新计划出价
	API_AD_BID_UPDATE = "/open_api/v1.0/qianchuan/ad/bid/update/"

	// API_AD_DETAIL_GET 获取计划详情（含创意信息）
	API_AD_DETAIL_GET = "/open_api/v1.0/qianchuan/ad/detail/get/"

	// API_AD_LIST_GET 获取账户下计划列表（不含创意）
	API_AD_LIST_GET = "/open_api/v1.0/qianchuan/ad/get/"

	// API_AD_REJECT_REASON 获取计划审核建议
	API_AD_REJECT_REASON = "/open_api/v1.0/qianchuan/ad/reject_reason/"

	// API_AD_ROI_GOAL_UPDATE 更新计划的支付ROI目标
	API_AD_ROI_GOAL_UPDATE = "/open_api/v1.0/qianchuan/roi/goal/update"

	// API_CREATIVE_STATUS_UPDATE 更新创意状态
	API_CREATIVE_STATUS_UPDATE = "/open_api/v1.0/qianchuan/creative/status/update/"

	// API_CREATIVE_GET 获取账户下创意列表
	API_CREATIVE_GET = "/open_api/v1.0/qianchuan/creative/get/"

	// API_CREATIVE_REJECT_REASON 获取创意审核建议
	API_CREATIVE_REJECT_REASON = "/open_api/v1.0/qianchuan/creative/reject_reason/"

	// API_PRODUCT_AVAILABLE_GET 获取可投商品列表接口
	API_PRODUCT_AVAILABLE_GET = "/open_api/v1.0/qianchuan/product/available/get/"

	// API_AWEME_AUTHORIZED_GET 获取千川账户下已授权抖音号
	API_AWEME_AUTHORIZED_GET = "/open_api/v1.0/qianchuan/aweme/authorized/get/"

	// API_FILE_IMAGE_AD 上传图片素材
	API_FILE_IMAGE_AD = "/open_api/2/file/image/ad/"

	// API_FILE_VIDEO_AD 上传视频素材
	API_FILE_VIDEO_AD = "/open_api/2/file/video/ad/"

	// API_FILE_IMAGE_GET 获取素材库的图片
	API_FILE_IMAGE_GET = "/open_api/2/file/image/get/"

	// API_FILE_VIDEO_GET 获取素材库的视频
	API_FILE_VIDEO_GET = "/open_api/2/file/video/get/"

	// API_FILE_VIDEO_AWEME_GET 获取抖音号下的视频
	API_FILE_VIDEO_AWEME_GET = "/open_api/v1.0/qianchuan/file/video/aweme/get/"

	// API_TOOLS_INDUSTRY_GET 获取行业列表
	API_TOOLS_INDUSTRY_GET = "/open_api/2/tools/industry/get/"

	// API_TOOLS_AWEME_CATEGORY_TOP_AUTHOR_GET 查询抖音类目下的推荐达人
	API_TOOLS_AWEME_CATEGORY_TOP_AUTHOR_GET = "/open_api/2/tools/aweme_category_top_author/get/"

	// API_TOOLS_AWEME_MULTI_LEVEL_CATEGORY_GET 查询抖音类目列表
	API_TOOLS_AWEME_MULTI_LEVEL_CATEGORY_GET = "/open_api/2/tools/aweme_multi_level_category/get/"

	// API_INTEREST_ACTION_ACTION_CATEGORY 行为类目查询
	API_INTEREST_ACTION_ACTION_CATEGORY = "/open_api/2/tools/interest_action/action/category/"

	// API_TOOLS_INTEREST_ACTION_ACTION_KEYWORD 行为关键词查询
	API_TOOLS_INTEREST_ACTION_ACTION_KEYWORD = "/open_api/2/tools/interest_action/action/keyword/"

	// API_TOOLS_INTEREST_ACTION_INTEREST_CATEGORY 兴趣类目查询
	API_TOOLS_INTEREST_ACTION_INTEREST_CATEGORY = "/open_api/2/tools/interest_action/interest/category/"

	// API_TOOLS_INTEREST_ACTION_INTEREST_KEYWORD 兴趣关键词查询
	API_TOOLS_INTEREST_ACTION_INTEREST_KEYWORD = "/open_api/2/tools/interest_action/interest/keyword/"

	// API_TOOLS_CREATIVE_WORD_SELECT 查询动态创意词包
	API_TOOLS_CREATIVE_WORD_SELECT = "/open_api/2/tools/creative_word/select/"

	// API_DMP_AUDIENCES_GET 查询创编可用人群
	API_DMP_AUDIENCES_GET = "/open_api/v1.0/qianchuan/dmp/audiences/get/"

	// API_WALLET_FINANCE_GET 获取账户钱包信息
	API_WALLET_FINANCE_GET = "/open_api/v1.0/qianchuan/finance/wallet/get/"

	//---------------千川巨量--------------------

	//查询抖音号id对应的达人信息
	API_TOOLS_AWEME_AUTHOR_INFO_GET = "/open_api/2/tools/aweme_author_info/get/"
)
