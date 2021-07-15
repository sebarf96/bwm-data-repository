package models

import "github.com/go-kit/kit/endpoint"

//Eps Structure that stores the endpoint
type Eps struct {
	CreateEP endpoint.Endpoint
	UpdateEP endpoint.Endpoint
	ReadEP   endpoint.Endpoint
	DeleteEP endpoint.Endpoint
}
