package wxpay

// 关闭订单
// document: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
type CloseOrderReq struct {
	Request

	OutTradeNo string `xml:"out_trade_no"`
}

func (w *CloseOrderReq) GateWay() string {
	return "https://api.mch.weixin.qq.com/pay/closeorder"
}

type CloseOrderResp struct {
	ReturnCode string `xml:"return_code,CDATA"`
	ReturnMsg  string `xml:"return_msg,omitempty,CDATA"`

	// 以下字段在return_code为SUCCESS的时候有返回
	AppId      string `xml:"appid,omitempty,CDATA"`
	MuchId     string `xml:"mch_id,omitempty,CDATA"`
	NonceStr   string `xml:"nonce_str,omitempty,CDATA"`
	Sign       string `xml:"sign,omitempty,CDATA"`
	ResultCode string `xml:"result_code,omitempty,CDATA"`
	ResultMsg  string `xml:"result_msg,omitempty,CDATA"`
	ErrCode    string `xml:"err_code,omitempty,CDATA"`
	ErrCodeDes string `xml:"err_code_des,omitempty,CDATA"`
}

func (w *CloseOrderResp) IsSuccess() bool {
	return (w.ReturnCode == "SUCCESS" && w.ResultCode == "SUCCESS")
}
