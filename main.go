package main

import (
	"errors"
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
	return result
}
func findDriverByName(drivers []Driver, name string) (Driver, error) {
	for _, driver := range drivers {
		if driver.Name == name {
			return driver, nil
		}

	}
	return Driver{}, errors.New("driver not found")
}

func findDriverByNumber(drivers []Driver, number int) (Driver, error) {
	for _, driver := range drivers {
		if driver.Number == number {
			return driver, nil
		}
	}
	return Driver{}, errors.New("driver not found")
}
func printDriver(driver Driver) {
	fmt.Printf("Driver: %s | Number: %d | Team: %s | Points: %d\n",
		driver.Name,
		driver.Number,
		driver.Team,
		driver.Points,
	)
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
	teamMcLaren := findTeamDrivers(drivers, "McLaren")
	driver, err := findDriverByName(drivers, "Lando Norris")
	if err != nil {
		fmt.Println(err)
		return
	}
	printDriver(driver)
	fmt.Print("Enter driver number: ")
	var number int
	fmt.Scan(&number)
	driver, err = findDriverByNumber(drivers, number)
	if err != nil {
		fmt.Println(err)
		return
	}
	printDriver(driver)
	fmt.Println("McLaren drivers:")
	for _, driver := range teamMcLaren {
		fmt.Printf("Driver: %s | Points: %d\n",
			driver.Name,
			driver.Points,
		)
	}
	fmt.Println("Championship leader:")
	fmt.Printf("Driver: %s | Team: %s | Points: %d\n",
		leader.Name,
		leader.Team,
		leader.Points,
	)
}
