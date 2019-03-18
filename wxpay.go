package wxpay

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

// 交易类型
const (
	// 统一下单
	JSAPI  = "JSAPI"  // 公众号支付
	NATIVE = "NATIVE" // 原生扫码支付
	APP    = "APP"    // app支付

	// 刷卡支付有单独的支付接口，不调用统一下单接口
	MICROPAY = "MICROPAY" // 刷卡支付
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
	MD5         = "MD5"
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

type ResponseResult interface {
	IsSuccess() bool
}

type WxPay struct {
	appId string // 微信支付分配的公共账号ID
	mchId string
	key   string
	*http.Client
}

func WxPayClient(appId, mchId, key string) *WxPay {
	return &WxPay{
		appId:  appId,
		mchId:  mchId,
		key:    key,
		Client: http.DefaultClient,
	}
}

// SignWithMD5 MD5签名
func (cli *WxPay) SignWithMD5(signStr string) string {
	signStr = fmt.Sprintf("%s&key=%s", signStr, cli.key)
	md5V := md5.Sum(Str2Bytes(signStr))
	return strings.ToUpper(hex.EncodeToString(md5V[:]))
}

// DoTlsRequest 双向证书
func (cli *WxPay) DoTlsRequest(p RequestParams, certPem, keyPem string) (*http.Response, error) {
	p.setAppId(cli.appId)
	p.setMchId(cli.mchId)

	signStr := p.SignStr()

	switch p.signType() {
	case "", MD5:
		p.setSign(cli.SignWithMD5(signStr))
	}

	b, err := xml.Marshal(p)
	if err != nil {
		return nil, err
	}

	var buf io.Reader
	buf = strings.NewReader(Bytes2Str(b))

	req, err := http.NewRequest("POST", p.GateWay(), buf)
	if err != nil {

		return nil, err
	}

	cliCrt, err := tls.LoadX509KeyPair(certPem, keyPem)
	if err != nil {

		return nil, err
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cliCrt},
		},
	}

	cli.Client = &http.Client{Transport: tr}
	return cli.Client.Do(req)
}

// DoRequest ...
func (cli *WxPay) DoRequest(p RequestParams) (*http.Response, error) {
	p.setAppId(cli.appId)
	p.setMchId(cli.mchId)

	signStr := p.SignStr()

	switch p.signType() {
	case "", MD5:
		p.setSign(cli.SignWithMD5(signStr))
	}

	b, err := xml.Marshal(p)
	if err != nil {
		return nil, err
	}

	var buf io.Reader
	buf = strings.NewReader(Bytes2Str(b))

	req, err := http.NewRequest("POST", p.GateWay(), buf)
	if err != nil {
		return nil, err
	}
	return cli.Do(req)
}

// ReadResponse 读取body
func (cli *WxPay) ReadResponse(resp *http.Response, data ResponseResult) error {
	b, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	err = xml.Unmarshal(b, data)
	return err
}

// AppPayNotification 通知
func (cli *WxPay) AppPayNotification(req *http.Request) (*WxAppPayNotification, error) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	var noti WxAppPayNotification
	err = xml.Unmarshal(b, &noti)
	if err != nil {
		return nil, err
	}

	if noti.ReturnCode != "SUCCESS" {
		return nil, errors.New("notification error")
	}

	sign := cli.SignWithMD5(signStr(ReflectStruct(noti)))
	if sign != noti.Sign {
		return nil, errors.New("签名错误")
	}

	return &noti, err
}
