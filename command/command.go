package command

import (
	"fmt"
	"net/http"
	"net/url"
)

const (
	urlTemplate = "http://%s/r?json=1&rapi=$%s"
	UriTemplate = "/r?json=1&rapi=$%s"
)

type Command struct {
	Type Type
	Value string
}

func (c Command) SendRequest(host string, parameters ...string) (*Response, error) {
	parametersString := fmt.Sprintf("%s", c.Type)

	for _, parameter := range parameters {
		parametersString += fmt.Sprintf(" %s", parameter)
	}

	parametersString = url.QueryEscape(parametersString)

	urlRequest := fmt.Sprintf(urlTemplate, host, parametersString)
	response, err := http.Get(urlRequest)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return newResponseFromHttp(response.Body)
}
