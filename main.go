package main

import "fmt"

func main(){
	fmt.Println("Receipt Printing Console App")

	mybill := createBill()

	promptOptions(&mybill)

	//fmt.Println(mybill)
}