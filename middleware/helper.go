package middleware

import (
	"fmt"

	"github.com/hchaudhari73/goAuth/config"
)

func getHomeEndpoint() (*string, error) {

	baseHttp, err := config.GetBaseHttpUrl()
	if err != nil {
		return nil, err
	}

	port, err := config.GetPort()
	if err != nil {
		return nil, err
	}

	// login endpoint
	endpoint := fmt.Sprintf("%s:%s", *baseHttp, *port)
	return &endpoint, nil
}

func getLoginEndpoint() (*string, error) {

	baseHttp, err := config.GetBaseHttpUrl()
	if err != nil {
		return nil, err
	}

	port, err := config.GetPort()
	if err != nil {
		return nil, err
	}

	// login endpoint
	endpoint := fmt.Sprintf("%s:%s/login", *baseHttp, *port)
	return &endpoint, nil
}
