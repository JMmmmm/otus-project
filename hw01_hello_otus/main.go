package main

func main() {
	a := make([]int, 0, 2)
	b := a
	a = append(a, 1)
	b = append(b, 2)
	println(a[0], b[0])

	//reversedMessage := stringutil.Reverse("Hello, OTUS!")
	//
	//fmt.Println(reversedMessage)
}
