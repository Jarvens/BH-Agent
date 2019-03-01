// date: 2019-02-28
package main

import (
	"encoding/json"
	"fmt"
	"github.com/Jarvens/BH-Agent/common"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
	"time"
)

func main() {
	event := "bh.order.base.btc_usdt.depth"
	fmt.Printf("拆分: %s\n", strings.Split(event, ".")[1])

	str1 := `{"userId":"1"}`
	//str,err:=json.Marshal(`{"userId":"1"}`)
	//if err!=nil{
	//	fmt.Printf("chucuo ")
	//}
	var dataMap map[string]interface{}
	json.Unmarshal([]byte(str1), &dataMap)
	fmt.Printf("json 转 Map: %v", dataMap["userId"])

	type UserInfo map[string]interface{}
	t := time.Now()
	key := "welcome goLang"
	userInfo := make(UserInfo)
	var expireTime int64 = 1000 * 60 * 10
	var tokenState string
	userInfo["userId"] = "000111"
	userInfo["exp"] = strconv.FormatInt(t.UTC().UnixNano(), 10)
	userInfo["iat"] = "0"
	tokenString := common.CreateToken(key, userInfo)
	fmt.Println("打印token string ", tokenString)
	claims, ok := common.ParseToken(tokenString, key)
	if ok {
		fmt.Printf("打印用户id: %s\n", claims.(jwt.MapClaims)["userId"])
		oldT, _ := strconv.ParseInt(claims.(jwt.MapClaims)["exp"].(string), 10, 64)
		ct := t.UTC().UnixNano()
		c := ct - oldT
		if c > expireTime {
			ok = false
			tokenState = "Token 已过期"
		} else {
			tokenState = "Token正常"
		}
	} else {
		tokenState = "Token无效"
	}
	fmt.Println(tokenState)
	fmt.Println(claims)

}
