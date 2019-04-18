package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Lotto struct {
	Lotto []LottoData `json:"Lotto649"`
}

type LottoData struct {
	Date    string `json:"Date"`
	Numbers []int  `json:"Numbers"`
}

func main() {

	fileLoc := "Lotto649.json"
	jsonFile, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	var lotto Lotto
	json.Unmarshal(jsonFile, &lotto)

	var s []int

	for i := 0; i < len(lotto.Lotto); i++ {
		for y := 0; y < len(lotto.Lotto[i].Numbers); y++ {
			s = append(s, lotto.Lotto[i].Numbers[y])
		}
	}

	fmt.Printf("%v", s)
}
