package wxpay

// 微信支付结果通知(支付回调)
// 文档： https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_7&index=3
type WxAppPayNotification struct {
	// 微信App支付回调
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg,omitempty"`

	// required
	AppId         string `xml:"appid"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code,omitempty"`
	Openid        string `xml:"openid"`
	TradeType     string `xml:"trade_type,omitempty"`
	BankType      string `xml:"bank_type,omitempty"`
	TotalFee      string `xml:"total_fee"`
	CashFee       string `xml:"cash_fee"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	TimeEnd       string `xml:"time_end"`

	DeviceInfo  string `xml:"device_info,omitempty"`
	ErrCode     string `xml:"err_code,omitempty"`
	ErrCodeDes  string `xml:"err_code_des,omitempty"`
	IsSubscribe string `xml:"is_subscribe,omitempty"`
	FeeType     string `xml:"fee_type,omitempty"`
	CashFeeType string `xml:"cash_fee_type,omitempty"`

	// CouponFee   string `xml:"coupon_fee"`
	// CouponCount string `xml:"coupon_count"`
	// CouponID   string `xml:"coupon_id_$n"`
	// CouponFee  string `xml:"coupon_fee_$n"`
	Attach string `xml:"attach"`
}
