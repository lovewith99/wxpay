package wxpay

import (
	"testing"
	//"encoding/xml"
	"fmt"
)

func TestReflectStruct(t *testing.T) {
	o := &WxPayUnifiedOrder{}

	//b, _ := xml.Marshal(o)
	//fmt.Println(Bytes2Str(b))
	fmt.Println(ReflectStruct(*o))
}

func TestRandString(t *testing.T) {
	fmt.Println(RandString(36))
}
