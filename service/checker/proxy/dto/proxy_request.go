package dto

type ProxyRequest struct {
	Target      string   `json:"target"`
	ShowContent bool     `json:"showContent"`
	Proxies     []string `json:"proxies"`
}
