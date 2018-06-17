package main

import (
	"fmt"
)

/*
References:
https://blog.golang.org/error-handling-and-go
https://blog.golang.org/defer-panic-and-recover
*/

func main() {
	fmt.Println("Start")
	err.SimpleErrorDemo()
	err.CustomErrorDemo()
	dpr.DeferDemo()
	dpr.PanicDemo()
	dpr.RecoverDemo()
	handle.AllInOneDemo()
	fmt.Println("End")
}