package wxpay

import (
	"encoding/xml"
	"testing"
)

const (
	appId = "xxx"
	mchId = "xxx"
	key   = "xxx"
)

func TestWxPayUnifiedOrder(t *testing.T) {
	o := &UnifiedOrder{}
	o.Body = "vip月会员"
	o.OutTradeNo = "2017071216411234567"
	o.TotalFee = 15
	o.SpbillCreateIp = "192.168.1.109"
	o.NotifyUrl = "https://www.baidu.com/"
	o.TradeType = APP
	// o.SignType = HMAC_SHA256

	cli, _ := NewWxPay(appId, mchId, key)

	var data UnifiedOrderResp
	if err := cli.Do(o, &data); err != nil {
		t.Error(err)
	}

	d := data.RequestData(cli)

	t.Log("resp: ", data)
	t.Log("data: ", d)

}

func TestAppPayNotify(t *testing.T) {
	data := []byte(`<xml><appid><![CDATA[xxx]]]></appid>
<attach><![CDATA[scene:1]]></attach>
<bank_type><![CDATA[CFT]]></bank_type>
<cash_fee><![CDATA[1]]></cash_fee>
<fee_type><![CDATA[CNY]]></fee_type>
<is_subscribe><![CDATA[N]]></is_subscribe>
<mch_id><![CDATA[xxx]]></mch_id>
<nonce_str><![CDATA[sep1dvwpcil447t8fo5zeum6wgtu3g1o]]></nonce_str>
<openid><![CDATA[oN8mm1R-mISqHgXCLIhNB4nA7dJ8]]></openid>
<out_trade_no><![CDATA[1896671908620014]]></out_trade_no>
<result_code><![CDATA[SUCCESS]]></result_code>
<return_code><![CDATA[SUCCESS]]></return_code>
<sign><![CDATA[02220B2BA6B26B5F14FE7667F564B6BA]]></sign>
<time_end><![CDATA[20180413163408]]></time_end>
<total_fee>1</total_fee>
<trade_type><![CDATA[APP]]></trade_type>
<transaction_id><![CDATA[xxx]]></transaction_id>
</xml>`)

	var noti AppPayNotification
	err := xml.Unmarshal(data, &noti)
	if err != nil {
		t.Error(err)
	}

}
