package wxpay

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"math/rand"
	"reflect"
	"sort"
	"strings"
	"time"
	"unsafe"
)

func Str2Bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	h := [3]uintptr{x[0], x[1], x[1]}

	return *(*[]byte)(unsafe.Pointer(&h))
}

func Bytes2Str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ReflectStruct(i interface{}) map[string]interface{} {
	t := reflect.TypeOf(i)
	v := reflect.ValueOf(i)

	m := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		tokens := strings.Split(t.Field(i).Tag.Get("xml"), ",")
		tag := tokens[0]
		val := v.Field(i)

		if isOmitEmpty(tokens) && isEmptyValue(val) {
			continue
		}
		m[tag] = val.Interface()
	}
	return m
}

func isOmitEmpty(tokens []string) bool {
	for _, e := range tokens[1:] {
		if e == "omitempty" {
			return true
		}
	}
	return false
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}

func signStr(m map[string]interface{}) string {
	delete(m, "sign")

	var keys = make([]string, 0, len(m))
	for k, _ := range m {
		k = strings.TrimSpace(k)
		if k == "xml" {
			continue
		}
		keys = append(keys, k)
	}

	sort.Strings(keys)
	var v = make([]string, len(keys))
	for i, k := range keys {
		v[i] = fmt.Sprintf("%s=%v", k, m[k])
	}

	return strings.Join(v, "&")
}

func RandString(l int) string {
	// l <= len(t)
	t := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd',
		'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's',
		't', 'u', 'v', 'w', 'x', 'y', 'z'}

	var buf bytes.Buffer
	rand.Seed(time.Now().Unix())
	for i := 0; i < l; i++ {
		r := rand.Intn(len(t))
		buf.WriteByte(t[r])
	}

	return buf.String()
}

func XmlToMap(xmlStr string) Params {
	params := make(Params)
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))

	var (
		key   string
		value string
	)

	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement: // 开始标签
			key = token.Name.Local
		case xml.CharData: // 标签内容
			content := string([]byte(token))
			value = content
		}
		if key != "xml" {
			if value != "\n" {
				params.SetString(key, value)
			}
		}
	}

	return params
}

func MapToXml(params Params) string {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range params {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(v)
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</xml>`)

	return buf.String()
}
