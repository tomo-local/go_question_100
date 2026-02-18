package main

import (
	"fmt"
)

func divmod(a int, b int) (int, int)  {
	quotient := a / b
	remainder := a % b
	return quotient, remainder
}

func main() {

	q, r := divmod(10,3)

	fmt.Printf("10 ÷ 3 = 商 %d, 余り %d\n", q, r)

	qOnly ,rNot := divmod(10, 2)
	fmt.Printf("商だけを取得: %d 余り：%d\n", qOnly, rNot)
}
