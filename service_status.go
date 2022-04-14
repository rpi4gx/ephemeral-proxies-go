package ephemeralproxies

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type AvailabilityTotal struct {
	NumberOfProxies int `json:"proxies"`
}

type Country struct {
	CountryISO      string `json:"country_iso"`
	NumberOfProxies int    `json:"proxies"`
}

type Availability struct {
	Total     AvailabilityTotal `json:"total"`
	Countries []Country         `json:"by_country"`
}

// ServiceStatus holds the /service_status response
type ServiceStatus struct {
	Availability Availability `json:"availability"`
}

type serviceStatusApiResponse struct {
	Success       bool          `json:"success"`
	ServiceStatus ServiceStatus `json:"service_status"`
	Message       string        `json:"message"`
}

func (ss *ServiceStatus) String() string {
	r, _ := json.MarshalIndent(ss, "", "    ")
	return string(r)
}

// GetServiceStatus returns service status of Ephemeral Proxies API
func GetServiceStatus(apiKey string) (*ServiceStatus, error) {
	url := "https://ephemeral-proxies.p.rapidapi.com/v1/service_status"
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

	var ss serviceStatusApiResponse
	if err := json.Unmarshal(body, &ss); err != nil {
		return nil, err
	}
	if !ss.Success {
		return nil, errors.New("api failure: " + ss.Message)
	}

	return &ss.ServiceStatus, nil
}
