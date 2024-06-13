package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{"Pie":3.99 ,"Apple": 5.99, "Orange":1.99},
		tip:   0,
	}

	return b
}

// Structs automatically gets dereferenced
// The format function can only be accessed using bill object now

func (b *bill) format() string {
	fs := "Bill Breakdown is:\n"

	var total float64 = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":",v)
		total += v
	}

	//add a tip

	fs += fmt.Sprintf("%-25v ...$%v\n","tip:",b.tip)
	total += b.tip
	// %-25v interesting to make alignment

	// total

	fs += fmt.Sprintf("%v ...$%0.2f", "total:", total)

	return fs
}

//update the tip

func (b *bill) updateTip(tip float64) {
	b.tip = tip


}

func (b *bill) addItem(name string, price float64){
	b.items[name] = price  
}