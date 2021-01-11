package main

import (
	"fmt"
	"log"
	"github.com/fsayed/calendar"
)

func main() {
	event := calendar.Event{}
	
	err := event.SetTitle("Birthday")
	if err != nil {
		log.Fatal(err)
	}


	err = event.Date.SetYear(2004)
	if err != nil {
		log.Fatal(err)
	}

	
	err = event.Date.SetMonth(1)
	if err != nil {
		log.Fatal(err)
	}

	err = event.Date.SetDay(10)
	if err != nil {
		log.Fatal(err)
	}


	fmt.Println(event.Title())
	fmt.Println(event.Year())
	fmt.Println(event.Date.Month())
	fmt.Println(event.Date.Day())

}
	
