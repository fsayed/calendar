package calendar

import(
	"errors"
	"unicode/utf8"
)

type Event struct {
	title string
	Date
}


type Date struct {
	year  int
	month int
	day   int
}



//Getter and Setter for the envent

func (e *Event) SetTitle(title string) error {
	if utf8.RuneCountInString(title) > 30 {
		return errors.New("Title is too long")
	}
	e.title = title
	return nil
}


func (e *Event) Title() string {
	return e.title
}


//create setter/getter for Date type

func (d *Date) SetYear(year int) error {
	if year < 0 {
		return errors.New("Invalid Year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("Invalid Month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("Invalid Day")
	}
	d.day = day
	return nil
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}

