package wxpay

// 关闭订单
// document: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_3
type WxPayCloseOrder struct {
	XMLName struct{} `xml:"xml"`

	AppId string `xml:"appid"`
	MchId string `xml:"mch_id"`
	OutTradeNo string `xml:"out_trade_no"`
	NonceStr string `xml:"nonce_str"`
	Sign string `xml:"sign"`
	SignType string `xml:"sign_type,omitempty"`
}

func (w *WxPayCloseOrder) GateWay() string {
	return "https://api.mch.weixin.qq.com/pay/closeorder"
}

func (w *WxPayCloseOrder) setAppId(appId string)  {
	w.AppId = appId
}

func (w *WxPayCloseOrder) setMchId(mchId string)  {
	w.MchId = mchId
}

func (w *WxPayCloseOrder) setSign(sign string) {
	w.Sign = sign
}

func (w *WxPayCloseOrder) SignStr() string {
	p := ReflectStruct(*w)

	return signStr(p)
}

func (w *WxPayCloseOrder) signType() string {
	return w.SignType
}

type WxPayCloseOrderResp struct {
	ReturnCode string `xml:"return_code,cdata"`
	ReturnMsg string `xml:"return_msg,omitempty,cdata"`

	// 以下字段在return_code为SUCCESS的时候有返回
	AppId string `xml:"appid,omitempty,cdata"`
	MuchId string `xml:"mch_id,omitempty,cdata"`
	NonceStr string `xml:"nonce_str,omitempty,cdata"`
	Sign string `xml:"sign,omitempty,cdata"`
	ResultCode string `xml:"result_code,omitempty,cdata"`
	ResultMsg string `xml:"result_msg,omitempty,cdata"`
	ErrCode string `xml:"err_code,omitempty,cdata"`
	ErrCodeDes string `xml:"err_code_des,omitempty,cdata"`
}

func (w *WxPayCloseOrderResp) IsSuccess() bool {
	return (w.ReturnCode == "SUCCESS" && w.ResultCode == "SUCCESS")
}
