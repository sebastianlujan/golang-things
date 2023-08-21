//just a sintax remainder, nothing important at all.

package main

import (
	"errors"
	"fmt"
)

type MessageToSend struct {
	Phone string
	Msg   string
	sender, recipient User
}

type User struct {
	Name string
	Number string
}

func main() {
	fmt.Println("Hello, 世界")
	test("hello,", " 世界")
	testIncrementSends(430, 25)
	testPoint(3, 4)
	testNames("Neils", "Bohr")
	testYearsUntilAdult(4)
	testCalculator(2, 3)

	numbers, msgs, _ := mockTestData("msgData")
	testMessageToSend(numbers, msgs)
}

func concat(s1 string, s2 string) string {
	return s1 + s2
}

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func testIncrementSends(sendsSoFar, sendsToAdd int) {
	sendsSoFar = incrementSends(sendsSoFar, sendsToAdd)
	fmt.Println("you've sent", sendsSoFar, "messages")
}

func incrementSends(sendsSoFar, sendsToAdd int) int {
	return sendsSoFar + sendsToAdd
}

func getPoint() (x, y int) {
	return 3, 4
}

func testPoint(x, y int) int {
	x, _ = getPoint()
	//ignore everything you don't use

	fmt.Println(x)
	return x
}

func setNames(name, lastName string) (x, y string) {
	return name, lastName
}

func testNames(name, lastName string) {
	firstName, _ := setNames(name, lastName)
	fmt.Println("Welcome", firstName)
}

func yearsUntilEvents(age int) (yearsUntilAdult, yearsUntilDrinking, yearsUntilCarRental int) {
	yearsUntilAdult = 18 - age
	yearsUntilDrinking = 21 - age
	yearsUntilCarRental = 25 - age

	if yearsUntilAdult < 0 {
		yearsUntilAdult = 0
	}
	if yearsUntilDrinking < 0 {
		yearsUntilDrinking = 0
	}
	if yearsUntilCarRental < 0 {
		yearsUntilCarRental = 0
	}

	return //implicit return
}

func testYearsUntilAdult(age int) {
	fmt.Println("Age", age)
	yAAdult, yADrinking, yACarRental := yearsUntilEvents(age)
	fmt.Println(yAAdult, yADrinking, yACarRental)
}

func calculator(a, b int) (int, int, error) {
	if b == 0 {
		return 0, 0, errors.New("Divide by Zero as oppenheimer")
	}
	return a * b, a / b, nil
}

func testCalculator(a, b int) {
	mult, div, err := calculator(a, b)
	fmt.Println(mult, div, err)
}

//I want to pass two lists , a numberList , and a msg list

func mockTestData(option string) (numberList, msgList []string, err error) {
	if option != "msgData" {
		return nil, nil, errors.New("msgData not found")
	}

	numberList = []string{"148255510981", "148255510982", "148255510983"}
	msgList = []string{"Thanks for signing up", "Data", "We're so excited to have you"}

	return
}

func testMessageToSend(numberList, msgList []string) {
	var messages []MessageToSend

	for i := 0; i < len(msgList); i++ {
		msg := MessageToSend{
			Phone: numberList[i],
			Msg:   msgList[i],
		}
		messages = append(messages, msg)
	}

	for _, msg := range messages {
		fmt.Printf("%s => %s \n", msg.Phone, msg.Msg)
	}


}
