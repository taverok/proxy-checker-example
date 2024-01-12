package dto

import "fmt"

type ProxyResponse struct {
	Status   int    `json:"status"`
	OK       bool   `json:"ok"`
	Content  string `json:"content"`
	Target   string `json:"target"`
	Proxy    string `json:"proxy"`
	Location string `json:"location"`
	Error    string `json:"error"`
}

func (it *ProxyResponse) SetError(err error) {
	it.Error = err.Error()
}

func (it *ProxyResponse) CliRepr() string {
	msg := fmt.Sprintf("Host: %s OK: %t", it.Target, it.OK)

	if it.Proxy != "" {
		msg += fmt.Sprintf(" Proxy:%s Location:%s", it.Proxy, it.Location)
	}
	if it.Error != "" {
		msg += fmt.Sprintf(" Error: %s StatusCode:%d", it.Error, it.Status)
	}

	if it.Content != "" {
		msg += fmt.Sprintln(it.Content)
	}

	msg += "\n"

	return msg
}
