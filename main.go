package main

import (
	"fmt"
	"strconv"
	"strings"
	"bufio"
	"os"
)

var (
	table = [][]string{
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "C"},
	}
)

func arabToRoman(arab int) string {
	if arab==0{
		panic("В системе римских цифр отсутствует 0")
	}
	const maxNumber = 100
	if arab <= 0 || arab > maxNumber {
		return "" 
	}
	var (
		roman = ""
		digit = 100
	)
	for i := 2; i >= 0; i-- {
		d := arab / digit
		roman += table[i][d]
		arab %= digit
		digit /= 10
	}
	return roman
}

func romanToArab(roman string) int {
	for i := 1; i <= 9; i++ {
		if table[0][i] == roman {
			return i
		} else if table[1][1] == roman{
			return 10
		} else if i==9{
			panic("Вывод ошибки, цифры принимаются от 1 до 10 включительно")
		} else if roman==strconv.Itoa(i)|| roman=="10"{
			panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
		}
	}
	return 10
}

func calc(operation string) int {
	var x,a,b int
	var operant string
	operation = strings.ReplaceAll(operation, " ", "")
	for i:=0;i<=len(operation)-1;i++{
		if strings.Count(string(operation[i]), "+")==1 || strings.Count(string(operation[i]), "-")==1 || strings.Count(string(operation[i]), "*")==1 || strings.Count(string(operation[i]), "/")==1 {
			x=i
			break
		} else if strings.Count(string(operation), "+")==2 || strings.Count(string(operation), "-")==2 || strings.Count(string(operation), "*")==2 || strings.Count(string(operation), "/")==2 {
			panic("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
		} else if i==len(operation)-1{
			panic("Вывод ошибки, так как строка не является математической операцией.")
		}
	}
	for i:=1;i<=9;i++{
		if operation[0:x]==strconv.Itoa(i) || operation[0:x]=="10" {
			a,b,operant = divide(operation,x)
			break
		}

	}
	if operant==""{
		a=romanToArab(operation[0:x])
		b=romanToArab(operation[x+1:len(operation)-2])
		if string(operation[x])=="-" && a<b{
			panic("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
		}
		operant=string(operation[x])
	}
	switch operant {
		case "+":
			return a + b
		case "-":
			return a - b
		case "*":
			return a * b
		case "/":
			return a / b
		}
	return 0
}

func divide(operation string,x int)(a int, b int, operant string){
	a, _ = strconv.Atoi(string(operation[0:x]))
	b, _ = strconv.Atoi(string(operation[x+1:len(operation)-2]))
	operant=string(operation[x])
	if a>10 || b>10{
		panic("Вывод ошибки, цифры принимаются от 1 до 10 включительно")
	} else {
		for i:=1;i<=9;i++{
			if b==i || b==10{
				break
			}else if i==9{
				panic("Вывод ошибки, так как используются одновременно разные системы счисления.")
			}
		}
	}
	return a, b , operant
} 

func main() {
	fmt.Println("Введите число операция число:")
	operation, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	result:=calc(operation)
	for i:=48;i<=57;i++{
		if operation[0]==byte(i){
			fmt.Println(result)
			break
		} else if i==57{
			fmt.Println(arabToRoman(result))
		}
	} 
}