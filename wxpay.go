package wxpay

import (
	"bytes"
	"crypto/tls"
	"encoding/xml"
	"errors"
	"net/http"
	"time"
)

type Client struct {
	appId string // 微信支付分配的公共账号ID
	mchId string
	key   string
	*http.Client
}

func (cli *Client) Do(req RequestIface, data interface{}) error {
	req.SetAppId(cli.appId)
	req.SetMchId(cli.mchId)
	req.SetNonceStr()
	SetSign(req, cli.key)

	bs, err := xml.Marshal(req)
	if err != nil {
		return err
	}

	httpreq, _ := http.NewRequest("POST", req.GateWay(), bytes.NewBuffer(bs))

	resp, err := cli.Client.Do(httpreq)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err = xml.NewDecoder(resp.Body).Decode(data); err != nil {
		return err
	}

	return nil
}

// AppPayNotification 通知
func (cli *Client) AppPayNotification(req *http.Request) (*AppPayNotification, error) {
	defer req.Body.Close()

	var noti AppPayNotification
	err := xml.NewDecoder(req.Body).Decode(&noti)
	if err != nil {
		return &noti, err
	}

	if !noti.VerifySign(cli.key) {
		err = errors.New("签名错误")
	}

	return &noti, err
}

func (cli *Client) MakeSign(hm map[string]interface{}, signType string) string {
	return MakeSign(hm, cli.key, signType)
}

func NewWxPay(appId, mchId, key string, opts ...func(*Client) error) (*Client, error) {
	cli := Client{
		appId: appId,
		mchId: mchId,
		key:   key,
		Client: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, f := range opts {
		if err := f(&cli); err != nil {
			return nil, err
		}
	}

	return &cli, nil
}

func WithTlsFile(certpem, keypem string) func(*Client) error {
	return func(cli *Client) error {
		cert, err := tls.LoadX509KeyPair(certpem, keypem)
		if err != nil {
			return err
		}

		cli.Client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		}
		return nil
	}
}
