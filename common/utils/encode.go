package utils

import "github.com/bytedance/sonic"

func EncodeIgnoreErr(obj interface{}) string {
	b, _ := sonic.MarshalString(obj)
	return b
}
