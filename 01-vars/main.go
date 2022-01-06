package main

import (
	"errors"
	"fmt"
)

func main() {
	//sayHello()
	//sayHello("fernando")
	//sayHello("maria", "fernando")
	checkError()
}

func sayHello(names ...string) {

	if len(names) == 0 {
		fmt.Println("nobody to greet")
		return
	}

	for _, n := range names {
		fmt.Printf("Hello %s\n", n)
	}
}

func checkError() {

	// if err := boom(0); err != nil {
	// 	fmt.Println("state :", err)
	// 	return
	// }

	state, err := boomC(1)

	if err != nil {
		fmt.Println("state :", err)
		return
	}

	fmt.Println(state)
}

func boom(code int) error {
	if code == 1 {
		return errors.New(" something is wrong ")
	}
	return nil
}

func boomC(code int) (string, error) {

	handler := func(err error) (string, error) {
		return "", err
	}

	if code == 1 {
		return handler(errors.New(" something is wrong "))
	}

	return "it's ok", nil
}
