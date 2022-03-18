package leetcode

type Bank struct {
	accounts []int64
	n        int
}

func Constructor(balance []int64) Bank {
	return Bank{
		accounts: balance,
		n:        len(balance),
	}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if account1 > this.n || account2 > this.n {
		return false
	}
	if this.accounts[account1-1] < money {
		return false
	}
	this.accounts[account1-1] -= money
	this.accounts[account2-1] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if account > this.n {
		return false
	}
	this.accounts[account-1] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if account > this.n {
		return false
	}
	if this.accounts[account-1] < money {
		return false
	}
	this.accounts[account-1] -= money
	return true
}
