package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var firstString, secondString, result string
	fmt.Println("Введите строки через пробел")
	fmt.Scanf("%s %s\n", &firstString, &secondString)
	countFirstString := utf8.RuneCountInString(firstString)
	countSecondString := utf8.RuneCountInString(secondString)
	switch {
	case countFirstString < countSecondString:
		result = firstString
	case countFirstString > countSecondString:
		result = secondString
	case countFirstString == countSecondString:
		result = createResultString(firstString, secondString, countFirstString, countSecondString)
	}
	fmt.Println(result)
}

func createResultString(firstString, secondString string, countFirstString, countSecondString int) (result string) {
	if countFirstString%2 == 0 {
		result = fmt.Sprintf("%s%s%s", string([]rune(firstString)[:countFirstString/2]), secondString, string([]rune(firstString)[countFirstString/2:]))
		fmt.Println("Ответ:", result)
	} else {
		result = fmt.Sprintf("%s%s%s", string([]rune(secondString)[:countSecondString/2]), firstString, string([]rune(secondString)[(countSecondString/2)+1:]))
		fmt.Println("Ответ:", result)
	}
	return
}

/*Вводятся 2 строки (любые) сравнить количество символов в строках и вывести только ту, в которой символов меньше.
Сразу предупреждаю, что буду вводить и сравнивать русскую строку с английской.
А если строки одинковые, то вывести одну в середине другой (если число символов нечётное, то заменить средний символо строкой)*/
