package submodule

import (
	"fmt"
	_ "log"
	"math/rand"
)

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a randomly selected message format by specifying
	// a random index for the slice of formats.
	return formats[rand.Intn(len(formats))]
}

func SubModule() {

	fmt.Println("submodule  init")
	// log.Fatalf("sub module log")

	var s string = randomFormat()
	message := fmt.Sprintf(s, "tang")
	fmt.Println(message)
}

func SubModule2() {
	fmt.Println("module 2 init")
}
