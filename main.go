package main

import "fmt"

func main() {
	mybill := newBill("mario's bill")

mybill.addItem("Pizza", 3490)
	mybill.updateTip(100)

	fmt.Println(mybill.format())

}
