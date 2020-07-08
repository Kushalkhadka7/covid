package controller

import (
	"context"
	"covid-19/constants"
	covid "covid-19/proto"
	util "covid-19/utils"
	"fmt"
	"log"
	"strings"

	"google.golang.org/grpc"
)

// Controller for covid-19 data.
type Controller struct {
	covid.CovidClient
}

//CovidController interface for covid.
type CovidController interface {
	HandleTotalCases(text string) error
	HandleCountryCases(text string) error
}

// NewCovid create new covid-19 client.
func NewCovid(conn *grpc.ClientConn) CovidController {
	return &Controller{
		covid.NewCovidClient(conn),
	}
}

// HandleTotalCases to handle overall data.
func (c *Controller) HandleTotalCases(text string) error {
	fmt.Printf("i am called %s", text)
	covidData, err := c.TotalData(context.Background(), &covid.Void{})
	if err != nil {
		return fmt.Errorf("Could not load data form server:%v", err)

	}

	log.Printf("%s", covidData.Total)

	return nil
}

// HandleCountryCases to handle data related to county.
func (c *Controller) HandleCountryCases(text string) error {
	switch text {
	case util.Find(constants.Countries, text):
		country := util.GetCountry(strings.Split(text, " "))

		covidData, err := c.CountryData(context.Background(), &covid.Country{Name: country})
		if err != nil {
			return fmt.Errorf("Could not load data form server:%v", err)
		}

		util.OutPutFormatter(covidData.Data[len(covidData.Data)-1])

	default:
		fmt.Println("hello i am default value")
	}

	return nil
}
