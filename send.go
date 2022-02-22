package tonweb

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type Send struct {
	*Client
}

func (s Send) SendBoc(b SendBocBody) (*Response, error) {
	url := s.url()
	url.Path = s.BasePath + "/sendBoc"

	body, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
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

func (s Send) SendQuery(b SendQueryBody) (*Response, error) {
	url := s.url()
	url.Path = s.BasePath + "/sendQuery"

	body, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
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

func (s Send) EstimateFee(b EstimateFeeBody) (*Response, error) {
	url := s.url()
	url.Path = s.BasePath + "/estimateFee"

	body, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	resp, err := s.client.Do(req)
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

type SendBocBody struct {
	Boc string `json:"boc"`
}

type SendQueryBody struct {
	Address  string `json:"address"`
	Body     string `json:"body"`
	InitCode string `json:"init_code"`
	InitData string `json:"init_data"`
}

type EstimateFeeBody struct {
	Address      string `json:"address"`
	Body         string `json:"body"`
	InitCode     string `json:"init_code"`
	InitData     string `json:"init_data"`
	IgnoreChksig bool   `json:"ignore_chksig"`
}
