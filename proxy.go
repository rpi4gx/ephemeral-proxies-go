package proxy

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// TODO: complete
type Proxy struct {
	Id         string
	Host       string
	Port       int
	Expires_at time.Time
}
type proxyApiResponse struct {
	Success bool   `json:"success"`
	Proxy   Proxy  `json:"proxy"`
	Message string `json:"message"`
}

func GetProxy(apiKey string) (*Proxy, error) {
	url := "https://ephemeral-proxies.p.rapidapi.com/v1/proxy"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "ephemeral-proxies.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", apiKey)

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

	return &p.Proxy, nil
}
