package main

import "time"

// creating a custom type

// initiated with the type keyword, then we give it a name starting with a capital letter.
type NewTypeName struct {
	Id          uint64
	Text        string
	DateCreated time.Time  `objectbox:"date type:int64 converter:timeInt64"`
}