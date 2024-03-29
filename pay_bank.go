package wxpay

// RespBase 返回结果公共字段
type RespBase struct {
	ReturnCode string `xml:"return_code,CDATA"`
	ReturnMsg  string `xml:"return_msg,omitempty,CDATA"`

	ResultCode string `xml:"result_code,omitempty,CDATA"`
	ErrCode    string `xml:"err_code,omitempty,CDATA"`
	ErrCodeDes string `xml:"err_code_des,omitempty,CDATA"`
	MchId      string `xml:"mch_id,omitempty,CDATA"`
}

func (r *RespBase) IsSuccess() bool {
	return (r.ReturnCode == "SUCCESS" && r.ResultCode == "SUCCESS")
}

// WxPayPublicKey 获取RSA加密公钥
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_7&index=4
type PublicKeyReq struct {
	Request
}

func (w *PublicKeyReq) GateWay() string {
	return "https://fraud.mch.weixin.qq.com/risk/getpublickey"
}

type PublicKeyResp struct {
	RespBase
	PubKey string `xml:"pub_key,omitempty,CDATA"`
}

// PayToWalletReq 微信企业付款到零钱包
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=14_2
type PayToWalletReq struct {
	XMLName struct{} `xml:"xml"`

	MchAppid string `xml:"mch_appid"` // 申请商户号的appid或商户号绑定的appid
	MchId    string `xml:"mchid"`     // 商户号
	NonceStr string `xml:"nonce_str"`
	Sign     string `xml:"sign"` // 签名

	PartnerTradeNo string `xml:"partner_trade_no,omitempty"`
	OpenId         string `xml:"openid,omitempty"`
	CheckName      string `xml:"check_name,omitempty"`   // NO_CHECK：不校验真实姓名 	FORCE_CHECK：强校验真实姓名
	ReUserName     string `xml:"re_user_name,omitempty"` // 收款用户姓名
	Amount         int    `xml:"amount,omitempty"`
	Desc           string `xml:"desc,omitempty"`             // 企业付款操作说明信息。必填
	SpbillCreateIp string `xml:"spbill_create_ip,omitempty"` // 终端IP
}

func (w *PayToWalletReq) GateWay() string {
	return "https://api.mch.weixin.qq.com/mmpaymkttransfers/promotion/transfers"
}

// PayToWalletResp 企业付款到零钱包返回参数
type PayToWalletResp struct {
	RespBase
	MchAppid   string `xml:"mch_appid,omitempty,CDATA"`
	DeviceInfo string `xml:"device_info,omitempty,CDATA"`
	NonceStr   string `xml:"nonce_str,omitempty,CDATA"`
	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	PartnerTradeNo string `xml:"partner_trade_no,omitempty,CDATA"`
	PaymentNo      string `xml:"payment_no,omitempty,CDATA"`
	PaymentTime    string `xml:"payment_time,omitempty,CDATA"`
}

// WxPayBank 微信付款到银行卡
// https://pay.weixin.qq.com/wiki/doc/api/tools/mch_pay.php?chapter=24_2
type PayToBankCardReq struct {
	XMLName struct{} `xml:"xml"`

	MchId          string `xml:"mch_id,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty"`
	Sign           string `xml:"sign,omitempty"`
	EncBankNo      string `xml:"enc_bank_no,omitempty"`
	EncTrueName    string `xml:"enc_true_name,omitempty"`
	BankCode       string `xml:"bank_code,omitempty"`
	Amount         int    `xml:"amount,omitempty"`
	Desc           string `xml:"desc,omitempty"`
}

// GateWay ...
func (w *PayToBankCardReq) GateWay() string {
	return "https://api.mch.weixin.qq.com/mmpaysptrans/pay_bank"
}

// PayToBankCardResp 微信付款到银行卡返回结果
type PayToBankCardResp struct {
	RespBase
	PartnerTradeNo string `xml:"partner_trade_no,omitempty,CDATA"`
	Amount         int    `xml:"amount,omitempty,CDATA"`
	NonceStr       string `xml:"nonce_str,omitempty,CDATA"`
	Sign           string `xml:"sign,omitempty,CDATA"`
	// 以下字段在return_code 和result_code都为SUCCESS的时候有返回
	PaymentNo string `xml:"payment_no,omitempty,CDATA"`
	CmmsAmt   int    `xml:"cmms_amt,omitempty,CDATA"`
}

// PayToWalletStatusQueryReq 查询付款到钱包状态
type PayToWalletStatusQueryReq struct {
	XMLName struct{} `xml:"xml"`

	AppId          string `xml:"appid"`
	MchId          string `xml:"mch_id,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty"`
	Sign           string `xml:"sign,omitempty"`
}

// GateWay ...
func (w *PayToWalletStatusQueryReq) GateWay() string {
	return "https://api.mch.weixin.qq.com/mmpaymkttransfers/gettransferinfo"
}

// PayToWalletStatusQueryResp 返回结果
type PayToWalletStatusQueryResp struct {
	RespBase
	PartnerTradeNo string `xml:"partner_trade_no,omitempty,CDATA"`
	DetailID       string `xml:"detail_id,omitempty,CDATA"`
	Status         string `xml:"status,omitempty,CDATA"`
	Reason         string `xml:"reason,omitempty,CDATA"`
	OpenID         string `xml:"openid,omitempty,CDATA"`
	TransferName   string `xml:"transfer_name,omitempty,CDATA"`
	PaymentAmount  int    `xml:"payment_amount,omitempty,CDATA"`
	TransferTime   string `xml:"transfer_time,omitempty,CDATA"`
	Desc           string `xml:"desc,omitempty,CDATA"`
}

// PayToBankCardStatusQueryReq 查询企业付款银行卡
type PayToBankCardStatusQueryReq struct {
	XMLName struct{} `xml:"xml"`

	MchId          string `xml:"mch_id,omitempty"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty"`
	NonceStr       string `xml:"nonce_str,omitempty"`
	Sign           string `xml:"sign,omitempty"`
}

// GateWay ...
func (b *PayToBankCardStatusQueryReq) GateWay() string {
	return "https://api.mch.weixin.qq.com/mmpaysptrans/query_bank"
}

// PayToBankCardStatusQueryResp 返回结果
type PayToBankCardStatusQueryResp struct {
	RespBase
	PaymentNo      string `xml:"payment_no,omitempty,CDATA"`
	PartnerTradeNo string `xml:"partner_trade_no,omitempty,CDATA"`
	BankNoMd5      string `xml:"bank_no_md5,omitempty,CDATA"`
	TrueNameMd5    string `xml:"true_name_md5,omitempty,CDATA"`
	Amount         int    `xml:"amount,omitempty,CDATA"`
	Status         string `xml:"status,omitempty,CDATA"`
	CmmsAmt        int    `xml:"cmms_amt,omitempty,CDATA"`
	CreateTime     string `xml:"create_time,omitempty,CDATA"`
	PaySuccTime    string `xml:"pay_succ_time,omitempty,CDATA"`
	Reason         string `xml:"reason,omitempty,CDATA"`
}
