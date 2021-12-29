// Code generated by "requestgen -method GET -responseType .APIResponse -responseDataField Data -type GetAllTickersRequest -url /api/v1/market/allTickers -responseDataType AllTickers"; DO NOT EDIT.

package kucoinapi

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
)

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (g *GetAllTickersRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for k, v := range params {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (g *GetAllTickersRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (g *GetAllTickersRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := g.GetParameters()
	if err != nil {
		return query, err
	}

	for k, v := range params {
		query.Add(k, fmt.Sprintf("%v", v))
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (g *GetAllTickersRequest) GetParametersJSON() ([]byte, error) {
	params, err := g.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

func (g *GetAllTickersRequest) Do(ctx context.Context) (*AllTickers, error) {

	// no body params
	var params interface{}
	query := url.Values{}

	req, err := g.client.NewRequest(ctx, "GET", "/api/v1/market/allTickers", query, params)
	if err != nil {
		return nil, err
	}

	response, err := g.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse APIResponse
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	var data AllTickers
	if err := json.Unmarshal(apiResponse.Data, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
