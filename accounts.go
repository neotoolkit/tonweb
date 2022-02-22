package tonweb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Accounts struct {
	*Client
}

func (a Accounts) GetAddressInformation(address string) (*Response, error) {
	url := a.url()
	url.Path = a.BasePath + "/getAddressInformation"
	url.RawQuery = fmt.Sprintf("address=%s", address)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := a.client.Do(&req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (a Accounts) GetExtendedAddressInformation(address string) (*Response, error) {
	url := a.url()
	url.Path = a.BasePath + "/getExtendedAddressInformation"
	url.RawQuery = fmt.Sprintf("address=%s", address)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := a.client.Do(&req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (a Accounts) GetWalletInformation(address string) (*Response, error) {
	url := a.url()
	url.Path = a.BasePath + "/getWalletInformation"
	url.RawQuery = fmt.Sprintf("address=%s", address)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := a.client.Do(&req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (a Accounts) GetTransactions(q GetTransactionsQuery) (*Response, error) {
	url := a.url()
	url.Path = a.BasePath + "/getTransactions"
	url.RawQuery = fmt.Sprintf("address=%s&limit=%v&lt=%v&hash=%v&to_lt=%v&archival=%v", q.Address, q.Limit, q.LogicalTime, q.Hash, q.ToLogicalTime, q.Archival)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := a.client.Do(&req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (a Accounts) GetAddressBalance(address string) (*Response, error) {
	url := a.url()
	url.Path = a.BasePath + "/getAddressBalance"
	url.RawQuery = fmt.Sprintf("address=%s", address)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := a.client.Do(&req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (a Accounts) GetAddressState(address string) (*Response, error) {
	url := a.url()
	url.Path = a.BasePath + "/getAddressState"
	url.RawQuery = fmt.Sprintf("address=%s", address)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := a.client.Do(&req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (a Accounts) PackAddress(address string) (*Response, error) {
	url := a.url()
	url.Path = a.BasePath + "/packAddress"
	url.RawQuery = fmt.Sprintf("address=%s", address)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := a.client.Do(&req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (a Accounts) UnpackAddress(address string) (*Response, error) {
	url := a.url()
	url.Path = a.BasePath + "/unpackAddress"
	url.RawQuery = fmt.Sprintf("address=%s", address)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := a.client.Do(&req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (a Accounts) DetectAddress(address string) (*Response, error) {
	url := a.url()
	url.Path = a.BasePath + "/detectAddress"
	url.RawQuery = fmt.Sprintf("address=%s", address)

	var req http.Request

	req.Method = http.MethodGet
	req.URL = url

	resp, err := a.client.Do(&req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Response

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

type GetTransactionsQuery struct {
	Address       string
	Limit         int
	LogicalTime   int
	Hash          string
	ToLogicalTime int
	Archival      bool
}
