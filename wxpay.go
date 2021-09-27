package wxpay

import (
	"bytes"
	"crypto/md5"
	"crypto/tls"
	"encoding/hex"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"
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

	_, err = cli.processResponseXml(string(b))

	if err != nil {
		return nil, err
	}

	var noti WxAppPayNotification
	err = xml.Unmarshal(b, &noti)
	if err != nil {
		return nil, err
	}

	return &noti, err
}

// 处理 API返回数据，转换成Map对象。return_code为SUCCESS时，验证签名。
func (cli *WxPay) processResponseXml(xmlStr string) (Params, error) {
	var returnCode string
	params := XmlToMap(xmlStr)
	if params.ContainsKey("return_code") {
		returnCode = params.GetString("return_code")
	} else {
		return params, errors.New("no return_code in XML")
	}
	if returnCode == "FAIL" {
		return params, nil
	} else if returnCode == "SUCCESS" {
		if cli.ValidSign(params) {
			return params, nil
		} else {
			return params, errors.New("invalid sign value in XML")
		}
	} else {
		return params, errors.New("return_code value is invalid in XML")
	}
}

// 验证签名
func (cli *WxPay) ValidSign(params Params) bool {
	if !params.ContainsKey("sign") {
		return false
	}
	return params.GetString("sign") == cli.Sign(params)
}

// 签名
func (cli *WxPay) Sign(params Params) string {
	var keys = make([]string, 0, len(params))
	for k := range params {
		if k != "sign" {
			keys = append(keys, k)
		}
	}

	sort.Strings(keys)

	var buf bytes.Buffer
	for _, k := range keys {
		if len(params.GetString(k)) > 0 {
			buf.WriteString(k)
			buf.WriteString(`=`)
			buf.WriteString(params.GetString(k))
			buf.WriteString(`&`)
		}
	}
	// 加入key作加密密钥
	buf.WriteString(`key=`)
	buf.WriteString(cli.key)

	// 默认用MD5签名
	dataMd5 := md5.Sum(buf.Bytes())
	str := hex.EncodeToString(dataMd5[:])
	return strings.ToUpper(str)
}
