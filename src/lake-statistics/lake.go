// Problem 1: Statistical Analysis
// Lake Pend Oreille weather data: https://lpo.dt.navy.mil

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	//"net/http"
	"os"
	"strconv"
)

func main() {
	// Load the TXT file
	f, err := os.Open("Environmental_Data_Deep_Moor_2015.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// fmt.Println(f) // DEBUG

	rdr := csv.NewReader(f)
	rdr.Comma = '\t'
	rdr.TrimLeadingSpace = true
	rows, err := rdr.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	var length = len(rows)-1
	fmt.Println("Total Records: ", length)  // DEBUG

	// Now use http GET, instead of the file
	// must set http_proxy environment variable to retrieve
	// res, err := http.Get("https://lpo.dt.navy.mil/data/DM/Environmental_Data_Deep_Moor_2015.txt")
	// fmt.Println("res.Body: ", res.Body)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// r := csv.NewReader(res.Body)
	// r.Comma = '\t'
	// r.TrimLeadingSpace = true
	// defer res.Body.Close()
	// rws, err := r.ReadAll()
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Total Records (http): ", len(rws)-1)  // DEBUG

	// get all wind speeds (column 7 (because date-time is considered one column)))
	windspeeds := make([]float64, length)
	airtemps := make([]float64, length)
	baro := make([]float64, length)
	for i, row := range rows {
		// skip the header row
		if i != 0 {
			val, err := strconv.ParseFloat(row[7], 64)
			if err != nil {
				panic(err)
			}
			windspeeds[i-1] = val

			valair, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				panic(err)
			}
			airtemps[i-1] = valair

			valbar, err := strconv.ParseFloat(row[2], 64)
			if err != nil {
				panic(err)
			}
			baro[i-1] = valbar
		}
	}
	
	// use the math package to get the mean and median wind speed
	fmt.Println("Air Temp:\t", mean(airtemps))
	fmt.Println("Barometric:\t", mean(baro))
	fmt.Println("Wind Speed:\t", mean(windspeeds))

}

func mean(data []float64) float64 {
	var sum = 0.0
	for _, val := range data {
		sum += val
	}
	return sum / float64(len(data))
}
