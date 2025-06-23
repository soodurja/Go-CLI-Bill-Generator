package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')
	input = strings.TrimSpace(input)

	return input, err
}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput("Create a new bill name: ", reader)

	b := newBill(name)
	fmt.Println("Created the bill -", b.name)

	return b
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)

	option, _ := getInput("Choose an option (a - add an item, s - save the bill, t - add a tip): ", reader)

	switch option {
	case "a":
		item, _ := getInput("Enter item name: ", reader)
		price, _ := getInput("Enter item price ($): ", reader)

		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("Please enter a number for the item price ($).")
			promptOptions(b)
			return
		}
		b.addItem(item, p)
		fmt.Printf("You added - %v: $%0.2f\n", item, p)
		promptOptions(b)

	case "s":
		b.saveBill()
		fmt.Println("Saved the bill -", b.name)

	case "t":
		tip, _ := getInput("Enter tip amount ($): ", reader)

		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("Please enter a number for the tip amount ($).")
			promptOptions(b)
			return
		}
		b.updateTip(t)
		fmt.Printf("You added - Tip: $%0.2f\n", t)
		promptOptions(b)

	default:
		fmt.Println("Please enter a valid option.")
		promptOptions(b)
	}
}

func main() {
	for {
		firstBill := createBill()
		promptOptions(firstBill)

		reader := bufio.NewReader(os.Stdin)

		response, _ := getInput("Do you want to create another bill? (y/n): ", reader)

		switch response {
		case "y":
			continue
		case "n":
			fmt.Println("Thank you for using the bill generator!")
			return
		default:
			fmt.Println("Please enter a valid option.")
			response, _ = getInput("Do you want to create another bill? (y/n): ", reader)
		}

	}
}
