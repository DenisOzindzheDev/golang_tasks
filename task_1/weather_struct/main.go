package main

import "fmt"

type Report struct {
	Sol         int
	Temperature temperature
	location    location
}
type celsius float64

type temperature struct {
	high, low celsius
}
type location struct {
	lat, long float64
}

func (t temperature) Average() celsius {
	return (t.high + t.low) / 2
}
func main() {
	locationIstanbul := location{lat: 0, long: 80}
	r := Report{
		Sol:         1000,
		Temperature: temperature{high: 20, low: 10},
		location:    locationIstanbul,
	}
	fmt.Println(r)
	fmt.Println(r.Temperature.Average())

}
