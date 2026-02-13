package main

import (
	"fmt"
	"time"
)

func main(){
	now := time.Now()
	layoutYYYYMMDD := "2006年01月02日"
	layout := "2006年01月02日 15時01分05秒"

	fmt.Printf("now time: %s\n", now.Format(layoutYYYYMMDD))

	t := time.Date(1998, 6, 24, 0, 0, 0, 0, time.Local)

	fmt.Printf("t time: %s\n", t.Format(layout))
}
