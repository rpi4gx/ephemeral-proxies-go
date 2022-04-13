package ephemeralproxies

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ProxyVisibility struct {
	Ip                string  `json:"ip"`
	Country           string  `json:"country"`
	CountryISO        string  `json:"country_iso"`
	IsCountryInEurope bool    `json:"country_eu"`
	Latitude          float32 `json:"latitude"`
	Longitude         float32 `json:"longitude"`
	Timezone          string  `json:"timezone"`
	ASN               string  `json:"asn"`
	ASNOrganization   string  `json:"asn_org"`
	ZipCode           string  `json:"zip_code"`
	RegionName        string  `json:"region_name"`
	RegionCode        string  `json:"region_code"`
	City              string  `json:"city"`
}

type ProxyFeaturesSupportedProtocols struct {
	Socks4 bool `json:"sock4"`
	Socks5 bool `json:"sock5"`
	Http   bool `json:"http"`
	Https  bool `json:"https"`
}

type ProxyFeatures struct {
	IsStatic           bool                            `json:"static"`
	SupportedProtocols ProxyFeaturesSupportedProtocols `json:"supported_protocols"`
}

type Proxy struct {
	Id             string          `json:"id"`
	Host           string          `json:"host"`
	Port           int             `json:"port"`
	ExpirationTime time.Time       `json:"expires_at"`
	WhitelistedIps []string        `json:"whitelisted_ips"`
	Visibility     ProxyVisibility `json:"visibility"`
	Features       ProxyFeatures   `json:"features"`
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

func processProxyApiResponse(req *http.Request) (*proxyApiResponse, error) {
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
	return &p, nil
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

	p, err := processProxyApiResponse(req)
	if err != nil {
		return nil, err
	}

	p.Proxy.apiKey = apiKey
	return &p.Proxy, nil
}

func (proxy *Proxy) ExtendExpirationTime() error {
	apiKey := proxy.apiKey
	url := "https://ephemeral-proxies.p.rapidapi.com/v1/extend_proxy"
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Host", "ephemeral-proxies.p.rapidapi.com")
	req.Header.Add("X-RapidAPI-Key", apiKey)

	query := req.URL.Query()
	query.Add("id", proxy.Id)
	req.URL.RawQuery = query.Encode()

	p, err := processProxyApiResponse(req)
	if err != nil {
		return err
	}

	*proxy = p.Proxy
	proxy.apiKey = apiKey
	return nil
}
