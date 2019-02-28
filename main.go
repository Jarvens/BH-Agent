// date: 2019-02-28
package main

import (
	"fmt"
	"strings"
)

func main() {
	event := "bh.order.base.btc_usdt.depth"
	fmt.Printf("拆分: %s", strings.Split(event, ".")[1])
}
