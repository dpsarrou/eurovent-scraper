package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	program := flag.String("program", "AC", "The eurovent program, eg. AC")
	productType := flag.String("type", "AC1/A/S/R", "The product specific type, eg. AC1/A/S/R")
	brandString := flag.String("brands", "CARRIER", "Comma separated list of brands eg. CARRIER,LG")
	flag.Parse()
	brands := strings.Split(*brandString, ",")

	c := Client{}
	e := Eurovent{c}

	total, err := e.GetTotalCount(*program, *productType, brands)
	if err != nil {
		log.Fatal(err)
	}

	if total > 0 {

		data, err := e.GetData(*program, *productType, brands, total)
		if err != nil {
			log.Fatal(err)
		}

		ids := []int{}
		for _, model := range data.Rows {
			ids = append(ids, model.ID)
		}

		csv, err := e.DataToCsv(*program, *productType, brands, ids)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(csv)
	}
}