package wxpay

import (
	"bytes"
	"math/rand"
	"reflect"
	"strings"
)

type RequestIface interface {
	GateWay() string
	GetSignType() string
	SetAppId(string)
	SetMchId(string)
	SetSignType(string)
	SetNonceStr()
	SetSign(string)
}

type Request struct {
	XMLName struct{} `xml:"xml"`

	AppId string `xml:"appid"`  // 公众账号id
	MchId string `xml:"mch_id"` // 商户号

	Sign     string `xml:"sign,omitempty"`      // 签名
	NonceStr string `xml:"nonce_str,omitempty"` // 随机字符串
	SignType string `xml:"sign_type,omitempty"` // 签名类型
}

func (req *Request) SetAppId(v string) {
	req.AppId = v
}

func (req *Request) SetMchId(v string) {
	req.MchId = v
}

func (req *Request) SetSign(v string) {
	req.Sign = v
}

func (req *Request) SetSignType(v string) {
	req.SignType = v
}

func (req *Request) GetSignType() string {
	return req.SignType
}

func (req *Request) SetNonceStr() {
	src := "0123456789abcdefghijklmnopqrstuvwxyz"

	var buf bytes.Buffer
	for i := 0; i < 32; i++ {
		j := rand.Uint32() % uint32(len(src))
		buf.WriteByte(src[j])
	}

	req.NonceStr = buf.String()
}

func SetSign(req RequestIface, key string) {
	if req.GetSignType() == "" {
		req.SetSignType(HMAC_SHA256)
	}

	hm := make(map[string]interface{})
	parseStruct(reflect.ValueOf(req), hm)

	req.SetSign(MakeSign(hm, key, req.GetSignType()))
}

func parseStruct(v reflect.Value, hm map[string]interface{}) {
	v = reflect.Indirect(v)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		switch t.Field(i).Type.Kind() {
		case reflect.Ptr, reflect.Struct:
			parseStruct(reflect.Indirect(v.Field(i)), hm)
		default:
			tag := t.Field(i).Tag.Get("xml")
			tks := strings.Split(tag, ",")
			val := v.Field(i)
			var omit bool
			if val.IsZero() {
				for i := 1; i < len(tks); i++ {
					if tks[i] == "omitempty" {
						omit = true
						break
					}
				}
			}
			if !omit {
				hm[tks[0]] = val.Interface()
			}
		}
	}
}
