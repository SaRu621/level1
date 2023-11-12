package main

import (
	"fmt"
	"time"
)

func Sleep(secs int) {
	<-time.After(time.Duration(secs) * time.Second)
}

func main() {
	fmt.Println("start")
	Sleep(1)
	fmt.Println("end")
}
