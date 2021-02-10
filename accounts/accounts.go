package accounts

import "errors"

//Account struct
type Account struct {
	owner string
	balance int
}

var errNoMoney = errors.New("Can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}

//Method

// Deposit x amound on your account
func (a *Account) Deposit(amount int){
	//a is receiver It's name rule is Account 's a to start 
	a.balance += amount
}

// Balance of your balance
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amout from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	} 
	a.balance -= amount
	return nil
}