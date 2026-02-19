package main

import "fmt"

func incrementer() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}

func main() {
	counter1 := incrementer()

	fmt.Printf("1回目：%d\n",counter1())
	fmt.Printf("2回目：%d\n",counter1())
	fmt.Printf("3回目：%d\n",counter1())
	fmt.Printf("4回目：%d\n",counter1())

	counter2 := incrementer()

	fmt.Printf("2-1回目：%d\n",counter2())
	fmt.Printf("2-2回目：%d\n",counter2())

}
