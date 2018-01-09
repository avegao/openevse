package command

import (
	"fmt"
	"net/http"
)

const (
	urlTemplate = "http://%s/r?json=1&rapi=$%s"
	UriTemplate = "/r?json=1&rapi=$%s"
)

type Command struct {
	Type Type
	Value string
}

func (c Command) SendRequest(host string) (*Response, error) {
	url := fmt.Sprintf(urlTemplate, host, c.Type)
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	return newResponseFromHttp(response.Body)
}