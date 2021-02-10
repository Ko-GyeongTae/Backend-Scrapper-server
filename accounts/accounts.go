package accounts

//Account struct
type Account struct {
	owner string
	balance int
}


// NewAccount creates Account
func NewAccount(owner string) *Account{
	account := Account{owner: owner, balance: 0}
	return &account
}

//Method
// Deposit x amound on your account
func (a Account) Deposit(amount int){
	//a is receiver It's name rule is Account 's a to start 
	a.balance += amount
}

// Balance of your balance
func (a Account) Balance() int {
	return a.balance
}