// В ЭТОМ ПРИМЕРЕ МЫ НЕ БУДЕМ МЕНЯТЬ ЯЧЕЙКУ ПАМЯТИ( БУДЕМ ССЫЛАТЬСЯ НА НАШИ ПЕРЕМЕННЫЕ ПРИ ПОМОЩИ УКАЗАТЕЛЕЙ. ИСПОЛЬЗУЯ (*) )

package main

import "fmt"

func main() {
	message := "variable " //0xc00004e250 адресс переменной
	fmt.Println(message)

	ChangeMessage(&message) //(reference - ссылка) тут мы ссылаемся на область в памяти (которая находится в message), при помощи амперсанда(&)
	fmt.Println(message)
}

func ChangeMessage(message *string) /*сюда мы передаём область в памяти*/ {
	*message += "from this function" /* 0xc00004e250 адресс переменной
	(*message - dereference) - по области в памяти мы обращаемся к значению */
	fmt.Println(message)
}
