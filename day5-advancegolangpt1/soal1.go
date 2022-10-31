package main

import (
	"fmt"
	"sort"
)

type Item struct {
	Id     int
	Barang string
	Harga  int
}

var listItems = []Item{
	{1, "Benih Lele", 50000},
	{2, "Pakan lele cap menara", 25000},
	{3, "Probiotik A", 75000},
	{4, "Probiotik Nila B", 10000},
	{5, "Pakan Nila", 20000},
	{6, "Benih Nila", 20000},
	{7, "Cupang", 5000},
	{8, "Benih Nila", 30000},
	{9, "Benih Cupang", 10000},
	{10, "Probiotik B", 10000},
}

func getItemWithAHundredThousandPrice(items []Item) []Item {
	var total int
	var res []Item

	sort.Slice(items, func(p, q int) bool {
		return items[p].Harga < items[q].Harga
	})

	for _, item := range items {
		total += item.Harga
		res = append(res, item)
		if total >= 100000 {
			break
		}
	}

	return res

}

func getItemWithTenThousandPrice(items []Item) []Item {
	var res []Item
	for _, item := range items {
		if item.Harga == 10000 {
			res = append(res, item)
		}
	}

	return res
}

func getCheapestItem(items []Item) {
	var cheapestItem Item
	for _, item := range items {
		if cheapestItem.Harga == 0 || item.Harga < cheapestItem.Harga {
			cheapestItem = item
		}
	}
	fmt.Printf("%s - %d \n", cheapestItem.Barang, cheapestItem.Harga)
}

func getMostExpensiveItem(items []Item) {
	var mostExpensiveItem Item
	for _, item := range items {
		if item.Harga > mostExpensiveItem.Harga {
			mostExpensiveItem = item
		}
	}

	fmt.Printf("%s - %d \n", mostExpensiveItem.Barang, mostExpensiveItem.Harga)

}

func createListItems(items []Item) {
	for _, item := range items {
		fmt.Printf("%s - %d \n", item.Barang, item.Harga)
	}
}

func main() {

	fmt.Println("Total produk dengan harga dibawah  Rp 100.000 :")
	createListItems(getItemWithAHundredThousandPrice(listItems))
	fmt.Println("")

	fmt.Println("Total produk dengan harga Rp 10.000 : ")
	createListItems(getItemWithTenThousandPrice(listItems))
	fmt.Println("")

	fmt.Print("Harga produk termahal : ")
	getMostExpensiveItem(listItems)

	fmt.Print("Harga produk termurah : ")
	getCheapestItem(listItems)

	fmt.Println("")

}
