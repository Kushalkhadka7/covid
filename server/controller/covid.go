package controller

import (
	"context"
	covid "covid-19/proto"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"

	"google.golang.org/grpc"
)

// covidServer.
type covidServer struct {
	listener net.Listener
}

// CServer interface to register new grpc server.
type CServer interface {
	NewGrpcServer() error
}

// NewServer return new grpc server with new http server listener.
func NewServer(listener net.Listener) CServer {

	return &covidServer{
		listener: listener,
	}
}

// NewGrpcServer registers the new grpc server.
func (cs *covidServer) NewGrpcServer() error {
	grpcServer := grpc.NewServer()

	covid.RegisterCovidServer(grpcServer, &covidServer{})

	if err := grpcServer.Serve(cs.listener); err != nil {
		return err
	}

	return nil
}

// SayHello grpc server method.
func (cs *covidServer) CountryData(ctx context.Context, country *covid.Country) (*covid.CountryResponseList, error) {
	url := "https://api.covid19api.com/live/country/" + country.Name

	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Error in fetching data form api:%v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse response body:%v", err)
	}
	defer res.Body.Close()

	var resposneData []*covid.CountryResposne
	if err := json.Unmarshal(body, &resposneData); err != nil {
		return nil, fmt.Errorf("Cannot un marshal body:%v", err)
	}

	return &covid.CountryResponseList{
		Data: resposneData,
	}, nil
}

// TotalData for overall response.
func (cs *covidServer) TotalData(ctx context.Context, void *covid.Void) (*covid.TotalResponse, error) {

	res, err := http.Get("https: //api.covid19api.com/stats")
	if err != nil {
		return nil, fmt.Errorf("Error fetching total data form api:%v", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("Cannot parse response body:%v", err)
	}
	defer res.Body.Close()

	var resposneData covid.TotalResponse
	if err := json.Unmarshal(body, &resposneData); err != nil {
		return nil, fmt.Errorf("Cannot un marshal body:%v", err)
	}

	return &resposneData, nil
}
