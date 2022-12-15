package main

import "fmt"

type Momento struct {
	Balance int
}

type BankAccount struct {
	balance int
	changes []*Momento
	current int
}

func NewBankAccount(balance int) *BankAccount {
	ba := &BankAccount{balance: balance}
	ba.changes = append(ba.changes, &Momento{Balance: balance})
	return ba
}

func (b *BankAccount) Deposit(amount int) *Momento {
	b.balance += amount
	m := &Momento{b.balance}
	b.changes = append(b.changes, m)
	b.current++
	fmt.Println("Deposited", amount, "balance is now", b.balance)
	return m
}

func (b *BankAccount) Restore(m *Momento) {
	if m != nil {
		b.balance = m.Balance
		b.changes = append(b.changes, m)
		b.current = len(b.changes) - 1
	}
}

func (b *BankAccount) String() string {
	return fmt.Sprintf("Balance = $%d, current = %d", b.balance, b.current)
}

func (b *BankAccount) Undo() *Momento {
	if b.current > 0 {
		b.current--
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func (b *BankAccount) Redo() *Momento {
	if b.current+1 < len(b.changes) {
		b.current++
		m := b.changes[b.current]
		b.balance = m.Balance
		return m
	}
	return nil
}

func main() {
	ba := NewBankAccount(100)
	ba.Deposit(50)
	ba.Deposit(25)
	fmt.Println(ba)
	ba.Undo()
	fmt.Println("Undo 1:", ba)
	ba.Undo()
	fmt.Println("Undo 2:", ba)
	ba.Redo()
	fmt.Println("Redo:", ba)
}

// Undo and Redo is a behavioral design pattern that lets you save and restore the previous state of an object without revealing the details of its implementation.
