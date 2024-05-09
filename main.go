package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// Функция конвертации римских чисел в арабские
func romToAr(romanNum string) int {
	var romanNumerals = map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
		"VII": 7, "VIII": 8, "IX": 9, "X": 10, "L": 50, "C": 100,
	}
	return romanNumerals[romanNum]
}

// Функция конвертации арабских чисел в римские с помощью StringBuilder
func arToRom(arabicNum int) string {
	var result strings.Builder

	// Определяем константы для римских цифр
	romanSymbols := []struct {
		value  int
		symbol string
	}{
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	// Записываем римские цифры
	for _, symbol := range romanSymbols {
		for arabicNum >= symbol.value {
			result.WriteString(symbol.symbol)
			arabicNum -= symbol.value
		}
	}

	return result.String()
}

// Function checks if the input string is a valid Roman numeral.
func isRoman(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) ||
			(unicode.ToUpper(char) != 'I' &&
				unicode.ToUpper(char) != 'V' &&
				unicode.ToUpper(char) != 'X') {
			return false
		}
	}
	return true
}

func main() {
	var a, b string
	var s string
	var err = ""
	fmt.Println("Введите ваш пример для решения римскими или арабскими цифрами\n" + "Введите выражение:")
	fmt.Fscanln(os.Stdin, &a, &s, &b, &err)

	if err != "" {
		panic("Вы ввели что-то лишнее, попробуйте \"число операнд число\"")
		return
	}
	if isRoman(a) != isRoman(b) {
		panic("Цифры должны быть одного типа (римские ИЛИ арабские)")
		return
	}

	// Преобразуем строки a и b в числа numA и numB соответственно,
	// используя romToAr для римских чисел и strconv.Atoi для арабских.
	var numA, numB int
	if isRoman(a) {
		numA = romToAr(a)
	} else {
		numA, _ = strconv.Atoi(a)
	}
	if isRoman(b) {
		numB = romToAr(b)
	} else {
		numB, _ = strconv.Atoi(b)
	}

	// Проверка на то, что числа входят в диапазон от 1 до 10
	// Этой проверкой убираем необходимость проверки при делении на 0
	if numB > 10 || numB < 1 || numA > 10 || numA < 1 {
		panic("Используйте только числа от 1 до 10")
		return
	}

	var result int

	// Вычисление результата
	switch s {
	case "+":
		result = numA + numB
	case "-":
		result = numA - numB
	case "*":
		result = numA * numB
	case "/":
		result = numA / numB
	default:
		panic("Неверный операнд")
		return
	}

	// Вывод результата
	if isRoman(a) && isRoman(b) {
		var e = arToRom(result)
		if e == "" {
			panic("Результат вычисления римских чисел меньше единицы!")
			return
		}
		fmt.Println("Результат:", arToRom(result))
	} else {
		fmt.Println("Результат:", result)
	}
}
