package wxpay

import (
	"fmt"
	"testing"
)

func TestWxPay_Do(t *testing.T) {
	o := &WxPayUnifiedOrder{}
	o.Body = "wawaji积分"
	o.OutTradeNo = "2017071216411234567"
	o.TotalFee = 15
	o.SpbillCreateIp = "192.168.1.109"
	o.NotifyUrl = "https://www.baidu.com/"
	o.TradeType = APP
	cli := WxPayClient("appid", "商户号", "密钥")

	resp, _ := cli.DoRequest(o)
	var data WxPayUnifiedOrderResp
	cli.ReadResponse(resp, &data)

	fmt.Println(data)
}
