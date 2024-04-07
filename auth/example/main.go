package main

import (
	"fmt"
	"github.com/246859/work/auth"
	"github.com/246859/work/user"
)

func main() {
	ok, err := auth.Verify(user.User{Name: "jack", Password: "jack123456"})
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v", ok)
}
