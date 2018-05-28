package wxpay

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"testing"
)

var data = []byte(`<xml><appid><![CDATA[wxac74ab90ef2da8ab]]></appid>
<attach><![CDATA[scene:1]]></attach>
<bank_type><![CDATA[CFT]]></bank_type>
<cash_fee><![CDATA[1]]></cash_fee>
<fee_type><![CDATA[CNY]]></fee_type>
<is_subscribe><![CDATA[N]]></is_subscribe>
<mch_id><![CDATA[1500441352]]></mch_id>
<nonce_str><![CDATA[sep1dvwpcil447t8fo5zeum6wgtu3g1o]]></nonce_str>
<openid><![CDATA[oN8mm1R-mISqHgXCLIhNB4nA7dJ8]]></openid>
<out_trade_no><![CDATA[1896671908620014]]></out_trade_no>
<result_code><![CDATA[SUCCESS]]></result_code>
<return_code><![CDATA[SUCCESS]]></return_code>
<sign><![CDATA[02220B2BA6B26B5F14FE7667F564B6BA]]></sign>
<time_end><![CDATA[20180413163408]]></time_end>
<total_fee>1</total_fee>
<trade_type><![CDATA[APP]]></trade_type>
<transaction_id><![CDATA[4200000069201804137584968101]]></transaction_id>
</xml>`)

func TestWxPay_AppPayNotification(t *testing.T) {
	appId := "wxac74ab90ef2da8ab"
	mchId := "1500441352"
	key := "d3dRCtTnw4pdRkreHC8JXuK63DgPuhOF"

	cli := &WxPay{
		appId:  appId,
		mchId:  mchId,
		key:    key,
		Client: http.DefaultClient,
	}

	var noti WxAppPayNotification
	err := xml.Unmarshal(data, &noti)
	if err != nil {
		fmt.Println("11111111: ", err)
	}

	//if noti.ReturnCode != "SUCCESS" {
	//	return nil, errors.New("notification error")
	//}
	fmt.Println("2222222222: ", noti.ReturnCode)

	sign := cli.SignWithMD5(signStr(ReflectStruct(noti)))
	fmt.Println("sign: ", sign)
	fmt.Println("noti.Sign: ", noti.Sign)
}
