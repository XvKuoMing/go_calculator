package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите арифметическую операцию")
		text, _ := reader.ReadString('\n')
		fmt.Println(parse_and_calculate(text))
	}

}

func parse_and_calculate(t string) string {
	parsed := strings.Fields(parse(t)) // разделяем строчку по пробелам
	if len(parsed) != 3 {
		if len(parsed) > 3 {
			panic("only 2 operands and 1 operator are awaited")
		}
		if len(parsed) == 1 {
			panic("no ariphmetic expression has been given")
		}
		if len(parsed) == 2 {
			panic("Incomplete expression")
		}
	}

	// declaring variables
	var operation string = parsed[1]

	// checking rome numbers
	var num1_is_rome bool = is_rome_number(parsed[0])
	var num2_is_rome bool = is_rome_number(parsed[2])
	var num1 int = 0
	var num2 int = 0

	if num1_is_rome || num2_is_rome {
		if num1_is_rome != num2_is_rome {
			panic("if rome number is given, other number must also be a rome number")
		} else {
			// translatin rome number into int
			num1 = rome2int(parsed[0])
			num2 = rome2int(parsed[2])
		}
	} else {
		_num1, _ := strconv.Atoi(parsed[0])
		_num2, _ := strconv.Atoi(parsed[2])
		num1 += _num1
		num2 += _num2
	}

	var res int
	// performing calculations
	if (0 < num1 && num1 <= 10) && (0 < num2 && num2 <= 10) {

		if operation == string("+") {
			res = num1 + num2
		}
		if operation == string("-") {
			res = num1 - num2
		}

		if operation == string("/") {
			res = num1 / num2
		}

		if operation == string("*") {
			res = num1 * num2
		}

		if num1_is_rome {
			return int2rome(res)
		} else {
			return fmt.Sprint(res)
		}

	} else {
		panic("Both numbers must be more than 0 and less (or equal) than 10")
	}
}

func parse(expr string) string {
	// парсит строчку, проверяет ее синтакси и удовлетворению условию (2 операнда 1 операнд)
	// и приводит к единому формату: число пробел операнд пробел число
	var charnum string = "0123456789IVX" // все допустимые числа
	var parsed string
	var need_number bool = true
	expr = strings.ReplaceAll(expr, " ", "") // deleting all whitespaces
	for index, char := range expr {
		if need_number {
			if strings.Contains(charnum, string(char)) {
				parsed += string(char)
				if (len(string(expr)) > index+1) && !strings.Contains(charnum, string(expr[index+1])) {
					need_number = false
				}
			}
		} else {
			if strings.Contains("+-/*", string(char)) {
				parsed += string(' ') + string(char) + string(' ')
				need_number = true
			}

		}

	}
	return parsed
}

func is_rome_number(r string) bool {
	const rome_numbers string = "XVI"
	var res bool = true
	for _, el := range r {
		if !strings.Contains(rome_numbers, string(el)) {
			res = false
		}
	}
	return res
}

func rome2int(r string) int {
	// translation of rome number to arabic
	if strings.Contains(r, "IIII") {
		panic("Incorrect romanian numeral: IIII")
	}

	var n int = 0
	for _, rome_number := range r {
		if string(rome_number) == string("I") {
			n += 1
		}
		if string(rome_number) == string("V") {
			n = 5 - n
		}
		if string(rome_number) == string("X") {
			n = 10 - n
		}
	}

	return n
}

func int2rome(m int) string {
	//translation of integer into rome
	// note that out maximum value if 100
	if m <= 0 {
		panic("the result is non-positive number, cannot calculate using romanian numerals")
	}

	num := [9]int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	sym := [9]string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res string // empty string
	var count int
	for index, div := range num {
		count = m / div
		for count > 0 {
			res += sym[index]
			count -= 1
		}
		m %= div
	}
	return res
}
