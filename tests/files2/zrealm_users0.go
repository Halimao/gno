package main

// SEND: 100gnot

import (
	"gno.land/r/users"
)

func main() {
	err := users.Register("", "jaekwon", "my profile")
	if err != nil {
		println(err.Error())
		panic(err)
	}
	println("done")
}

// Error:
// insufficient payment
