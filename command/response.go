package command

import (
	"io/ioutil"
	"encoding/json"
	"io"
	"strings"
)

const (
	SuccessResponse = "$OK"
	FailureResponse = "$NK"
)

type Response struct {
	Command  string `json:"cmd"`
	Response string `json:"ret"`
}

func newResponseFromHttp(httpResponse io.ReadCloser) (response *Response, err error) {
	body, err := ioutil.ReadAll(httpResponse)

	if err != nil {
		return
	}

	err = json.Unmarshal(body, &response)

	if err != nil {
		return
	}

	response.Response = strings.Split(response.Response, "^")[0]

	return
}
