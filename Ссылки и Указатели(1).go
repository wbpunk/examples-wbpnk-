// В ЭТОМ ПРИМЕРЕ МЫ БУДЕМ МЕНЯТЬ ЯЧЕЙКУ ПАМЯТИ

package main

import "fmt"

func main() {
	message := "variable "

	PrintMessage(message) //ВПИСАВ СЮДА ЭТУ ФУНКЦИЮ, В ТЕРМИНАЛЕ ОТОБРАЗИТЬСЯ НАША ПЕРЕЗАПИСАННАЯ ПЕРМЕННАЯ

	fmt.Println(message) //ТАК КАК ТУТ МЫ НЕ ВЫЗЫВАЕМ НАШУ ФУНКЦИЮ, А ПРОСТО ВЫВОДИМ С ПОМОЩЬЮ FMT, В ТЕРМИНАЛЕ У НАС БУДЕТ ПЕРВОНАЧАЛЬНАЯ ПЕРЕМЕННАЯ
}

func PrintMessage(message string) { //В ЭТОЙ ФУНКЦИИ ПЕРЕМЕННАЯ message ПЕРЕЗАПИСЫВАЕТСЯ
	message += "from this function" // С ПОМОЩЬЮ += ДОБАВЛЯЕМ К НАШЕЙ ПЕРЕМЕННОЙ СО СТРОКОЙ "???" ЕЩЕ ОДНУ СТРОКУ "from this function"(тоесть перезаписываем ее)
	fmt.Println(message)
}
