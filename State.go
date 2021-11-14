package main

type State int32

const (
	Released State = iota
	Wanted   State = iota
	Held     State = iota
)

//this defines an enum type for representing state
