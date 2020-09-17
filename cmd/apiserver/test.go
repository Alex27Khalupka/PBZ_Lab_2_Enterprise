package main

import (
	"fmt"
	"time"
)

func main(){
	const shortForm = "2006-01-02"
	t, _ := time.Parse(shortForm, "2013-02-03")
	fmt.Println(t)
}