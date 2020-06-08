package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func main() {
	args := os.Args
	var canShuffel bool
	var timeOut int
	if len(args) > 1 {
		if strings.Contains(args[1], "limit") {
			if args[1] == "limit" {
				timeOut = 30
			} else {
				rep := strings.Replace(args[1], "limit=", "", 1)
				timeOut, _ = strconv.Atoi(rep)
			}
		}

		if len(args) >= 2 {
			for _, key := range args {
				if key == "shuffel" {
					canShuffel = true
				}
			}
		}
	}

	csvFile, err := os.Open("quiz.csv")
	if err != nil {
		fmt.Printf("Failed to open CSV file Error: %v", err)
		return
	}

	csvReader, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Printf("Failed to read csv data Error: %v", err)
		return
	}

	if canShuffel {
		rand.Seed(time.Now().UnixNano())
		rand.Shuffle(len(csvReader), func(i, j int) {
			csvReader[i], csvReader[j] = csvReader[j], csvReader[i]
		})
	}
	var wait sync.WaitGroup
	var success, fail int

	go func() {

		for _, quiz := range csvReader {
			var result string
			fmt.Printf("what %v, sir?,", quiz[0])
			_, err = fmt.Scanf("%v\n", &result)
			if err != nil {
				fmt.Printf("Failed to read input from user Error: %v", err)
				return
			}
			if result == quiz[1] {
				success++
			} else {
				fail++
			}
		}
		wait.Done()
	}()

	if timeOut > 0 {
		duration := time.Duration(timeOut) * time.Second
		time.Sleep(duration)
	} else {
		wait.Wait()
	}
	fmt.Printf("\nScored %v out of %v\n", success, len(csvReader))

}
