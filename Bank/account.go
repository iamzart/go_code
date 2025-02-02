package account

import "fmt"

type Savingaccount struct {
	Account
}

type Creditaccount struct {
	Account
}

type Account struct {
	Accountnumber string
	Balance       int
	Accounttype   string
}

//这个接口包含两个方法
type Accountoperations interface {
	Deposit(amount int)
	Withdraw(amount int)
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
	fmt.Printf("余额为： %d\n", a.Balance)
}

func (a *Account) Withdraw(amount int) {
	if a.Balance <= amount {
		fmt.Println("余额不足")
		return
	}
	a.Balance -= amount
	fmt.Printf("余额为： %d\n", a.Balance)
}

func (s *Savingaccount) Deposit(amount int) {
	s.Balance += amount
	fmt.Printf("余额为： %d\n", s.Balance)
}

func (s *Savingaccount) Withdraw(amount int) {
	if s.Balance <= amount {
		fmt.Println("余额不足")
		return
	}
	s.Balance -= amount
	fmt.Printf("余额为： %d\n", s.Balance)
}

func (c *Creditaccount) Deposit(amount int) {
	c.Balance += amount
	fmt.Printf("余额为： %d\n", c.Balance)
}

func (c *Creditaccount) Withdraw(amount int) {
	if c.Balance <= amount {
		fmt.Println("余额不足")
		return
	}
	c.Balance -= amount
	fmt.Printf("余额为： %d\n", c.Balance)
}
