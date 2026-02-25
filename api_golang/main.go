package main

import "fmt"

const PI = 3.14;

func main() {
    fmt.Println("Hello Invest Tracking 🚀")

    name := "John";
    listName := [5]string{"John", "Jane", "Jim", "Jill", "Jack"};
    fmt.Println("Hello", name);
    fmt.Println("List Name", listName[0]);
    fmt.Println("PI", PI);
    fmt.Println(len(listName));
    fmt.Println(cap(listName));
}
