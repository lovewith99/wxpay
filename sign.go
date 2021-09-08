package wxpay

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
	"sort"
	"strings"
)

func MakeSign(hm map[string]interface{}, key, signType string) string {
	keys := make([]string, 0, len(hm))
	for k := range hm {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var buf bytes.Buffer
	for _, k := range keys {
		v := hm[k]
		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(k)
		buf.WriteByte('=')
		switch s := v.(type) {
		case string:
			buf.WriteString(s)
		default:
			buf.WriteString(fmt.Sprintf("%v", s))
		}
	}

	var h hash.Hash
	buf.WriteString("&key=")
	buf.WriteString(key)
	switch signType {
	case MD5:
		h = md5.New()
	case HMAC_SHA256:
		h = hmac.New(sha256.New, []byte(key))
	}
	h.Write(buf.Bytes())

	return strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
}
