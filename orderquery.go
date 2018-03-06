package wxpay

// 订单查询
// document: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_2
type WxPayOrderQuery struct {
	XMLName struct{}  `xml:"xml"`

	AppId string `xml:"appid"`
	MchId string `xml:"mch_id"`

	TransactionId string `xml:"transaction_id,omitempty"`
	OutTradeNo string `xml:"out_trade_no,omitempty"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`
	SignType string `xml:"sign_type,omitempty"`
}

func (w *WxPayOrderQuery) GateWay() string {
	return "https://api.mch.weixin.qq.com/pay/orderquery"
}

func (w *WxPayOrderQuery) setAppId(appId string)  {
	w.AppId = appId
}

func (w *WxPayOrderQuery) setMchId(mchId string)  {
	w.MchId = mchId
}

func (w *WxPayOrderQuery) setSign(sign string)  {
	w.Sign = sign
}

func (w *WxPayOrderQuery) SignStr() string {
	p := ReflectStruct(*w)

	return signStr(p)
}

func (w *WxPayOrderQuery) signType() string {
	return w.SignType
}

type WxPayOrderQueryResp struct {
	ReturnCode string `xml:"return_code,CDATA"`
	ReturnMsg string `xml:"return_msg,omitempty,CDATA"`

	//以下字段在return_code为SUCCESS的时候有返回
	AppId string `xml:"appid,omitempty,CDATA"`
	MchId string `xml:"mch_id,omitempty,CDATA"`
	NonceStr string `xml:"nonce_str,omitempty,CDATA"`
	Sign string `xml:"sign,omitempty,CDATA"`
	ResultCode string `xml:"result_code,omitempty,CDATA"`
	ErrCode string `xml:"err_code,omitempty,CDATA"`
	ErrCodeDes string `xml:"err_code_des,omitempty,CDATA"`

	// 以下字段在return_code 、result_code、trade_state都为SUCCESS时有返回
	// ，如trade_state不为 SUCCESS，则只返回out_trade_no（必传）和attach（选传）
	DeviceInfo string `xml:"device_info,omitempty,CDATA"`
	OpenId string `xml:"openid,omitempty,CDATA"`
	IsSubscribe string `xml:"is_subscribe,omitempty,CDATA"`
	TradeType string `xml:"trade_type,omitempty,CDATA"`
	TradeState string `xml:"trade_state,omitempty,CDATA"`
	BankType string `xml:"bank_type,omitempty,CDATA"`
	TotalFee int `xml:"total_fee,omitempty,CDATA"`
	SettlementTotalFee int `xml:"settlement_total_fee,omitempty,CDATA"`
	FeeType string `xml:"fee_type,omitempty,CDATA"`
	CashFee int `xml:"cash_fee,omitempty,CDATA"`
	CashFeeType string `xml:"cash_fee_type,omitempty,CDATA"`
	CouponFee int `xml:"coupon_fee,omitempty,CDATA"`
	CouponCount int `xml:"coupon_count,omitempty,CDATA"`
	CouponTypeN string `xml:"coupon_type_$n,omitempty,CDATA"`
	CouponIdN  string `xml:"coupon_id_$n,omitempty,CDATA"`
	CouponFeeN int `xml:"coupon_fee_$n,omitempty,CDATA"`
	TransactionId string `xml:"transaction_id,omitempty,CDATA"`
	OutTradeNo string `xml:"out_trade_no,omitempty,CDATA"`
	Attach string `xml:"attach,omitempty,CDATA"`
	TimeEnd string `xml:"time_end,omitempty,CDATA"`
	TradeStateDesc string `xml:"trade_state_desc,omitempty,CDATA"`
}

func (w *WxPayOrderQueryResp) IsSuccess() bool  {
	return (w.ReturnCode == "SUCCESS" && w.ResultCode == "SUCCESS")
}