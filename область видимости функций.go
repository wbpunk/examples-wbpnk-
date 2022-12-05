package main

import "fmt"

func main() {
	//ФУНКЦИЯ ЧЕРЕЗ ПЕРЕМЕННУЮ:
	hello := sayHello("Alex")
	printMessage(hello)

	// ВЫЗОВ ФУНКЦИЙ:
	Print()
	Cycle()
	printMessage("hello")
	printMessage("im")
	printMessage("Alex")
	PrintNumber(10)
	PrintNumber(30)
	PrintNumber(20)

}

// ФУНКЦИИ:
func Print() {
	fmt.Println("yooyoyoy")
	fmt.Println("suuuuup")
}
func Cycle() {
	var a int = 20
	for i := 1; i <= 10; i++ {
		if a >= 10 {
			fmt.Println("nooo")
		}

	}
}
func printMessage(message string) {
	fmt.Println(message)

}
func PrintNumber(number int) {
	fmt.Println(number)
}

// ВОЗВРАТ ЗНАЧЕНИЯ:(RETURN)
func sayHello(name string) string {
	return "Привет, " + name
}
