package wxpay

import "reflect"

// 微信支付结果通知(支付回调)
// 文档： https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_7&index=3
type AppPayNotification struct {
	// 微信App支付回调
	ReturnCode string `xml:"return_code,CDATA"`
	ReturnMsg  string `xml:"return_msg,omitempty,CDATA"`

	// required
	AppId         string `xml:"appid,CDATA"`
	MchId         string `xml:"mch_id,CDATA"`
	NonceStr      string `xml:"nonce_str,CDATA"`
	Sign          string `xml:"sign,CDATA"`
	ResultCode    string `xml:"result_code,omitempty,CDATA"`
	Openid        string `xml:"openid,CDATA"`
	TradeType     string `xml:"trade_type,omitempty,CDATA"`
	BankType      string `xml:"bank_type,omitempty,CDATA"`
	TotalFee      string `xml:"total_fee,CDATA"`
	CashFee       string `xml:"cash_fee,CDATA"`
	TransactionId string `xml:"transaction_id,CDATA"`
	OutTradeNo    string `xml:"out_trade_no,CDATA"`
	TimeEnd       string `xml:"time_end,CDATA"`

	DeviceInfo  string `xml:"device_info,omitempty,CDATA"`
	ErrCode     string `xml:"err_code,omitempty,CDATA"`
	ErrCodeDes  string `xml:"err_code_des,omitempty,CDATA"`
	IsSubscribe string `xml:"is_subscribe,omitempty,CDATA"`
	FeeType     string `xml:"fee_type,omitempty,CDATA"`
	CashFeeType string `xml:"cash_fee_type,omitempty,CDATA"`

	// CouponFee   string `xml:"coupon_fee"`
	// CouponCount string `xml:"coupon_count"`
	// CouponID   string `xml:"coupon_id_$n"`
	// CouponFee  string `xml:"coupon_fee_$n"`
	Attach string `xml:"attach,CDATA"`
}

func (v *AppPayNotification) IsSuccess() bool {
	return v.ReturnCode == "SUCCESS"
}

func (c *AppPayNotification) VerifySign(key string) bool {
	hm := make(map[string]interface{})
	parseStruct(reflect.ValueOf(c), hm)

	return MakeSign(hm, key, "") == c.Sign
}
