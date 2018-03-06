package wxpay

import (
	"testing"
	"fmt"
)

func TestWxPay_Do(t *testing.T) {
	o := &WxPayUnifiedOrder{}
	o.Body = "wawaji积分"
	o.OutTradeNo = "2017071216411234567"
	o.TotalFee = 15
	o.SpbillCreateIp = "192.168.1.109"
	o.NotifyUrl = "https://www.baidu.com/"
	o.TradeType = APP
	cli := WxPayClient("wx06ea77aa9672474ee", "1385757472", "7c01cd5266b24989bd58127c490fa568")

	resp, _ := cli.DoRequest(o)
	var data WxPayUnifiedOrderResp
	cli.ReadResponse(resp, &data)

	fmt.Println(data)
}
