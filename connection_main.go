// date: 2019-03-01
package main

import (
	"encoding/json"
	"fmt"
	"github.com/Jarvens/BH-Agent/common"
)

type TestConnection struct {
	ChanMap map[string]interface{}
}

func main() {
	test := TestConnection{}
	event := []string{"bh.order", "bh.account", "bh.chat"}
	moduleMap := make(map[string]interface{})
	moduleMap["account"] = event
	moduleMap["order"] = event
	connMap := make(map[string]interface{})
	connMap["conn1"] = moduleMap
	chanType := make(map[string]interface{})
	chanType["web"] = connMap
	chanType["mobile"] = connMap
	userMap := make(map[string]interface{})
	userMap["123456"] = chanType
	test.ChanMap = userMap
	data, _ := json.Marshal(test)
	fmt.Println(string(data))

	str := []string{"a", "b", "c"}
	v, _ := common.Contain("a", str)
	fmt.Println(v)
}
