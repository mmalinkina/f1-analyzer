package main

import (
	"errors"
	"fmt"
	"sort"
)

type Driver struct {
	Name   string
	Team   string
	Number int
	Points int
	Wins   int
}

func findLeader(drivers []Driver) (Driver, error) {
	if len(drivers) == 0 {
		return Driver{}, errors.New("no drivers available")
	}
	leader := drivers[0]
	for _, driver := range drivers {
		if isBetterDriver(driver, leader) {
			leader = driver
		}
	}
	return leader, nil
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

func findDriverByNumber(driversByNumber map[int]Driver, number int) (Driver, error) {
	driver, ok := driversByNumber[number]
	if ok {
		return driver, nil
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

func printStandings(drivers []Driver) {
	for i, driver := range drivers {
		fmt.Printf("%d. %s --- %d\n",
			i+1,
			driver.Name,
			driver.Points)
	}
}
func sortDriversByChampionship(drivers []Driver) {
	sort.Slice(drivers, func(i, j int) bool {
		return isBetterDriver(drivers[i], drivers[j])
	})
}
func isBetterDriver(a Driver, b Driver) bool {
	if a.Points != b.Points {
		return a.Points > b.Points
	}
	return a.Wins > b.Wins
}

func main() {
	fmt.Println("F1 Analyzer is starting...")
	drivers := []Driver{
		{Name: "Lando Norris",
			Team:   "McLaren",
			Number: 1,
			Points: 188,
			Wins:   3,
		},
		{Name: "Max Verstappen",
			Team:   "Red Bull",
			Number: 3,
			Points: 188,
			Wins:   5,
		},
		{
			Name:   "Charles Leclerc",
			Team:   "Ferrari",
			Number: 16,
			Points: 191,
			Wins:   2,
		},
		{
			Name:   "Oscar Piastri",
			Team:   "McLaren",
			Number: 81,
			Points: 128,
			Wins:   4,
		},
	}
	driversByNumber := make(map[int]Driver)

	for _, driver := range drivers {
		driversByNumber[driver.Number] = driver
	}
	sortDriversByChampionship(drivers)
	printStandings(drivers)

	leader, err := findLeader(drivers)
	if err != nil {
		fmt.Println(err)
		return
	}
	teamMcLaren := findTeamDrivers(drivers, "McLaren")
	driver, err := findDriverByName(drivers, "Lando Norris")
	if err != nil {
		fmt.Println(err)
		return
	}
	printDriver(driver)
	for {

		fmt.Print("Enter driver number: ")
		var number int
		_, err := fmt.Scan(&number)
		if err != nil {
			fmt.Println(err)
			continue
		}

		driver, err = findDriverByNumber(driversByNumber, number)

		if err != nil {
			fmt.Println(err)
			continue
		}
		printDriver(driver)
		break
	}
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
