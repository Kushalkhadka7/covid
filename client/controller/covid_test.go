package controller

import (
	"fmt"
	"testing"

	"log"
)

func TestHandleTotalCases(t *testing.T) {

	tests := []struct {
		name string
		text string
	}{
		{name: "first", text: "total"},
		{name: "second", text: "total deaths"},
		{name: "third", text: "total cases"},
	}

	for _, value := range tests {
		t.Run(value.name, func(t *testing.T) {
			err := HandleTotalCases(value.text)
			if err != nil {
				log.Fatalf("error in testing data:%v", err)
			}
		})
	}

}

func TestHandleCountryCases(t *testing.T) {
	fmt.Println("hello world")
}
