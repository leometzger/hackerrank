package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(int(1800), time.January, 0, 12, 0, 0, 0, time.Local)

	fmt.Println(t.Add(time.Duration(256) * time.Hour * 24).Format("02.01.2006"))
}
