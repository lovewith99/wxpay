package wxpay

import (
	"net/http"
	"fmt"
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// 交易类型
const (
	// 统一下单
	JSAPI = "JSAPI" // 公众号支付
	NATIVE  = "NATIVE" // 原生扫码支付
	APP  = "APP" // app支付

	// 刷卡支付有单独的支付接口，不调用统一下单接口
	MICROPAY = "MICROPAY"  // 刷卡支付
)

// 货币类型
const (
	CNY = "CNY" // 人民币
)

// 银行类型
const (
	ICBC_DEBIT = "ICBC_DEBIT" // 工商银行(借记卡)
)

// 签名类型
const (
	MD5 = "MD5"
	HMAC_SHA256 = "HMAC-SHA256"
)

type RequestParams interface {
	GateWay() string
	setAppId(string)
	setMchId(string)
	setSign(string)
	SignStr() string
	signType() string
}

type WxPay struct {
	appId string       // 微信支付分配的公共账号ID
	mchId string
	key   string
	*http.Client
}

func WxPayClient(appId, mchId, key string) *WxPay {
	return &WxPay{
		appId: appId,
		mchId: mchId,
		key:   key,
		Client: http.DefaultClient,
	}
}

func (cli *WxPay) SignWithMD5(signStr string) string {
	signStr = fmt.Sprintf("%s&key=%s", signStr, cli.key)

	fmt.Println(signStr)
	md5V := md5.Sum(Str2Bytes(signStr))
	return strings.ToUpper(hex.EncodeToString(md5V[:]))
}

func (cli *WxPay) Do(p RequestParams)  {
	p.setAppId(cli.appId)
	p.setMchId(cli.mchId)

	signStr := p.SignStr()

	switch p.signType() {
	case "", MD5:
		p.setSign(cli.SignWithMD5(signStr))
	}
}