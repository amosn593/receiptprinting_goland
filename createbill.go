package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getUserInput(prompt string, r *bufio.Reader) (string, error){
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func promptOptions(b *Bill){
	reader := bufio.NewReader(os.Stdin)

	option, _ := getUserInput("Choose Option (a - addItem, s - saveBill, t - addTip): ", reader)

	switch option {
	case "a":
		name, _ := getUserInput("Item Name: ", reader)

		price, _ := getUserInput("Item Price ($): ", reader)

		priceInFloat64, err := strconv.ParseFloat(price, 64)

		if err != nil{
			fmt.Println("Price must a valid number...")
			promptOptions(b)
		}

		b.AddItems(name, priceInFloat64)
		
		fmt.Println("Item Added...",name, price)

		promptOptions(b)
	case "s":
		fmt.Println("You chose to save the bill...\n ", b.format())
		b.Save()
	case "t":
		tip, _ := getUserInput("Enter Tip Amount ($): ", reader)

		tipInFloat64, err := strconv.ParseFloat(tip, 64)

		if err != nil{
			fmt.Println("Tip must a valid number...")
			promptOptions(b)
		}

		b.UpdateTip(tipInFloat64)

		fmt.Println("Tip Added...", tipInFloat64)

		promptOptions(b)

	default:
		fmt.Println("That is not a valid option...")
		promptOptions(b)
	}
	
}

func createBill() Bill {
	reader := bufio.NewReader(os.Stdin)

	name, _ := getUserInput("Create a new Bill, Name: ", reader)
	
	name = strings.TrimSpace(name)

	newbill := NewBill(name)

	fmt.Println("Created Bill... ", newbill.Name)

	return  newbill

}