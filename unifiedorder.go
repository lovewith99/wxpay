package wxpay

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"
)

// UnifiedOrder 统一下单
// document: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
type UnifiedOrder struct {
	// XMLName struct{} `xml:"xml"`
	// AppId string `xml:"appid"`  // 公众账号id
	// MchId string `xml:"mch_id"` // 商户号
	// Sign  string `xml:"sign"`   // 签名
	Request

	DeviceInfo     string `xml:"device_info,omitempty"`      // 设备号
	Body           string `xml:"body"`                       // 商品描述
	Detail         string `xml:"detail,omitempty,CDATA"`     // 商品详情
	Attach         string `xml:"attach,omitempty"`           // 附加数据
	OutTradeNo     string `xml:"out_trade_no"`               // 商户订单号
	FeeType        string `xml:"fee_type,omitempty"`         // 标价币种
	TotalFee       int    `xml:"total_fee"`                  // 订单总金额,单位分
	SpbillCreateIp string `xml:"spbill_create_ip"`           // 终端IP
	TimeStart      string `xml:"time_start,omitempty"`       // 交易起始时间: 20091225091010
	TimeExpire     string `xml:"time_expire,omitempty"`      // 交易结束时间: 20091227091010
	GoodsTag       string `xml:"goods_tag,omitempty"`        // 订单优惠标记
	NotifyUrl      string `xml:"notify_url"`                 // 通知地址
	TradeType      string `xml:"trade_type"`                 // 交易类型	JSAPI: 公众号支付 NATIVE: 扫码支付 APP: APP支付
	ProductId      string `xml:"product_id,omitempty"`       // 商品ID
	LimitPay       string `xml:"limit_pay,omitempty"`        // 指定支付方式,上传此参数no_credit--可限制用户不能使用信用卡支付
	OpenId         string `xml:"openid,omitempty"`           // 用户标识
	SceneInfo      string `xml:"scene_info,omitempty,CDATA"` // 场景信息
}

func (w *UnifiedOrder) GateWay() string {
	return "https://api.mch.weixin.qq.com/pay/unifiedorder"
}

func (w *UnifiedOrder) SetSceneInfo(id, name, areaCode, address string) error {
	// 设置统一支付场景信息
	var sceneInfo = make(map[string]string)

	if id != "" {
		sceneInfo["id"] = id
	}

	if name != "" {
		sceneInfo["name"] = name
	}

	if areaCode != "" {
		sceneInfo["area_code"] = areaCode
	}

	if address != "" {
		sceneInfo["address"] = address
	}

	b, err := json.Marshal(sceneInfo)
	if err != nil {
		return err
	}

	w.SceneInfo = string(b)
	return nil
}

type UnifiedOrderResp struct {
	ReturnCode string `xml:"return_code,CDATA"`
	ReturnMsg  string `xml:"return_msg,omitempty,CDATA"`

	// 仅当return_code 为SUCCESS的时候有返回
	AppId      string `xml:"appid,omitempty,CDATA"`
	MchId      string `xml:"mch_id,omitempty,CDATA"`
	DeviceInfo string `xml:"device_info,omitempty,CDATA"`
	NonceStr   string `xml:"nonce_str,omitempty,CDATA"`
	Sign       string `xml:"sign,omitempty,CDATA"`
	ResultCode string `xml:"result_code,omitempty,CDATA"`
	ErrCode    string `xml:"err_code,omitempty,CDATA"`
	ErrCodeDes string `xml:"err_code_des,omitempty,CDATA"`

	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	TradeType string `xml:"trade_type,omitempty,CDATA"`
	PrepayId  string `xml:"prepay_id,omitempty,CDATA"`
	CodeUrl   string `xml:"code_url,omitempty,CDATA"`
}

func (w *UnifiedOrderResp) IsSuccess() bool {
	return (w.ReturnCode == "SUCCESS" && w.ResultCode == "SUCCESS")
}

// RequestData app统一下单支付参数 https://pay.weixin.qq.com/wiki/doc/api/app/app.php?chapter=9_12
// 公众号支付 https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=7_7&index=6
func (w *UnifiedOrderResp) RequestData(cli *Client) map[string]interface{} {
	var hm map[string]interface{}
	switch w.TradeType {
	case "APP":
		hm = map[string]interface{}{
			"appid":     w.AppId,
			"partnerid": w.MchId,
			"prepayid":  w.PrepayId,
			"package":   "Sign=WXPay",
			"noncestr":  w.NonceStr,
			"timestamp": strconv.FormatInt(time.Now().Unix(), 10),
		}
	case "JSAPI":
		hm = map[string]interface{}{
			"appId":     w.AppId,
			"timeStamp": strconv.FormatInt(time.Now().Unix(), 10),
			"nonceStr":  w.NonceStr,
			"package":   fmt.Sprintf("prepay_id=%s", w.PrepayId),
			"signType":  MD5,
		}
	}

	sign := MakeSign(hm, cli.key, MD5)

	if w.TradeType == "APP" {
		return map[string]interface{}{
			"appId":        w.AppId,
			"partnerId":    w.MchId,
			"prepayId":     w.PrepayId,
			"packageValue": "Sign=WXPay",
			"nonceStr":     w.NonceStr,
			"timeStamp":    hm["timestamp"],
			"sign":         sign,
		}
	}

	if w.TradeType == "JSAPI" {
		return map[string]interface{}{
			"appId":     w.AppId,
			"timeStamp": hm["timeStamp"],
			"nonceStr":  w.NonceStr,
			"package":   hm["package"],
			"signType":  hm["signType"],
			"paySign":   sign,
		}
	}
	return nil
}
