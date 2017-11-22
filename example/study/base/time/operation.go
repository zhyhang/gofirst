package main

import (
	"time"
	"fmt"
)

func main() {
	// add/sub duration
	tn := time.Now()
	tna24 := tn.Add(time.Hour * 24)
	tns24 := tn.Add(-24 * time.Hour)
	tnaddt := tn.AddDate(1, 0, 0)
	beforen := tns24.Before(tn)
	aftern := tna24.After(tn)
	fmt.Printf("now is: %v\n", tn)
	fmt.Printf("add 24 hours: %v\n", tna24)
	fmt.Printf("sub 24 hours: %v\n", tns24)
	fmt.Printf("add 1 year:: %v\n", tnaddt)
	fmt.Printf("%v is before %v? %v\n", tns24, tn, beforen)
	fmt.Printf("%v is after %v? %v\n", tna24, tn, aftern)

	// differ between times
	difftime := tn.Sub(tns24)
	fmt.Printf("now - yesterday: %v\n", difftime)
	fmt.Printf("now - yesterday: hour(%v)=minute(%v)=(%v)=ns(%v)\n",
		difftime.Hours(), difftime.Minutes(), difftime.Seconds(), difftime.Nanoseconds())

	// get hour minute...
	fmt.Printf("year:%v, month:%d, day:%v, minute:%v, seconds:%v, nsec:%d\n", tn.Year(), tn.Month(), tn.Day(),
		tn.Minute(), tn.Second(), tn.Nanosecond())

	// get Date, Clock
	y, m, d := tn.Date()
	fmt.Printf("from Date(), year:%d, month:%d, day:%d\n", y, m, d)
	h, mm, s := tn.Clock()
	fmt.Printf("from Clock(), hour:%d, minute:%d, second:%d\n", h, mm, s)

	// time with Equal() method to decide to one time same as another
	tn.Equal(tn)

	// shorthand method
	fmt.Printf("since yesterday %v\n", time.Since(tns24).Hours())
	fmt.Printf("until tomorrow %v\n", time.Until(tna24).Hours())

	//year days have gone of the specify instant
	fmt.Printf("this year days: %d\n", tn.YearDay())
	fmt.Printf("year %d days: %d\n", 2016, time.Date(2016, 12, 31, 0, 0, 0, 0, time.Local).YearDay())

	// leap year
	y = tn.Year()
	fmt.Printf("%d is leap year?%v/%v(method1/method2)\n", y, isLeapYear(y), isLeapYear1(y))
	y = 2016
	fmt.Printf("%d is leap year?%v/%v(method1/method2)\n", y, isLeapYear(y), isLeapYear1(y))
	y = 400
	fmt.Printf("%d is leap year?%v/%v(method1/method2)\n", y, isLeapYear(y), isLeapYear1(y))
	y = 2100
	fmt.Printf("%d is leap year?%v/%v(method1/method2)\n", y, isLeapYear(y), isLeapYear1(y))
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func isLeapYear1(y int) bool {

	year := time.Date(y, time.December, 31, 0, 0, 0, 0, time.Local)
	days := year.YearDay()

	if days > 365 {
		return true
	} else {
		return false
	}
}
