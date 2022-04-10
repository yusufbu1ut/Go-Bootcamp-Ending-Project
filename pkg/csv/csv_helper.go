package csv

import (
	"encoding/csv"
	"github.com/yusufbu1ut/Go-Bootcamp-Ending-Project/internal/domain/category"
	"log"
	"os"
	"strconv"
	"sync"
)

//ReadCsvWithWorkerPool func takes file path and open it after reads and takes returns items with channel
func ReadCsvWithWorkerPool(path string) chan category.Category {

	linesChan := make(chan []string)
	resultsChan := make(chan category.Category)

	var places []int
	//open file
	f, err := os.Open(path)
	if err != nil {
		log.Printf("File cant open, error: %s", err.Error())
		return nil
	}
	defer f.Close()

	//take lines
	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		log.Printf("File cant read, error: %s", err.Error())
		return nil
	}
	//read first line checks is there what program need
	places = findCategoryItemsPlaces(lines[0])

	wg := new(sync.WaitGroup)
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go convertToCategoryItem(linesChan, resultsChan, wg, places)
	}

	go func() {
		//reading lines
		isFirstRow := true
		for _, line := range lines {
			if isFirstRow {
				isFirstRow = false
				continue
			}

			linesChan <- line
		}
		close(linesChan)
	}()

	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	return resultsChan
}

//convertToCategoryItem func listens linesChan and if there is a new item checks it and adds resultChan
func convertToCategoryItem(linesChan <-chan []string, resultsChan chan<- category.Category, wg *sync.WaitGroup, places []int) {
	defer wg.Done()

	for c := range linesChan {
		//checking line items positions and content
		if len(c) < 2 {
			log.Printf("Item not created; %s \n", c)
			continue
		}
		code, err := strconv.Atoi(c[places[1]])
		if err != nil {
			log.Printf("Item not created; %s \n", c)
			continue
		}
		if code <= 0 {
			log.Printf("Item not created; %s \n", c)
			continue
		}
		description := ""
		if len(c) >= 2 {
			description = c[places[2]]
		}
		categoryItem := category.NewCategory(c[places[0]], uint(code), description)
		resultsChan <- *categoryItem
	}
}

//findCategoryItemsPlaces this function finds the places for category' columns if there is more columns in csv file
//csv file should contain columns named as name,code,description
func findCategoryItemsPlaces(line []string) []int {
	places := make([]int, 3)
	for i, c := range line {
		if c == "name" {
			places[0] = i
		}
		if c == "code" {
			places[1] = i
		}
		if c == "description" {
			places[2] = i
		}
	}
	return places
}
