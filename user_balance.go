package ephemeralproxies

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Balance struct {
	ConsumedMegabytes int `json:"consumed_megabytes"`
	LimitMegabytes    int `json:"limit_megabytes"`
}

type userBalanceApiResponse struct {
	Success bool    `json:"success"`
	Balance Balance `json:"balance"`
	Message string  `json:"message"`
}

func (balance *Balance) String() string {
	r, _ := json.MarshalIndent(balance, "", "    ")
	return string(r)
}

// GetUserBalance returns the monthly balance left for the user
func GetUserBalance(apiKey string, proxyType ProxyType) (*Balance, error) {
	if proxyType != Residential {
		return nil, errors.New("Proxy type " + proxyType.String() + " does not support retrieving user's balance")
	}
	url := "https://ephemeral-proxies.p.rapidapi.com/v2/" + proxyType.String() + "/balance"
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

	var balanceRespone userBalanceApiResponse
	if err := json.Unmarshal(body, &balanceRespone); err != nil {
		return nil, err
	}
	if !balanceRespone.Success {
		return nil, errors.New("api failure: " + balanceRespone.Message)
	}

	return &balanceRespone.Balance, nil
}
