package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	reversedMessage := stringutil.Reverse("Hello, OTUS!")

	fmt.Println(reversedMessage)
}
