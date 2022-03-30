package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	var Input1 string
	var Input2 string
	fmt.Println("Введите 1-ю строку:")
	fmt.Scanln(&Input1)
	fmt.Println("Введите 2-ю строку:")
	fmt.Scanln(&Input2)
	Output1 := utf8.RuneCountInString(Input1)
	Output2 := utf8.RuneCountInString(Input2)

	switch {

	case Output1 < Output2:
		fmt.Println(Input1)
	case Output1 > Output2:
		fmt.Println(Input2)
	case Output1 == Output2:
		if Output1%2 == 0 { //Чётное
			Output3 := Output1 / 2
			q := string([]rune(Input1)[:Output3]) + string([]rune(Input2)) + string([]rune(Input1)[Output3:])
			fmt.Println("Ответ:", q)
		}
		if Output1%2 == 1 { //Нечётное
			Output4 := Output2 / 2
			Output5 := Output4 + 1
			q := string([]rune(Input2)[:Output4]) + string([]rune(Input1)) + string([]rune(Input2)[Output5:])
			fmt.Println("Ответ:", q)

		}
	}
}

/*Задание:
Вводятся 2 строки (любые) сравнить количество символов в строках и вывести только ту, в которой символов меньше.
Сразу предупреждаю, что буду вводить и сравнивать русскую строку с английской.
А если строки одинковые, то вывести одну в середине другой (если число символов нечётное, то заменить средний символо строкой)*/
