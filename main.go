package main

import "fmt"

func mains() { //main
	print("ISTA")
	fmt.Print(" MAR")
	fmt.Print(" All A Round The World\n")
	fmt.Print(" 1\n", " 2\n", " 3\n")
	fmt.Println("Belajar Golang ")
	fmt.Println("1000", " 2000", " 3000")
	// comment golang
	/*
	 * multiline
	 */
	var nama string = "Kamu"
	var status string = "Jos"
	var data string
	var (
		a_username string
		b_password string
	)
	a_username = "usernam"
	b_password = "passwor"
	fmt.Println(nama, " ", status, " ", data)
	fmt.Print(a_username, b_password+"\n\n")

	var a string = "FIZZ"
	var b string = "BUZZ"
	var count int = 15
	for i := 1; i <= count; i++ {
		if i%3 == 0 && i%5 == 0 {
			println(a + b)
		} else if i == 3 || i%3 == 0 {
			println(a)
		} else if i == 5 || i%5 == 0 {
			println(b)
		} else {
			println(i)
		}
	}

}
