package module_2

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*Напишите функцию sumInt, принимающую переменное количество аргументов
типа int, и возвращающую количество полученных функцией аргументов и их сумму*/

func SumInt(nums ...int) (int, int) {
	sum := 0

	for _, num := range nums {
		sum += num
	}
	return len(nums), sum
}

/*Напишите "функцию голосования", возвращающую то значение (0 или 1),
которое среди значений ее аргументов x, y, z встречается чаще*/

func Vote(x int, y int, z int) int {
	if x+y+z <= 1 {
		return 0
	}

	return 1
}

/*Последовательность Фибоначчи определена следующим образом: φ1=1,
φ2=1, φn=φn-1+φn-2 при n>1. Начало ряда Фибоначчи выглядит следующим
образом: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... Напишите функцию,
которая по указанному натуральному n возвращает φn*/

func Fibonacci(NumberInTheFibonacciSequence int) int {
	var firstNumber, secondNumber = 1, 1
	var result int

	fmt.Scan(&NumberInTheFibonacciSequence)

	for searchIndex := 2; searchIndex < NumberInTheFibonacciSequence; searchIndex++ {
		result = firstNumber + secondNumber

		firstNumber = secondNumber

		secondNumber = result
	}

	if NumberInTheFibonacciSequence == 1 || NumberInTheFibonacciSequence == 2 {
		return 1
	}

	return result
}

//Поменяйте местами значения переменных на которые ссылаются указатели.

func ExchangeOfMeanings(x1 *int, x2 *int) {
	*x1, *x2 = *x2, *x1
	fmt.Print(*x1, *x2)
}

/*На вход подается строка. Нужно определить, является ли она правильной или нет.
Правильная строка начинается с заглавной буквы и заканчивается точкой. Если
строка правильная - вывести Right иначе - вывести Wrong*/

func CheckStringConditions() {
	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	runes := []rune(text)

	if unicode.IsUpper(runes[0]) && runes[len(runes)-1] == '.' {
		fmt.Println("Right")
		return
	}

	fmt.Println("Wrong")
}

/*На вход подается строка. Нужно определить, является ли она палиндромом.
Если строка является палиндромом - вывести Палиндром иначе - вывести Нет.
 Все входные строчки в нижнем регистре*/

func PalindromeOrNot() {
	var mainString string

	fmt.Scanln(&mainString)

	runeString := []rune(mainString)

	for index := range runeString {
		if runeString[index] != runeString[len(runeString)-index-1] {
			fmt.Print("Нет")
			return
		}
	}

	fmt.Println("Палиндром")
}

/* На вход дается строка, из нее нужно сделать другую строку,
оставив только нечетные символы (считая с нуля)*/

func DoNewString() {
	var mainString, resultString string

	fmt.Scan(&mainString)
	runes := []rune(mainString)

	for index := range runes {
		if index%2 != 0 {
			resultString += string(runes[index])
		}
	}

	fmt.Print(resultString)
}

/*Дается строка. Нужно удалить все символы, которые встречаются
более одного раза и вывести получившуюся строку*/

func CheckRepeatedChar() {
	var mainString string

	fmt.Scan(&mainString)

	var result []rune
	for _, char := range mainString {
		if strings.Count(mainString, string(char)) == 1 {
			result = append(result, char)
		}
	}

	fmt.Print(string(result))
}

/*Ваша задача сделать проверку подходит ли пароль вводимый пользователем
под заданные требования. Длина пароля должна быть не менее 5 символов, он
должен содержать только арабские цифры и буквы латинского алфавита. На вход
подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok",
иначе вывести "Wrong password"*/

func PasswordIsCorrect() {
	var (
		passwordString string
		alphabet       string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
		flag           bool   = true
	)

	fmt.Scan(&passwordString)

	passwordRune := []rune(passwordString)

	if len(passwordRune) < 5 {
		flag = false
		fmt.Print("Wrong password")
		return
	}

	for checkIndex := range passwordRune {
		if !strings.Contains(alphabet, string(passwordRune[checkIndex])) {
			flag = false
			break
		}

		if unicode.IsSpace(passwordRune[checkIndex]) {
			flag = false
			break
		}

		if !unicode.IsLetter(passwordRune[checkIndex]) && !unicode.IsDigit(passwordRune[checkIndex]) {
			flag = false
			break
		}
	}

	if flag {
		fmt.Print("Ok")
		return
	}

	fmt.Print("Wrong password")
}

/*На вход подаются a и b - катеты прямоугольного треугольника.
Нужно найти длину гипотенузы*/

func FindHypotenuse() {
	var leg1, leg2 float64

	fmt.Scan(&leg1, &leg2)

	fmt.Print(math.Hypot(leg1, leg2))
}

/*Дана строка, содержащая только английские буквы (большие и маленькие).
Добавить символ ‘*’ (звездочка) между буквами (перед первой буквой и
после последней символ ‘*’ добавлять не нужно)*/

func addCharacterBetweenCharacters() {
	var mainString string

	fmt.Scan(&mainString)

	resultString := strings.Replace(mainString, "", "*", -1)

	fmt.Print(strings.Trim(resultString, string(resultString[0])))
}

/*Дана строка, содержащая только арабские цифры. Найти
и вывести наибольшую цифру*/

func FindLargestDigitInString() {
	var (
		mainString string
		max        int
	)

	fmt.Scan(&mainString)

	runeString := []rune(mainString)

	for searchIndex := 0; searchIndex < len(runeString); searchIndex++ {
		if runeString[searchIndex] >= '0' && runeString[searchIndex] <= '9' {

			digit, _ := strconv.Atoi(string(runeString[searchIndex]))

			if digit > max {
				max = digit
			}

			if max == 9 {
				fmt.Print(max)
				return
			}
		}
	}

	fmt.Print(max)
}

/*На вход подается целое число. Необходимо возвести в квадрат 
каждую цифру числа и вывести получившееся число. */

func DigitalMultiplication() {
	var mainString, resultString string

	fmt.Scan(&mainString)
	runeString := []rune(mainString)

	for multiplicationIndex := range runeString {
		if runeString[multiplicationIndex] >= '0' && runeString[multiplicationIndex] <= '9' {
			digit, _ := strconv.Atoi(string(runeString[multiplicationIndex]))
			resultString += strconv.Itoa(digit * digit)
		}
	}

	fmt.Print(resultString)
}