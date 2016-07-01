package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Now()
	fmt.Println(t)
	fmt.Printf("%4d-%02d-%02d\n", t.Year(), t.Month(), t.Day())

	t = time.Now().UTC()
	fmt.Println(t)
	fmt.Println(time.Now())

	week = 7 * 24 * 60 * 60 * 1e9 //nanosec
	week_from_now := t.Add(week)
	fmt.Println(week_from_now)

	fmt.Println(t.Format(time.RFC822))
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t.Format("02 Jan 2006 15:04"))

	s := t.Format("20060102")
	fmt.Println(t, "=>", s)
}
