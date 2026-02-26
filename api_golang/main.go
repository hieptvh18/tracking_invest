package main

import "fmt"

const PI = 3.14

type Brand struct {
	Name  string
	Phone string
}

func main() {
	fmt.Println("Hello Invest Tracking 🚀")

	name := "John"
	listName := [5]string{"John", "Jane", "Jim", "Jill", "Jack"}
	fmt.Println("Hello", name)
	fmt.Println("List Name", listName[0])
	fmt.Println("PI", PI)
	fmt.Println(len(listName))
	fmt.Println(cap(listName))
	fmt.Println(sum(1, 2))
	fmt.Println(sayHi())
	fmt.Println(mapNumber())
	fmt.Println(demoPointer())

	shop := Brand{"iphone", "Adndroi"}

	fmt.Println(shop.Name)
}

func sum(a int, b int) int {
	return a + b
}

func sayHi() string {
	return "Hello"
}

func mapNumber() int {
	total := 0
	for i := 1; i <= 10; i++ {
		total += i
	}

	return total
}

func demoPointer() int {
	a, b := 1, 1000

	// defined pointer p = a vì trỏ về cùng 1 vùng nhớ
	p := &a
	*p = b

	return a
}
