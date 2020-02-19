package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界", time.Now().Format("2006.01.02.15.04.05.000"))
	now := time.Now()
	fmt.Println(now.Format("2006_01_02_15_04_05.csv"))
}
