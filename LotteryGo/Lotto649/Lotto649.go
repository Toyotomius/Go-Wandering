package lotto649

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

// Lotto : Maps to initial json array containing the rest of the data.
type Lotto struct {
	Lotto []LottoData `json:"Lotto649"`
}

// LottoData : Pulls the date and winning numbers from the lottery json file.
type LottoData struct {
	Date    string `json:"Date"`
	Numbers []int  `json:"Numbers"`
}

func frequencyCount(nums []int) map[int]int {
	numFreq := make(map[int]int)
	for _, num := range nums {
		numFreq[num]++
	}
	return numFreq
}

//Lotto649 : Created for testing goroutines & multiple go files.
func Lotto649(messages chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
	defer func() {
		fmt.Println("Lotto649 time elapsed:", time.Since(start))
	}()
	fileLoc := "Lotto649.json"
	writeFileLoc := "Lotto649Singles.txt"

	f, err := os.Create(writeFileLoc)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	jsonFile, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	var lotto Lotto
	json.Unmarshal(jsonFile, &lotto)

	var numList []int

	// Flattens the number arrays from the json into one array for comparison
	for i := 0; i < len(lotto.Lotto); i++ {
		for y := 0; y < len(lotto.Lotto[i].Numbers); y++ {
			numList = append(numList, lotto.Lotto[i].Numbers[y])
		}
	}

	// numMap returns unsorted (unless fmt is used, then key) map of number:frequency pairs. Ex: map[1:423 2:440 3:450 4:460 5:425 6:426 7:459]
	numMap := frequencyCount(numList)

	// struct to create slice of key/value pairs.
	type kvLotto struct {
		Key   int
		Value int
	}

	// Creates slice to be sorted and assigns key/value pairs from map
	var sortedFrequency []kvLotto
	for k, v := range numMap {
		sortedFrequency = append(sortedFrequency, kvLotto{k, v})
	}
	// Returns sorted slice where the value of any given index is larger. Output: [{31 501} {45 492} {40 491} {34 481}] for k/v pairs.
	sort.Slice(sortedFrequency, func(i, j int) bool {
		return sortedFrequency[i].Value > sortedFrequency[j].Value
	})
	//Creates a stringbuilder and uses Sprintf to compose the strings correctly for writing.
	var sb strings.Builder
	for _, kv := range sortedFrequency {
		s := fmt.Sprintf("Number: %d, Frequency: %d\n", kv.Key, kv.Value)
		sb.WriteString(s)

	}
	s := sb.String()
	f.WriteString(s)
	messages <- "Done Lotto649"
}
