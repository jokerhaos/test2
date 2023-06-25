package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strings"
	// "github.com/zeromicro/go-zero/core/threading"
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
		case int, uint, int16, int32, int64:
			sortedParams = append(sortedParams, fmt.Sprintf("%s=%d", key, v))
		case float64, float32:
			sortedParams = append(sortedParams, fmt.Sprintf("%s=%f", key, v))
		default:
			sortedParams = append(sortedParams, key+"="+value.(string))
		}

	}
	return strings.Join(sortedParams, "&")
}

func (s *SignAlgo) Sign(data map[string]interface{}, secret []byte) string {
	str := s.SortParam(data)
	str += "&key=" + string(secret)
	hash := md5.Sum(secret)
	return hex.EncodeToString(hash[:])
}

func (s *SignAlgo) Verify(data map[string]interface{}, secret []byte) bool {
	signature := data["sign"].(string)
	delete(data, "sign")
	return s.Sign(data, secret) == signature
}

// PKCS5Padding 对数据进行PKCS5填充
func PKCS5Padding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// PKCS5Unpadding 去除PKCS5填充
func PKCS5Unpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("PKCS5 unpadding error: data is empty")
	}
	unpadding := int(data[length-1])
	if unpadding > length {
		return nil, errors.New("PKCS5 unpadding error: invalid padding size")
	}
	return data[:length-unpadding], nil
}

// ZeroPadding 使用ZeroPadding填充数据
func ZeroPadding(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{0}, padding)
	return append(data, padText...)
}

// ZeroUnpadding 去除ZeroPadding填充数据
func ZeroUnpadding(data []byte) ([]byte, error) {
	length := len(data)
	if length == 0 {
		return nil, errors.New("ZeroUnpadding error: data is empty")
	}
	unpadding := 0
	for i := length - 1; i >= 0; i-- {
		if data[i] == 0 {
			unpadding++
		} else {
			break
		}
	}
	if unpadding == 0 {
		return nil, errors.New("ZeroUnpadding error: no padding bytes found")
	}
	return data[:length-unpadding], nil
}

func (s *SignAlgo) EnAES(data map[string]interface{}, secret []byte, iv []byte) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 创建AES-128加密器
	block, err := aes.NewCipher(secret)
	if err != nil {
		return "", err
	}

	// 检查IV长度是否为16字节，如果不足则进行补零操作
	if len(iv) < aes.BlockSize {
		padding := make([]byte, aes.BlockSize-len(iv))
		iv = append(iv, padding...)
	}

	// 创建CBC模式的加密器
	mode := cipher.NewCBCEncrypter(block, iv)

	// 填充原始数据
	paddedData := PKCS5Padding(bytes, aes.BlockSize)

	// 创建加密缓冲区
	encrypted := make([]byte, len(paddedData))

	// 加密数据
	mode.CryptBlocks(encrypted, paddedData)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (s *SignAlgo) DeAES(decrypt string, secret []byte, iv []byte) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(decrypt)
	if err != nil {
		return nil, err
	}

	// 创建AES-128解密器
	block, err := aes.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	// 检查IV长度是否为16字节，如果不足则进行补零操作
	if len(iv) < aes.BlockSize {
		padding := make([]byte, aes.BlockSize-len(iv))
		iv = append(iv, padding...)
	}

	// 创建CBC模式的解密器
	mode := cipher.NewCBCDecrypter(block, iv)

	// 创建解密缓冲区
	decrypted := make([]byte, len(decoded))

	// 解密数据
	mode.CryptBlocks(decrypted, decoded)

	// 去除填充数据
	unpaddedData, err := PKCS5Unpadding(decrypted)

	if err != nil {
		return nil, err
	}

	return unpaddedData, nil
}

func (s *SignAlgo) EncryptDES(data map[string]interface{}, secret []byte, iv []byte) (string, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	// 创建DES加密器
	block, err := des.NewCipher(secret)
	if err != nil {
		return "", err
	}

	// 检查IV长度是否为8字节，如果不足则进行补零操作
	if len(iv) < des.BlockSize {
		padding := make([]byte, des.BlockSize-len(iv))
		iv = append(iv, padding...)
	}

	// 创建CBC模式的加密器
	mode := cipher.NewCBCEncrypter(block, iv)

	// 填充原始数据
	paddedData := ZeroPadding(bytes, des.BlockSize)

	// 创建加密缓冲区
	encrypted := make([]byte, len(paddedData))

	// 加密数据
	mode.CryptBlocks(encrypted, paddedData)

	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (s *SignAlgo) DecryptDES(decrypt string, secret []byte, iv []byte) ([]byte, error) {
	decoded, err := base64.StdEncoding.DecodeString(decrypt)
	if err != nil {
		return nil, err
	}

	// 创建并返回一个使用DES算法的cipher.Block接口
	block, err := des.NewCipher(secret)
	if err != nil {
		return nil, err
	}

	// 检查IV长度是否为8字节，如果不足则进行补零操作
	if len(iv) < des.BlockSize {
		padding := make([]byte, des.BlockSize-len(iv))
		iv = append(iv, padding...)
	}

	// 创建CBC模式的解密器
	mode := cipher.NewCBCDecrypter(block, iv)

	// 创建解密缓冲区
	decrypted := make([]byte, len(decoded))

	// 解密数据
	mode.CryptBlocks(decrypted, decoded)

	// 去除填充数据
	unpaddedData, err := ZeroUnpadding(decrypted)

	if err != nil {
		return nil, err
	}

	return unpaddedData, nil
}

func toJSONString(data map[string]interface{}) string {
	bytes, _ := json.Marshal(data)
	stringData := string(bytes)
	return stringData
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

	secret := []byte("aaaaaaaabbbbbbbb")

	sortedParams := sign.SortParam(param)
	println("Sorted Params:", sortedParams)

	signature := sign.Sign(param, secret)
	println("Signature:", signature)

	param["sign"] = signature
	verified := sign.Verify(param, secret)
	println("Verified:", verified)

	encrypted, _ := sign.EnAES(param, secret, make([]byte, 0))
	println("Encrypted:", encrypted)

	decrypted, _ := sign.DeAES(encrypted, secret, make([]byte, 0))

	println("Decrypted:", string(decrypted))

	secret2 := []byte("12345678")

	desEncrypted, _ := sign.EncryptDES(param, secret2, make([]byte, 0))
	println("desEncrypted:", desEncrypted)

	desDecrypted, _ := sign.DecryptDES(desEncrypted, secret2, make([]byte, 0))

	println("desDecrypted:", string(desDecrypted))

	// threading.GoSafe()
}
