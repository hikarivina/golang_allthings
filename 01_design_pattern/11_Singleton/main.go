package main

import (
	"bufio"
	"fmt"
	"os"
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

// sync.Once init() -- thread safety
// laziness

func readData(path string) (map[string]int, error) {
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// exPath := filepath.Dir(ex)

	// file, err := os.Open(exPath + path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	result := map[string]int{}

	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		// fmt.Println(k)

		v, _ := strconv.Atoi(scanner.Text())

		// fmt.Println(v)
		result[k] = v
	}

	return result, nil

}

var once sync.Once
var instance *singletonDatabase

func GetSingletonDatabase() *singletonDatabase {

	once.Do(func() {
		caps, e := readData("./capitals.txt")
		db := singletonDatabase{caps}
		if e == nil {
			db.capitals = caps
		}

		instance = &db
	})

	return instance
}

func GetTotalPopulationEx(db Database, cities []string) int {
	result := 0
	for _, city := range cities {
		result += db.GetPopulation(city)
	}
	return result
}

func GetTotalPopulation(cities []string) int {
	result := 0
	for _, city := range cities {
		result += GetSingletonDatabase().GetPopulation(city)
	}
	return result
}

type DummyDatabase struct {
	dummyData map[string]int
}

func (d *DummyDatabase) GetPopulation(name string) int {
	if len(d.dummyData) == 0 {
		d.dummyData = map[string]int{
			"alpha": 1,
			"beta":  2,
			"gamma": 3}
	}
	return d.dummyData[name]
}

func main() {

	// db := GetSingletonDatabase()
	// pop := db.GetPopulation("Seoul")
	// fmt.Println("Pop of Seoul = ", pop)

	// cities := []string{"Seoul", "Mexico City"}
	// tp := GetTotalPopulation(cities)
	// ok := tp == (17500000 + 17400000)
	// fmt.Println(ok)

	names := []string{"alpha", "gamma"}
	tp := GetTotalPopulationEx(&DummyDatabase{}, names)
	fmt.Println(tp == 4)
}
