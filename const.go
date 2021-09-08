package wxpay

// 交易类型
const (
	// 统一下单
	JSAPI  = "JSAPI"  // 公众号支付
	NATIVE = "NATIVE" // 原生扫码支付
	APP    = "APP"    // app支付

	// 刷卡支付有单独的支付接口，不调用统一下单接口
	MICROPAY = "MICROPAY" // 刷卡支付
)

// 货币类型
const (
	CNY = "CNY" // 人民币
)

// 银行类型
const (
	ICBC_DEBIT = "ICBC_DEBIT" // 工商银行(借记卡)
)

// 签名类型
const (
	MD5         = "MD5"
	HMAC_SHA256 = "HMAC-SHA256"
)
