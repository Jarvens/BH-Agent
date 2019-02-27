// date: 2019-02-27
package common

import "strings"

type BihaiSymbol struct {
	BihaiUsdtSymbol []string
	BihaiBtcSymbol  []string
	BihaiEthSymbol  []string
	BihaiSymbols    []string
}

func NewBihaiSymbol() *BihaiSymbol {
	bihai := new(BihaiSymbol)
	return bihai.defaultInit()
}

func (b *BihaiSymbol) defaultInit() *BihaiSymbol {
	for _, symbol := range append(CommonUsdt, BihaiUsdt...) {
		b.BihaiUsdtSymbol = append(b.BihaiUsdtSymbol, strings.ToLower(strings.Replace(symbol, "_", "", -1)))
	}

	for _, symbol := range append(CommonBtc, BihaiBtc...) {
		b.BihaiBtcSymbol = append(b.BihaiBtcSymbol, strings.ToLower(strings.Replace(symbol, "_", "", -1)))
	}

	for _, symbol := range append(CommonEth, BihaiEth...) {
		b.BihaiEthSymbol = append(b.BihaiEthSymbol, strings.ToLower(strings.Replace(symbol, "_", "", -1)))
	}
	b.BihaiSymbols = append(b.BihaiUsdtSymbol, append(b.BihaiBtcSymbol, b.BihaiEthSymbol...)...)
	return b
}

func (b *BihaiSymbol) symbolTransfer(symbol string) string {
	isExist1, _ := Contain(symbol, b.BihaiUsdtSymbol)

	if isExist1 {
		return strings.ToUpper(strings.Replace(symbol, "usdt", "_usdt", -1))
	}

	isExist2, _ := Contain(symbol, b.BihaiBtcSymbol)
	if isExist2 {
		return strings.ToUpper(strings.Replace(symbol, "btc", "_btc", -1))
	}

	isExist3, _ := Contain(symbol, b.BihaiEthSymbol)
	if isExist3 {
		return strings.ToUpper(strings.Replace(symbol, "eth", "_eth", -1))
	}

	return ""
}
