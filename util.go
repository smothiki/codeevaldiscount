package main

import (
	"bufio"
	"os"
	"strings"
)

//Hasgcd calulates if two numbers has gcd
func Hasgcd(a, b int) bool {
	if gcd(a, b) != 1 {
		return true
	}
	return false
}

func gcd(a, b int) int {
	if a == b {
		return a
	}
	if a > b {
		return gcd(a-b, b)
	}
	return gcd(a, b-a)
}

//Wordcount calculates the number of words in item that matches chars
func Wordcount(item, chars string) int {
	pl := strings.Split(item, "")
	count := 0
	for i := range pl {
		if strings.Contains(chars, pl[i]) {
			count++
		}
	}
	return count
}

//Parsedatafile parses test data and give a pc data slice
func Parsedatafile() []PCdata {

	pc := []PCdata{}
	file, err := os.Open("testdata")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ";")
		product, consumer := line[1], line[0]
		lines := PCdata{products: strings.Split(product, ","), customers: strings.Split(consumer, ",")}
		pc = append(pc, lines)
	}

	return pc
}
