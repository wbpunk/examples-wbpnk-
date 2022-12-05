package main

import (
	"errors"
	"fmt"
)

// calendar(example)
func main() {
	fmt.Println(prediction("вт"))
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "понедельник(((", nil
	case "вт":
		return "вторник(((", nil
	case "ср":
		return "средаааа(((", nil
	case "чт":
		return "четвергг)))", nil
	case "пт":
		return "пятница))))", nil
	default:
		return "невалидный день недели", errors.New("invalid day of the week")
	}
}
