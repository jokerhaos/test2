package main

import (
	"crypto/aes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
)

type SignAlgo struct{}

func (s *SignAlgo) SortParam(param map[string]interface{}) string {
	keys := make([]string, 0, len(param))
	for key := range param {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var sortedParams []string
	for _, key := range keys {

		value := param[key]

		if key == "sign" || value == nil || value == "" {
			continue
		}

		switch v := value.(type) {
		case int:
			sortedParams = append(sortedParams, fmt.Sprintf("%s=%d", key, v))
		default:
			sortedParams = append(sortedParams, key+"="+value.(string))
		}

	}
	return strings.Join(sortedParams, "&")
}

func (s *SignAlgo) Sign(data map[string]interface{}, secret string) string {
	str := s.SortParam(data)
	str += "&key=" + secret
	hash := md5.Sum([]byte(str))
	return hex.EncodeToString(hash[:])
}

func (s *SignAlgo) Verify(data map[string]interface{}, secret string) bool {
	signature := data["sign"].(string)
	delete(data, "sign")
	return s.Sign(data, secret) == signature
}

func (s *SignAlgo) EnAES(data map[string]interface{}, secret string, iv string) string {
	str := toJSONString(data)
	cipher, _ := aes.NewCipher([]byte(secret))
	encrypted := make([]byte, len(str))
	cipher.Encrypt(encrypted, []byte(str))
	return base64.StdEncoding.EncodeToString(encrypted)
}

func (s *SignAlgo) DeAES(decrypt string, secret string, iv string) string {
	decoded, _ := base64.StdEncoding.DecodeString(decrypt)
	cipher, _ := aes.NewCipher([]byte(secret))
	decrypted := make([]byte, len(decoded))
	cipher.Decrypt(decrypted, []byte(decoded))
	return string(decrypted)
}

func toJSONString(data map[string]interface{}) string {
	// 将 map 转为 JSON 字符串
	// 这里需要根据具体的 JSON 序列化库来实现
	// 可以使用 Go 标准库的 encoding/json 或者第三方库如 jsoniter 等
	// 这里只是一个简单的示例，具体实现需要根据具体需求进行调整
	return ""
}

func main() {
	// 使用示例
	sign := &SignAlgo{}

	param := map[string]interface{}{
		"a":    1,
		"b":    "b",
		"c":    2.0,
		"d":    "d",
		"sign": "xxx",
	}

	secret := "secret_key"

	sortedParams := sign.SortParam(param)
	println("Sorted Params:", sortedParams)

	signature := sign.Sign(param, secret)
	println("Signature:", signature)

	verified := sign.Verify(param, secret)
	println("Verified:", verified)

	encrypted := sign.EnAES(param, secret, "iv")
	println("Encrypted:", encrypted)

	decrypted := sign.DeAES(encrypted, secret, "iv")
	println("Decrypted:", decrypted)
}
