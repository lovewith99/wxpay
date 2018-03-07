package wxpay

// 微信支付结果通知(支付回调)
// 文档： https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_7&index=3
type WxAppPayNotification struct {
	// 微信App支付回调
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`

	// required
	AppId         string `xml:"appid"`
	MchId         string `xml:"mch_id"`
	NonceStr      string `xml:"nonce_str"`
	Sign          string `xml:"sign"`
	ResultCode    string `xml:"result_code"`
	Openid        string `xml:"openid"`
	TradeType     string `xml:"trade_type"`
	BankType      string `xml:"bank_type"`
	TotalFee      string `xml:"total_fee"`
	CashFee       string `xml:"cash_fee"`
	TransactionId string `xml:"transaction_id"`
	OutTradeNo    string `xml:"out_trade_no"`
	TimeEnd       string `xml:"time_end"`

	DeviceInfo  string `xml:"device_info"`
	ErrCode     string `xml:"err_code"`
	ErrCodeDes  string `xml:"err_code_des"`
	IsSubscribe string `xml:"is_subscribe"`
	FeeType     string `xml:"fee_type"`
	CashFeeType string `xml:"cash_fee_type"`
	//CouponFee string `xml:"coupon_fee"`
	//CouponCount string `xml:"coupon_count"`
	//CouponId_$n string `xml:"coupon_id_$n"`
	//CouponFee_$n string `xml:"coupon_fee_$n"`
	Attach string `xml:"attach"`
}

