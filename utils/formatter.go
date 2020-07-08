package util

import (
	covid "covid-19/proto"
	"fmt"
)

// OutPutFormatter formats the data in a particular format to std out.
func OutPutFormatter(data *covid.CountryResposne) {
	fmt.Printf("************************ %s *****************\n\n", data.GetCountry())
	fmt.Printf("Total Cases: %d\n", data.GetConfirmed())
	fmt.Printf("Total Deaths: %d\n", data.GetDeaths())
	fmt.Printf("Total Recovered: %d\n", data.GetRecovered())
}
