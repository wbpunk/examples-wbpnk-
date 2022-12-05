package main

import (
	"fmt"
	"time"
)

func main() {
	type Person struct {
		Name string
		Age  int
	}
	Vlad := Person{"Vlad", 30}
	fmt.Println(Vlad)

	//АНОНИМНЫЕ СТРУКТУРЫ:
	unnamedStruct := struct {
		Name, LastName, BirthDate string
	}{
		Name:      "???",
		LastName:  "???",
		BirthDate: fmt.Sprintf("%s", time.Now()), //(fmt.Sprintf("%d"- digital(число),"%s"-string(строка))-КОНВЕРТАЦИЯ В ТИПЫ
	}
	fmt.Println(unnamedStruct)
}
