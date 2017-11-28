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
	ReturnCode string `xml:"return_code,cdata"`
	ReturnMsg string `xml:"return_msg,omitempty,cdata"`

	//以下字段在return_code为SUCCESS的时候有返回
	AppId string `xml:"appid,omitempty,cdata"`
	MchId string `xml:"mch_id,omitempty,cdata"`
	NonceStr string `xml:"nonce_str,omitempty,cdata"`
	Sign string `xml:"sign,omitempty,cdata"`
	ResultCode string `xml:"result_code,omitempty,cdata"`
	ErrCode string `xml:"err_code,omitempty,cdata"`
	ErrCodeDes string `xml:"err_code_des,omitempty,cdata"`

	// 以下字段在return_code 、result_code、trade_state都为SUCCESS时有返回
	// ，如trade_state不为 SUCCESS，则只返回out_trade_no（必传）和attach（选传）
	DeviceInfo string `xml:"device_info,omitempty,cdata"`
	OpenId string `xml:"openid,omitempty,cdata"`
	IsSubscribe string `xml:"is_subscribe,omitempty,cdata"`
	TradeType string `xml:"trade_type,omitempty,cdata"`
	TradeState string `xml:"trade_state,omitempty,cdata"`
	BankType string `xml:"bank_type,omitempty,cdata"`
	TotalFee int `xml:"total_fee,omitempty,cdata"`
	SettlementTotalFee int `xml:"settlement_total_fee,omitempty,cdata"`
	FeeType string `xml:"fee_type,omitempty,cdata"`
	CashFee int `xml:"cash_fee,omitempty,cdata"`
	CashFeeType string `xml:"cash_fee_type,omitempty,cdata"`
	CouponFee int `xml:"coupon_fee,omitempty,cdata"`
	CouponCount int `xml:"coupon_count,omitempty,cdata"`
	CouponTypeN string `xml:"coupon_type_$n,omitempty,cdata"`
	CouponIdN  string `xml:"coupon_id_$n,omitempty,cdata"`
	CouponFeeN int `xml:"coupon_fee_$n,omitempty,cdata"`
	TransactionId string `xml:"transaction_id,omitempty,cdata"`
	OutTradeNo string `xml:"out_trade_no,omitempty,cdata"`
	Attach string `xml:"attach,omitempty,cdata"`
	TimeEnd string `xml:"time_end,omitempty,cdata"`
	TradeStateDesc string `xml:"trade_state_desc,omitempty,cdata"`
}

func (w *WxPayOrderQueryResp) IsSuccess() bool  {
	return (w.ReturnCode == "SUCCESS" && w.ResultCode == "SUCCESS")
}