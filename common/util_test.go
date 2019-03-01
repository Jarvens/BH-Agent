// date: 2019-03-01
package common

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//测试Token创建
func TestCreateToken(t *testing.T) {
	type UserInfo map[string]interface{}
	userInfo := make(UserInfo)
	userInfo["userId"] = "123456"
	userInfo["exp"] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
	userInfo["name"] = "张三"
	key := "BH-Agent"
	token := CreateToken(key, userInfo)
	fmt.Println("token: ", token)
}

//token创建性能压测
func BenchmarkCreateToken(b *testing.B) {
	for i := 0; i < b.N; i++ {
		type UserInfo map[string]interface{}
		userInfo := make(UserInfo)
		userInfo["userId"] = "123456"
		userInfo["exp"] = strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
		userInfo["name"] = "张三"
		key := "BH-Agent"
		token := CreateToken(key, userInfo)
		fmt.Println("Benchmark token: ", token)
	}
}

//contain 测试
func TestContain(t *testing.T) {
	symbol := "eth"
	symbolMap := make(map[string]string)
	symbolMap["usdt"] = "usdt"
	symbolMap["xrp"] = "xrp"
	symbolMap["btc"] = "btc"
	valid, _ := Contain(symbol, symbolMap)
	fmt.Println("valid: ", valid)
}
