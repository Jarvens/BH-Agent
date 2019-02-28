// date: 2019-02-28
package main

import (
	"encoding/json"
	"fmt"
	"strings"
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

}
