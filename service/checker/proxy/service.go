package proxy

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/taverok/proxy-checker-example/service/checker/config"
	"github.com/taverok/proxy-checker-example/service/checker/proxy/dto"
)

type Service struct {
	Cfg  *config.Config
	Repo *Repo
}

func (it *Service) CheckRequest(r *dto.ProxyRequest) ([]*dto.ProxyResponse, error) {
	result := make([]*dto.ProxyResponse, 0, len(r.Proxies))

	if len(r.Proxies) == 0 {
		check, err := it.Check(r.Target, r.ShowContent, nil)
		if err != nil {
			check.Error = err.Error()
		}

		result = append(result, check)
	}

	for _, host := range r.Proxies {
		proxy, err := it.Repo.GetByHost(host)
		if err != nil {
			return nil, err
		}

		check, err := it.Check(r.Target, r.ShowContent, proxy)
		if err != nil {
			check.Error = err.Error()
		}

		result = append(result, check)
	}

	return result, nil
}

func (it *Service) Check(target string, showContent bool, p *Proxy) (*dto.ProxyResponse, error) {
	target = normalizeTarget(target)
	client := http.Client{}

	result := &dto.ProxyResponse{
		Target: target,
	}

	if p != nil {
		client.Transport = &http.Transport{
			Proxy: http.ProxyURL(&url.URL{
				Scheme: p.Protocol,
				User:   url.UserPassword(p.Login, p.Password),
				Host:   fmt.Sprintf("%s:%d", p.IP, p.Port),
			}),
		}
		result.Proxy = p.Domain
		result.Location = p.Location
	}

	uri, err := url.Parse(target)
	if err != nil {
		return nil, err
	}

	req := http.Request{
		Method: http.MethodGet,
		URL:    uri,
	}

	res, err := client.Do(&req)
	if err != nil {
		result.SetError(err)
	}
	rawBody, err := io.ReadAll(res.Body)
	if err != nil {
		result.SetError(err)
	}
	defer res.Body.Close()

	result.Status = res.StatusCode
	if showContent {
		result.Content = string(rawBody)
	}
	if res.StatusCode < 500 {
		result.OK = true
	}

	return result, nil
}

func normalizeTarget(target string) string {
	if !strings.HasPrefix(target, "http") {
		target = "http://" + target
	}

	return target
}
