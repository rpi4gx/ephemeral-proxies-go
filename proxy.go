package proxy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// # TODO: complete
type ProxyVisibility struct {
	Ip         string `json:"ip"`
	Country    string `json:"country"`
	CountryISO string `json:"country_iso"`
}

type Proxy struct {
	Id             string          `json:"id"`
	Host           string          `json:"host"`
	Port           int             `json:"port"`
	ExpirationTime time.Time       `json:"expirest_at"` // TODO: fix time parsing
	Visibility     ProxyVisibility `json:"visibility"`
	apiKey         string
}
type proxyApiResponse struct {
	Success bool   `json:"success"`
	Proxy   Proxy  `json:"proxy"`
	Message string `json:"message"`
}

func (p *Proxy) String() string {
	r, _ := json.MarshalIndent(p, "", "    ")
	return string(r)
}

func NewProxy(apiKey string) (*Proxy, error) {
	return NewProxyWithOptions(apiKey, []string{}, "")
}

func NewProxyWithOptions(apiKey string, countriesISO []string, extraWhitelistIp string) (*Proxy, error) {
	url := "https://ephemeral-proxies.p.rapidapi.com/v1/proxy"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "ephemeral-proxies.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", apiKey)

	if len(countriesISO) > 0 {
		query := req.URL.Query()
		query.Add("countries", strings.Join(countriesISO, ","))
		req.URL.RawQuery = query.Encode()
	}

	if len(extraWhitelistIp) > 0 {
		query := req.URL.Query()
		query.Add("whitelist_ip", extraWhitelistIp)
		req.URL.RawQuery = query.Encode()
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.New("http response: " + strconv.Itoa(res.StatusCode))
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var p proxyApiResponse
	if err := json.Unmarshal(body, &p); err != nil {
		return nil, err
	}
	if !p.Success {
		return nil, errors.New("api failure: " + p.Message)
	}

	p.Proxy.apiKey = apiKey
	return &p.Proxy, nil
}

func (p *Proxy) ExtendExpirationTime() error {
	return nil
}
