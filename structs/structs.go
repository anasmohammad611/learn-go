package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func test(m messageToSend) {
	fmt.Printf("Sending message: '%s' to: %v\n", m.message, m.phoneNumber)
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		phoneNumber: 148255510981,
		message:     "Thanks for signing up",
	})
	test(messageToSend{
		phoneNumber: 148255510982,
		message:     "Love to have you aboard!",
	})
	test(messageToSend{
		phoneNumber: 148255510983,
		message:     "We're so excited to have you",
	})

	// Functions in Structs
	testAuth(authenticationInfo{
		username: "Google",
		password: "12345",
	})
	testAuth(authenticationInfo{
		username: "Bing",
		password: "98765",
	})
	testAuth(authenticationInfo{
		username: "DDG",
		password: "76921",
	})
}

type authenticationInfo struct {
	username string
	password string
}

func (auth authenticationInfo) getBasicAuth() string {
	return auth.username + ":" + auth.password
}

// don't touch below this line

func testAuth(authInfo authenticationInfo) {
	fmt.Println(authInfo.getBasicAuth())
	fmt.Println("====================================")
}
