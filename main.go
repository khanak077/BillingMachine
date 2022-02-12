package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var quant string
	var product string
	var number int
	var tray = []int{}

	for true {
		fmt.Println("---------MENU----------")
		fmt.Print("C: coffee, 40rs\n D: dosa, 80 rs\n T: tomato soup, 20rs\n I : idli , 20rs\n V : vada, 25rs.\n B: bhature &chhole, 30rs\n P: paneer pakoda, 30rs\n M: manchurian, 90rs\n H: hakka noodle, 70rs.\n F: French fries, 30rs\n J: jalebi, 30rs\n L: lemonade, 15rs\n S: spring roll, 20rs\n")
		fmt.Println("-----------------------")
		fmt.Println("Press 'END' For closing the day.")
		fmt.Println("Please place the order: ")
		fmt.Scan(&quant)
		quant = strings.ToUpper(quant)

		if quant != "END" {
			number, _ = strconv.Atoi(quant)

			fmt.Scan(&product)
			product = strings.ToUpper(product)

			bill := billing(number, product)

			fmt.Println("--> --> --> --> --> --> -->")
			fmt.Println("TOTAL BILL :- ", bill)
			fmt.Println("--> --> --> --> --> --> -->")
			tray = append(tray, bill)

		} else {
			EarnedAmount(tray)
		}

	}
}

func EarnedAmount(earn []int) {
	var sum int = 0

	for i := 0; i < len(earn); i++ {
		sum = earn[i] + sum
	}
	fmt.Println("Total Income for the day is : ", sum)
	os.Exit(0)
}

func billing(quantity int, order string) int {
	point := map[string]int{
		"C": 40,
		"D": 80,
		"T": 20,
		"I": 20,
		"V": 25,
		"B": 30,
		"P": 30,
		"M": 90,
		"H": 70,
		"F": 30,
		"J": 30,
		"L": 15,
		"S": 20,
	}
	total := quantity * int(point[order])
	return int(total)
}
