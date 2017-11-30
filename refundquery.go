package wxpay

// 查询退款
// document: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_5
type WxPayRefundQuery struct {
	XMLName struct{} `xml:"xml"`

	AppId string `xml:"appid"`
	MchId string `xml:"mch_id"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`
	SignType string `xml:"sign_type,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty"`
	OutTradeNo string `xml:"out_trade_no,omitempty"`
	OutRefundNo string `xml:"out_refund_no,omitempty"`
	RefundId string `xml:"refund_id,omitempty"`
	Offset int `xml:"offset,omitempty"`
}

func (w *WxPayRefundQuery) GateWay() string {
	return "https://api.mch.weixin.qq.com/pay/refundquery"
}

func (w *WxPayRefundQuery) setAppId(appId string)  {
	w.AppId = appId
}

func (w *WxPayRefundQuery) setMchId(mchId string)  {
	w.MchId = mchId
}

func (w *WxPayRefundQuery) setSign(sign string)  {
	w.Sign = sign
}

func (w *WxPayRefundQuery) SingStr() string {
	p := ReflectStruct(*w)

	return signStr(p)
}

func (w *WxPayRefundQuery) signType() string {
	return w.SignType
}

type WxPayRefundQueryResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg,omitempty"`

	// 以下字段在return_code为SUCCESS的时候有返回
	ResultCode string `xml:"result_code,omitempty"`
	ErrCode string `xml:"err_code,omitempty"`
	ErrCodeDes string `xml:"err_code_des,omitempty"`
	AppId string `xml:"appid,omitempty"`
	MchId string `xml:"mch_id,omitempty"`
	NonceStr string `xml:"nonce_str,omitempty"`
	Sign string `xml:"sign,omitempty"`
	TotalRefundCount int `xml:"total_refund_count,omitempty"`
	TransactionId string `xml:"transaction_id,omitempty"`
	OutTradeNo string `xml:"out_trade_no,omitempty"`
	TotalFee int `xml:"total_fee,omitempty"`
	SettlementTotalFee int `xml:"settlement_total_fee,omitempty"`
	FeeType string `xml:"fee_type,omitempty"`
	CashFee int `xml:"cash_type,omitempty"`
	RefundCount int `xml:"refund_count,omitempty"`
	OutRefundNoN string `xml:"out_refund_no_$n,omitempty"`
	RefundIdN string `xml:"refund_id_$n,omitempty"`
	RefundChannelN string `xml:"refund_channel_$n,omitempty"`
	RefundFeeN int `xml:"refund_fee_$n,omitempty"`
	SettlementRefundFeeN int `xml:"settlement_refund_fee_$n,omitempty"`
	CouponTypeNM string `xml:"coupon_type_$n_$m,omitempty"`
	CouponRefundFeeN int `xml:"coupon_refund_fee_$n,omitempty"`
	CouponRefundCountN int `xml:"coupon_refund_count_$n,omitempty"`
	CouponRefundIdNM string `xml:"coupon_refund_id_$n_$m,omitempty"`
	CouponRefundFeeNM string `xml:"coupon_refund_fee_$n_$m,omitempty"`
	RefundStatusN string `xml:"refund_status_$n,omitempty"`
	RefundAccountN string `xml:"refund_account_$n,omitempty"`
	RefundRecvAccountN string `xml:"refund_recv_account_$n,omitempty"`
	RefundSuccessTimeN string `xml:"refund_success_time_$n,omitempty"`
}

func (w *WxPayRefundQueryResp) IsSuccess() bool {
	return (w.ReturnCode == "SUCCESS" && w.ResultCode == "SUCCESS")
}
