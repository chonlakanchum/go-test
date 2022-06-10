package main

import "fmt"

type Formular interface {
	GetSpeed() float64
}

type FirstSpeedFormular struct {
	EarlySpeed, Acceleration, Time float64
}

type SecondSpeedFormular struct {
	EarlySpeed, Acceleration, Distance float64
}

func (s FirstSpeedFormular) GetSpeed() float64 {
	return s.EarlySpeed + (s.Acceleration * s.Time)
}

func (s SecondSpeedFormular) GetSpeed() float64 {
	return (s.EarlySpeed * s.EarlySpeed) + (2 * s.Acceleration * s.Distance)
}

func main() {
	first := FirstSpeedFormular{100, 5, 20}
	second := SecondSpeedFormular{10, 5, 100}

	allSpeed := []Formular{first, second}
	for _, speed := range allSpeed {
		fmt.Println("GetSpeed ", speed.GetSpeed())
	}
}
