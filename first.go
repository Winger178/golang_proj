// first
package main

import (
	"fmt"
	//"reflect"
	"strconv"
	"strings"
)

func main() {

	input := ""
	fmt.Scan(&input)
	result, result_string := calculate(input)
	//Сравнение флажков Арабского и Римского ответа
	if result_string == "Арабский ответ" {
		fmt.Println(result)
	} else if result_string == "Римский ответ" {
		fmt.Println(convertToRoman(result))
	} else {
		fmt.Println(result, result_string)
	}

}

// Функция для конвертирования числа в римскую цифру
func convertToRoman(num int) string {
	// Определение символов и их соответствующих значений
	romanSymbols := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	romanValues := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}

	// Переменная для хранения результата
	roman := ""

	// Проход по символам и значениям
	for i := 0; i < len(romanSymbols); i++ {
		// Пока число больше или равно текущему значению,
		// добавляем символ к римской цифре и вычитаем значение из числа
		for num >= romanValues[i] {
			roman += romanSymbols[i]
			num -= romanValues[i]
		}
	}

	return roman
}

// функция для конвертировнаия римского в арабское
func convertToArabic(roman string) int {
	// Определение символов и их соответствующих значений
	romanSymbols := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	arabicValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Поиск римского числа в слайсе символов
	for i := 0; i < len(romanSymbols); i++ {
		if roman == romanSymbols[i] {
			return arabicValues[i]
		}
	}

	return -1 // Возвращаем -1, если римское число не найдено
}

// Функция для калькулирования
func calculate(input string) (int, string) {
	// Удаляем все пробелы из строки
	input = strings.ReplaceAll(input, " ", "")

	// Разделяем строку на операнды и оператор
	operatorIndex := strings.IndexAny(input, "+-*/")
	leftOperandStr := input[:operatorIndex]
	operator := input[operatorIndex]
	rightOperandStr := input[operatorIndex+1:]

	// Преобразуем операнды из строкового формата в числовой

	leftOperand, _ := strconv.Atoi(leftOperandStr)
	rightOperand, _ := strconv.Atoi(rightOperandStr)

	var result int = 0
	//if reflect.TypeOf(leftOperand) == int && reflect.TypeOf(rightOperand) == int {
	// Вычисляем результат в зависимости от оператора
	if 1 <= leftOperand && leftOperand <= 10 && 1 <= rightOperand && rightOperand <= 10 {

		switch operator {
		case '+':
			result = leftOperand + rightOperand
		case '-':
			result = leftOperand - rightOperand
		case '*':
			result = leftOperand * rightOperand
		case '/':
			result = leftOperand / rightOperand
		}
		return result, "Арабский ответ"
	}
	//}

	//if reflect.TypeOf(convertToArabic(leftOperand)) == int && reflect.TypeOf(convertToArabic(rightOperand)) == int {
	//if 1 <= leftOperand && leftOperand <= 10 && 1 <= rightOperand && rightOperand <= 10 {
	if 1 <= convertToArabic(leftOperandStr) && convertToArabic(leftOperandStr) <= 10 && 1 <= convertToArabic(rightOperandStr) && convertToArabic(rightOperandStr) <= 10 {
		switch operator {
		case '+':
			result = convertToArabic(leftOperandStr) + convertToArabic(rightOperandStr)
		case '-':
			result = convertToArabic(leftOperandStr) - convertToArabic(rightOperandStr)
		case '*':
			result = convertToArabic(leftOperandStr) * convertToArabic(rightOperandStr)
		case '/':
			result = convertToArabic(leftOperandStr) / convertToArabic(rightOperandStr)
		}
		if result > 0 {
			return result, "Римский ответ"
		} else {
			return result, "Не бывает римских чисел меньше либо равны нулю!"
		}
	}
	//}
	//}
	return result, "Неправильно введеный формат данных"
}
