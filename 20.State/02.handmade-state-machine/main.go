package main

import (
	"bufio"
	"os"
	"strconv"
)

type State int

const (
	OffHook State = iota
	Connecting
	Connected
	OnHold
	OnHook
)

func (s State) String() string {
	names := [...]string{
		"OffHook",
		"Connecting",
		"Connected",
		"OnHold",
		"OnHook",
	}
	if s < OffHook || s > OnHook {
		return "Unknown"
	}
	return names[s]
}

type Trigger int

const (
	CallDialed Trigger = iota
	HungUp
	CallConnected
	PlacedOnHold
	TakenOffHold
	LeftMessage
)

func (t Trigger) String() string {
	names := [...]string{
		"CallDialed",
		"HungUp",
		"CallConnected",
		"PlacedOnHold",
		"TakenOffHold",
		"LeftMessage",
	}
	if t < CallDialed || t > LeftMessage {
		return "Unknown"
	}
	return names[t]
}

type TriggerResult struct {
	Trigger Trigger
	State   State
}

var rules = map[State][]TriggerResult{
	OffHook: {
		{CallDialed, Connecting},
	},
	Connecting: {
		{HungUp, OnHook},
		{CallConnected, Connected},
	},
	Connected: {
		{LeftMessage, OnHook},
		{HungUp, OffHook},
		{PlacedOnHold, OnHold},
	},
	OnHold: {
		{TakenOffHold, Connected},
		{HungUp, OffHook},
	},
}

func main() {
	state, existState := OffHook, OnHook
	for state != existState {
		println("The phone is currently", state.String())
		println("Select a trigger:")

		for i := 0; i < len(rules[state]); i++ {
			t := rules[state][i].Trigger
			println(i, ".", t.String())
		}

		input, _, _ := bufio.NewReader(os.Stdin).ReadLine()
		i, _ := strconv.Atoi(string(input))

		tr := rules[state][i]
		state = tr.State
	}
	println("We are done using the phone")
}

// State Machine in Go
