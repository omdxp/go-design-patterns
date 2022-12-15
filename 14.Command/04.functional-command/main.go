package main

import "fmt"

var overdraftLimit = -500

type BankAccount struct {
	balance int
}

func (b *BankAccount) Deposit(amount int) {
	b.balance += amount
	fmt.Println("Deposited", amount, ", balance is now", b.balance)
}

func (b *BankAccount) Withdraw(amount int) bool {
	if b.balance-amount >= overdraftLimit {
		b.balance -= amount
		fmt.Println("Withdrew", amount, ", balance is now", b.balance)
		return true
	}
	return false
}

type Command interface {
	Call()
	Undo()
	Succeeded() bool
	SetSucceeded(bool)
}

type Action int

const (
	Deposit Action = iota
	Withdraw
)

type BankAccountCommand struct {
	account *BankAccount
	action  Action
	amount  int
	success bool
}

func NewBankAccountCommand(account *BankAccount, action Action, amount int) *BankAccountCommand {
	return &BankAccountCommand{account, action, amount, false}
}

func (c *BankAccountCommand) Call() {
	switch c.action {
	case Deposit:
		c.account.Deposit(c.amount)
		c.success = true
	case Withdraw:
		c.success = c.account.Withdraw(c.amount)
	}
}

func (c *BankAccountCommand) Undo() {
	if !c.success {
		return
	}
	switch c.action {
	case Deposit:
		c.account.Withdraw(c.amount)
	case Withdraw:
		c.account.Deposit(c.amount)
	}
}

func (c *BankAccountCommand) Succeeded() bool {
	return c.success
}

func (c *BankAccountCommand) SetSucceeded(success bool) {
	c.success = success
}

type CompositeBankAccountCommand struct {
	commands []Command
}

func NewCompositeBankAccountCommand(commands []Command) *CompositeBankAccountCommand {
	return &CompositeBankAccountCommand{commands}
}

func (c *CompositeBankAccountCommand) Call() {
	for _, cmd := range c.commands {
		cmd.Call()
	}
}

func (c *CompositeBankAccountCommand) Undo() {
	for i := range c.commands {
		c.commands[len(c.commands)-1-i].Undo() // reverse order
	}
}

func (c *CompositeBankAccountCommand) Succeeded() bool {
	for _, cmd := range c.commands {
		if !cmd.Succeeded() {
			return false
		}
	}
	return true
}

func (c *CompositeBankAccountCommand) SetSucceeded(success bool) {
	for _, cmd := range c.commands {
		cmd.SetSucceeded(success)
	}
}

type MoneyTransferCommand struct {
	CompositeBankAccountCommand
	from, to *BankAccount
	amount   int
}

func NewMoneyTransferCommand(from, to *BankAccount, amount int) *MoneyTransferCommand {
	c := &MoneyTransferCommand{from: from, to: to, amount: amount}
	c.commands = append(c.commands,
		NewBankAccountCommand(from, Withdraw, amount),
		NewBankAccountCommand(to, Deposit, amount))
	return c
}

func (c *MoneyTransferCommand) Call() {
	ok := true
	for _, cmd := range c.commands {
		if ok { // only execute if previous succeeded
			cmd.Call()
			ok = cmd.Succeeded()
		} else { // if previous failed, mark all as failed
			cmd.SetSucceeded(false)
		}
	}
}

func main() {
	ba := &BankAccount{}
	var commands []func()
	commands = append(commands, func() {
		ba.Deposit(100)
	}, func() {
		ba.Withdraw(25)
	})
	for _, cmd := range commands {
		cmd()
	}
	fmt.Println(ba)
}

// Functional Command
