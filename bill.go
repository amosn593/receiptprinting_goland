package main

import (
	"fmt"
	"os"
)

type Bill struct {
	Name  string
	Items map[string]float64
	Tip   float64
}

func NewBill(name string) Bill {
	bill := Bill{
		Name:  name,
		Items: map[string]float64{}, // Initialize the map, or map[string]float64{}
		Tip:   0,
	}
	return bill
}

// Receiver function on bill
//Format Bill
func (b Bill) format() string {
	fs := "Bill BreakDown: \n"
	var total float64 = 0

	for k, v := range b.Items {
		fs += fmt.Sprintf("%-25v ...$%0.2f \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f \n", "tip:", b.Tip)

	//total
	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.Tip)

	return  fs;
}

//Update Tip
func (b *Bill) UpdateTip(tip float64){
	b.Tip = tip
}

//Add Items to the bill
func (b *Bill) AddItems(name string, price float64){
	b.Items[name] = price
}

//Save The Bill to a text file
func (b *Bill) Save(){
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.Name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}

	fmt.Println("Bill saved to a file... ", b.Name+".txt")
}