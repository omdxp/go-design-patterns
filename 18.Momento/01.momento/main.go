package main

type Momento struct {
	Balance int
}

type BankAccount struct {
	balance int
}

func NewBankAccount(balance int) (*BankAccount, *Momento) {
	return &BankAccount{balance: balance}, &Momento{Balance: balance}
}

func (b *BankAccount) Deposit(amount int) *Momento {
	b.balance += amount
	return &Momento{Balance: b.balance}
}

func (b *BankAccount) Restore(m *Momento) {
	b.balance = m.Balance
}

func main() {
	ba, m0 := NewBankAccount(100)
	m1 := ba.Deposit(50)
	m2 := ba.Deposit(25)
	println(ba.balance)
	ba.Restore(m1)
	println(ba.balance)
	ba.Restore(m2)
	println(ba.balance)
	ba.Restore(m0)
	println(ba.balance)
}

// Momento is a behavioral design pattern that lets you save and restore the previous state of an object without revealing the details of its implementation.
