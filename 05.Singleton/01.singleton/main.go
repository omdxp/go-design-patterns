package main

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

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

func main() {
	db := GetSingletonDatabase()
	pop := db.GetPopulation("Seoul")
	println(pop)
}

// Singleton is a creational design pattern that lets you ensure that a class has only one instance, while providing a global access point to this instance.
