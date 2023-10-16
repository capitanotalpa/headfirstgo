package main

import (
	"fmt"
	"log"
)

func find(item string, slice []string) bool {
	for _, sliceItem := range slice {
		if item == sliceItem {
			return true
		}
	}
	return false
}

type Refridgerator []string

func (r Refridgerator) Open() {
	fmt.Println("Opening fridge...")
}

func (r Refridgerator) Close() {
	fmt.Println("Closing fridge...")
}

func (r Refridgerator) FindFood(food string) error {
	r.Open()
	defer r.Close()
	if find(food, r) {
		fmt.Println("Found", food)
	} else {
		return fmt.Errorf("%s not found", food)
	}
	return nil
}

func main() {
	fridge := Refridgerator{"Milk", "Pizza", "Salsa"}
	for _, food := range []string{"Milk", "Bananas"} {
		err := fridge.FindFood(food)
		if err != nil {
			log.Fatal(err)
		}
	}
}
