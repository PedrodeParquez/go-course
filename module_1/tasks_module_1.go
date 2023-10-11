package module_1

import (
	"fmt"
	"strconv"
)

/* Часовая стрелка повернулась с начала суток на d градусов.
Определите, сколько сейчас целых часов h и целых минут m*/

func TimeByDegrees() {
	var time int

	fmt.Scan(&time)

	hours := time / 30
	min := 2 * (time % 30)

	fmt.Print("It is ", hours, " hours ", min, " minutes.")
}

// Требуется определить, является ли данный год високосным.

func LeapYearOrNot() {
	var year int

	fmt.Scan(&year)

	if year <= 0 || year > 10000 {
		fmt.Println("Incorrect input! Try again.")
		LeapYearOrNot()
	}

	if year%400 == 0 || (year%100 != 0 && year%4 == 0) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

//Поменять 2 числа с указанными индексами местами в массиве

func SwapTwoElements() {
	var firstNumber, secondNumber uint8
	var workArray [10]uint8

	for inputIndex := 0; inputIndex < len(workArray); inputIndex++ {
		fmt.Scan(&workArray[inputIndex])
	}

	for perIndex := 0; perIndex < 3; perIndex++ {
		fmt.Scan(&firstNumber, &secondNumber)

		tempVar := workArray[firstNumber]

		workArray[firstNumber] = workArray[secondNumber]

		workArray[secondNumber] = tempVar
	}

	for outputIndex := 0; outputIndex < len(workArray); outputIndex++ {
		fmt.Print(workArray[outputIndex], " ")
	}
}

/*Подаются 5 числовых значений, которые записываются в массив.
Необходимо найти максимальное число в этом массиве*/

func MaxNumberArray() {
	var number, maxNumber int
	array := [5]int{}

	for inputIndex := 0; inputIndex < 5; inputIndex++ {
		fmt.Scan(&number)
		array[inputIndex] = number
	}

	for searchIndex := range array {
		if searchIndex == 0 {
			maxNumber = array[searchIndex]
		}

		if array[searchIndex] > maxNumber {
			maxNumber = array[searchIndex]
		}
	}

	fmt.Print(maxNumber)
}

/*Дан массив, состоящий из целых чисел. Нумерация элементов
начинается с 0. Напишите программу, которая выведет элементы
массива, индексы которых четны*/

func FindElementsEvenIndices() {
	var amountNumbers int

	fmt.Scan(&amountNumbers)

	var array []int = make([]int, int(amountNumbers))

	for inputIndex := range array {
		fmt.Scan(&array[inputIndex])
	}

	for searchIndex := range array {
		if searchIndex%2 == 0 {
			fmt.Print(array[searchIndex], " ")
		}
	}
}

/*Дана последовательность, состоящая из целых чисел. Напишите
программу, которая подсчитывает количество положительных чисел
среди элементов последовательности*/

func FindAmountPositiveNumbers() {
	var positiveNumbers, amountNumbers int

	fmt.Scan(&amountNumbers)

	var array []int = make([]int, int(amountNumbers))

	for inputIndex := range array {
		fmt.Scan(&array[inputIndex])
	}

	for searchIndex := range array {
		if array[searchIndex] > 0 {
			positiveNumbers++
		}
	}

	fmt.Println(positiveNumbers)
}

/*Идёт k-я секунда суток. Определите, сколько целых
часов h и  целых минут m прошло с начала суток.
Например, если k=13257=3*3600+40*60+57, то h=3 и m=40*/

func TimeFromSeconds() {
	var dailySecond int

	fmt.Scan(&dailySecond)

	numberOfHours := dailySecond / 3600

	numberOfMinutes := (dailySecond - numberOfHours*3600) / 60

	fmt.Print("It is ", numberOfHours, " hours ", numberOfMinutes, " minutes.")
}

/*По данному числу N распечатайте все целые значения
степени двойки, не превосходящие N, в порядке возрастания*/

func PowersTwo() {
	var (
		number int
		power  = 1
	)

	fmt.Scan(&number)

	for power <= number {
		fmt.Print(power, " ")
		power *= 2
	}
}

/*Дано натуральное число N. Выведите его представление в двоичном виде*/

func BinaryFormat() {
	var number int64

	fmt.Scan(&number)

	binaryNumber := strconv.FormatInt(number, 2)

	fmt.Println(binaryNumber)
}

/*Из натурального числа удалить заданную цифру*/

func DelSelectNumber() {
	var digit, number, result string

	fmt.Scan(&number, &digit)

	for searchIndex := range number {
		if string(number[searchIndex]) != digit {
			result += string(number[searchIndex])
		}
	}

	fmt.Print(result)
}

//По данному числу определить его цифровой корень

func FindDigitRoot() {
	var number, result int

	fmt.Scan(&number)

	for number > 9 {
		result = number % 10
		number /= 10
		number += result
	}

	fmt.Print(number)
}

/*Дано натуральное число A > 1. Определите, каким по счету числом Фибоначчи
оно является, то есть выведите такое число n, что φn=A. Если А не является
числом Фибоначчи, выведите число -1*/

func FibonacciNumber() {
	var fibonacciNumber, result int
	firstNumber, secondNumber := 1, 1

	fmt.Scan(&fibonacciNumber)

	for searchIndex := 1; searchIndex < fibonacciNumber; searchIndex++ {
		result = firstNumber + secondNumber

		firstNumber = secondNumber
		secondNumber = result

		if result == fibonacciNumber {
			fmt.Print(searchIndex + 2)
			return
		}
	}

	fmt.Print("-1")
}

/*По данному числу n закончите фразу "На лугу пасется..." одним из возможных
продолжений: "n коров", "n корова", "n коровы", правильно склоняя слово "корова"*/

func DeclensionWordCow() {
	var numberCows int

	fmt.Scan(&numberCows)

	switch {
	case numberCows == 1:
		fmt.Print(numberCows, " korova")
	case (numberCows > 1 && numberCows < 5):
		fmt.Print(numberCows, " korovy")
	case numberCows >= 5 && numberCows <= 20:
		fmt.Print(numberCows, " korov")
	case numberCows > 20 && numberCows < 100 && numberCows%10 == 1:
		fmt.Print(numberCows, " korova")
	case numberCows > 21 && numberCows < 100 && numberCows%10 > 1 && numberCows%10 < 5:
		fmt.Print(numberCows, " korovy")
	case numberCows > 24 && numberCows < 100:
		fmt.Print(numberCows, " korov")
	default:
		fmt.Print("Invalid input. Try again!")
	}
}