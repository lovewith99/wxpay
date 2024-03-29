package wxpay

// 申请退款
// document: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_4
type RefundReq struct {
	Request

	TransactionId string `xml:"transaction_id,omitempty"`
	OutTradeNo    string `xml:"out_trade_no,omitempty"`
	OutRefundNo   string `xml:"out_refund_no,omitempty"`
	TotalFee      int    `xml:"total_fee"`
	RefundFee     int    `xml:"refund_fee"`
	RefundFeeType string `xml:"refund_fee_type,omitempty"`
	RefundDesc    string `xml:"refund_desc,omitempty"`
	RefundAccount string `xml:"refund_account,omitempty"`
}

func (w *RefundReq) GateWay() string {
	return "https://api.mch.weixin.qq.com/secapi/pay/refund"
}

type RefundResp struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg,omitempty"`

	// 以下字段在return_code为SUCCESS的时候有返回
	ResultCode          string `xml:"result_code,omitempty"`
	ErrCode             string `xml:"err_code,omitempty"`
	ErrCodeDes          string `xml:"err_code_des,omitempty"`
	AppId               string `xml:"appid,omitempty"`
	MchId               string `xml:"mch_id,omitempty"`
	NonceStr            string `xml:"nonce_str,omitempty"`
	Sign                string `xml:"sign,omitempty"`
	TransactionId       string `xml:"transaction_id,omitempty"`
	OutTradeNo          string `xml:"out_trade_no,omitempty"`
	OutRefundNo         string `xml:"out_refund_no,omitempty"`
	RefundId            string `xml:"refund_id,omitempty"`
	RefundFee           int    `xml:"refund_fee,omitempty"`
	SettlementRefundFee int    `xml:"settlement_refund_fee,omitempty"`
	TotalFee            int    `xml:"total_fee,omitempty"`
	SettlementTotalFee  int    `xml:"settlement_total_fee,omitempty"`
	FeeType             string `xml:"fee_type,omitempty"`
	CashFee             int    `xml:"cash_fee,omitempty"`
	CashFeeType         string `xml:"cash_fee_type,omitempty"`
	CashRefundFee       int    `xml:"cash_refund_fee,omitempty"`
	CouponTypeN         string `xml:"coupon_type_$n,omitempty"`
	CouponRefundFee     int    `xml:"coupon_refund_fee,omitempty"`
	CouponRefundFeeN    int    `xml:"coupon_refund_fee_$n,omitempty"`
	CouponRefundCount   int    `xml:"coupon_refund_count,omitempty"`
	CouponRefundIdN     string `xml:"coupon_refund_id_$n,omitempty"`
}

func (w *RefundResp) IsSuccess() bool {
	return (w.ReturnCode == "SUCCESS" && w.ResultCode == "SUCCESS")
}
