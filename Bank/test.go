package main

import (
	"gocode/project01/main/Bank/account"
)

func main() {
	savac := &account.Savingaccount{
		Account: account.Account{
			Accountnumber: "00001234556",
			Balance:       50000,
			Accounttype:   "saving",
		},
	}
	savac.Deposit(10000)
} //savac是一个结构体指针变量
