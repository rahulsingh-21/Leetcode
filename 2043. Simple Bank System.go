package main

type Bank struct {
	accounts map[int]int64
}

func ConstructorIII(balance []int64) Bank {
	m := make(map[int]int64)

	for i := 0; i < len(balance); i++ {
		m[i+1] = balance[i]
	}
	return Bank{m}
}

func (this *Bank) Transfer(account1 int, account2 int, money int64) bool {
	if _, ok := this.accounts[account1]; !ok {
		return false
	}
	if _, ok := this.accounts[account2]; !ok {
		return false
	}

	if this.accounts[account1] >= money {
		this.accounts[account1] -= money
	} else {
		return false
	}
	this.accounts[account2] += money
	return true
}

func (this *Bank) Deposit(account int, money int64) bool {
	if _, ok := this.accounts[account]; !ok {
		return false
	}
	this.accounts[account] += money
	return true
}

func (this *Bank) Withdraw(account int, money int64) bool {
	if _, ok := this.accounts[account]; !ok {
		return false
	}
	if this.accounts[account] >= money {
		this.accounts[account] -= money
	} else {
		return false
	}
	return true
}

/**
 * Your Bank object will be instantiated and called as such:
 * obj := Constructor(balance);
 * param_1 := obj.Transfer(account1,account2,money);
 * param_2 := obj.Deposit(account,money);
 * param_3 := obj.Withdraw(account,money);
 */
