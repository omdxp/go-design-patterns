package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

type Database interface {
	GetPopulation(name string) int
}

type singletonDatabase struct {
	capitals map[string]int
}

func (db *singletonDatabase) GetPopulation(name string) int {
	return db.capitals[name]
}

// init() is called before main() and is used to initialize the package.
// sync.Once is used to ensure that the initialization code is executed only once.
// laziness is a property of the singleton pattern.

var once sync.Once
var instance *singletonDatabase

func readData(path string) (map[string]int, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	file, err := os.Open(absPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	result := map[string]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		capital := scanner.Text()
		scanner.Scan()
		population, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		result[capital] = population
	}
	return result, nil
}

func GetSingletonDatabase() *singletonDatabase {
	once.Do(func() {
		caps, err := readData("./capitals.txt")
		db := singletonDatabase{caps}
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

func GetTotalPopulationEx(cities []string, db Database) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (db *DummyDatabase) GetPopulation(name string) int {
	if len(db.dummyData) == 0 {
		db.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3,
		}
	}
	return db.dummyData[name]
}

func main() {
	names := []string{"alpha", "gamma"}
	pop := GetTotalPopulationEx(names, &DummyDatabase{})
	println(pop == 4)
}
