//package main
//
//import (
//	"fmt"
//)
//
//func (e email) cost() float64 {
//	if e.isSubscribed {
//		return float64(len(e.body)) * 0.01
//	} else {
//		return float64(len(e.body)) * 0.05
//	}
//}
//
//func (e email) print() {
//	fmt.Println(e.body)
//}
//
//// don't touch below this line
//
//type expense interface {
//	cost() float64
//}
//
//type printer interface {
//	print()
//}
//
//type email struct {
//	isSubscribed bool
//	body         string
//}
//
//func print(p printer) {
//	p.print()
//}
//
//func test(e expense, p printer) {
//	fmt.Printf("Printing with cost: $%.2f ...\n", e.cost())
//	p.print()
//	fmt.Println("====================================")
//}
//
//func main() {
//	e := email{
//		isSubscribed: true,
//		body:         "hello there",
//	}
//	test(e, e)
//	e = email{
//		isSubscribed: false,
//		body:         "I want my money back",
//	}
//	test(e, e)
//	e = email{
//		isSubscribed: true,
//		body:         "Are you free for a chat?",
//	}
//	test(e, e)
//	e = email{
//		isSubscribed: false,
//		body:         "This meeting could have been an email",
//	}
//	test(e, e)
//}

package main

import (
	"fmt"
)

func getExpenseReport(e expense) (string, float64) {
	//one way to do it
	//em, ok := e.(email)
	//if ok {
	//	return em.toAddress, e.cost()
	//}
	//sms, ok := e.(sms)
	//if ok {
	//	return sms.toPhoneNumber, e.cost()
	//}
	//return "", 0.0

	switch v := e.(type) {
	case email:
		return v.toAddress, e.cost()
	case sms:
		return v.toPhoneNumber, e.cost()
	default:
		return "", 0.0
	}
}

// don't touch below this line

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	})
	test(invalid{})
}
