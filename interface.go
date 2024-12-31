package main

import (
	"fmt"
)

type PayMethod interface {
	Acount
	Pay(amount int) bool
}

type Acount interface {
	GetBalance() int
}

// CreditCard 结构体实现 PaymentMethod 接口
type CreditCard struct {
	balance int
	limit   int
}

func (c *CreditCard) Pay(amount int) bool {
	if (c.balance + amount) <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("信用卡支付失败: 超出额度")
	return false
}

func (c *CreditCard) GetBalance() int {
	return c.balance
}

// DebitCard 结构体实现 PaymentMethod 接口
type DebitCard struct {
	balance int
}

func (c *DebitCard) Pay(amount int) bool {
	if c.balance >= amount {
		c.balance -= amount
		fmt.Printf("借记卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败: 余额不足")
	return false
}

func (c *DebitCard) GetBalance() int {
	return c.balance
}

// 使用 PaymentMethod 接口的函数
func purchaseItem(p PayMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功，剩余余额: %d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

func main() {
	creditCard := CreditCard{balance: 1000, limit: 2000}
	debitCard := DebitCard{balance: 500}

	purchaseItem(&creditCard, 100)
	purchaseItem(&debitCard, 200)

}
