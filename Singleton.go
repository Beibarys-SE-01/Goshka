package main

import (
	"bufio"
	"os"
	"strconv"
	"sync"
)

type Database interface {
	GetNumber(name string) int
}

type singletonDatabase struct {
	numbers map[string]int
}

func (db *singletonDatabase) GetNumber(name string) int {
	return db.numbers[name]
}

func readData(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	result := map[string]int{}
	for scanner.Scan() {
		k := scanner.Text()
		scanner.Scan()
		v, _ := strconv.Atoi(scanner.Text())
		result[k] = v
	}
	return result, nil
}

var once sync.Once

var instance Database

func GetSingletonDatabase() Database {
	once.Do(func() {
		numbers, _ := readData("./db.txt")
		db := singletonDatabase{numbers}
		instance = &db
	})
	return instance
}
