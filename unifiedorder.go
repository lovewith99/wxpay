package wxpay

import (
	"encoding/json"
)

// 统一下单
// document: https://pay.weixin.qq.com/wiki/doc/api/jsapi.php?chapter=9_1
type WxPayUnifiedOrder struct {
	XMLName   struct{} `xml:"xml"`

	AppId string `xml:"appid"`  // 公众账号id
	MchId string `xml:"mch_id"` // 商户号
	Sign string `xml:"sign"`      // 签名

	DeviceInfo string `xml:"device_info,omitempty"` // 设备号
	NonceStr string `xml:"nonce_str"` // 随机字符串
	SignType string `xml:"sign_type,omitempty"` // 签名类型
	Body string `xml:"body"`          // 商品描述
	Detail string `xml:"detail,omitempty,CDATA"` // 商品详情
	Attach string `xml:"attach,omitempty"` // 附加数据
	OutTradeNo string `xml:"out_trade_no"` // 商户订单号
	FeeType string `xml:"fee_type,omitempty"` // 标价币种
	TotalFee int `xml:"total_fee"`       // 订单总金额,单位分
	SpbillCreateIp string `xml:"spbill_create_ip"` // 终端IP
	TimeStart string `xml:"time_start,omitempty"` // 交易起始时间: 20091225091010
	TimeExpire string `xml:"time_expire,omitempty"` // 交易结束时间: 20091227091010
	GoodsTag string `xml:"goods_tag,omitempty"` // 订单优惠标记
	NotifyUrl string `xml:"notify_url"` // 通知地址
	TradeType string `xml:"trade_type"` // 交易类型
	ProductId string `xml:"product_id,omitempty"` // 商品ID
	LimitPay string `xml:"limit_pay,omitempty"` // 指定支付方式,上传此参数no_credit--可限制用户不能使用信用卡支付
	OpenId string `xml:"openid,omitempty"` // 用户标识
	SceneInfo string `xml:"scene_info,omitempty,CDATA"` // 场景信息
}

func (w *WxPayUnifiedOrder) GateWay() string {
	return "https://api.mch.weixin.qq.com/pay/unifiedorder"
}

func (w *WxPayUnifiedOrder) setAppId(appId string)  {
	w.AppId = appId
}

func (w *WxPayUnifiedOrder) setMchId(mchId string)  {
	w.MchId = mchId
}

func (w *WxPayUnifiedOrder) setSign(sign string)  {
	w.Sign = sign
}

func (w *WxPayUnifiedOrder) SignStr() string {
	w.NonceStr = RandString(32)
	p := ReflectStruct(*w)

	return signStr(p)
}

func (w *WxPayUnifiedOrder) signType() string  {
	return w.SignType
}

func (w *WxPayUnifiedOrder) SetSceneInfo(id, name, areaCode, address string) error {
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

	w.SceneInfo = Bytes2Str(b)
	return nil
}

type WxPayUnifiedOrderResp struct {
	ReturnCode string `xml:"return_code,CDATA"`
	ReturnMsg string `xml:"return_msg,omitempty,CDATA"`

	// 仅当return_code 为SUCCESS的时候有返回
	AppId string `xml:"appid,omitempty,CDATA"`
	MchId string `xml:"mch_id,omitempty,CDATA"`
	DeviceInfo string `xml:"device_info,omitempty,CDATA"`
	NonceStr string `xml:"nonce_str,omitempty,CDATA"`
	Sign string `xml:"sign,omitempty,CDATA"`
	ResultCode string `xml:"result_code,omitempty,CDATA"`
	ErrCode string `xml:"err_code,omitempty,CDATA"`
	ErrCodeDes string `xml:"err_code_des,omitempty,CDATA"`

	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	TradeType string `xml:"trade_type,omitempty,CDATA"`
	PrepayId string `xml:"prepay_id,omitempty,CDATA"`
	CodeUrl string `xml:"code_url,omitempty,CDATA"`
}

func (w *WxPayUnifiedOrderResp) IsSuccess() bool {
	return (w.ReturnCode == "SUCCESS" && w.ResultCode == "SUCCESS")
}
