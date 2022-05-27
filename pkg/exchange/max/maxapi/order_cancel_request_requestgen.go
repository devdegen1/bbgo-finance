// Code generated by "requestgen -method POST -url v2/order/delete -type OrderCancelRequest -responseType .Order"; DO NOT EDIT.

package max

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"regexp"
)

func (o *OrderCancelRequest) Id(id uint64) *OrderCancelRequest {
	o.id = &id
	return o
}

func (o *OrderCancelRequest) ClientOrderID(clientOrderID string) *OrderCancelRequest {
	o.clientOrderID = &clientOrderID
	return o
}

// GetQueryParameters builds and checks the query parameters and returns url.Values
func (o *OrderCancelRequest) GetQueryParameters() (url.Values, error) {
	var params = map[string]interface{}{}

	query := url.Values{}
	for _k, _v := range params {
		query.Add(_k, fmt.Sprintf("%v", _v))
	}

	return query, nil
}

// GetParameters builds and checks the parameters and return the result in a map object
func (o *OrderCancelRequest) GetParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}
	// check id field -> json key id
	if o.id != nil {
		id := *o.id

		// assign parameter of id
		params["id"] = id
	} else {
	}
	// check clientOrderID field -> json key client_oid
	if o.clientOrderID != nil {
		clientOrderID := *o.clientOrderID

		// assign parameter of clientOrderID
		params["client_oid"] = clientOrderID
	} else {
	}

	return params, nil
}

// GetParametersQuery converts the parameters from GetParameters into the url.Values format
func (o *OrderCancelRequest) GetParametersQuery() (url.Values, error) {
	query := url.Values{}

	params, err := o.GetParameters()
	if err != nil {
		return query, err
	}

	for _k, _v := range params {
		if o.isVarSlice(_v) {
			o.iterateSlice(_v, func(it interface{}) {
				query.Add(_k+"[]", fmt.Sprintf("%v", it))
			})
		} else {
			query.Add(_k, fmt.Sprintf("%v", _v))
		}
	}

	return query, nil
}

// GetParametersJSON converts the parameters from GetParameters into the JSON format
func (o *OrderCancelRequest) GetParametersJSON() ([]byte, error) {
	params, err := o.GetParameters()
	if err != nil {
		return nil, err
	}

	return json.Marshal(params)
}

// GetSlugParameters builds and checks the slug parameters and return the result in a map object
func (o *OrderCancelRequest) GetSlugParameters() (map[string]interface{}, error) {
	var params = map[string]interface{}{}

	return params, nil
}

func (o *OrderCancelRequest) applySlugsToUrl(url string, slugs map[string]string) string {
	for _k, _v := range slugs {
		needleRE := regexp.MustCompile(":" + _k + "\\b")
		url = needleRE.ReplaceAllString(url, _v)
	}

	return url
}

func (o *OrderCancelRequest) iterateSlice(slice interface{}, _f func(it interface{})) {
	sliceValue := reflect.ValueOf(slice)
	for _i := 0; _i < sliceValue.Len(); _i++ {
		it := sliceValue.Index(_i).Interface()
		_f(it)
	}
}

func (o *OrderCancelRequest) isVarSlice(_v interface{}) bool {
	rt := reflect.TypeOf(_v)
	switch rt.Kind() {
	case reflect.Slice:
		return true
	}
	return false
}

func (o *OrderCancelRequest) GetSlugsMap() (map[string]string, error) {
	slugs := map[string]string{}
	params, err := o.GetSlugParameters()
	if err != nil {
		return slugs, nil
	}

	for _k, _v := range params {
		slugs[_k] = fmt.Sprintf("%v", _v)
	}

	return slugs, nil
}

func (o *OrderCancelRequest) Do(ctx context.Context) (*Order, error) {

	params, err := o.GetParameters()
	if err != nil {
		return nil, err
	}
	query := url.Values{}

	apiURL := "v2/order/delete"

	req, err := o.client.NewAuthenticatedRequest(ctx, "POST", apiURL, query, params)
	if err != nil {
		return nil, err
	}

	response, err := o.client.SendRequest(req)
	if err != nil {
		return nil, err
	}

	var apiResponse Order
	if err := response.DecodeJSON(&apiResponse); err != nil {
		return nil, err
	}
	return &apiResponse, nil
}
