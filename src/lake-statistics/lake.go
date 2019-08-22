// Problem 1: Statistical Analysis
// Lake Pend Oreille weather data: https://lpo.dt.navy.mil

package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	//"net/http"
)

func main() {
	// Load the TXT file
	f, err := os.Open("Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(f) // DEBUG

	rdr := csv.NewReader(bufio.NewReader(f))
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Total Records: ", len(rows)-1)  // DEBUG

	// Now use http GET, instead of the file
	// res, err := http.Get("https://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2019.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// rdr := csv.NewReader(res.Body)
	// rdr.Comma = '\t'
	// rdr.TrimLeadingSpace = true
	// defer res.Body.Close()
	// rows, err := rdr.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Total Records: ", len(rows)-1)  // DEBUG
}
