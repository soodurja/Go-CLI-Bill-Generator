package main

import (
	"fmt"
	"os"
	"strings"
)

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// Makes new bills
func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}

	return b
}

// Format the bill
func (b *bill) format() string {
	var total float64 = 0
	fs := "Bill breakdown: \n"

	var entries []string
	maxEntryLength := 0

	for k, v := range b.items {
		itemPrice := fmt.Sprintf("%0.2f", v)
		entry := fmt.Sprintf("%-25v $%10v\n", k+":", itemPrice)
		entries = append(entries, entry)
		total += v

		if len(entry) > maxEntryLength {
			maxEntryLength = len(entry)
		}
	}

	tipPrice := fmt.Sprintf("%0.2f", b.tip)
	tipEntry := fmt.Sprintf("%-25v $%10v\n", "Tip:", tipPrice)

	totalPrice := fmt.Sprintf("%0.2f", total+b.tip)
	totalEntry := fmt.Sprintf("%-25v $%10v", "Total:", totalPrice)

	if len(totalEntry) > maxEntryLength {
		maxEntryLength = len(totalEntry)
	}
	if len(tipEntry) > maxEntryLength {
		maxEntryLength = len(tipEntry)
	}

	// List all items
	for _, addEntry := range entries {
		fs += addEntry
	}

	// Add the separator
	fs += strings.Repeat("-", maxEntryLength) + "\n"

	// Add the tip
	fs += tipEntry

	// Add the separator
	fs += strings.Repeat("-", maxEntryLength) + "\n"

	// Add the total
	fs += totalEntry

	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip += tip
}

func (b *bill) addItem(item string, price float64) {
	b.items[item] = price
}

func (b *bill) saveBill() {
	data := []byte(b.format())

	err := os.WriteFile("bills/"+b.name+".txt", data, 0644)

	if err != nil {
		panic(err)
	}
	fmt.Println("The bill was successfully saved to file.")
}
