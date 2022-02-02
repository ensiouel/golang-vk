package vk

import (
	"encoding/json"
	"fmt"
)

type Error struct {
	Code          int    `json:"error_code"`
	Message       string `json:"error_msg"`
	RequestParams []struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	} `json:"request_params"`
}

func (error Error) Error() string {
	return fmt.Sprintf("Error code: %d, Error msg: %s\n", error.Code, error.Message)
}

type Response struct {
	Response json.RawMessage `json:"response"`
	Error    Error           `json:"error"`
}

func (response *Response) Unmarshal(v interface{}) error {
	return json.Unmarshal(response.Response, v)
}
