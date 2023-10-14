package module_3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

/*Проверить содержиться ли данный ключ в map, если да, то вывести его,
если нет то добавить ключ в map и вывести его*/

// Функция из условия задачи
func work(x int) int

func AddingAndComparingKeys() {
	mapNumbers := make(map[int]int)
	var number int

	for i := 0; i < 10; i++ {
		fmt.Scan(&number)

		if _, inMap := mapNumbers[number]; !inMap {
			mapNumbers[number] = work(number)
		}

		fmt.Print(mapNumbers[number], " ")
	}
}

/*Требуется исправить cityPopulation, чтобы в ней была сохранена
информация только о городах из группы groupCity[100], то есть
сравнивать ключ cityPopulation cо значениями в groupCity с ключом 100*/

func mapsFix() {
	groupCity := map[int][]string{
		10:   []string{"Село", "Деревня", "ПГТ"},
		100:  []string{"Город", "Станица"},
		1000: []string{"Мегаполис", "Миллионник"},
	}

	cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	}

	var found bool

	for city := range cityPopulation {
		found = false

		for _, validCity := range groupCity[100] {
			if validCity == city {
				found = true
				break
			}
		}
		if !found {
			delete(cityPopulation, city)
		}
	}
}

/*Нужно получить 2 строки, очистить их от всего кроме цифр,
конвертировать в int64 и вернуть сумму получившихся чисел*/

func Adding(num1 string, num2 string) int64 {
	num1, num2 = extractNumbers(num1), extractNumbers(num2)

	number1, err := strconv.ParseInt(num1, 10, 64)
	if err != nil {
		return 0
	}

	number2, err := strconv.ParseInt(num2, 10, 64)
	if err != nil {
		return 0
	}

	return number1 + number2
}

func extractNumbers(str string) string {
	var resultString string

	for _, char := range str {
		if unicode.IsDigit(char) {
			resultString += string(char)
		}
	}

	return resultString
}

/*На стандартный ввод вы получаете 2 таких вещественных числа,
в качестве результата требуется вывести частное от деления первого
числа на второе с точностью до четырех знаков после "запятой".
Пример получаемых данных: 1 149,6088607594936;1 179,0666666666666*/

func ConversionReceivedData() {
	var number1, number2 string

	reader := bufio.NewReader(os.Stdin)
	inputString, _, _ := reader.ReadLine()

	str := string(inputString)
	str = strings.Replace(strings.Replace(str, " ", "", -1), ",", ".", -1)

	for searchIndex, char := range str {
		if char == ';' {
			number2 = str[searchIndex+1:]
			break
		}

		number1 += string(char)
	}

	dividend, err := strconv.ParseFloat(number1, 64)
	if err != nil {
		fmt.Println("Ошибка преобразовании:", err)
	}

	divider, err := strconv.ParseFloat(number2, 64)
	if err != nil {
		fmt.Println("Ошибка преобразовании:", err)
	}

	fmt.Printf("%.4f", dividend/divider)
}

/*Анонимная функция на вход получает целое положительное число (uint),
возвращает число того же типа в котором отсутствуют нечетные
цифры и цифра 0. Если же получившееся число равно 0, то
возвращается число 100. Цифры в возвращаемом числе имеют
тот же порядок, что и в исходном числе*/

func AnonymousFunc() {
	var number uint

	fn := func(number uint) uint {
		if number == 0 {
			return 100
		}

		result := uint(0)
		multiplier := uint(1)

		for number > 0 {
			digit := number % 10

			if digit%2 == 0 && digit != 0 {
				result += digit * multiplier
				multiplier *= 10
			}

			number /= 10
		}

		if result == 0 {
			return 100
		}

		return result
	}

	fmt.Scan(&number)

	fmt.Print(fn(number))
}

/*На стандартный ввод вы получаете строку, состоящую ровно из 10
цифр: 0 или 1 (порядок 0/1 случайный). Ваша задача считать эту строку любым
возможным способом и создать на основе этой строки объект объявленного вами
на первом этапе типа: надеюсь, вы понимаете, что строка символизирует емкость
батарейки: 0 - это "опустошенная" часть, а 1 - "заряженная"*/

type Battery struct {
	Capacity string
}

func (b Battery) String() string {
	numberZeros := strings.Repeat("0", strings.Count(b.Capacity, "0"))
	numberUnits := strings.Repeat("1", strings.Count(b.Capacity, "1"))

	numberZeros = strings.Replace(numberZeros, "0", " ", -1)
	numberUnits = strings.Replace(numberUnits, "1", "X", -1)

	return "[" + numberZeros + numberUnits + "]"
}

func FindOutCharge() {
	var input string
	fmt.Scan(&input)

	batteryForTest := Battery{Capacity: input}

	fmt.Print(batteryForTest)
}
