package main

import (
	"fmt"
)

type Driver struct {
	Name   string
	Team   string
	Number int
	Points int
}

func findLeader(drivers []Driver) Driver {
	leader := drivers[0]
	for _, driver := range drivers {
		if driver.Points > leader.Points {
			leader = driver
		}
	}
	return leader
}
func findTeamDrivers(drivers []Driver, team string) []Driver {
	result := []Driver{}
	for _, driver := range drivers {
    if driver.Team == team {
		result = append(result, driver)

}
}

func main() {
	fmt.Println("F1 Analyzer is starting...")
	lando := Driver{
		Name:   "Lando Norris",
		Team:   "McLaren",
		Number: 1,
		Points: 188,
	}

	max := Driver{
		Name:   "Max Verstappen",
		Team:   "Red Bull",
		Number: 3,
		Points: 188,
	}

	charles := Driver{
		Name:   "Charles Leclerc",
		Team:   "Ferrari",
		Number: 16,
		Points: 191,
	}

	oscar := Driver{
		Name:   "Oscar Piastri",
		Team:   "McLaren",
		Number: 81,
		Points: 128,
	}
	drivers := []Driver{lando, max, charles, oscar}
	leader := findLeader(drivers)
	fmt.Printf("Driver: %s | Team: %s | Points: %d\n",
		leader.Name,
		leader.Team,
		leader.Points,
	)
}
