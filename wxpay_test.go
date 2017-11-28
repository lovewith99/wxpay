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
	cli := WxPayClient("wx06ea7739672474ee", "1485757472", "7c01cd5266b24989bd58127c49046568")

	cli.Do(o)
	fmt.Println(o.Sign)
}
