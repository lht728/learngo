package main

import "fmt"

type ZhiFuBao struct {
	// 支付宝
}

type WeChat struct {
	// 微信支付
}

type Payer interface {
	Pay(int64)
}

// Pay 支付宝的支付方法
func (z *ZhiFuBao) Pay(amount int64) {
	fmt.Printf("使用支付宝付款%.2f元。\n", float64(amount))
}

// Pay 微信的支付方法
func (w *WeChat) Pay(amount int64) {
	fmt.Printf("使用微信付款%.2f元。\n", float64(amount))
}

// Checkout 结账
func Checkout(obj Payer, amount int64) {
	obj.Pay(amount)
}

func main() {
	z := ZhiFuBao{}
	Checkout(&z, 100)
	w := WeChat{}
	Checkout(&w, 5)
}
