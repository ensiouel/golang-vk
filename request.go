package vk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/url"
)

type Request struct {
	session *Session
	params  url.Values
	method  string
}

func (request *Request) Set(key string, value interface{}) *Request {
	request.params.Set(key, fmt.Sprintf("%v", value))
	return request
}

func (request *Request) Del(key string) *Request {
	request.params.Del(key)
	return request
}

func (request *Request) Execute(v interface{}) error {
	if request.params.Get("v") == "" {
		request.params.Set("v", APIVersion)
	}
	request.params.Set("access_token", request.session.token)

	resp, err := request.session.client.Get(fmt.Sprintf(APIAddress, request.method, request.params.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var vkResponse Response
	err = json.Unmarshal(body, &vkResponse)
	if err != nil {
		return err
	}
	
	if vkResponse.Error.Code != 0 {
		return vkResponse.Error
	}

	return vkResponse.Unmarshal(&v)
}
